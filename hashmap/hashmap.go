package hashmap

import "unsafe"

type hashmap[K comparable[K], V any] struct {
	length   int
	serial   uint32
	capacity uint32
	indexes  []uint32
	buckets  []element[K, V]
}

func newHashMap[K comparable[K], V any](capacity ...uint32) *hashmap[K, V] {
	var size uint32 = 8
	if len(capacity) > 0 {
		size = 1
		for capacity[0] > size {
			size <<= 1
		}
	}

	var m = &hashmap[K, V]{
		length:   0,
		capacity: size,
		buckets:  make([]element[K, V], size+1),
		indexes:  make([]uint32, size),
	}
	return m
}

func (c *hashmap[K, V]) Len() int {
	return c.length
}

func (c *hashmap[K, V]) hash(key interface{}) uint32 {
	switch key.(type) {
	case *string:
		var data = *(*[]byte)(unsafe.Pointer(key.(*string)))
		return hashKey(data)
	case *int:
		var x = *(key.(*int))
		return uint32(x ^ (x >> 32))
	}
	return 0
}

func (c *hashmap[K, V]) NextID() uint32 {
	c.serial++
	return c.serial
}

func (c *hashmap[K, V]) ForEach(fn func(key K, val V)) {
	for i, _ := range c.buckets {
		var item = &c.buckets[i]
		if item.Ptr == 0 {
			continue
		}
		fn(item.Data.Key, item.Data.Val)
	}
}

func (c *hashmap[K, V]) findElement(key *K, ele *element[K, V]) (result, tail *element[K, V]) {
	if *key == ele.Data.Key {
		return ele, nil
	}
	if ele.NextPtr != 0 {
		return c.findElement(key, &c.buckets[ele.NextPtr])
	} else {
		return nil, ele
	}
}

func (c *hashmap[K, V]) Get(key K) (val V, exist bool) {
	var idx1 = c.hash(&key) & (c.capacity - 1)
	var idx2 = c.indexes[idx1]
	if idx2 == 0 {
		return val, false
	}
	if result, _ := c.findElement(&key, &c.buckets[idx2]); result != nil {
		return result.Data.Val, true
	}
	return val, false
}

// Delete key exist return true
func (c *hashmap[K, V]) Delete(key K) bool {
	var idx1 = c.hash(&key) % c.capacity
	var idx2 = c.indexes[idx1]
	if idx2 == 0 {
		return false
	}

	var dst = &c.buckets[idx2]
	var next *element[K, V]
	if dst.NextPtr != 0 {
		next = &c.buckets[dst.NextPtr]
	}
	deleted, empty := c.deleteElement(&key, nil, dst, next)
	if deleted {
		c.length--
	}
	if empty {
		c.indexes[idx1] = 0
	}
	return deleted
}

func (c *hashmap[K, V]) deleteElement(key *K, prev, cur, next *element[K, V]) (deleted, empty bool) {
	if cur == nil {
		return false, false
	}

	if cur.Data.Key == *key {
		if prev != nil && next != nil {
			prev.NextPtr = next.Ptr
			cur.Ptr = 0
			cur.NextPtr = 0
		} else if prev != nil && next == nil {
			prev.NextPtr = 0
			cur.Ptr = 0
			cur.NextPtr = 0
		} else if prev == nil && next != nil {
			*cur = *next
			next.Ptr = 0
			next.NextPtr = 0
		} else {
			empty = true
			cur.Ptr = 0
			cur.NextPtr = 0
		}
		deleted = true
		return
	}

	var last *element[K, V]
	if next != nil && next.NextPtr != 0 {
		last = &c.buckets[next.NextPtr]
	}
	return c.deleteElement(key, cur, next, last)
}

func (c *hashmap[K, V]) Set(key K, val V) {
	var hashCode = c.hash(&key)
	var idx1 = hashCode & (c.capacity - 1)
	var idx2 = c.indexes[idx1]
	if idx2 == 0 {
		var cursor = c.NextID()
		if cursor > c.capacity {
			c.incr()
			c.Set(key, val)
			return
		}

		c.indexes[idx1] = cursor
		c.buckets[cursor] = element[K, V]{
			Ptr:     cursor,
			NextPtr: 0,
			Data: entry[K, V]{
				HashCode: hashCode,
				Key:      key,
				Val:      val,
			},
		}
		c.length++
		return
	}

	var dst = &c.buckets[idx2]
	result, tail := c.findElement(&key, dst)
	if result != nil {
		result.Data.Val = val
	} else {
		var cursor = c.NextID()
		if cursor > c.capacity {
			c.incr()
			c.Set(key, val)
			return
		}

		tail.NextPtr = cursor
		c.buckets[cursor] = element[K, V]{
			Ptr: cursor,
			Data: entry[K, V]{
				HashCode: hashCode,
				Key:      key,
				Val:      val,
			},
		}
		c.length++
	}
}

func (c *hashmap[K, V]) incr() {
	var old = *c
	var m = newHashMap[K, V](2 * old.capacity)
	for i, _ := range old.buckets {
		var item = &old.buckets[i]
		if item.Ptr == 0 {
			continue
		}
		m.setByIncr(&item.Data)
	}
	*c = *m
}

func (c *hashmap[K, V]) setByIncr(pair *entry[K, V]) {
	var idx1 = pair.HashCode & (c.capacity - 1)
	var idx2 = c.indexes[idx1]
	if idx2 == 0 {
		var cursor = c.NextID()
		c.indexes[idx1] = cursor
		c.buckets[cursor] = element[K, V]{
			Ptr:     cursor,
			NextPtr: 0,
			Data:    *pair,
		}
		c.length++
		return
	}

	var dst = &c.buckets[idx2]
	result, tail := c.findElement(&pair.Key, dst)
	if result != nil {
		result.Data.Val = pair.Val
	} else {
		var cursor = c.NextID()
		tail.NextPtr = cursor
		c.buckets[cursor] = element[K, V]{
			Ptr:  cursor,
			Data: *pair,
		}
		c.length++
	}
}