package hashmap

import (
	"github.com/cespare/xxhash"
	"github.com/lxzan/dao"
	"github.com/lxzan/dao/rapid"
	"github.com/lxzan/dao/vector"
	"unsafe"
)

type HashMap[K dao.Hashable, V any] struct {
	caps        uint32  // caps=2^n
	b           uint64  // b=size-1
	load_factor float64 // load_factor=1.0
	indexes     []rapid.EntryPoint
	storage     *rapid.Rapid[K, V]
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
		load_factor: 1.0,
		caps:        caps[0],
		b:           uint64(caps[0]) - 1,
		indexes:     make([]rapid.EntryPoint, caps[0], caps[0]),
		storage:     rapid.New[K, V](caps[0]),
	}
}

// Len get the length of hashmap
func (c *HashMap[K, V]) Len() int {
	return c.storage.Length
}

// LoadFactor set the load_factor of hashmap
// if length/capacity>x, it will grow
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

// Set insert a element into the hashmap
// if key exists, value will be replaced
func (c *HashMap[K, V]) Set(key K, val V) (replaced bool) {
	c.increase()
	var idx = c.getIndex(&key)
	return c.storage.Push(&c.indexes[idx], key, val)
}

// Get search if hashmap contains the key
func (c *HashMap[K, V]) Get(key K) (val V, exist bool) {
	var idx = c.getIndex(&key)
	for i := c.storage.Begin(c.indexes[idx].Head); !c.storage.End(i); i = c.storage.Next(i) {
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
	for i := c.storage.Begin(entrypoint.Head); !c.storage.End(i); i = c.storage.Next(i) {
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

// Keys get all the keys of the hashmap, construct it as a dynamic array and return it
func (c *HashMap[K, V]) Keys() *vector.Vector[K] {
	var keys = vector.New[K](0, c.Len())
	c.ForEach(func(key K, val V) {
		keys.Push(key)
	})
	return keys
}

// Values get all the values of the hashmap, construct it as a dynamic array and return it
func (c *HashMap[K, V]) Values() *vector.Vector[V] {
	var values = vector.New[V](0, c.Len())
	c.ForEach(func(key K, val V) {
		values.Push(val)
	})
	return values
}
