package heap

import (
	"github.com/lxzan/dao/algo"
	"github.com/lxzan/dao/types/cmp"
)

type (
	Element[K comparable, V any] struct {
		index int
		key   K
		value V
	}

	// HashHeap 带哈希索引的四叉堆, 简称哈希堆
	HashHeap[K comparable, V any] struct {
		heap    *underlyingHeap[K, V]
		mapping map[K]*Element[K, V]
	}
)

func (c *Element[K, V]) Key() K { return c.key }

func (c *Element[K, V]) Value() V { return c.value }

// NewHashHeap 创建一个哈希堆
func NewHashHeap[K comparable, V any](less cmp.LessFunc[V]) *HashHeap[K, V] {
	return &HashHeap[K, V]{
		heap:    &underlyingHeap[K, V]{lessFunc: less},
		mapping: make(map[K]*Element[K, V]),
	}
}

// Len 返回元素数量
func (c *HashHeap[K, V]) Len() int { return len(c.mapping) }

// Set 设置键值
func (c *HashHeap[K, V]) Set(k K, v V) {
	ele, ok := c.mapping[k]
	if !ok {
		ele = &Element[K, V]{key: k, value: v}
		c.mapping[k] = ele
		c.heap.Push(ele)
		return
	}
	c.heap.Update(ele, v)
}

// Get 查询
func (c *HashHeap[K, V]) Get(k K) (V, bool) {
	v, ok := c.mapping[k]
	return v.value, ok
}

// Delete 删除一个元素
func (c *HashHeap[K, V]) Delete(k K) {
	ele, ok := c.mapping[k]
	if !ok {
		return
	}
	c.heap.Delete(ele.index)
	delete(c.mapping, k)
}

// Top 返回堆顶元素
func (c *HashHeap[K, V]) Top() *Element[K, V] {
	if c.Len() > 0 {
		return c.heap.data[0]
	}
	return nil
}

// Pop 弹出堆顶元素
func (c *HashHeap[K, V]) Pop() *Element[K, V] {
	if ele := c.Top(); ele != nil {
		c.Delete(ele.key)
		return ele
	}
	return nil
}

// Range 遍历堆元素
func (c *HashHeap[K, V]) Range(f func(ele *Element[K, V]) bool) {
	for _, ele := range c.heap.data {
		if !f(ele) {
			return
		}
	}
}

type underlyingHeap[K comparable, V any] struct {
	data     []*Element[K, V]
	lessFunc func(a, b V) bool
}

func (c *underlyingHeap[K, V]) Len() int {
	return len(c.data)
}

func (c *underlyingHeap[K, V]) Push(ele *Element[K, V]) {
	ele.index = c.Len()
	c.data = append(c.data, ele)
	c.up(c.Len() - 1)
}

func (c *underlyingHeap[K, V]) up(i int) {
	var j = (i - 1) >> 2
	if i >= 1 && c.lessFunc(c.data[i].value, c.data[j].value) {
		c.swap(i, j)
		c.up(j)
	}
}

func (c *underlyingHeap[K, V]) swap(i, j int) {
	c.data[i].index, c.data[j].index = c.data[j].index, c.data[i].index
	c.data[i], c.data[j] = c.data[j], c.data[i]
}

func (c *underlyingHeap[K, V]) Pop() (ele *Element[K, V]) {
	var n = c.Len()
	switch n {
	case 0:
	case 1:
		ele = c.data[0]
		c.data = c.data[:0]
	default:
		ele = c.data[0]
		c.swap(0, n-1)
		c.data = c.data[:n-1]
		c.down(0, n-1)
	}
	return
}

func (c *underlyingHeap[K, V]) down(i, n int) {
	var base = i << 2
	var index = base + 1
	if index >= n {
		return
	}

	var end = algo.Min(base+4, n-1)
	for j := base + 2; j <= end; j++ {
		if c.lessFunc(c.data[j].value, c.data[index].value) {
			index = j
		}
	}

	if c.lessFunc(c.data[index].value, c.data[i].value) {
		c.swap(i, index)
		c.down(index, n)
	}
}

func (c *underlyingHeap[K, V]) Update(ele *Element[K, V], value V) {
	ele.value = value
	var down bool
	if ele.index == 0 {
		down = true
	} else {
		var i = ele.index >> 2
		var p = c.data[i]
		down = c.lessFunc(p.value, ele.value)
	}

	if down {
		c.down(ele.index, c.Len())
	} else {
		c.up(ele.index)
	}
}

func (c *underlyingHeap[K, V]) Delete(i int) {
	if i == 0 {
		c.Pop()
		return
	}

	var n = c.Len()
	var down = c.lessFunc(c.data[i].value, c.data[n-1].value)
	c.swap(i, n-1)
	c.data = c.data[:n-1]
	if i < n-1 {
		if down {
			c.down(i, n-1)
		} else {
			c.up(i)
		}
	}
}
