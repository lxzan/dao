package hashmap

import (
	"github.com/lxzan/dao"
)

type HashMap[K dao.Hashable, V any] struct {
	m map[K]V
}

// New instantiates a hashmap
// at most one param, means initial capacity
func New[K dao.Hashable, V any](caps ...uint32) *HashMap[K, V] {
	if len(caps) == 0 {
		caps = []uint32{4}
	}

	return &HashMap[K, V]{
		m: make(map[K]V, caps[0]),
	}
}

// Len get the length of hashmap
func (c *HashMap[K, V]) Len() int {
	return len(c.m)
}

// Set insert a element into the hashmap
// if key exists, value will be replaced
func (c *HashMap[K, V]) Set(key K, val V) {
	c.m[key] = val
}

// Get search if hashmap contains the key
func (c *HashMap[K, V]) Get(key K) (val V, exist bool) {
	val, exist = c.m[key]
	return
}

// Delete delete a element if the key exists
func (c *HashMap[K, V]) Delete(key K) {
	delete(c.m, key)
}

func (c *HashMap[K, V]) ForEach(f func(key K, val V) bool) {
	for k, v := range c.m {
		if !f(k, v) {
			break
		}
	}
}

// Keys get all the keys of the hashmap, construct it as a dynamic array and return it
func (c *HashMap[K, V]) Keys() []K {
	var keys = make([]K, 0, len(c.m))
	c.ForEach(func(k K, v V) bool {
		keys = append(keys, k)
		return true
	})
	return keys
}

// Values get all the values of the hashmap, construct it as a dynamic array and return it
func (c *HashMap[K, V]) Values() []V {
	var vals = make([]V, 0, len(c.m))
	c.ForEach(func(k K, v V) bool {
		vals = append(vals, v)
		return true
	})
	return vals
}
