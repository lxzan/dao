package rapid

import "github.com/lxzan/dao"

type (
	Pointer uint32

	EntryPoint struct {
		Head Pointer
		Tail Pointer
	}

	Iterator[K dao.Hashable[K], V any] struct {
		Ptr     Pointer
		PrevPtr Pointer
		NextPtr Pointer
		Key     K
		Data    V
	}
)

func (c *Iterator[K, V]) Reset() {
	c.Ptr = 0
	c.NextPtr = 0
}

type Rapid[K dao.Hashable[K], V any] struct {
	Length     int
	Serial     uint32
	Recyclable array_stack // do not recycle head
	Buckets    []Iterator[K, V]
}

func New[K dao.Hashable[K], V any](size ...uint32) *Rapid[K, V] {
	if len(size) == 0 {
		size = []uint32{8}
	}
	return &Rapid[K, V]{
		Serial:     1,
		Recyclable: []Pointer{},
		Buckets:    make([]Iterator[K, V], size[0]+1),
		Length:     0,
	}
}

func (c Rapid[K, V]) Begin(entrypoint EntryPoint) *Iterator[K, V] {
	return &c.Buckets[entrypoint.Head]
}

func (c Rapid[K, V]) Next(iter *Iterator[K, V]) *Iterator[K, V] {
	return &c.Buckets[iter.NextPtr]
}

func (c Rapid[K, V]) End(iter *Iterator[K, V]) bool {
	return iter.Ptr == 0
}

// NextID apply a pointer
func (c *Rapid[K, V]) NextID() Pointer {
	if c.Recyclable.Len() > 0 {
		return c.Recyclable.Pop()
	}

	var result = c.Serial
	if result >= uint32(len(c.Buckets)) {
		var ele Iterator[K, V]
		c.Buckets = append(c.Buckets, ele)
	}
	c.Serial++
	return Pointer(result)
}

// Push append an element with unique check
func (c *Rapid[K, V]) Push(entrypoint *EntryPoint, key *K, val *V) (replaced bool) {
	var head = &c.Buckets[entrypoint.Head]
	if head.Ptr == 0 {
		head.Ptr = entrypoint.Head
		head.Key = *key
		head.Data = *val
		c.Length++
		return false
	}
	for i := head; !c.End(i); i = c.Next(i) {
		if *key == i.Key {
			i.Data = *val
			return true
		}
	}

	var cursor = c.NextID()
	var tail = &c.Buckets[entrypoint.Tail]
	tail.NextPtr = cursor
	entrypoint.Tail = cursor
	var target = &c.Buckets[cursor]
	target.Ptr = cursor
	target.Key = *key
	target.Data = *val
	target.PrevPtr = tail.Ptr
	c.Length++
	return false
}

// Append append an element without unique check
func (c *Rapid[K, V]) Append(entrypoint *EntryPoint, key *K, val *V) {
	var head = &c.Buckets[entrypoint.Head]
	if head.Ptr == 0 {
		head.Ptr = entrypoint.Head
		head.Key = *key
		head.Data = *val
		c.Length++
		return
	}

	var cursor = c.NextID()
	var tail = &c.Buckets[entrypoint.Tail]
	tail.NextPtr = cursor
	entrypoint.Tail = cursor
	var target = &c.Buckets[cursor]
	target.Ptr = cursor
	target.Key = *key
	target.Data = *val
	target.PrevPtr = tail.Ptr
	c.Length++
}

// Delete do not delete in loop if no break
func (c *Rapid[K, V]) Delete(entrypoint *EntryPoint, target *Iterator[K, V]) (deleted bool) {
	var head = c.Buckets[entrypoint.Head]
	if head.Ptr == 0 || target == nil || target.Ptr == 0 {
		return false
	}

	c.Length--
	if target.NextPtr == 0 {
		if target.PrevPtr != 0 {
			var prev = &c.Buckets[target.PrevPtr]
			prev.NextPtr = 0
			entrypoint.Tail = prev.Ptr
			c.Recyclable.Push(target.Ptr)
		}
		target.Reset()
		return true
	}

	var next = &c.Buckets[target.NextPtr]
	c.Recyclable.Push(next.Ptr)
	next.Ptr = target.Ptr
	next.PrevPtr = target.PrevPtr
	*target = *next
	next.Reset()
	if target.NextPtr == 0 {
		entrypoint.Tail = target.Ptr
	}
	return true
}

func (c *Rapid[K, V]) Find(entrypoint EntryPoint, key *K) (result *Iterator[K, V], exist bool) {
	if entrypoint.Head == 0 {
		return nil, false
	}
	for i := c.Begin(entrypoint); !c.End(i); i = c.Next(i) {
		if *key == i.Key {
			return i, true
		}
	}
	return nil, false
}
