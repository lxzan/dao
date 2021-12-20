package rapid

import "github.com/lxzan/dao"

type (
	Pointer uint32

	EntryPoint struct {
		Head Pointer
		Tail Pointer
	}

	Iterator[T any] struct {
		Ptr     Pointer
		PrevPtr Pointer
		NextPtr Pointer
		Data    T
	}
)

func (c *Iterator[T]) Reset() {
	c.Ptr = 0
	c.NextPtr = 0
}

type Rapid[T dao.Equaler[T]] struct {
	Serial     uint32
	Recyclable array_stack // do not recycle head
	Buckets    []Iterator[T]
	Length     int
}

func New[T dao.Equaler[T]](size ...uint32) *Rapid[T] {
	if len(size) == 0 {
		size = []uint32{8}
	}
	return &Rapid[T]{
		Serial:     1,
		Recyclable: []Pointer{},
		Buckets:    make([]Iterator[T], size[0]+1),
		Length:     0,
	}
}

func (c Rapid[T]) Begin(entrypoint EntryPoint) *Iterator[T] {
	return &c.Buckets[entrypoint.Head]
}

func (c Rapid[T]) Next(iter *Iterator[T]) *Iterator[T] {
	return &c.Buckets[iter.NextPtr]
}

func (c Rapid[T]) End(iter *Iterator[T]) bool {
	return iter.Ptr == 0
}

// NextID apply a pointer
func (c *Rapid[T]) NextID() Pointer {
	if c.Recyclable.Len() > 0 {
		return c.Recyclable.Pop()
	}

	var result = c.Serial
	if result >= uint32(len(c.Buckets)) {
		var ele Iterator[T]
		c.Buckets = append(c.Buckets, ele)
	}
	c.Serial++
	return Pointer(result)
}

// Push append an element with unique check
func (c *Rapid[T]) Push(entrypoint *EntryPoint, data *T) (replaced bool) {
	var head = &c.Buckets[entrypoint.Head]
	if head.Ptr == 0 {
		head.Ptr = entrypoint.Head
		head.Data = *data
		c.Length++
		return false
	}

	for i := head; !c.End(i); i = c.Next(i) {
		if i.Data.Equal(data) {
			i.Data = *data
			return true
		}
	}

	var cursor = c.NextID()
	var tail = &c.Buckets[entrypoint.Tail]
	tail.NextPtr = cursor
	entrypoint.Tail = cursor
	var target = &c.Buckets[cursor]
	target.Ptr = cursor
	target.Data = *data
	target.PrevPtr = tail.Ptr
	c.Length++
	return false
}

// Append append an element without unique check
func (c *Rapid[T]) Append(entrypoint *EntryPoint, data *T) {
	var head = &c.Buckets[entrypoint.Head]
	if head.Ptr == 0 {
		head.Ptr = entrypoint.Head
		head.Data = *data
		c.Length++
		return
	}

	var cursor = c.NextID()
	var tail = &c.Buckets[entrypoint.Tail]
	tail.NextPtr = cursor
	entrypoint.Tail = cursor
	var target = &c.Buckets[cursor]
	target.Ptr = cursor
	target.Data = *data
	target.PrevPtr = tail.Ptr
	c.Length++
}

// Delete do not delete in loop if no break
func (c *Rapid[T]) Delete(entrypoint *EntryPoint, target *Iterator[T]) (deleted bool) {
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

func (c *Rapid[T]) Find(entrypoint EntryPoint, data *T) (result *Iterator[T], exist bool) {
	if entrypoint.Head == 0 {
		return nil, false
	}
	for i := c.Begin(entrypoint); !c.End(i); i = c.Next(i) {
		if i.Data.Equal(data) {
			return i, true
		}
	}
	return nil, false
}
