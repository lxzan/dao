package hashmap

import (
	"github.com/lxzan/dao"
	"github.com/lxzan/dao/internal/hash"
	"github.com/lxzan/dao/internal/utils"
	"github.com/lxzan/dao/vector"
)

type Iterator[K dao.Hashable, V any] struct {
	Key    K
	Value  V
	broken bool
}

func (c *Iterator[K, V]) Break() {
	c.broken = true
}

type HashMap[K dao.Hashable, V any] struct {
	cap     uint32       // cap=2^n
	mod     uint64       // b=cap-1
	indexes []Pointer    // list header pointer
	storage *mList[K, V] // list data
}

// New instantiates a hashmap
// at most one param, means initial capacity
func New[K dao.Hashable, V any](caps ...uint32) *HashMap[K, V] {
	if len(caps) == 0 {
		caps = []uint32{4}
	}

	var size = caps[0]
	var capacity = uint32(4)
	for capacity < size {
		capacity *= 2
	}

	return &HashMap[K, V]{
		cap:     capacity,
		mod:     uint64(capacity) - 1,
		indexes: make([]Pointer, capacity, capacity),
		storage: newMList[K, V](size),
	}
}

// Len get the length of hashmap
func (c *HashMap[K, V]) Len() int {
	return c.storage.Length
}

func (c *HashMap[K, V]) hash(key any) uint64 {
	var hashcode uint64
	switch val := key.(type) {
	case string:
		hashcode = hash.HashBytes64(utils.S2B(val))
	case uint64:
		hashcode = val
	case uint:
		hashcode = uint64(val)
	case uint32:
		hashcode = uint64(val)
	case uint16:
		hashcode = uint64(val)
	case uint8:
		hashcode = uint64(val)
	case int64:
		hashcode = uint64(val)
	case int:
		hashcode = uint64(val)
	case int32:
		hashcode = uint64(val)
	case int16:
		hashcode = uint64(val)
	case int8:
		hashcode = uint64(val)
	default:
		panic("key type not supported")
	}
	return hashcode
}

func (c *HashMap[K, V]) grow() {
	if c.storage.Recyclable.Len() == 0 && int(c.storage.Serial) >= cap(c.storage.Buckets) {
		var m = New[K, V](c.cap * 2)
		for i := 1; i < int(c.storage.Serial); i++ {
			var item = &c.storage.Buckets[i]
			if item.Ptr > 0 {
				m.Set(item.Key, item.Value)
			}
		}
		*c = *m
	}
}

// Set insert a element into the hashmap
// if key exists, value will be replaced
func (c *HashMap[K, V]) Set(key K, val V) {
	c.grow()
	var hashCode = c.hash(key)
	var idx = hashCode & c.mod
	c.storage.Push(&c.indexes[idx], key, val, hashCode)
}

// Get search if hashmap contains the key
func (c *HashMap[K, V]) Get(key K) (val V, exist bool) {
	var hashCode = c.hash(key)
	var idx = hashCode & c.mod
	for i := c.storage.Begin(c.indexes[idx]); !c.storage.End(i); i = c.storage.Next(i) {
		if i.HashCode == hashCode && i.Key == key {
			return i.Value, true
		}
	}
	return val, false
}

// Delete delete a element if the key exists
func (c *HashMap[K, V]) Delete(key K) (deleted bool) {
	var hashCode = c.hash(key)
	var idx = hashCode & c.mod
	for i := c.storage.Begin(c.indexes[idx]); !c.storage.End(i); i = c.storage.Next(i) {
		if i.HashCode == hashCode && i.Key == key {
			return c.storage.Delete(&c.indexes[idx], i)
		}
	}
	return false
}

func (c *HashMap[K, V]) ForEach(fn func(iter *Iterator[K, V])) {
	var iter = &Iterator[K, V]{}
	for i := 1; i < int(c.storage.Serial); i++ {
		var item = &c.storage.Buckets[i]
		if item.Ptr > 0 {
			iter.Key = item.Key
			iter.Value = item.Value
			fn(iter)
			if iter.broken {
				return
			}
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
