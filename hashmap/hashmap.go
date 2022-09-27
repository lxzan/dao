package hashmap

import (
	"github.com/lxzan/dao"
	"github.com/lxzan/dao/internal/hash"
	"github.com/lxzan/dao/rapid"
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
	loadFactor float64 // loadFactor=1.0
	cap        uint32  // cap=2^n
	b          uint64  // b=size-1
	length     int     // length of indexes
	indexes    []rapid.Pointer
	storage    *rapid.Rapid[K, V]
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

	var capacity = caps[0]
	return &HashMap[K, V]{
		loadFactor: 1.0,
		cap:        capacity,
		b:          uint64(capacity) - 1,
		indexes:    make([]rapid.Pointer, capacity, capacity),
		storage:    rapid.New[K, V](capacity),
	}
}

// Len get the length of hashmap
func (c *HashMap[K, V]) Len() int {
	return c.length + c.storage.Length
}

func (c *HashMap[K, V]) getIndex(key any) uint64 {
	var hashcode uint64
	switch val := key.(type) {
	case *string:
		hashcode = hash.HashBytes64(*(*[]byte)(unsafe.Pointer(val)))
	case *uint64:
		hashcode = *val
	case *uint:
		hashcode = uint64(*val)
	case *uint32:
		hashcode = uint64(*val)
	case *uint16:
		hashcode = uint64(*val)
	case *uint8:
		hashcode = uint64(*val)
	case *int64:
		hashcode = uint64(*val)
	case *int:
		hashcode = uint64(*val)
	case *int32:
		hashcode = uint64(*val)
	case *int16:
		hashcode = uint64(*val)
	case *int8:
		hashcode = uint64(*val)
	default:
		panic("key type not supported")
	}
	return hashcode & c.b
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
	return c.storage.Push(&c.indexes[idx], key, val)
}

// Get search if hashmap contains the key
func (c *HashMap[K, V]) Get(key K) (val V, exist bool) {
	var idx = c.getIndex(&key)
	for i := c.storage.Begin(c.indexes[idx]); !c.storage.End(i); i = c.storage.Next(i) {
		if i.Key == key {
			return i.Value, true
		}
	}
	return val, false
}

// Delete delete a element if the key exists
func (c *HashMap[K, V]) Delete(key K) (deleted bool) {
	var idx = c.getIndex(&key)
	var entrypoint = &c.indexes[idx]
	for i := c.storage.Begin(*entrypoint); !c.storage.End(i); i = c.storage.Next(i) {
		if i.Key == key {
			return c.storage.Delete(entrypoint, i)
		}
	}
	return false
}

func (c *HashMap[K, V]) ForEach(fn func(iter *Iterator[K, V])) {
	var iter = &Iterator[K, V]{}

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
