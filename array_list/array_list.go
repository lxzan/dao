package array_list

import (
	"github.com/lxzan/dao"
	"github.com/lxzan/dao/algorithm"
)

type ArrayList[T any] []T

// New param1: size, param2: cap
func New[T any](size ...int) *ArrayList[T] {
	var length, capacity int
	switch len(size) {
	case 0, 1:
		capacity = 3
	default:
		length, capacity = size[0], size[1]
	}

	v := ArrayList[T](make([]T, length, capacity))
	return &v
}

func NewFromSlice[T any](s []T) *ArrayList[T] {
	v := ArrayList[T](s)
	return &v
}

func (c *ArrayList[T]) Elem() ArrayList[T] {
	return *c
}

func (c *ArrayList[T]) Len() int {
	return len(*c)
}

func (c *ArrayList[T]) Get(i int) T {
	return (*c)[i]
}

func (c *ArrayList[T]) Push(eles ...T) {
	*c = append(*c, eles...)
}

func (c *ArrayList[T]) RPop() T {
	var n = c.Len()
	var result = (*c)[n-1]
	*c = (*c)[:n-1]
	return result
}

func (c *ArrayList[T]) LPop() T {
	var result = (*c)[0]
	*c = (*c)[1:]
	return result
}

func (c *ArrayList[T]) Swap(i, j int) {
	(*c)[i], (*c)[j] = (*c)[j], (*c)[i]
}

func (c *ArrayList[T]) Reverse() *ArrayList[T] {
	var n = c.Len()
	for i := 0; i < n/2; i++ {
		c.Swap(i, n-i-1)
	}
	return c
}

func (c *ArrayList[T]) Delete(i int) {
	var n = c.Len()
	for j := i; j < n-1; j++ {
		(*c)[j] = (*c)[j+1]
	}
	if i >= 0 && i < n {
		*c = (*c)[:n-1]
	}
}

func (c *ArrayList[T]) Reset() {
	*c = (*c)[:0]
}

func (c *ArrayList[T]) Sort(cmp func(a, b T) dao.Ordering) *ArrayList[T] {
	algorithm.Sort(*c, cmp)
	return c
}

func (c *ArrayList[T]) ForEach(fn func(i int, v T) bool) {
	for index, value := range *c {
		if !fn(index, value) {
			return
		}
	}
}

func (c *ArrayList[T]) Range(i, j int) *ArrayList[T] {
	v := (*c)[i:j]
	return &v
}
