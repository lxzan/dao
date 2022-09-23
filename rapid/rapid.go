package rapid

type (
	Pointer uint32

	EntryPoint struct {
		Head Pointer
		Tail Pointer
	}

	Iterator[K comparable, V any] struct {
		Ptr     Pointer
		PrevPtr Pointer
		NextPtr Pointer
		Key     K
		Value   V
	}
)

type Rapid[K comparable, V any] struct {
	Length     int
	Serial     uint32
	Recyclable array_stack // do not recycle head
	Buckets    []Iterator[K, V]
}

func New[K comparable, V any](size uint32) *Rapid[K, V] {
	return &Rapid[K, V]{
		Serial:     1,
		Recyclable: []Pointer{},
		Buckets:    make([]Iterator[K, V], size+1),
		Length:     0,
	}
}

func (c *Rapid[K, V]) Collect(ptr Pointer) {
	var node = &c.Buckets[ptr]
	node.Ptr = 0
	node.NextPtr = 0
	node.PrevPtr = 0
	c.Recyclable.Push(ptr)
}

func (c *Rapid[K, V]) Begin(headPointer Pointer) *Iterator[K, V] {
	return &c.Buckets[headPointer]
}

func (c *Rapid[K, V]) Next(iter *Iterator[K, V]) *Iterator[K, V] {
	return &c.Buckets[iter.NextPtr]
}

func (c *Rapid[K, V]) End(iter *Iterator[K, V]) bool {
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

// Append append an element without unique check
func (c *Rapid[K, V]) Append(entrypoint *EntryPoint, key K, value V) {
	if entrypoint.Head == 0 {
		var ptr = c.NextID()
		entrypoint.Head = ptr
		entrypoint.Tail = ptr
	}

	var head = &c.Buckets[entrypoint.Head]
	if head.Ptr == 0 {
		c.Length++
		c.Buckets[entrypoint.Head] = Iterator[K, V]{
			Ptr:     entrypoint.Head,
			PrevPtr: 0,
			NextPtr: 0,
			Key:     key,
			Value:   value,
		}
		return
	}

	var cursor = c.NextID()
	var tail = &c.Buckets[entrypoint.Tail]
	tail.NextPtr = cursor
	entrypoint.Tail = cursor
	c.Buckets[cursor] = Iterator[K, V]{
		Ptr:     cursor,
		PrevPtr: tail.Ptr,
		NextPtr: 0,
		Key:     key,
		Value:   value,
	}
	c.Length++
	return
}

// Push append an element with unique check
func (c *Rapid[K, V]) Push(entrypoint *EntryPoint, key K, value V) (replaced bool) {
	if entrypoint.Head == 0 {
		var ptr = c.NextID()
		entrypoint.Head = ptr
		entrypoint.Tail = ptr
	}

	var head = &c.Buckets[entrypoint.Head]
	if head.Ptr == 0 {
		c.Length++
		c.Buckets[entrypoint.Head] = Iterator[K, V]{
			Ptr:     entrypoint.Head,
			PrevPtr: 0,
			NextPtr: 0,
			Key:     key,
			Value:   value,
		}
		return false
	}

	for i := c.Begin(entrypoint.Head); !c.End(i); i = c.Next(i) {
		if i.Key == key {
			i.Value = value
			return true
		}
	}

	var cursor = c.NextID()
	var tail = &c.Buckets[entrypoint.Tail]
	tail.NextPtr = cursor
	entrypoint.Tail = cursor
	c.Buckets[cursor] = Iterator[K, V]{
		Ptr:     cursor,
		PrevPtr: tail.Ptr,
		NextPtr: 0,
		Key:     key,
		Value:   value,
	}
	c.Length++
	return false
}

// Delete do not delete in loop if no break
func (c *Rapid[K, V]) Delete(entrypoint *EntryPoint, target *Iterator[K, V]) (deleted bool) {
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

func (c *Rapid[K, V]) Find(entrypointer Pointer, key K) (value V, exist bool) {
	for i := c.Begin(entrypointer); !c.End(i); i = c.Next(i) {
		if i.Key == key {
			return i.Value, true
		}
	}
	return value, false
}
