package multi_queue

type queueElement struct {
	Next *queueElement
	Data Pointer
}

type queue struct {
	length int
	head   *queueElement
	tail   *queueElement
}

func (c *queue) Len() int {
	return c.length
}

func (c *queue) Push(v Pointer) {
	var ele = &queueElement{
		Data: v,
	}
	if c.length == 0 {
		c.head = ele
		c.tail = ele
	} else {
		c.tail.Next = ele
		c.tail = ele
	}
	c.length++
}

func (c *queue) Pop() Pointer {
	switch c.length {
	case 0:
		return 0
	case 1:
		var result = c.head.Data
		c.length = 0
		c.head = nil
		c.tail = nil
		return result
	default:
		var result = c.head.Data
		c.length--
		c.head = c.head.Next
		return result
	}
}
