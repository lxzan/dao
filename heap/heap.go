package heap

import (
	"github.com/lxzan/dao"
)

// 最小堆
func MinHeap[T dao.Comparable](a, b T) dao.Ordering {
	if a > b {
		return dao.Greater
	} else if a < b {
		return dao.Less
	} else {
		return dao.Equal
	}
}

// 最大堆
func MaxHeap[T dao.Comparable](a, b T) dao.Ordering {
	return -1 * MinHeap(a, b)
}

// New 新建一个堆
// Create a new heap
func New[T any](cap int, cmp func(a, b T) dao.Ordering) *Heap[T] {
	return &Heap[T]{
		Data: make([]T, 0, cap),
		Cmp:  cmp,
	}
}

// Init 将切片初始化为一个堆
// Initialize the slice to a heap
func Init[T any](arr []T, cmp func(a, b T) dao.Ordering) *Heap[T] {
	var h = &Heap[T]{
		Data: arr,
		Cmp:  cmp,
	}
	var n = len(arr)
	for i := 1; i < n; i++ {
		h.Up(i)
	}
	return h
}

type Heap[T any] struct {
	Data []T
	Cmp  func(a, b T) dao.Ordering
}

func (c *Heap[T]) Len() int {
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
	if j >= 0 && c.Cmp(c.Data[i], c.Data[j]) == dao.Less {
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
	var k = 2*i + 2
	var x = -1
	if j < n {
		x = j
	}
	if k < n && c.Cmp(c.Data[k], c.Data[j]) == dao.Less {
		x = k
	}
	if x != -1 && c.Cmp(c.Data[x], c.Data[i]) == dao.Less {
		c.Swap(i, x)
		c.Down(x, n)
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

// Front 访问堆顶元素
// Accessing the top element of the heap
func (c *Heap[T]) Front() T {
	return c.Data[0]
}
