package slice

import (
	"github.com/lxzan/dao"
	"github.com/lxzan/dao/algorithm"
)

type Iterator[T any] struct {
	next  bool
	Index int
	Value T
}

func (this *Iterator[T]) Break() {
	this.next = false
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

func (this Slice[T]) Len() int {
	return len(this)
}

func (this *Slice[T]) Push(eles ...T) {
	*this = append(*this, eles...)
}

func (this *Slice[T]) RPop() T {
	var n = this.Len()
	var result = (*this)[n-1]
	*this = (*this)[:n-1]
	return result
}

func (this *Slice[T]) LPop() T {
	var result = (*this)[0]
	*this = (*this)[1:]
	return result
}

// Head return first element
func (this Slice[T]) Front() T {
	return this[0]
}

// Tail return first element
func (this Slice[T]) Back() T {
	return this[this.Len()-1]
}

func (this *Slice[T]) Swap(i, j int) {
	(*this)[i], (*this)[j] = (*this)[j], (*this)[i]
}

func (this *Slice[T]) Reverse() *Slice[T] {
	var n = this.Len()
	for i := 0; i < n/2; i++ {
		this.Swap(i, n-i-1)
	}
	return this
}

func (this *Slice[T]) Delete(i int) {
	var n = this.Len()
	for j := i; j < n-1; j++ {
		(*this)[j] = (*this)[j+1]
	}
	if i >= 0 && i < n {
		*this = (*this)[:n-1]
	}
}

func (this *Slice[T]) Clear() {
	*this = (*this)[:0]
}

func (this *Slice[T]) Sort(cmp func(a, b T) dao.Ordering) *Slice[T] {
	algorithm.Sort(*this, cmp)
	return this
}

func (this Slice[T]) ForEach(fn func(iter *Iterator[T])) {
	var n = this.Len()
	var iter = &Iterator[T]{next: true}
	for i := 0; i < n && iter.next; i++ {
		iter.Index = i
		iter.Value = this[i]
		fn(iter)
	}
}

func (this *Slice[T]) Range(i, j int) Slice[T] {
	return (*this)[i:j]
}

func (this *Slice[T]) Unique(cmp func(a, b T) dao.Ordering) *Slice[T] {
	var n = this.Len()
	if n == 0 {
		return this
	}

	this.Sort(cmp)
	var results = New[T](0, n)
	results.Push((*this)[0])
	for i := 1; i < n; i++ {
		if cmp((*this)[i], (*this)[i-1]) != dao.Equal {
			results.Push((*this)[i])
		}
	}
	*this = results
	return this
}

func (this *Slice[T]) Filter(fn func(ele T) bool) *Slice[T] {
	var results = New[T](0, this.Len())
	for i := range *this {
		if fn((*this)[i]) {
			results.Push((*this)[i])
		}
	}
	return &results
}

func (this *Slice[T]) Map(fn func(ele T) T) *Slice[T] {
	for i := range *this {
		(*this)[i] = fn((*this)[i])
	}
	return this
}
