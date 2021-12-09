package mqueue

type (
	element[T comparable[T]] struct {
		Ptr     Pointer
		PrevPtr Pointer
		NextPtr Pointer
		TailPtr Pointer
		Data    T
	}

	mqueue[T comparable[T]] struct {
		length  int
		serial  uint32
		buckets []*element[T]
		unused  *queue
	}
)

func new_mqueue[T comparable[T]](cap ...uint32) *mqueue[T] {
	var capacity uint32 = 8
	if len(cap) > 0 {
		capacity = cap[0]
	}
	return &mqueue[T]{
		length:  0,
		serial:  1,
		buckets: make([]*element[T], 1, capacity+1),
		unused:  array_stack(make[]uint32, 0)),
	}
}

func (c *mqueue[T]) Get(ptr Pointer) *element[T] {
	return c.buckets[ptr]
}

func (c *mqueue[T]) Next(ele *element[T]) *element[T] {
	return c.buckets[ele.NextPtr]
}

func (c *mqueue[T]) End(ele *element[T]) bool {
	return ele.NextPtr == 0
}

func (c *mqueue[T]) Empty(ele *element[T]) bool {
	return ele == nil || ele.Ptr == 0
}

func (c *mqueue[T]) Len() int {
	return c.length
}

// NextID apply pointer
func (c *mqueue[T]) NextID() Pointer {
	if ptr := c.unused.Pop(); ptr != 0 {
		return ptr
	}

	var ele element
[T]
c.buckets = append(c.buckets, &ele)
result := c.serial
c.serial++
return Pointer(result)
}

// Push append an element
func (c *mqueue[T]) Push(headPtr Pointer, v T) {
	c.length++
	var head = c.buckets[headPtr]
	if head.Ptr == 0 {
		head.Ptr = headPtr
		head.TailPtr = headPtr
		head.Data = v
	} else {
		var cursor = c.NextID()
		var ele = c.buckets[cursor]
		var tail = c.buckets[head.TailPtr]
		head.TailPtr = cursor
		tail.NextPtr = cursor
		ele.PrevPtr = tail.Ptr
		ele.Ptr = cursor
		ele.Data = v
	}
}

// DeleteOne @return true: target exist and deleted
func (c *mqueue[T]) DeleteOne(headPtr Pointer, key T) bool {
	target, exist := c.Find(headPtr, key)
	if !exist {
		return false
	}
	c.DeleteTarget(headPtr, target)
	return true
}

func (c *mqueue[T]) DeleteAll(headPtr Pointer, key T) (affected int) {
	var eles = make([]*element[T], 0)
	for_each[element[T]](c, headPtr, func(ele *element[T]) (next bool) {
		if ele.Data == key {
			eles = append(eles, ele)
		}
		return true
	})
	for _, item := range eles {
		c.DeleteTarget(headPtr, item)
	}
	return len(eles)
}

func (c *mqueue[T]) doDeleteTarget(headPtr Pointer, cur *element[T], key T) (ele *element[T]) {
	if cur == nil || cur.Ptr == 0 || cur.Data != key {
		return nil
	}

	if cur.PrevPtr != 0 && cur.NextPtr != 0 {
		c.buckets[cur.PrevPtr].NextPtr = cur.NextPtr
		ele = cur
	} else if cur.PrevPtr != 0 && cur.NextPtr == 0 {
		c.buckets[cur.PrevPtr].NextPtr = 0
		c.buckets[headPtr].TailPtr = cur.PrevPtr
		ele = cur
	} else if cur.PrevPtr == 0 && cur.NextPtr != 0 {
		var tailPtr = cur.TailPtr
		var next = c.buckets[cur.NextPtr]
		*cur = *next
		cur.TailPtr = tailPtr
		ele = next
	} else {
		ele = cur
	}
	return ele
}

// DeleteTarget delete the element
func (c *mqueue[T]) DeleteTarget(headPtr Pointer, target *element[T]) {
	if dst := c.doDeleteTarget(headPtr, target, target.Data); dst != nil {
		c.unused.Push(dst.Ptr) // recycle the pointer
		dst.Ptr = 0
		dst.NextPtr = 0
		dst.PrevPtr = 0
		c.length--
	}
}

func (c *mqueue[T]) ForEach(headPtr Pointer, fn func(ele *element [T]) (next bool)) {
	for_each[element[T]](c, headPtr, fn)
}

func (c *mqueue[T]) Find(headPtr Pointer, key T) (result *element[T], exist bool) {
	for_each[element[T]](c, headPtr, func(ele *element[T]) (next bool) {
		if key == ele.Data {
			result = ele
			exist = true
			return false
		}
		return true
	})
	return
}
