package mqueue

type array_stack []uint32

func (c array_stack) Len() int {
	return len(c)
}

func (c *array_stack) Push(v uint32) {
	*c = append(*c, v)
}

func (c *array_stack) Pop() uint32 {
	var n = c.Len()
	if n >= 1 {
		var result = (*c)[n-1]
		*c = (*c)[:n-1]
		return result
	}
	return 0
}
