package hashmap

import (
	"github.com/lxzan/dao/internal/hash"
	"github.com/lxzan/dao/rapid"
	"github.com/lxzan/dao/slice"
	"unsafe"
)

const max_uint32 = 2<<32 - 1

type Iterator[K comparable, V any] struct {
	Key  K
	Val  V
	next bool
}

func (this *Iterator[K, V]) Break() {
	this.next = false
}

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

func (this *HashMap[K, V]) Len() int {
	return this.storage.Length
}

func (this *HashMap[K, V]) LoadFactor(x float64) *HashMap[K, V] {
	this.load_factor = x
	return this
}

func (this *HashMap[K, V]) Hash(key K) uint32 {
	switch unsafe.Sizeof(key) {
	case 16:
		data := *(*[]byte)(unsafe.Pointer(&key))
		return hash.NewFnv32(data)
	case 8:
		var x = *(*uint64)(unsafe.Pointer(&key))
		return uint32(x & max_uint32)
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
func (this *HashMap[K, V]) Set(key K, val V) (replaced bool) {
	this.increase()
	var hashCode = this.Hash(key)
	var idx = hashCode & (this.size - 1)
	var entrypoint = &this.indexes[idx]
	if entrypoint.Head == 0 {
		var ptr = this.storage.NextID()
		entrypoint.Head = ptr
		entrypoint.Tail = ptr
		this.storage.Buckets[entrypoint.Head] = rapid.Iterator[Pair[K, V]]{
			Ptr:  ptr,
			Data: Pair[K, V]{hashCode: hashCode, Key: key, Val: val},
		}
		this.storage.Length++
		return false
	}

	for i := &this.storage.Buckets[entrypoint.Head]; !this.storage.End(i); i = this.storage.Next(i) {
		if i.Data.Key == key {
			i.Data.Val = val
			return true
		}
	}

	var cursor = this.storage.NextID()
	var tail = &this.storage.Buckets[entrypoint.Tail]
	tail.NextPtr = cursor
	entrypoint.Tail = cursor
	this.storage.Buckets[cursor] = rapid.Iterator[Pair[K, V]]{
		Ptr:     cursor,
		PrevPtr: tail.Ptr,
		Data:    Pair[K, V]{hashCode: hashCode, Key: key, Val: val},
	}
	this.storage.Length++
	return false
}

// find one
func (this *HashMap[K, V]) Get(key K) (val V, exist bool) {
	var hashCode = this.Hash(key)
	var idx = hashCode & (this.size - 1)
	for i := this.storage.Begin(this.indexes[idx]); !this.storage.End(i); i = this.storage.Next(i) {
		if i.Data.Key == key {
			return i.Data.Val, true
		}
	}
	return val, exist
}

// delete one
func (this *HashMap[K, V]) Delete(key K) (deleted bool) {
	var hashCode = this.Hash(key)
	var idx = hashCode & (this.size - 1)
	var entrypoint = this.indexes[idx]
	if entrypoint.Head == 0 {
		return false
	}

	for i := this.storage.Begin(entrypoint); !this.storage.End(i); i = this.storage.Next(i) {
		if i.Data.Key == key {
			return this.storage.Delete(&entrypoint, i)
		}
	}
	return false
}

// update directly
func (this *HashMap[K, V]) ForEach(fn func(iter *Iterator[K, V])) {
	var n = len(this.storage.Buckets)
	var iter = &Iterator[K, V]{
		next: true,
	}
	for i := 1; i < n; i++ {
		if !iter.next {
			break
		}

		var item = &this.storage.Buckets[i]
		if item.Ptr != 0 {
			iter.Key = item.Data.Key
			iter.Val = item.Data.Val
			fn(iter)
		}
	}
}

func (this *HashMap[K, V]) Keys() slice.Slice[K] {
	var keys = make([]K, 0, this.Len())
	this.ForEach(func(iter *Iterator[K, V]) {
		keys = append(keys, iter.Key)
	})
	return keys
}

func (this *HashMap[K, V]) Values() slice.Slice[V] {
	var values = make([]V, 0, this.Len())
	this.ForEach(func(iter *Iterator[K, V]) {
		values = append(values, iter.Val)
	})
	return values
}
