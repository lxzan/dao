package rapid

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

func (c *Rapid[T]) Collect(ptr Pointer) {
	var node = &c.Buckets[ptr]
	node.Ptr = 0
	node.NextPtr = 0
	node.PrevPtr = 0
	c.Recyclable.Push(ptr)
}

type Rapid[T any] struct {
	Length     int
	Serial     uint32
	Recyclable array_stack // do not recycle head
	Buckets    []Iterator[T]
	Equal      func(a, b *T) bool
}

func New[T any](size uint32, eq func(a, b *T) bool) *Rapid[T] {
	return &Rapid[T]{
		Serial:     1,
		Recyclable: []Pointer{},
		Buckets:    make([]Iterator[T], size+1),
		Length:     0,
		Equal:      eq,
	}
}

func (c Rapid[T]) Begin(entrypoint *EntryPoint) *Iterator[T] {
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
	if entrypoint.Head == 0 {
		var ptr = c.NextID()
		entrypoint.Head = ptr
		entrypoint.Tail = ptr
	}

	var head = &c.Buckets[entrypoint.Head]
	if head.Ptr == 0 {
		c.Length++
		c.Buckets[entrypoint.Head] = Iterator[T]{
			Ptr:     entrypoint.Head,
			PrevPtr: 0,
			NextPtr: 0,
			Data:    *data,
		}
		return false
	}

	for i := c.Begin(entrypoint); !c.End(i); i = c.Next(i) {
		if c.Equal(&i.Data, data) {
			i.Data = *data
			return true
		}
	}

	var cursor = c.NextID()
	var tail = &c.Buckets[entrypoint.Tail]
	tail.NextPtr = cursor
	entrypoint.Tail = cursor
	c.Buckets[cursor] = Iterator[T]{
		Ptr:     cursor,
		PrevPtr: tail.Ptr,
		NextPtr: 0,
		Data:    *data,
	}
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

	// delete last node
	if target.NextPtr == 0 && target.PrevPtr == 0 {
		entrypoint.Head = 0
		entrypoint.Tail = 0
		c.Collect(target.Ptr)
		return true
	}

	// delete head
	if target.PrevPtr == 0 {
		var next = &c.Buckets[target.NextPtr]
		entrypoint.Head = next.Ptr
		next.PrevPtr = 0
		c.Collect(target.Ptr)
		return true
	}

	// delete tail
	if target.NextPtr == 0 {
		var prev = &c.Buckets[target.PrevPtr]
		entrypoint.Tail = prev.Ptr
		prev.NextPtr = 0
		c.Collect(target.Ptr)
		return true
	}

	var prev = &c.Buckets[target.PrevPtr]
	var next = &c.Buckets[target.NextPtr]
	next.PrevPtr = prev.Ptr
	prev.NextPtr = next.Ptr
	c.Collect(target.Ptr)
	return true
}

func (c *Rapid[T]) Find(entrypoint *EntryPoint, data *T) (result *Iterator[T], exist bool) {
	if entrypoint.Head == 0 {
		return nil, false
	}
	for i := c.Begin(entrypoint); !c.End(i); i = c.Next(i) {
		if c.Equal(&i.Data, data) {
			return i, true
		}
	}
	return nil, false
}
