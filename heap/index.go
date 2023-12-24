package heap

import (
	"github.com/lxzan/dao/algorithm"
	"github.com/lxzan/dao/internal/utils"
	"github.com/lxzan/dao/types/cmp"
)

type (
	Element[K cmp.Ordered, V any] struct {
		index int
		key   K
		Value V
	}

	IndexedHeap[K cmp.Ordered, V any] struct {
		bits     int
		ways     int
		data     []*Element[K, V]
		lessFunc func(a, b K) bool
	}
)

// Key 获取排序Key
func (c *Element[K, V]) Key() K {
	return c.key
}

// Index 获取索引
// index==-1 表示元素已被删除, 不允许做更新或删除操作
func (c *Element[K, V]) Index() int {
	return c.index
}

func (c *Element[K, V]) delete() {
	c.index = -1
}

// NewIndexedHeap 新建索引堆
// @ways 分叉数, ways=pow(2,n)
// @lessFunc 比较函数, 可以传空指针, 默认为最小堆
func NewIndexedHeap[K cmp.Ordered, V any](ways uint32, lessFunc cmp.LessFunc[K]) *IndexedHeap[K, V] {
	var c = new(IndexedHeap[K, V])
	c.setWays(ways)
	c.lessFunc = lessFunc
	if c.lessFunc == nil {
		c.lessFunc = cmp.Less[K]
	}
	return c
}

// SetForkNumber 设置分叉数
func (c *IndexedHeap[K, V]) setWays(n uint32) *IndexedHeap[K, V] {
	n = algorithm.SelectValue(n == 0, Quadratic, n)
	if !utils.IsBinaryNumber(n) {
		panic("incorrect number of ways")
	}
	c.ways = int(n)
	c.bits = utils.GetBinaryExponential(c.ways)
	return c
}

// Len 获取元素数量
func (c *IndexedHeap[K, V]) Len() int {
	return len(c.data)
}

// Reset 重置堆
func (c *IndexedHeap[K, V]) Reset() {
	c.data = c.data[:0]
}

// SetCap 设置预分配容量
func (c *IndexedHeap[K, V]) SetCap(n int) *IndexedHeap[K, V] {
	c.data = make([]*Element[K, V], 0, n)
	return c
}

func (c *IndexedHeap[K, V]) swap(i, j int) {
	c.data[i].index, c.data[j].index = c.data[j].index, c.data[i].index
	c.data[i], c.data[j] = c.data[j], c.data[i]
}

func (c *IndexedHeap[K, V]) less(i, j int) bool {
	return c.lessFunc(c.data[i].key, c.data[j].key)
}

func (c *IndexedHeap[K, V]) up(i int) {
	for i > 0 {
		var j = (i - 1) >> c.bits
		if !c.less(i, j) {
			return
		}

		c.swap(i, j)
		i = j
	}
}

func (c *IndexedHeap[K, V]) down(i int) {
	var n = c.Len()
	for {
		var base = i << c.bits
		var index = base + 1
		if index >= n {
			return
		}

		var end = algorithm.Min(base+c.ways, n-1)
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

// Push 追加元素
func (c *IndexedHeap[K, V]) Push(key K, value V) *Element[K, V] {
	ele := &Element[K, V]{key: key, Value: value}
	ele.index = c.Len()
	c.data = append(c.data, ele)
	c.up(c.Len() - 1)
	return ele
}

// Pop 弹出堆顶元素
func (c *IndexedHeap[K, V]) Pop() (ele *Element[K, V]) {
	var n = c.Len()
	switch n {
	case 0:
		return ele
	case 1:
		ele = c.data[0]
		c.data = c.data[:0]
	default:
		ele = c.data[0]
		c.swap(0, n-1)
		c.data = c.data[:n-1]
		c.down(0)
	}
	ele.delete()
	return ele
}

// UpdateKeyByIndex 通过索引更新排序Key
func (c *IndexedHeap[K, V]) UpdateKeyByIndex(index int, key K) {
	ele := c.data[index]
	var down = c.lessFunc(ele.key, key)
	ele.key = key
	if down {
		c.down(ele.index)
	} else {
		c.up(ele.index)
	}
}

// GetByIndex 通过索引获取元素
func (c *IndexedHeap[K, V]) GetByIndex(index int) *Element[K, V] {
	return c.data[index]
}

// DeleteByIndex 通过索引删除元素
func (c *IndexedHeap[K, V]) DeleteByIndex(index int) {
	if index == 0 {
		c.Pop()
		return
	}

	var n = c.Len()
	var down = c.less(index, n-1)
	c.swap(index, n-1)
	c.data[n-1].delete()
	c.data = c.data[:n-1]
	if index < n-1 {
		if down {
			c.down(index)
		} else {
			c.up(index)
		}
	}
}

// Top 获取堆顶元素
func (c *IndexedHeap[K, V]) Top() *Element[K, V] {
	return c.data[0]
}

// Range 遍历
func (c *IndexedHeap[K, V]) Range(f func(ele *Element[K, V]) bool) {
	for _, v := range c.data {
		if !f(v) {
			return
		}
	}
}

// Clone 拷贝索引堆副本
func (c *IndexedHeap[K, V]) Clone() *IndexedHeap[K, V] {
	var v = *c
	v.data = utils.Clone(c.data)
	return &v
}
