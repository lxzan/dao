package heap

import "github.com/lxzan/dao"

func MinHeap[T dao.Comparable](a, b T) dao.Ordering {
	if a > b {
		return dao.Greater
	} else if a < b {
		return dao.Less
	} else {
		return dao.Equal
	}
}

func MaxHeap[T dao.Comparable](a, b T) dao.Ordering {
	return -1 * MinHeap(a, b)
}

func New[T any](cap int, cmp func(a, b T) dao.Ordering) *Heap[T] {
	return &Heap[T]{
		Data: make([]T, 0, cap),
		Cmp:  cmp,
	}
}

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
	if j < n && c.Cmp(c.Data[j], c.Data[i]) == dao.Less {
		c.Swap(i, j)
		c.Down(j, n)
	}
	var k = 2*i + 2
	if k < n && c.Cmp(c.Data[k], c.Data[i]) == dao.Less {
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

func (c *Heap[T]) Find(target T) (result T, exist bool) {
	var q = find_param[T]{
		Length: c.Len(),
		Target: target,
		Result: result,
		Exist:  false,
	}
	c.do_find(0, &q)
	return q.Result, q.Exist
}

type find_param[T any] struct {
	Length int
	Target T
	Result T
	Exist  bool
}

func (c *Heap[T]) do_find(i int, q *find_param[T]) {
	if q.Exist {
		return
	}

	if c.Cmp(c.Data[i], q.Target) == dao.Equal {
		q.Result = c.Data[i]
		q.Exist = true
		return
	}

	var j = 2*i + 1
	if j < q.Length && c.Cmp(c.Data[j], q.Target) != dao.Greater {
		c.do_find(j, q)
	}

	var k = 2*i + 2
	if k < q.Length && c.Cmp(c.Data[k], q.Target) != dao.Greater {
		c.do_find(k, q)
	}
}
