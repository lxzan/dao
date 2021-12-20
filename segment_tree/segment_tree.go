package segment_tree

import (
	"github.com/lxzan/dao"
)

type Operate uint8

const (
	Operate_Create Operate = 0
	Operate_Query  Operate = 1
	Operate_Update Operate = 2
)

type Interface[T any, S any] interface {
	Init(op Operate, x T) S
	Merge(a, b S) S
}

type Element[T dao.Number[T]] struct {
	left     int
	right    int
	son      *Element[T]
	daughter *Element[T]
	data     Schema[T]
}

type SegmentTree[T dao.Number[T], S Interface[T, S]] struct {
	root *Element[T]
	arr  []T
}

func New[T dao.Number[T], S Interface[T, S]](arr []T) *SegmentTree[T, S] {
	var obj = &SegmentTree[T, S]{
		root: &Element[T]{
			left:  0,
			right: len(arr) - 1,
		},
		arr: arr,
	}
	obj.build(obj.root)
	return obj
}

func (c *SegmentTree[T, S]) build(cur *Element[T]) {
	if cur.left == cur.right {
		cur.data = cur.data.Init(Operate_Create, c.arr[cur.left])
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
	cur.data = cur.data.Merge(cur.son.data, cur.daughter.data)
}

func (c *SegmentTree[T, S]) Query(left int, right int) Schema[T] {
	var result Schema[T]
	result = result.Init(Operate_Query, c.arr[left])
	c.doQuery(c.root, left, right, &result)
	return result
}

func (c *SegmentTree[T, S]) doQuery(cur *Element[T], left int, right int, result *Schema[T]) {
	if cur.left >= left && cur.right <= right {
		*result = result.Merge(*result, cur.data)
	} else if !(cur.left > right || cur.right < left) {
		c.doQuery(cur.son, left, right, result)
		c.doQuery(cur.daughter, left, right, result)
	}
}

func (c *SegmentTree[T, S]) Update(i int, v T) {
	c.arr[i] = v
	c.rebuild(c.root, i)
}

func (c *SegmentTree[T, S]) rebuild(cur *Element[T], i int) {
	if !(i >= cur.left && i <= cur.right) {
		return
	}

	if cur.left == cur.right && cur.left == i {
		cur.data = cur.data.Init(Operate_Update, c.arr[cur.left])
		return
	}

	c.rebuild(cur.son, i)
	c.rebuild(cur.daughter, i)
	cur.data = cur.data.Merge(cur.son.data, cur.daughter.data)
}
