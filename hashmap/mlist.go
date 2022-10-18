package hashmap

const bucketSize = 4

type (
	Pointer uint32

	bucket[K comparable, V any] struct {
		Len     int
		Ptr     Pointer
		PrevPtr Pointer
		NextPtr Pointer
		Keys    [bucketSize]K
		Values  [bucketSize]V
	}

	element[K comparable, V any] struct {
		Key      K
		Value    V
		HashCode uint64
	}
)

func (c *bucket[K, V]) Full() bool {
	return c.Len == bucketSize
}

type mList[K comparable, V any] struct {
	EmptyKey   K
	EmptyValue V
	Length     int
	Serial     uint32
	Recyclable arrayStack // do not recycle head
	Buckets    []bucket[K, V]
}

func newMList[K comparable, V any](size uint32) *mList[K, V] {
	return &mList[K, V]{
		Serial:     1,
		Recyclable: nil,
		Buckets:    make([]bucket[K, V], size+1, size+1),
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
	c.Buckets[ptr] = bucket[K, V]{}
	c.Recyclable.Push(ptr)
}

func (c *mList[K, V]) Begin(ptr Pointer) *bucket[K, V] {
	return &c.Buckets[ptr]
}

func (c *mList[K, V]) Next(iter *bucket[K, V]) *bucket[K, V] {
	return &c.Buckets[iter.NextPtr]
}

func (c *mList[K, V]) End(iter *bucket[K, V]) bool {
	return iter.Ptr == 0
}

// Push append an bucket[] with unique check
func (c *mList[K, V]) Push(entrypoint *Pointer, key K, value V) (replaced bool) {
	if *entrypoint == 0 {
		*entrypoint = c.NextID()
	}

	var head = &c.Buckets[*entrypoint]
	if head.Ptr == 0 {
		c.Length++
		head.Ptr = *entrypoint
		head.PrevPtr = 0
		head.NextPtr = 0
		head.Len = 1
		head.Keys[0] = key
		head.Values[0] = value
		return false
	}

	i := c.Begin(*entrypoint)
	for ; !c.End(i); i = c.Next(i) {
		for j := 0; j < i.Len; j++ {
			if i.Keys[j] == key {
				i.Values[j] = value
				return true
			}
		}

		if i.NextPtr == 0 {
			c.Length++
			if !i.Full() {
				i.Keys[i.Len] = key
				i.Values[i.Len] = value
				i.Len++
			} else {
				var cursor = c.NextID()
				c.Buckets[i.Ptr].NextPtr = cursor
				var dst = &c.Buckets[cursor]
				dst.Ptr = cursor
				dst.PrevPtr = i.Ptr
				dst.NextPtr = 0
				dst.Keys[0] = key
				dst.Values[0] = value
				dst.Len = 1
			}
		}
	}
	return false
}

// Delete do not delete in loop if no break
func (c *mList[K, V]) Delete(entrypoint *Pointer, target *bucket[K, V]) (deleted bool) {
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

//func (c *mList[K, V]) Find(entrypoint Pointer, key K) (value V, exist bool) {
//	for i := c.Begin(entrypoint); !c.End(i); i = c.Next(i) {
//		if i.Key == key {
//			return i.Value, true
//		}
//	}
//	return value, false
//}

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
