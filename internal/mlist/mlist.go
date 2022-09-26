package mlist

type (
	Pointer uint32

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
	Recyclable Stack
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

func (c *Rapid[K, V]) Reset() {
	c.Length = 0
	c.Serial = 1
	c.Recyclable = c.Recyclable[:0]
	c.Buckets = c.Buckets[:1]
}

func (c *Rapid[K, V]) Collect(ptr Pointer) {
	var node = &c.Buckets[ptr]
	node.Ptr = 0
	node.NextPtr = 0
	node.PrevPtr = 0
	c.Recyclable.Push(ptr)
}

func (c *Rapid[K, V]) Begin(ptr Pointer) *Iterator[K, V] {
	return &c.Buckets[ptr]
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

// Push append an element with unique check
func (c *Rapid[K, V]) Push(entrypoint *Iterator[K, V], key K, value V) (replaced bool) {
	if entrypoint.NextPtr == 0 {
		var ptr = c.NextID()
		entrypoint.NextPtr = ptr
		entrypoint.PrevPtr = ptr
	}

	var head = &c.Buckets[entrypoint.NextPtr]
	if head.Ptr == 0 {
		c.Length++
		head.Ptr = entrypoint.NextPtr
		head.PrevPtr = 0
		head.NextPtr = 0
		head.Key = key
		head.Value = value
		return false
	}

	for i := c.Begin(entrypoint.NextPtr); !c.End(i); i = c.Next(i) {
		if i.Key == key {
			i.Value = value
			return true
		}
	}

	var cursor = c.NextID()
	var tail = &c.Buckets[entrypoint.PrevPtr]
	tail.NextPtr = cursor
	entrypoint.PrevPtr = cursor
	var dst = &c.Buckets[cursor]
	dst.Ptr = cursor
	dst.PrevPtr = tail.Ptr
	dst.NextPtr = 0
	dst.Key = key
	dst.Value = value
	c.Length++
	return false
}

// Delete do not delete in loop if no break
func (c *Rapid[K, V]) Delete(entrypoint *Iterator[K, V], target *Iterator[K, V]) (deleted bool) {
	if entrypoint.NextPtr == 0 {
		return false
	}

	c.Length--

	// delete last node
	if target.NextPtr == 0 && target.PrevPtr == 0 {
		entrypoint.NextPtr = 0
		c.Collect(target.Ptr)
		return true
	}

	// delete head
	if target.PrevPtr == 0 {
		var next = &c.Buckets[target.NextPtr]
		entrypoint.NextPtr = next.Ptr
		next.PrevPtr = 0
		c.Collect(target.Ptr)
		return true
	}

	// delete tail
	if target.NextPtr == 0 {
		var prev = &c.Buckets[target.PrevPtr]
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

type Stack []Pointer

func (c *Stack) Len() int {
	return len(*c)
}

func (c *Stack) Push(v Pointer) {
	*c = append(*c, v)
}

func (c *Stack) Pop() Pointer {
	var n = c.Len()
	if n >= 1 {
		var result = (*c)[n-1]
		*c = (*c)[:n-1]
		return result
	}
	return 0
}
