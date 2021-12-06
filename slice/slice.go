package slice

import "github.com/lxzan/dao/types"

type slice[T any] []T

// New param1: size, param2: cap
func New[T any](sizes ...int) slice[T] {
	var n = len(sizes)
	var size = 0
	var capicity = 3
	if n >= 1 {
		size = sizes[0]
	}
	if n >= 2 {
		capicity = sizes[1]
	}
	var v = make([]T, size, capicity)
	return slice[T](v)
}

func NewFrom[T any](arr []T) slice[T] {
	return slice[T](arr)
}

func (c slice[T]) Len() int {
	return len(c)
}

func (c *slice[T]) Push(ele T) {
	*c = append(*c, ele)
}

func (c *slice[T]) RPop() T {
	var n = c.Len()
	var result = (*c)[n-1]
	*c = (*c)[:n-1]
	return result
}

func (c *slice[T]) LPop() T {
	var result = (*c)[0]
	*c = (*c)[1:]
	return result
}

// Head return first element
func (c slice[T]) Head() T {
	return c[0]
}

// Tail return first element
func (c slice[T]) Tail() T {
	return c[c.Len()-1]
}

func (c *slice[T]) Swap(i, j int) {
	(*c)[i], (*c)[j] = (*c)[j], (*c)[i]
}

func (c *slice[T]) Reverse() {
	var n = c.Len()
	for i := 0; i < n/2; i++ {
		c.Swap(i, n-i-1)
	}
}

func (c *slice[T]) Delete(i int) {
	var n = c.Len()
	for j := i; j < n-1; j++ {
		(*c)[j] = (*c)[j+1]
	}
	*c = (*c)[:n-1]
}

func (c slice[T]) Get(i int) T {
	return c[i]
}

func (c *slice[T]) Clear() {
	*c = (*c)[:0]
}

func (c *slice[T]) Sort(cmp func(a, b T) types.Ordering) *slice[T] {
	c.quickSort(0, c.Len()-1, cmp)
	return c
}

func (c *slice[T]) quickSort(left int, right int, cmp func(a, b T) types.Ordering) {
	if left >= right {
		return
	}
	var index = left
	var mid = c.getMedium(left, right, cmp)
	c.Swap(mid, left)
	for i := left + 1; i <= right; i++ {
		var order = cmp(c.Get(i), c.Get(left))
		if order == types.Less || (order == types.Equal && i%2 == 0) {
			index++
			if i != index {
				c.Swap(index, i)
			}
		}
	}
	c.Swap(left, index)
	c.quickSort(left, index-1, cmp)
	c.quickSort(index+1, right, cmp)
}

func (c *slice[T]) getMedium(left int, right int, cmp func(a, b T) types.Ordering) int {
	var mid = (left + right) / 2
	if cmp(c.Get(right), c.Get(mid))+cmp(c.Get(mid), c.Get(left)) != 0 {
		return mid
	}
	if cmp(c.Get(right), c.Get(left))+cmp(c.Get(left), c.Get(mid)) != 0 {
		return left
	}
	return right
}

func (c slice[T]) ForEach(fn func(index int, value T)) {
	var n = c.Len()
	for i := 0; i < n; i++ {
		fn(i, c[i])
	}
}

func (c *slice[T]) Range(i, j int) slice[T] {
	return (*c)[i:j]
}

func (c *slice[T]) Filter(fn func(ele T) bool) *slice[T] {
	var collection = New()
	for i, _ := range *c {
		if fn((*c)[i]) {
			collection.Push((*c)[i])
		}
	}
	return &collection
}

func (c *slice[T]) Map(fn func(ele T) T) *slice[T] {
	for i, _ := range *c {
		(*c)[i] = fn((*c)[i])
	}
	return c
}
