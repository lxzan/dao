package hashmap

import (
	"github.com/cespare/xxhash"
	"github.com/lxzan/dao"
	"github.com/lxzan/dao/rapid"
	"github.com/lxzan/dao/vector"
	"unsafe"
)

//type Pair[K dao.Hashable, V any] struct {
//	hashCode uint32
//	Key      K
//	Val      V
//}

type HashMap[K dao.Hashable, V any] struct {
	load_factor float64 // load_factor=1.0
	size        uint32  // cap=2^n
	b           uint64  // b=size-1
	indexes     []rapid.EntryPoint
	storage     *rapid.Rapid[K, V]
}

// vol = size*load_factor
func New[K dao.Hashable, V any](size ...uint32) *HashMap[K, V] {
	if len(size) == 0 {
		size = []uint32{8}
	} else {
		var vol uint32 = 8
		for vol < size[0] {
			vol <<= 1
		}
		size[0] = vol
	}
	return &HashMap[K, V]{
		load_factor: 1.0,
		size:        size[0],
		b:           uint64(size[0]) - 1,
		indexes:     make([]rapid.EntryPoint, size[0], size[0]),
		storage:     rapid.New[K, V](size[0]),
	}
}

func (c *HashMap[K, V]) Len() int {
	return c.storage.Length
}

func (c *HashMap[K, V]) LoadFactor(x float64) *HashMap[K, V] {
	c.load_factor = x
	return c
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

// insert with unique check
func (c *HashMap[K, V]) Set(key K, val V) (replaced bool) {
	c.increase()
	var idx = c.getIndex(&key)
	return c.storage.Push(&c.indexes[idx], key, val)
}

// find one
func (c *HashMap[K, V]) Get(key K) (val V, exist bool) {
	var idx = c.getIndex(&key)
	for i := c.storage.Begin(&c.indexes[idx]); !c.storage.End(i); i = c.storage.Next(i) {
		if i.Key == key {
			return i.Value, true
		}
	}
	return val, false
}

// delete one
func (c *HashMap[K, V]) Delete(key K) (deleted bool) {
	var idx = c.getIndex(&key)
	var entrypoint = &c.indexes[idx]
	for i := c.storage.Begin(entrypoint); !c.storage.End(i); i = c.storage.Next(i) {
		if i.Key == key {
			return c.storage.Delete(entrypoint, i)
		}
	}
	return false
}

func (c *HashMap[K, V]) ForEach(fn func(key K, val V)) {
	for _, item := range c.storage.Buckets {
		if item.Ptr != 0 {
			fn(item.Key, item.Value)
		}
	}
}

func (c *HashMap[K, V]) Keys() *vector.Vector[K] {
	var keys = vector.New[K](0, c.Len())
	c.ForEach(func(key K, val V) {
		keys.Push(key)
	})
	return keys
}

func (c *HashMap[K, V]) Values() *vector.Vector[V] {
	var values = vector.New[V](0, c.Len())
	c.ForEach(func(key K, val V) {
		values.Push(val)
	})
	return values
}
