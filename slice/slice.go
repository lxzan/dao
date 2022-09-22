package slice

import (
	"github.com/lxzan/dao"
	"github.com/lxzan/dao/algorithm"
)

type (
	Slice[T any] []T

	Iterator[T any] struct {
		i, n  int
		Value T
	}
)

// New param1: size, param2: cap
func New[T any](sizes ...int) Slice[T] {
	var n = len(sizes)
	var size = 0
	var capicity = 3
	if n >= 1 {
		size = sizes[0]
		capicity += size
	}
	if n >= 2 {
		capicity = sizes[1]
	}
	return make([]T, size, capicity)
}

func NewFrom[T any](arr []T) Slice[T] {
	return arr
}

func (c *Slice[T]) Begin() *Iterator[T] {
	var iter = &Iterator[T]{
		i: 0,
		n: len(*c),
	}
	if iter.n > 0 {
		iter.Value = (*c)[0]
	}
	return iter
}

func (c *Slice[T]) Next(iter *Iterator[T]) *Iterator[T] {
	iter.i++
	if iter.i < iter.n {
		iter.Value = (*c)[iter.i]
	}
	return iter
}

func (c *Slice[T]) End(iter *Iterator[T]) bool {
	return iter.i >= iter.n
}

func (c *Slice[T]) Len() int {
	return len(*c)
}

func (c *Slice[T]) Push(eles ...T) {
	*c = append(*c, eles...)
}

func (c *Slice[T]) RPop() T {
	var n = c.Len()
	var result = (*c)[n-1]
	*c = (*c)[:n-1]
	return result
}

func (c *Slice[T]) LPop() T {
	var result = (*c)[0]
	*c = (*c)[1:]
	return result
}

// Head return first element
func (c Slice[T]) Front() T {
	return c[0]
}

// Tail return first element
func (c Slice[T]) Back() T {
	return c[c.Len()-1]
}

func (c *Slice[T]) Swap(i, j int) {
	(*c)[i], (*c)[j] = (*c)[j], (*c)[i]
}

func (c *Slice[T]) Reverse() *Slice[T] {
	var n = c.Len()
	for i := 0; i < n/2; i++ {
		c.Swap(i, n-i-1)
	}
	return c
}

func (c *Slice[T]) Delete(i int) {
	var n = c.Len()
	for j := i; j < n-1; j++ {
		(*c)[j] = (*c)[j+1]
	}
	if i >= 0 && i < n {
		*c = (*c)[:n-1]
	}
}

func (c *Slice[T]) Clear() {
	*c = (*c)[:0]
}

func (c *Slice[T]) Sort(cmp func(a, b T) dao.Ordering) *Slice[T] {
	algorithm.Sort(*c, cmp)
	return c
}

func (c *Slice[T]) ForEach(fn func(iter *Iterator[T])) {
	for i := c.Begin(); !c.End(i); i = c.Next(i) {
		fn(i)
	}
}

func (c *Slice[T]) Range(i, j int) Slice[T] {
	return (*c)[i:j]
}

func (c *Slice[T]) Unique(cmp func(a, b T) dao.Ordering) *Slice[T] {
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
	*c = results
	return c
}

func (c *Slice[T]) Filter(fn func(ele T) bool) *Slice[T] {
	var results = New[T](0, c.Len())
	for i := range *c {
		if fn((*c)[i]) {
			results.Push((*c)[i])
		}
	}
	return &results
}

func (c *Slice[T]) Map(fn func(ele T) T) *Slice[T] {
	for i := range *c {
		(*c)[i] = fn((*c)[i])
	}
	return c
}
