package slice

import (
	"github.com/lxzan/dao"
	"github.com/lxzan/dao/algorithm"
)

type Iterator[T any] struct {
	length int
	Index  int
	Value  T
}

type Slice[T any] []T

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

func (c Slice[T]) Begin() *Iterator[T] {
	return &Iterator[T]{
		length: len(c),
		Index:  0,
		Value:  c[0],
	}
}

func (c Slice[T]) Next(iter *Iterator[T]) *Iterator[T] {
	iter.Index++
	if iter.Index < iter.length {
		iter.Value = c[iter.Index]
	}
	return iter
}

func (c Slice[T]) End(iter *Iterator[T]) bool {
	return iter.Index >= iter.length
}

func (c Slice[T]) Len() int {
	return len(c)
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

func (c *Slice[T]) Reverse() {
	var n = c.Len()
	for i := 0; i < n/2; i++ {
		c.Swap(i, n-i-1)
	}
}

func (c *Slice[T]) Delete(i int) {
	var n = c.Len()
	for j := i; j < n-1; j++ {
		(*c)[j] = (*c)[j+1]
	}
	*c = (*c)[:n-1]
}

func (c *Slice[T]) Clear() {
	*c = (*c)[:0]
}

func (c *Slice[T]) Sort(cmp func(a, b T) dao.Ordering) *Slice[T] {
	algorithm.Sort(*c, cmp)
	return c
}

func (c Slice[T]) ForEach(fn func(index int, value T)) {
	var n = c.Len()
	for i := 0; i < n; i++ {
		fn(i, c[i])
	}
}

func (c *Slice[T]) Range(i, j int) Slice[T] {
	return (*c)[i:j]
}

func (c *Slice[T]) Filter(fn func(ele T) bool) *Slice[T] {
	var results = New[T](0, c.Len())
	for i, _ := range *c {
		if fn((*c)[i]) {
			results.Push((*c)[i])
		}
	}
	return &results
}

func (c *Slice[T]) Map(fn func(ele T) T) *Slice[T] {
	for i, _ := range *c {
		(*c)[i] = fn((*c)[i])
	}
	return c
}
