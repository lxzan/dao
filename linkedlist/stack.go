package linkedlist

type stack[T any] struct {
	length int
	head   *element[T]
}

type element[T any] struct {
	next  *element[T]
	Data T
}

func (c *stack[T]) Len() int {
	return c.length
}

func (c *stack[T]) Push(v T) {
	var ele = &element[T]{Data: v}
	if c.length > 0 {
		ele.next = c.head
		c.head = ele
	} else {
		c.head = ele
	}
	c.length++
}

func (c *stack[T]) Front() *element[T] {
	return c.head
}

func (c *stack[T]) Pop() *element[T] {
	switch c.length {
	case 0:
		return nil
	case 1:
		var result = c.head
		c.head = nil
		c.length = 0
		return result
	default:
		var result = c.head
		c.head = c.head.next
		c.length--
		return result
	}
}
