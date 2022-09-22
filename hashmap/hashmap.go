package hashmap

import (
	"github.com/lxzan/dao/internal/hash"
	"github.com/lxzan/dao/rapid"
	"github.com/lxzan/dao/slice"
	"math"
	"unsafe"
)

type Pair[K comparable, V any] struct {
	hashCode uint32
	Key      K
	Val      V
}

type HashMap[K comparable, V any] struct {
	load_factor float64 // load_factor=1.0
	size        uint32  // cap=2^n
	indexes     []rapid.EntryPoint
	storage     *rapid.Rapid[Pair[K, V]]
}

// vol = size*load_factor
func New[K comparable, V any](size ...uint32) *HashMap[K, V] {
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
		indexes:     make([]rapid.EntryPoint, size[0], size[0]),
		storage: rapid.New(size[0], func(a, b *Pair[K, V]) bool {
			return a.Key == b.Key
		}),
	}
}

func (c *HashMap[K, V]) Len() int {
	return c.storage.Length
}

func (c *HashMap[K, V]) LoadFactor(x float64) *HashMap[K, V] {
	c.load_factor = x
	return c
}

func (c *HashMap[K, V]) Hash(key K) uint32 {
	switch unsafe.Sizeof(key) {
	case 16:
		data := *(*[]byte)(unsafe.Pointer(&key))
		return hash.NewFnv32(data)
	case 8:
		var x = *(*uint64)(unsafe.Pointer(&key))
		return uint32(x & math.MaxUint32)
	case 4:
		return *(*uint32)(unsafe.Pointer(&key))
	case 2:
		var x = *(*uint16)(unsafe.Pointer(&key))
		return uint32(x)
	default:
		var x = *(*uint8)(unsafe.Pointer(&key))
		return uint32(x)
	}
}

// insert with unique check
func (c *HashMap[K, V]) Set(key K, val V) (replaced bool) {
	c.increase()
	var hashCode = c.Hash(key)
	var idx = hashCode & (c.size - 1)
	var entrypoint = &c.indexes[idx]
	if entrypoint.Head == 0 {
		var ptr = c.storage.NextID()
		entrypoint.Head = ptr
		entrypoint.Tail = ptr
		c.storage.Buckets[entrypoint.Head] = rapid.Iterator[Pair[K, V]]{
			Ptr:  ptr,
			Data: Pair[K, V]{hashCode: hashCode, Key: key, Val: val},
		}
		c.storage.Length++
		return false
	}

	for i := &c.storage.Buckets[entrypoint.Head]; !c.storage.End(i); i = c.storage.Next(i) {
		if i.Data.Key == key {
			i.Data.Val = val
			return true
		}
	}

	var cursor = c.storage.NextID()
	var tail = &c.storage.Buckets[entrypoint.Tail]
	tail.NextPtr = cursor
	entrypoint.Tail = cursor
	c.storage.Buckets[cursor] = rapid.Iterator[Pair[K, V]]{
		Ptr:     cursor,
		PrevPtr: tail.Ptr,
		Data:    Pair[K, V]{hashCode: hashCode, Key: key, Val: val},
	}
	c.storage.Length++
	return false
}

// find one
func (c *HashMap[K, V]) Get(key K) (val V, exist bool) {
	var hashCode = c.Hash(key)
	var idx = hashCode & (c.size - 1)
	for i := c.storage.Begin(c.indexes[idx]); !c.storage.End(i); i = c.storage.Next(i) {
		if i.Data.Key == key {
			return i.Data.Val, true
		}
	}
	return val, exist
}

// delete one
func (c *HashMap[K, V]) Delete(key K) (deleted bool) {
	var hashCode = c.Hash(key)
	var idx = hashCode & (c.size - 1)
	var entrypoint = c.indexes[idx]
	if entrypoint.Head == 0 {
		return false
	}

	for i := c.storage.Begin(entrypoint); !c.storage.End(i); i = c.storage.Next(i) {
		if i.Data.Key == key {
			return c.storage.Delete(&entrypoint, i)
		}
	}
	return false
}

func (c *HashMap[K, V]) ForEach(fn func(p *Pair[K, V])) {
	var arr = slice.Slice[rapid.Iterator[Pair[K, V]]](c.storage.Buckets)
	for i := arr.Begin(); !arr.End(i); i = arr.Next(i) {
		if i.Value.Ptr != 0 {
			fn(&i.Value.Data)
		}
	}
}

func (c *HashMap[K, V]) Keys() slice.Slice[K] {
	var keys = make([]K, 0, c.Len())
	c.ForEach(func(iter *Pair[K, V]) {
		keys = append(keys, iter.Key)
	})
	return keys
}

func (c *HashMap[K, V]) Values() slice.Slice[V] {
	var values = make([]V, 0, c.Len())
	c.ForEach(func(iter *Pair[K, V]) {
		values = append(values, iter.Val)
	})
	return values
}
