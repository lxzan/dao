package hashmap

import (
	"github.com/lxzan/dao"
	"github.com/lxzan/dao/internal/hash"
	"github.com/lxzan/dao/rapid"
	"github.com/lxzan/dao/slice"
	"unsafe"
)

type Pair[K dao.Hashable[K], V any] struct {
	hashCode uint32
	Key      K
	Val      V
}

func (c Pair[K, V]) Equal(x *Pair[K, V]) bool {
	return c.Key == x.Key
}

type HashMap[K dao.Hashable[K], V any] struct {
	load_factor float64 // load_factor=1.0
	size        uint32  // cap=2^n
	indexes     []rapid.EntryPoint
	storage     *rapid.Rapid[Pair[K, V]]
}

// vol = size*load_factor
func New[K dao.Hashable[K], V any](size ...uint32) *HashMap[K, V] {
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
		storage: rapid.New[Pair[K, V]](size[0], func(a, b *Pair[K, V]) bool {
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

func (c *HashMap[K, V]) Hash(key *K) uint32 {
	switch unsafe.Sizeof(*key) {
	case 16:
		data := *(*[]byte)(unsafe.Pointer(key))
		return hash.NewFnv32(data)
	case 8:
		var x = *(*uint64)(unsafe.Pointer(key))
		return uint32(x & (2<<32 - 1))
	case 4:
		return *(*uint32)(unsafe.Pointer(key))
	case 2:
		var x = *(*uint16)(unsafe.Pointer(key))
		return uint32(x)
	default:
		var x = *(*uint8)(unsafe.Pointer(key))
		return uint32(x)
	}
}

// insert with unique check
func (c *HashMap[K, V]) Set(key K, val V) (replaced bool) {
	c.increase()
	var hashCode = c.Hash(&key)
	var idx = hashCode & (c.size - 1)
	var entrypoint = &c.indexes[idx]
	if entrypoint.Head == 0 {
		var ptr = c.storage.NextID()
		entrypoint.Head = ptr
		entrypoint.Tail = ptr
	}

	//var pair = Pair[K, V]{hashCode: hashCode, Key: key, Val: val}
	//replaced = c.storage.Push(entrypoint, &pair)
	//return replaced

	var head = &c.storage.Buckets[entrypoint.Head]
	if head.Ptr == 0 {
		head.Ptr = entrypoint.Head
		head.Data = Pair[K, V]{hashCode: hashCode, Key: key, Val: val}
		c.storage.Length++
		return false
	}

	for i := head; !c.storage.End(i); i = c.storage.Next(i) {
		if i.Data.Key == key {
			i.Data.Val = val
			return true
		}
	}

	var cursor = c.storage.NextID()
	var tail = &c.storage.Buckets[entrypoint.Tail]
	tail.NextPtr = cursor
	entrypoint.Tail = cursor
	var target = &c.storage.Buckets[cursor]
	target.Ptr = cursor
	target.Data = Pair[K, V]{hashCode: hashCode, Key: key, Val: val}
	target.PrevPtr = tail.Ptr
	c.storage.Length++
	return false
}

// find one
func (c *HashMap[K, V]) Get(key K) (val V, exist bool) {
	var hashCode = c.Hash(&key)
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
	var hashCode = c.Hash(&key)
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

// update directly
func (c *HashMap[K, V]) ForEach(fn func(key K, val V) (continued bool)) {
	var n = len(c.storage.Buckets)
	for i := 1; i < n; i++ {
		var item = &c.storage.Buckets[i]
		if item.Ptr != 0 {
			if !fn(item.Data.Key, item.Data.Val) {
				break
			}
		}
	}
}

func (c *HashMap[K, V]) Keys() slice.Slice[K] {
	var keys = make([]K, 0)
	c.ForEach(func(key K, val V) bool {
		keys = append(keys, key)
		return true
	})
	return keys
}

func (c *HashMap[K, V]) Values() slice.Slice[V] {
	var values = make([]V, 0)
	c.ForEach(func(key K, val V) bool {
		values = append(values, val)
		return true
	})
	return values
}
