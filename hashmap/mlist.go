package hashmap

type (
	Pointer uint32

	element[K comparable, V any] struct {
		Ptr      Pointer
		PrevPtr  Pointer
		NextPtr  Pointer
		HashCode uint64
		Key      K
		Value    V
	}
)

type mList[K comparable, V any] struct {
	EmptyKey   K
	EmptyValue V
	Length     int
	Serial     uint32
	Recyclable arrayStack // do not recycle head
	Buckets    []element[K, V]
}

func newMList[K comparable, V any](size uint32) *mList[K, V] {
	return &mList[K, V]{
		Serial:     1,
		Recyclable: nil,
		Buckets:    make([]element[K, V], size+1, size+1),
		Length:     0,
	}
}

// NextID apply a pointer
func (c *mList[K, V]) NextID() Pointer {
	if c.Recyclable.Len() > 0 {
		return c.Recyclable.Pop()
	}

	var ptr = c.Serial
	c.Serial++
	return Pointer(ptr)
}

func (c *mList[K, V]) Collect(ptr Pointer) {
	var bucket = &c.Buckets[ptr]
	bucket.Ptr = 0
	bucket.NextPtr = 0
	bucket.PrevPtr = 0
	bucket.Key = c.EmptyKey
	bucket.Value = c.EmptyValue
	bucket.HashCode = 0
	c.Recyclable.Push(ptr)
}

func (c *mList[K, V]) Begin(ptr Pointer) *element[K, V] {
	return &c.Buckets[ptr]
}

func (c *mList[K, V]) Next(iter *element[K, V]) *element[K, V] {
	return &c.Buckets[iter.NextPtr]
}

func (c *mList[K, V]) End(iter *element[K, V]) bool {
	return iter.Ptr == 0
}

// Push append an element[] with unique check
func (c *mList[K, V]) Push(entrypoint *Pointer, key K, value V, hashCode uint64) (replaced bool) {
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
		head.HashCode = hashCode
		return false
	}

	for i := c.Begin(*entrypoint); !c.End(i); i = c.Next(i) {
		if i.HashCode == hashCode && i.Key == key {
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
			dst.HashCode = hashCode
			c.Length++
			break
		}
	}
	return false
}

// Delete do not delete in loop if no break
func (c *mList[K, V]) Delete(entrypoint *Pointer, target *element[K, V]) (deleted bool) {
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
