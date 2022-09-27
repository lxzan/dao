package hashmap

type (
	Pointer uint32

	Iterator[K comparable, V any] struct {
		broken  bool
		ptr     Pointer
		prevPtr Pointer
		nextPtr Pointer
		Key     K
		Value   V
	}
)

func (c *Iterator[K, V]) Break() {
	c.broken = true
}

type mList[K comparable, V any] struct {
	Length     int
	Serial     uint32
	Recyclable arrayStack // do not recycle head
	Buckets    []Iterator[K, V]
}

func newMList[K comparable, V any](size uint32) *mList[K, V] {
	return &mList[K, V]{
		Serial:     1,
		Recyclable: []Pointer{},
		Buckets:    make([]Iterator[K, V], size+1),
		Length:     0,
	}
}

func (c *mList[K, V]) Collect(ptr Pointer) {
	c.Buckets[ptr] = Iterator[K, V]{}
	c.Recyclable.Push(ptr)
}

func (c *mList[K, V]) Begin(ptr Pointer) *Iterator[K, V] {
	return &c.Buckets[ptr]
}

func (c *mList[K, V]) Next(iter *Iterator[K, V]) *Iterator[K, V] {
	return &c.Buckets[iter.nextPtr]
}

func (c *mList[K, V]) End(iter *Iterator[K, V]) bool {
	return iter.ptr == 0
}

// NextID apply a pointer
func (c *mList[K, V]) NextID() Pointer {
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
func (c *mList[K, V]) Push(entrypoint *Pointer, key K, value V) (replaced bool) {
	if *entrypoint == 0 {
		*entrypoint = c.NextID()
	}
	var head = &c.Buckets[*entrypoint]
	if head.ptr == 0 {
		c.Length++
		head.ptr = *entrypoint
		head.prevPtr = 0
		head.nextPtr = 0
		head.Key = key
		head.Value = value
		return false
	}

	for i := c.Begin(*entrypoint); !c.End(i); i = c.Next(i) {
		if i.Key == key {
			i.Value = value
			return true
		}
		if i.nextPtr == 0 {
			var cursor = c.NextID()
			c.Buckets[i.ptr].nextPtr = cursor
			var dst = &c.Buckets[cursor]
			dst.ptr = cursor
			dst.prevPtr = i.ptr
			dst.nextPtr = 0
			dst.Key = key
			dst.Value = value
			c.Length++
			break
		}
	}
	return false
}

// Delete do not delete in loop if no break
func (c *mList[K, V]) Delete(entrypoint *Pointer, target *Iterator[K, V]) (deleted bool) {
	var head = c.Buckets[*entrypoint]
	if head.ptr == 0 || target.ptr == 0 {
		return false
	}

	c.Length--

	// delete last node
	if target.nextPtr == 0 && target.prevPtr == 0 {
		*entrypoint = 0
		c.Collect(target.ptr)
		return true
	}

	// delete head
	if target.prevPtr == 0 {
		var next = &c.Buckets[target.nextPtr]
		*entrypoint = next.ptr
		next.prevPtr = 0
		c.Collect(target.ptr)
		return true
	}

	// delete tail
	if target.nextPtr == 0 {
		var prev = &c.Buckets[target.prevPtr]
		prev.nextPtr = 0
		c.Collect(target.ptr)
		return true
	}

	var prev = &c.Buckets[target.prevPtr]
	var next = &c.Buckets[target.nextPtr]
	next.prevPtr = prev.ptr
	prev.nextPtr = next.ptr
	c.Collect(target.ptr)
	return true
}

func (c *mList[K, V]) Find(entrypoint Pointer, key K) (value V, exist bool) {
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
