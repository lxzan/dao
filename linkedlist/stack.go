package linkedlist

func NewStack[T any]() *Stack[T] {
	return new(Stack[T])
}

type Stack[T any] struct {
	length int
	head   *Iterator[T]
}

func (c *Stack[T]) Next(iter *Iterator[T]) *Iterator[T] {
	return iter.next
}

func (c *Stack[T]) Begin() *Iterator[T] {
	return c.head
}

func (c *Stack[T]) End(iter *Iterator[T]) bool {
	return iter == nil
}

func (c *Stack[T]) Len() int {
	return c.length
}

func (c *Stack[T]) Clear() {
	c.head = nil
	c.length = 0
}

func (c *Stack[T]) Push(values ...T) {
	for _, v := range values {
		var ele = new(Iterator[T])
		ele.Data = v
		if c.length > 0 {
			ele.next = c.head
			c.head = ele
		} else {
			c.head = ele
		}
		c.length++
	}
}

func (c *Stack[T]) Front() *Iterator[T] {
	return c.head
}

func (c *Stack[T]) Pop() *Iterator[T] {
	if c.length == 0 {
		return nil
	}
	var result = c.head
	c.head = c.head.next
	c.length--
	return result
}
