package heap

import "github.com/lxzan/dao"

func MinHeap[T dao.Comparable[T]](a, b T) bool {
	return a < b
}

func MaxHeap[T dao.Comparable[T]](a, b T) bool {
	return a > b
}

func New[T any](cap int, less func(a, b T) bool) *Heap[T] {
	return &Heap[T]{
		Data: make([]T, 0, cap),
		Less: less,
	}
}

func Init[T any](arr []T, less func(a, b T) bool) *Heap[T] {
	var h = &Heap[T]{
		Data: arr,
		Less: less,
	}
	var n = len(arr)
	for i := 1; i < n; i++ {
		h.Up(i)
	}
	return h
}

type Heap[T any] struct {
	Data []T
	Less func(a, b T) bool
}

func (c Heap[T]) Len() int {
	return len(c.Data)
}

func (c *Heap[T]) Swap(i, j int) {
	c.Data[i], c.Data[j] = c.Data[j], c.Data[i]
}

func (c *Heap[T]) Push(eles ...T) {
	for _, item := range eles {
		c.Data = append(c.Data, item)
		c.Up(c.Len() - 1)
	}
}

func (c *Heap[T]) Up(i int) {
	var j = (i - 1) / 2
	if j >= 0 && c.Less(c.Data[i], c.Data[j]) {
		c.Swap(i, j)
		c.Up(j)
	}
}

func (c *Heap[T]) Pop() T {
	var n = c.Len()
	var result = c.Data[0]
	c.Data[0] = c.Data[n-1]
	c.Data = c.Data[:n-1]
	c.Down(0, n-1)
	return result
}

func (c *Heap[T]) Down(i, n int) {
	var j = 2*i + 1
	if j < n && c.Less(c.Data[j], c.Data[i]) {
		c.Swap(i, j)
		c.Down(j, n)
	}
	var k = 2*i + 2
	if k < n && c.Less(c.Data[k], c.Data[i]) {
		c.Swap(i, k)
		c.Down(k, n)
	}
}

func (c *Heap[T]) Sort() []T {
	var n = c.Len()
	if n >= 2 {
		for i := n - 1; i >= 2; i-- {
			c.Swap(0, i)
			c.Down(0, i)
		}
		c.Swap(0, 1)
	}
	return c.Data
}
