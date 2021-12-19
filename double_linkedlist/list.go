package double_linkedlist

func New[T any]() *List[T] {
	return new(List[T])
}

// safe delete in loop
type Iterator[T any] struct {
	prev, next *Iterator[T]
	Data       T
}

type List[T any] struct {
	length int
	head   *Iterator[T]
	tail   *Iterator[T]
}

func (c List[T]) Begin() *Iterator[T] {
	return c.head
}

func (c *List[T]) Next(iter *Iterator[T]) *Iterator[T] {
	return iter.next
}

func (c List[T]) End(iter *Iterator[T]) bool {
	return iter == nil
}

func (c List[T]) Len() int {
	return c.length
}

func (c List[T]) Clear() {
	c.head = nil
	c.tail = nil
	c.length = 0
}

func (c *List[T]) Front() *Iterator[T] {
	return c.head
}

func (c *List[T]) Back() *Iterator[T] {
	return c.tail
}

func (c *List[T]) RPush(values ...T) {
	for _, v := range values {
		var ele = new(Iterator[T])
		ele.Data = v
		if c.length > 0 {
			c.tail.next = ele
			ele.prev = c.tail
			c.tail = ele
		} else {
			c.head = ele
			c.tail = ele
		}
		c.length++
	}
}

func (c *List[T]) LPush(values ...T) {
	for _, v := range values {
		var ele = new(Iterator[T])
		ele.Data = v
		if c.length > 0 {
			ele.next = c.head
			c.head.prev = ele
			c.head = ele
		} else {
			c.head = ele
			c.tail = ele
		}
		c.length++
	}
}

func (c *List[T]) LPop() *Iterator[T] {
	if c.length == 0 {
		return nil
	}
	var result = c.head
	c.head = c.head.next
	result.next = nil
	if c.head != nil {
		c.head.prev = nil
	}
	if c.length == 1 {
		c.tail = nil
	}
	c.length--
	return result
}

func (c *List[T]) RPop() *Iterator[T] {
	if c.length == 0 {
		return nil
	}

	var result = c.tail
	c.tail = c.tail.prev
	result.prev = nil
	if c.tail != nil {
		c.tail.next = nil
	}
	if c.length == 1 {
		c.head = nil
	}
	c.length--
	return result
}

// Delete it's safe delete in loop
func (c *List[T]) Delete(iter *Iterator[T]) {
	var prev = iter.prev
	var next = iter.next
	if prev != nil && next != nil {
		prev.next = next
		next.prev = prev
	} else if prev != nil && next == nil {
		prev.next = nil
		c.tail = prev
	} else if prev == nil && next != nil {
		next.prev = nil
		c.head = next
	} else {
		c.head = nil
		c.tail = nil
	}
	c.length--
}

func (c *List[T]) InsertAfter(iter *Iterator[T], v T) {
	var next = iter.next
	var cur = new(Iterator[T])
	cur.prev = iter
	cur.next = next
	cur.Data = v
	iter.next = cur
	if next != nil {
		next.prev = cur
	}
	c.length++
}

func (c *List[T]) InsertBefore(iter *Iterator[T], v T) {
	var prev = iter.prev
	var cur = new(Iterator[T])
	cur.prev = prev
	cur.next = iter
	cur.Data = v
	iter.prev = cur
	if prev != nil {
		prev.next = cur
	}
	c.length++
}
