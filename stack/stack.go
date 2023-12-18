package stack

type Stack[T any] []T

func New[T any](capacity uint32) *Stack[T] {
	s := Stack[T](make([]T, 0, capacity))
	return &s
}

func (c *Stack[T]) Reset() {
	*c = (*c)[:0]
}

func (c *Stack[T]) Len() int {
	return len(*c)
}

func (c *Stack[T]) Push(v T) {
	*c = append(*c, v)
}

func (c *Stack[T]) Pop() (value T) {
	n := c.Len()
	switch n {
	case 0:
		return
	default:
		value = (*c)[n-1]
		*c = (*c)[:n-1]
		return
	}
}

func (c *Stack[T]) Range(f func(value T) bool) {
	for _, item := range *c {
		if !f(item) {
			return
		}
	}
}
