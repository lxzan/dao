package hashmap

import (
	"github.com/cespare/xxhash"
	"github.com/lxzan/dao"
	"github.com/lxzan/dao/internal/mlist"
	"github.com/lxzan/dao/vector"
	"unsafe"
)

type Iterator[K dao.Hashable, V any] struct {
	broken bool
	Key    K
	Value  V
}

func (c *Iterator[K, V]) Break() {
	c.broken = true
}

type HashMap[K dao.Hashable, V any] struct {
	loadFactor float64 // load_factor=1.0
	cap        uint32  // cap=2^n
	b          uint64  // b=size-1
	length     int
	indexes    []mlist.Iterator[K, V]
	storage    *mlist.Rapid[K, V]
}

// New instantiates a hashmap
// at most one param, means initial capacity
func New[K dao.Hashable, V any](caps ...uint32) *HashMap[K, V] {
	if len(caps) == 0 {
		caps = []uint32{8}
	} else {
		var vol uint32 = 8
		for vol < caps[0] {
			vol <<= 1
		}
		caps[0] = vol
	}
	return &HashMap[K, V]{
		loadFactor: 1.0,
		cap:        caps[0],
		b:          uint64(caps[0]) - 1,
		indexes:    make([]mlist.Iterator[K, V], caps[0], caps[0]),
		storage:    mlist.New[K, V](caps[0]),
	}
}

// Reset reset the hashmap
//func (c *HashMap[K, V]) Reset() {
//	var temp = rapid.EntryPoint{}
//	for i, _ := range c.indexes {
//		c.indexes[i] = temp
//	}
//	c.storage.Reset()
//}

// Len get the length of hashmap
func (c *HashMap[K, V]) Len() int {
	return c.length + c.storage.Length
}

// Cap get the capacity of hashmap
func (c *HashMap[K, V]) Cap() int {
	return cap(c.storage.Buckets) - 1
}

func (c *HashMap[K, V]) getIndex(key any) uint64 {
	switch val := key.(type) {
	case *string:
		return (xxhash.Sum64(*(*[]byte)(unsafe.Pointer(val)))) & c.b
	case *uint64:
		return *val & c.b
	case *uint:
		return uint64(*val) & c.b
	case *uint32:
		return uint64(*val) & c.b
	case *uint16:
		return uint64(*val) & c.b
	case *uint8:
		return uint64(*val) & c.b
	case *int64:
		return uint64(*val) & c.b
	case *int:
		return uint64(*val) & c.b
	case *int32:
		return uint64(*val) & c.b
	case *int16:
		return uint64(*val) & c.b
	case *int8:
		return uint64(*val) & c.b
	default:
		panic("key type not supported")
	}
}

func (c *HashMap[K, V]) grow() {
	if float64(c.storage.Length)/float64(c.cap) > c.loadFactor {
		var m = New[K, V](c.cap * 2)
		c.ForEach(func(iter *Iterator[K, V]) {
			m.Set(iter.Key, iter.Value)
		})
		*c = *m
	}
}

// Set insert a element into the hashmap
// if key exists, value will be replaced
func (c *HashMap[K, V]) Set(key K, val V) (replaced bool) {
	c.grow()
	var idx = c.getIndex(&key)
	var header = &c.indexes[idx]
	if header.Ptr == 0 {
		c.length++
		header.Ptr = 1
		header.Key = key
		header.Value = val
		return false
	}
	if header.Key == key {
		header.Value = val
		return true
	}
	return c.storage.Push(header, key, val)
}

// Get search if hashmap contains the key
func (c *HashMap[K, V]) Get(key K) (val V, exist bool) {
	var idx = c.getIndex(&key)
	var header = &c.indexes[idx]
	if header.Key == key {
		return header.Value, true
	}

	for i := c.storage.Begin(header.NextPtr); !c.storage.End(i); i = c.storage.Next(i) {
		if i.Key == key {
			return i.Value, true
		}
	}
	return val, false
}

// Delete delete a element if the key exists
func (c *HashMap[K, V]) Delete(key K) (deleted bool) {
	var idx = c.getIndex(&key)
	var header = &c.indexes[idx]
	if header.Key == key {
		if header.NextPtr == 0 {
			header.Ptr = 0
			c.length--
		} else {
			var dst = &c.storage.Buckets[header.NextPtr]
			header.NextPtr = dst.NextPtr
			header.Key = dst.Key
			header.Value = dst.Value
			c.storage.Collect(dst.Ptr)
			c.storage.Length--
		}
		return true
	}

	for i := c.storage.Begin(header.NextPtr); !c.storage.End(i); i = c.storage.Next(i) {
		if i.Key == key {
			return c.storage.Delete(header, i)
		}
	}
	return false
}

func (c *HashMap[K, V]) ForEach(fn func(iter *Iterator[K, V])) {
	var iter = &Iterator[K, V]{}
	for _, item := range c.indexes {
		if iter.broken {
			return
		}
		if item.Ptr > 0 {
			iter.Key = item.Key
			iter.Value = item.Value
			fn(iter)
		}
	}

	for i := 1; i < int(c.storage.Serial); i++ {
		if iter.broken {
			return
		}

		var item = &c.storage.Buckets[i]
		if item.Ptr > 0 {
			iter.Key = item.Key
			iter.Value = item.Value
			fn(iter)
		}
	}
}

// Keys get all the keys of the hashmap, construct it as a dynamic array and return it
func (c *HashMap[K, V]) Keys() *vector.Vector[K] {
	var keys = vector.New[K](0, c.Len())
	c.ForEach(func(iter *Iterator[K, V]) {
		keys.Push(iter.Key)
	})
	return keys
}

// Values get all the values of the hashmap, construct it as a dynamic array and return it
func (c *HashMap[K, V]) Values() *vector.Vector[V] {
	var values = vector.New[V](0, c.Len())
	c.ForEach(func(iter *Iterator[K, V]) {
		values.Push(iter.Value)
	})
	return values
}
