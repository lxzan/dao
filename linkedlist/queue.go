package linkedlist

type queue[T any] struct {
	length int
	head   *element[T]
	tail   *element[T]
}

type element[T any] struct {
	next  *element[T]
	Value T
}

func (c *queue[T]) Len() int {
	return c.length
}

func (c *queue[T]) Push(v T) {
	var ele = &element[T]{Value: v}
	if c.length > 0 {
		c.tail.next = ele
		c.tail = ele
		c.length++
	} else {
		c.head = ele
		c.tail = ele
		c.length++
	}
}

func (c *queue[T]) Front() *element[T] {
	return c.head
}

func (c *queue[T]) Pop() *element[T] {
	switch c.length {
	case 0:
		return nil
	case 1:
		var result = c.head
		c.head = nil
		c.tail = nil
		c.length = 0
		return result
	default:
		var result = c.head
		c.head = c.head.next
		c.length--
		return result
	}
}
