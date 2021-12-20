package linkedlist

type Iterator[T any] struct {
	next *Iterator[T]
	Data T
}

func NewQueue[T any]() *Queue[T] {
	return new(Queue[T])
}

type Queue[T any] struct {
	length int
	head   *Iterator[T]
	tail   *Iterator[T]
}

func (c *Queue[T]) Clear() {
	c.head = nil
	c.tail = nil
	c.length = 0
}

func (c *Queue[T]) Begin() *Iterator[T] {
	return c.head
}

func (c *Queue[T]) Next(iter *Iterator[T]) *Iterator[T] {
	return iter.next
}

func (c *Queue[T]) End(iter *Iterator[T]) bool {
	return iter == nil
}

func (c *Queue[T]) Len() int {
	return c.length
}

func (c *Queue[T]) Push(values ...T) {
	for _, v := range values {
		var ele = &Iterator[T]{Data: v}
		if c.length > 0 {
			c.tail.next = ele
			c.tail = ele
		} else {
			c.head = ele
			c.tail = ele
		}
		c.length++
	}
}

func (c *Queue[T]) Front() *Iterator[T] {
	return c.head
}

func (c *Queue[T]) Pop() *Iterator[T] {
	if c.length == 0 {
		return nil
	}
	var result = c.head
	c.head = c.head.next
	c.length--
	if c.length == 0 {
		c.tail = nil
	}
	return result
}
