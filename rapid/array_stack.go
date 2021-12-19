package rapid

type array_stack []Pointer

func (c array_stack) Len() int {
	return len(c)
}

func (c *array_stack) Push(v Pointer) {
	*c = append(*c, v)
}

func (c *array_stack) Pop() Pointer {
	var n = c.Len()
	if n >= 1 {
		var result = (*c)[n-1]
		*c = (*c)[:n-1]
		return result
	}
	return 0
}
