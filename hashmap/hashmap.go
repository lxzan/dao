package hashmap

import (
	"github.com/lxzan/dao"
	"github.com/lxzan/dao/internal/hash"
	"github.com/lxzan/dao/rapid"
	"unsafe"
)

const (
	offset32 = 2166136261
	prime32  = 16777619
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
		storage:     rapid.New[Pair[K, V]](size[0]),
	}
}

func (c *HashMap[K, V]) Len() int {
	return c.storage.Length
}

func (c *HashMap[K, V]) SetLoadFactor(x float64) *HashMap[K, V] {
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
func (c *HashMap[K, V]) Insert(key K, val V) (replaced bool) {
	c.increase()
	var hashCode = c.Hash(&key)
	var idx = hashCode & (c.size - 1)
	var entrypoint = &c.indexes[idx]
	if entrypoint.Head == 0 {
		var ptr = c.storage.NextID()
		entrypoint.Head = ptr
		entrypoint.Tail = ptr
	}
	var data = Pair[K, V]{
		hashCode: hashCode,
		Key:      key,
		Val:      val,
	}
	replaced = c.storage.Push(entrypoint, &data)
	return replaced
}

// find one
func (c *HashMap[K, V]) Find(key K) (val V, exist bool) {
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
func (c *HashMap[K, V]) ForEach(fn func(item *Pair[K, V]) (continued bool)) {
	var n = len(c.storage.Buckets)
	for i := 1; i < n; i++ {
		var item = &c.storage.Buckets[i]
		if item.Ptr != 0 {
			if !fn(&item.Data) {
				break
			}
		}
	}
}

func (c *HashMap[K, V]) Keys() []K {
	var keys = make([]K, 0)
	c.ForEach(func(item *Pair[K, V]) bool {
		keys = append(keys, item.Key)
		return true
	})
	return keys
}

func (c *HashMap[K, V]) Values() []V {
	var values = make([]V, 0)
	c.ForEach(func(item *Pair[K, V]) bool {
		values = append(values, item.Val)
		return true
	})
	return values
}
