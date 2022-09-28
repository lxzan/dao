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

type MList[K comparable, V any] struct {
	Length     int
	Serial     uint32
	Recyclable arrayStack // do not recycle head
	Buckets    []Iterator[K, V]
}

func NewMList[K comparable, V any](size uint32) *MList[K, V] {
	return &MList[K, V]{
		Serial:     1,
		Recyclable: []Pointer{},
		Buckets:    make([]Iterator[K, V], size+1),
		Length:     0,
	}
}

func (c *MList[K, V]) Collect(ptr Pointer) {
	c.Buckets[ptr] = Iterator[K, V]{}
	c.Recyclable.Push(ptr)
}

func (c *MList[K, V]) Begin(ptr Pointer) *Iterator[K, V] {
	return &c.Buckets[ptr]
}

func (c *MList[K, V]) Next(iter *Iterator[K, V]) *Iterator[K, V] {
	return &c.Buckets[iter.NextPtr]
}

func (c *MList[K, V]) End(iter *Iterator[K, V]) bool {
	return iter.Ptr == 0
}

// NextID apply a pointer
func (c *MList[K, V]) NextID() Pointer {
	if c.Recyclable.Len() > 0 {
		return c.Recyclable.Pop()
	}

	var ptr = c.Serial
	if ptr >= uint32(len(c.Buckets)) {
		var ele Iterator[K, V]
		c.Buckets = append(c.Buckets, ele)
	}
	c.Serial++
	return Pointer(ptr)
}

// Push append an Iterator[] with unique check
func (c *MList[K, V]) Push(entrypoint *Pointer, key K, value V) (replaced bool) {
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
func (c *MList[K, V]) Delete(entrypoint *Pointer, target *Iterator[K, V]) (deleted bool) {
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

func (c *MList[K, V]) Find(entrypoint Pointer, key K) (value V, exist bool) {
	for i := c.Begin(entrypoint); !c.End(i); i = c.Next(i) {
		if i.Key == key {
			return i.Value, true
		}
	}
	return value, false
}

type arrayStack []Pointer

func (c *arrayStack) Len() int {
	return len(*c)
}

func (c *arrayStack) Push(v Pointer) {
	*c = append(*c, v)
}

func (c *arrayStack) Pop() Pointer {
	var n = c.Len()
	if n >= 1 {
		var result = (*c)[n-1]
		*c = (*c)[:n-1]
		return result
	}
	return 0
}
