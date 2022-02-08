package heap

import "github.com/lxzan/dao"

func MinHeap[T dao.Comparable[T]](a, b T) dao.Ordering {
	if a > b {
		return dao.Greater
	} else if a < b {
		return dao.Less
	} else {
		return dao.Equal
	}
}

func MaxHeap[T dao.Comparable[T]](a, b T) dao.Ordering {
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

func (this Heap[T]) Len() int {
	return len(this.Data)
}

func (this *Heap[T]) Swap(i, j int) {
	this.Data[i], this.Data[j] = this.Data[j], this.Data[i]
}

func (this *Heap[T]) Push(eles ...T) {
	for _, item := range eles {
		this.Data = append(this.Data, item)
		this.Up(this.Len() - 1)
	}
}

func (this *Heap[T]) Up(i int) {
	var j = (i - 1) / 2
	if j >= 0 && this.Cmp(this.Data[i], this.Data[j]) == dao.Less {
		this.Swap(i, j)
		this.Up(j)
	}
}

func (this *Heap[T]) Pop() T {
	var n = this.Len()
	var result = this.Data[0]
	this.Data[0] = this.Data[n-1]
	this.Data = this.Data[:n-1]
	this.Down(0, n-1)
	return result
}

func (this *Heap[T]) Down(i, n int) {
	var j = 2*i + 1
	if j < n && this.Cmp(this.Data[j], this.Data[i]) == dao.Less {
		this.Swap(i, j)
		this.Down(j, n)
	}
	var k = 2*i + 2
	if k < n && this.Cmp(this.Data[k], this.Data[i]) == dao.Less {
		this.Swap(i, k)
		this.Down(k, n)
	}
}

func (this *Heap[T]) Sort() []T {
	var n = this.Len()
	if n >= 2 {
		for i := n - 1; i >= 2; i-- {
			this.Swap(0, i)
			this.Down(0, i)
		}
		this.Swap(0, 1)
	}
	return this.Data
}

func (this *Heap[T]) Find(target T) (result T, exist bool) {
	var q = find_param[T]{
		Length: this.Len(),
		Target: target,
		Result: result,
		Exist:  false,
	}
	this.do_find(0, &q)
	return q.Result, q.Exist
}

type find_param[T any] struct {
	Length int
	Target T
	Result T
	Exist  bool
}

func (this Heap[T]) do_find(i int, q *find_param[T]) {
	if q.Exist {
		return
	}

	if this.Cmp(this.Data[i], q.Target) == dao.Equal {
		q.Result = this.Data[i]
		q.Exist = true
		return
	}

	var j = 2*i + 1
	if j < q.Length && this.Cmp(this.Data[j], q.Target) != dao.Greater {
		this.do_find(j, q)
	}

	var k = 2*i + 2
	if k < q.Length && this.Cmp(this.Data[k], q.Target) != dao.Greater {
		this.do_find(k, q)
	}
}
