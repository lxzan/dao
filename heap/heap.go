package heap

import (
	"github.com/lxzan/dao/algorithm"
	"github.com/lxzan/dao/internal/utils"
	"github.com/lxzan/dao/types/cmp"
)

const (
	Binary = 2

	Quadratic = 4

	Octal = 8
)

// New 新建一个最小四叉堆
// Create a new minimum quadratic heap
func New[T cmp.Ordered]() *Heap[T] { return NewWithForks(Quadratic, cmp.Less[T]) }

// NewWithForks 新建堆
// @forks 分叉数, forks=pow(2,n)
// @lessFunc 比较函数
func NewWithForks[T any](forks uint32, lessFunc cmp.LessFunc[T]) *Heap[T] {
	h := &Heap[T]{lessFunc: lessFunc}
	h.setForkNumber(forks)
	return h
}

type Heap[T any] struct {
	bits     int
	forks    int
	data     []T
	lessFunc func(a, b T) bool
}

// SetCap 设置预分配容量
func (c *Heap[T]) SetCap(n int) *Heap[T] {
	c.data = make([]T, 0, n)
	return c
}

// setForkNumber 设置分叉数
func (c *Heap[T]) setForkNumber(n uint32) *Heap[T] {
	n = algorithm.SelectValue(n == 0, Quadratic, n)
	if !utils.IsBinaryNumber(n) {
		panic("incorrect number of forks")
	}
	c.forks = int(n)
	c.bits = utils.GetBinaryExponential(c.forks)
	return c
}

// Len 获取元素数量
func (c *Heap[T]) Len() int {
	return len(c.data)
}

func (c *Heap[T]) less(i, j int) bool {
	return c.lessFunc(c.data[i], c.data[j])
}

func (c *Heap[T]) swap(i, j int) {
	c.data[i], c.data[j] = c.data[j], c.data[i]
}

func (c *Heap[T]) up(i int) {
	for i > 0 {
		var j = (i - 1) >> c.bits
		if !c.less(i, j) {
			return
		}

		c.swap(i, j)
		i = j
	}
}

func (c *Heap[T]) down(i int) {
	var n = c.Len()
	for {
		var base = i << c.bits
		var index = base + 1
		if index >= n {
			return
		}

		var end = algorithm.Min(base+c.forks, n-1)
		for j := base + 2; j <= end; j++ {
			if c.less(j, index) {
				index = j
			}
		}

		if !c.less(index, i) {
			return
		}

		c.swap(i, index)
		i = index
	}
}

// Reset 重置堆
func (c *Heap[T]) Reset() {
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
		c.down(0)
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
	v.data = utils.Clone(c.data)
	return &v
}

// UnWrap 解包, 返回底层数组
func (c *Heap[T]) UnWrap() []T {
	return c.data
}
