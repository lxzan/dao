package hashmap

import (
	"github.com/lxzan/dao"
	"github.com/lxzan/dao/internal/hash"
	"github.com/lxzan/dao/rapid"
	"github.com/lxzan/dao/vector"
	"unsafe"
)

type Pair[K dao.Hashable, V any] struct {
	hashCode uint32
	Key      K
	Val      V
}

type HashMap[K dao.Hashable, V any] struct {
	load_factor float64 // load_factor=1.0
	size        uint32  // cap=2^n
	indexes     []rapid.EntryPoint
	storage     *rapid.Rapid[Pair[K, V]]
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

func (c *HashMap[K, V]) Hash(key interface{}) uint32 {
	switch key.(type) {
	case string:
		var s = key.(string)
		return hash.NewFnv32(*(*[]byte)(unsafe.Pointer(&s)))
	case uint64:
		return uint32(key.(uint64))
	case uint32:
		return key.(uint32)
	case uint16:
		return uint32(key.(uint16))
	case uint8:
		return uint32(key.(uint8))
	case int64:
		return uint32(uint64(key.(int64)) & uint64(c.size-1))
	case int32:
		return uint32(key.(int32))
	case int16:
		return uint32(key.(int16))
	case int8:
		return uint32(key.(int8))
	default:
		panic("key type not supported")
	}
}

// insert with unique check
func (c *HashMap[K, V]) Set(key K, val V) (replaced bool) {
	c.increase()
	var hashCode = c.Hash(key)
	var idx = hashCode & (c.size - 1)
	var entrypoint = &c.indexes[idx]
	var data = &Pair[K, V]{hashCode: hashCode, Key: key, Val: val}
	return c.storage.Push(entrypoint, data)
}

// find one
func (c *HashMap[K, V]) Get(key K) (val V, exist bool) {
	var hashCode = c.Hash(key)
	var idx = hashCode & (c.size - 1)
	for i := c.storage.Begin(&c.indexes[idx]); !c.storage.End(i); i = c.storage.Next(i) {
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
	var entrypoint = &c.indexes[idx]
	for i := c.storage.Begin(entrypoint); !c.storage.End(i); i = c.storage.Next(i) {
		if i.Data.Key == key {
			return c.storage.Delete(entrypoint, i)
		}
	}
	return false
}

func (c *HashMap[K, V]) ForEach(fn func(key K, val V)) {
	for _, item := range c.storage.Buckets {
		if item.Ptr != 0 {
			fn(item.Data.Key, item.Data.Val)
		}
	}
}

func (c *HashMap[K, V]) Keys() vector.Vector[K] {
	var keys = make([]K, 0, c.Len())
	c.ForEach(func(key K, val V) {
		keys = append(keys, key)
	})
	return keys
}

func (c *HashMap[K, V]) Values() vector.Vector[V] {
	var values = make([]V, 0, c.Len())
	c.ForEach(func(key K, val V) {
		values = append(values, val)
	})
	return values
}
