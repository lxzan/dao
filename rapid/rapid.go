package rapid

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
	c.Buckets[ptr] = Iterator[K, V]{}
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
func (c *Rapid[K, V]) Push(entrypoint *Pointer, key K, value V) (replaced bool) {
	if *entrypoint == 0 {
		*entrypoint = c.NextID()
	}
	var head = &c.Buckets[*entrypoint]
	if head.Ptr == 0 {
		c.Length++
		head.Ptr = *entrypoint
		head.PrevPtr = 0
		head.NextPtr = 0
		head.Key = key
		head.Value = value
		return false
	}

	for i := c.Begin(*entrypoint); !c.End(i); i = c.Next(i) {
		if i.Key == key {
			i.Value = value
			return true
		}
		if i.NextPtr == 0 {
			var cursor = c.NextID()
			c.Buckets[i.Ptr].NextPtr = cursor
			var dst = &c.Buckets[cursor]
			dst.Ptr = cursor
			dst.PrevPtr = i.Ptr
			dst.NextPtr = 0
			dst.Key = key
			dst.Value = value
			c.Length++
			break
		}
	}
	return false
}

// Delete do not delete in loop if no break
func (c *Rapid[K, V]) Delete(entrypoint *Pointer, target *Iterator[K, V]) (deleted bool) {
	var head = c.Buckets[*entrypoint]
	if head.Ptr == 0 || target.Ptr == 0 {
		return false
	}

	c.Length--

	// delete last node
	if target.NextPtr == 0 && target.PrevPtr == 0 {
		*entrypoint = 0
		c.Collect(target.Ptr)
		return true
	}

	// delete head
	if target.PrevPtr == 0 {
		var next = &c.Buckets[target.NextPtr]
		*entrypoint = next.Ptr
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

func (c *Rapid[K, V]) Find(entrypoint Pointer, key K) (value V, exist bool) {
	for i := c.Begin(entrypoint); !c.End(i); i = c.Next(i) {
		if i.Key == key {
			return i.Value, true
		}
	}
	return value, false
}
