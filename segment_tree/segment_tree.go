package segment_tree

import (
	"github.com/lxzan/dao"
	"github.com/lxzan/dao/algorithm"
)

type Schema[T dao.Number[T]] struct {
	MaxValue T
	MinValue T
	Sum      T
}

type Element[T dao.Number[T]] struct {
	left     int
	right    int
	son      *Element[T]
	daughter *Element[T]
	data     Schema[T]
}

type SegmentTree[T dao.Number[T]] struct {
	root *Element[T]
	arr  []T
}

func New[T dao.Number[T]](arr []T) *SegmentTree[T] {
	var obj = &SegmentTree[T]{
		root: &Element[T]{
			left:  0,
			right: len(arr) - 1,
		},
		arr: arr,
	}
	obj.build(obj.root)
	return obj
}

func (c *SegmentTree[T]) build(cur *Element[T]) {
	if cur.left == cur.right {
		cur.data = Schema[T]{
			MaxValue: c.arr[cur.left],
			MinValue: c.arr[cur.left],
			Sum:      c.arr[cur.left],
		}
		return
	}

	var mid = (cur.left + cur.right) / 2
	cur.son = &Element[T]{
		left:  cur.left,
		right: mid,
	}
	cur.daughter = &Element[T]{
		left:  mid + 1,
		right: cur.right,
	}
	c.build(cur.son)
	c.build(cur.daughter)
	cur.data = Schema[T]{
		MaxValue: algorithm.Max[T](cur.son.data.MaxValue, cur.daughter.data.MaxValue),
		MinValue: algorithm.Min[T](cur.son.data.MinValue, cur.daughter.data.MinValue),
		Sum:      cur.son.data.Sum + cur.daughter.data.Sum,
	}
}

func (c *SegmentTree[T]) Query(left int, right int) Schema[T] {
	var result = Schema[T]{
		MaxValue: c.arr[left],
		MinValue: c.arr[left],
		Sum:      0,
	}
	c.doQuery(c.root, left, right, &result)
	return result
}

func (c *SegmentTree[T]) doQuery(cur *Element[T], left int, right int, result *Schema[T]) {
	if cur.left >= left && cur.right <= right {
		result.Sum += cur.data.Sum
		result.MaxValue = algorithm.Max[T](result.MaxValue, cur.data.MaxValue)
		result.MinValue = algorithm.Min[T](result.MinValue, cur.data.MinValue)
	} else if !(cur.left > right || cur.right < left) {
		c.doQuery(cur.son, left, right, result)
		c.doQuery(cur.daughter, left, right, result)
	}
}

func (c *SegmentTree[T]) Update(i int, v T) {
	c.arr[i] = v
	c.rebuild(c.root, i)
}

func (c *SegmentTree[T]) rebuild(cur *Element[T], i int) {
	if !(i >= cur.left && i <= cur.right) {
		return
	}

	if cur.left == cur.right && cur.left == i {
		cur.data = Schema[T]{
			MaxValue: c.arr[cur.left],
			MinValue: c.arr[cur.left],
			Sum:      c.arr[cur.left],
		}
		return
	}

	c.rebuild(cur.son, i)
	c.rebuild(cur.daughter, i)
	cur.data = Schema[T]{
		MaxValue: algorithm.Max[T](cur.son.data.MaxValue, cur.daughter.data.MaxValue),
		MinValue: algorithm.Min[T](cur.son.data.MinValue, cur.daughter.data.MinValue),
		Sum:      cur.son.data.Sum + cur.daughter.data.Sum,
	}
}
