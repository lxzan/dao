package segment_tree

type Operate uint8

const (
	OperateCreate Operate = 0
	OperateQuery  Operate = 1
	OperateUpdate Operate = 2
)

type Element[T any, S any] struct {
	left     int
	right    int
	son      *Element[T, S]
	daughter *Element[T, S]
	data     S
}

type SegmentTree[T any, S any] struct {
	init  func(op Operate, x T) S
	merge func(a, b S) S
	root  *Element[T, S]
	arr   []T
}

func New[T any, S any](arr []T, init func(op Operate, x T) S, merge func(a, b S) S) *SegmentTree[T, S] {
	var obj = &SegmentTree[T, S]{
		init:  init,
		merge: merge,
		root: &Element[T, S]{
			left:  0,
			right: len(arr) - 1,
		},
		arr: arr,
	}
	obj.build(obj.root)
	return obj
}

func (c *SegmentTree[T, S]) build(cur *Element[T, S]) {
	if cur.left == cur.right {
		cur.data = c.init(OperateCreate, c.arr[cur.left])
		return
	}

	var mid = (cur.left + cur.right) / 2
	cur.son = &Element[T, S]{
		left:  cur.left,
		right: mid,
	}
	cur.daughter = &Element[T, S]{
		left:  mid + 1,
		right: cur.right,
	}
	c.build(cur.son)
	c.build(cur.daughter)
	cur.data = c.merge(cur.son.data, cur.daughter.data)
}

// Query 查询
func (c *SegmentTree[T, S]) Query(left int, right int) S {
	var result S
	result = c.init(OperateQuery, c.arr[left])
	c.doQuery(c.root, left, right, &result)
	return result
}

func (c *SegmentTree[T, S]) doQuery(cur *Element[T, S], left int, right int, result *S) {
	if cur.left >= left && cur.right <= right {
		*result = c.merge(*result, cur.data)
	} else if !(cur.left > right || cur.right < left) {
		c.doQuery(cur.son, left, right, result)
		c.doQuery(cur.daughter, left, right, result)
	}
}

// Update 更新
func (c *SegmentTree[T, S]) Update(i int, v T) {
	c.arr[i] = v
	c.rebuild(c.root, i)
}

func (c *SegmentTree[T, S]) rebuild(cur *Element[T, S], i int) {
	if !(i >= cur.left && i <= cur.right) {
		return
	}

	if cur.left == cur.right && cur.left == i {
		cur.data = c.init(OperateUpdate, c.arr[cur.left])
		return
	}

	c.rebuild(cur.son, i)
	c.rebuild(cur.daughter, i)
	cur.data = c.merge(cur.son.data, cur.daughter.data)
}
