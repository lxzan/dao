package vector

import (
	"github.com/lxzan/dao"
	"github.com/lxzan/dao/algorithm"
)

type (
	Vector[T any] []T

	Iterator[T any] struct {
		i, n  int
		Value T
	}
)

// New param1: size, param2: cap
func New[T any](size ...int) *Vector[T] {
	var length, capacity int
	switch len(size) {
	case 0, 1:
		capacity = 3
	default:
		length, capacity = size[0], size[1]
	}

	v := Vector[T](make([]T, length, capacity))
	return &v
}

func NewFromSlice[T any](s []T) *Vector[T] {
	v := Vector[T](s)
	return &v
}

func (c *Vector[T]) Elem() Vector[T] {
	return *c
}

func (c *Vector[T]) Begin() *Iterator[T] {
	var iter = &Iterator[T]{
		i: 0,
		n: len(*c),
	}
	if iter.n > 0 {
		iter.Value = (*c)[0]
	}
	return iter
}

func (c *Vector[T]) Next(iter *Iterator[T]) *Iterator[T] {
	iter.i++
	if iter.i < iter.n {
		iter.Value = (*c)[iter.i]
	}
	return iter
}

func (c *Vector[T]) End(iter *Iterator[T]) bool {
	return iter.i >= iter.n
}

func (c *Vector[T]) Len() int {
	return len(*c)
}

func (c *Vector[T]) Get(i int) T {
	return (*c)[i]
}

func (c *Vector[T]) Push(eles ...T) {
	*c = append(*c, eles...)
}

func (c *Vector[T]) RPop() T {
	var n = c.Len()
	var result = (*c)[n-1]
	*c = (*c)[:n-1]
	return result
}

func (c *Vector[T]) LPop() T {
	var result = (*c)[0]
	*c = (*c)[1:]
	return result
}

func (c *Vector[T]) Swap(i, j int) {
	(*c)[i], (*c)[j] = (*c)[j], (*c)[i]
}

func (c *Vector[T]) Reverse() *Vector[T] {
	var n = c.Len()
	for i := 0; i < n/2; i++ {
		c.Swap(i, n-i-1)
	}
	return c
}

func (c *Vector[T]) Delete(i int) {
	var n = c.Len()
	for j := i; j < n-1; j++ {
		(*c)[j] = (*c)[j+1]
	}
	if i >= 0 && i < n {
		*c = (*c)[:n-1]
	}
}

func (c *Vector[T]) Reset() {
	*c = (*c)[:0]
}

func (c *Vector[T]) Sort(cmp func(a, b T) dao.Ordering) *Vector[T] {
	algorithm.Sort(*c, cmp)
	return c
}

func (c *Vector[T]) ForEach(fn func(index int, value T)) {
	for i, v := range *c {
		fn(i, v)
	}
}

func (c *Vector[T]) Range(i, j int) *Vector[T] {
	v := (*c)[i:j]
	return &v
}

func (c *Vector[T]) Unique(cmp func(a, b T) dao.Ordering) *Vector[T] {
	var n = c.Len()
	if n == 0 {
		return c
	}

	c.Sort(cmp)
	var results = New[T](0, n)
	results.Push((*c)[0])
	for i := 1; i < n; i++ {
		if cmp((*c)[i], (*c)[i-1]) != dao.Equal {
			results.Push((*c)[i])
		}
	}
	return results
}

func (c *Vector[T]) Filter(fn func(ele T) bool) *Vector[T] {
	var results = New[T](0, c.Len())
	for i := range *c {
		if fn((*c)[i]) {
			results.Push((*c)[i])
		}
	}
	return results
}

func (c *Vector[T]) Map(fn func(ele T) T) *Vector[T] {
	for i := range *c {
		(*c)[i] = fn((*c)[i])
	}
	return c
}
