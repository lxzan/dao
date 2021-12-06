package hashmap

import (
	"hash/fnv"
	"reflect"
	"unsafe"
)

type comparable[T any] interface{ type int8, int16, int32, int64, int, uint8, uint16, uint32, uint64, uint, string }

type entry[K comparable[K], V any] struct {
	Key K
	Val V
}

type hashmap[K comparable[K], V any] struct {
	length  uint32
	cap     uint32
	buckets []queue[entry[K, V]]
}

func newMap[K comparable[K], V any](cap uint32) *hashmap[K, V] {
	return &hashmap[K, V]{
		length:  0,
		cap:     cap,
		buckets: make([]queue[entry[K, V]], cap),
	}
}

func (c *hashmap[K, V]) hash(key *K) uint32 {
	switch reflect.TypeOf(key).Kind() {
	case reflect.String:
		var h = fnv.New32()
		var data = *(*[]byte)(unsafe.Pointer(key))
		h.Write(data)
		return h.Sum32()
	case reflect.Int:
		return uint32(*(*int)(unsafe.Pointer(key)))
	}
	return 0
}

func (c *hashmap[K, V]) findElement(key *K, ele *element[entry[K, V]]) *element[entry[K, V]] {
	if ele.Value.Key == *key {
		return ele
	} else if ele.next != nil {
		return c.findElement(key, ele.next)
	} else {
		return nil
	}
}

func (c *hashmap[K, V]) Set(key K, val V) {
	var dst = &c.buckets[c.hash(&key)%c.cap]
	var p = entry[K, V]{Key: key, Val: val}
	if dst.length == 0 {
		var ele = &element[entry[K, V]]{Value: p}
		dst.head = ele
		dst.tail = ele
		dst.length++
		c.length++
		return
	}

	if ele := c.findElement(&key, dst.head); ele != nil {
		ele.Value = p
	} else {
		var node = &element[entry[K, V]]{Value: p}
		dst.tail.next = node
		dst.tail = node
		dst.length++
		c.length++
	}

	//dst.Push()
}

func (c *hashmap[K, V]) Get(key K) (result V, exist bool) {
	var dst = &c.buckets[c.hash(&key)%c.cap]
	if ele := c.findElement(&key, dst.head); ele != nil {
		return ele.Value.Val, true
	}
	return result, false
}

func (c *hashmap[K, V]) ForEach(fn func(entry[K, V])) {
	for _, item := range c.buckets {
		if item.length == 0 {
			continue
		}
		item.doForEach(item.head, func(e entry[K, V]) {
			fn(e)
		})
	}
}
