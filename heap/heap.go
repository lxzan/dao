package heap

import "github.com/lxzan/dao"

const (
	Binary = 2

	Quadratic = 4

	Octal = 8
)

// New 新建一个堆
// Create a new heap
func New[T any](less dao.LessFunc[T]) *Heap[T] {
	h := &Heap[T]{cmp: less}
	h.SetForkNumber(Quadratic)
	return h
}

type Heap[T any] struct {
	bits  uint32
	forks int
	data  []T
	cmp   func(a, b T) bool
}

// SetCap 设置预分配容量
func (c *Heap[T]) SetCap(n uint32) *Heap[T] {
	c.data = make([]T, 0, n)
	return c
}

// SetForkNumber 设置分叉数
func (c *Heap[T]) SetForkNumber(n uint32) *Heap[T] {
	c.forks = int(n)
	switch n {
	case Quadratic, Binary:
		c.bits = n / 2
	case Octal:
		c.bits = 3
	default:
		panic("incorrect number of forks")
	}
	return c
}

// Len 获取元素数量
func (c *Heap[T]) Len() int {
	return len(c.data)
}

func (c *Heap[T]) less(i, j int) bool {
	return c.cmp(c.data[i], c.data[j])
}

func (c *Heap[T]) swap(i, j int) {
	c.data[i], c.data[j] = c.data[j], c.data[i]
}

func (c *Heap[T]) up(i int) {
	var j = (i - 1) >> c.bits
	if i >= 1 && c.less(i, j) {
		c.swap(i, j)
		c.up(j)
	}
}

func (c *Heap[T]) down(i, n int) {
	var base = i << c.bits
	var index = base + 1
	if index >= n {
		return
	}

	for j := base + 2; j <= base+c.forks && j < n; j++ {
		if c.less(j, index) {
			index = j
		}
	}

	if c.less(index, i) {
		c.swap(i, index)
		c.down(index, n)
	}
}

// Reset 重置堆
func (c *Heap[T]) Reset() {
	clear(c.data)
	c.data = c.data[:0]
}

// Push 追加元素
func (c *Heap[T]) Push(v T) {
	c.data = append(c.data, v)
	c.up(c.Len() - 1)
}

// Pop 弹出堆顶元素
func (c *Heap[T]) Pop() (ele T) {
	var n = c.Len()
	switch n {
	case 0:
	case 1:
		ele = c.data[0]
		c.data = c.data[:0]
	default:
		ele = c.data[0]
		c.data[0] = c.data[n-1]
		c.data = c.data[:n-1]
		c.down(0, n-1)
	}
	return
}

// Top 获取堆顶元素
func (c *Heap[T]) Top() T {
	return c.data[0]
}

// Range 遍历
func (c *Heap[T]) Range(f func(index int, value T) bool) {
	for i, v := range c.data {
		if !f(i, v) {
			return
		}
	}
}

func (c *Heap[T]) Clone() *Heap[T] {
	var v = *c
	v.data = make([]T, len(c.data))
	copy(v.data, c.data)
	return &v
}
