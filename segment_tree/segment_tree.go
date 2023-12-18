package segment_tree

type Operate uint8

const (
	OperateCreate Operate = 0
	OperateQuery  Operate = 1
	OperateUpdate Operate = 2
)

type (
	Initer[T any] interface {
		Init(op Operate) T
	}

	Merger[T any] interface {
		Merge(T) T
	}
)

type Element[S Merger[S], T Initer[S]] struct {
	left     int
	right    int
	son      *Element[S, T]
	daughter *Element[S, T]
	data     S
}

type SegmentTree[S Merger[S], T Initer[S]] struct {
	root *Element[S, T]
	arr  []T
}

func New[S Merger[S], T Initer[S]](arr []T) *SegmentTree[S, T] {
	var obj = &SegmentTree[S, T]{
		root: &Element[S, T]{
			left:  0,
			right: len(arr) - 1,
		},
		arr: arr,
	}
	obj.build(obj.root)
	return obj
}

func (c *SegmentTree[S, T]) build(cur *Element[S, T]) {
	if cur.left == cur.right {
		cur.data = c.arr[cur.left].Init(OperateCreate)
		return
	}

	var mid = (cur.left + cur.right) / 2
	cur.son = &Element[S, T]{
		left:  cur.left,
		right: mid,
	}
	cur.daughter = &Element[S, T]{
		left:  mid + 1,
		right: cur.right,
	}
	c.build(cur.son)
	c.build(cur.daughter)
	cur.data = cur.son.data.Merge(cur.daughter.data)
}

// Query 查询 left <= index <= right 区间
func (c *SegmentTree[S, T]) Query(left int, right int) S {
	var result S
	result = c.arr[left].Init(OperateQuery)
	c.doQuery(c.root, left, right, &result)
	return result
}

func (c *SegmentTree[S, T]) doQuery(cur *Element[S, T], left int, right int, result *S) {
	if cur.left >= left && cur.right <= right {
		*result = cur.data.Merge(*result)
	} else if !(cur.left > right || cur.right < left) {
		c.doQuery(cur.son, left, right, result)
		c.doQuery(cur.daughter, left, right, result)
	}
}

// Update 更新
func (c *SegmentTree[S, T]) Update(i int, v T) {
	c.arr[i] = v
	c.rebuild(c.root, i)
}

func (c *SegmentTree[S, T]) rebuild(cur *Element[S, T], i int) {
	if !(i >= cur.left && i <= cur.right) {
		return
	}

	if cur.left == cur.right && cur.left == i {
		cur.data = c.arr[cur.left].Init(OperateUpdate)
		return
	}

	c.rebuild(cur.son, i)
	c.rebuild(cur.daughter, i)
	cur.data = cur.son.data.Merge(cur.daughter.data)
}
