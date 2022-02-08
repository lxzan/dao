package segment_tree

type Operate uint8

const (
	Operate_Create Operate = 0
	Operate_Query  Operate = 1
	Operate_Update Operate = 2
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

func (this *SegmentTree[T, S]) build(cur *Element[T, S]) {
	if cur.left == cur.right {
		cur.data = this.init(Operate_Create, this.arr[cur.left])
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
	this.build(cur.son)
	this.build(cur.daughter)
	cur.data = this.merge(cur.son.data, cur.daughter.data)
}

func (this *SegmentTree[T, S]) Query(left int, right int) S {
	var result S
	result = this.init(Operate_Query, this.arr[left])
	this.doQuery(this.root, left, right, &result)
	return result
}

func (this *SegmentTree[T, S]) doQuery(cur *Element[T, S], left int, right int, result *S) {
	if cur.left >= left && cur.right <= right {
		*result = this.merge(*result, cur.data)
	} else if !(cur.left > right || cur.right < left) {
		this.doQuery(cur.son, left, right, result)
		this.doQuery(cur.daughter, left, right, result)
	}
}

func (this *SegmentTree[T, S]) Update(i int, v T) {
	this.arr[i] = v
	this.rebuild(this.root, i)
}

func (this *SegmentTree[T, S]) rebuild(cur *Element[T, S], i int) {
	if !(i >= cur.left && i <= cur.right) {
		return
	}

	if cur.left == cur.right && cur.left == i {
		cur.data = this.init(Operate_Update, this.arr[cur.left])
		return
	}

	this.rebuild(cur.son, i)
	this.rebuild(cur.daughter, i)
	cur.data = this.merge(cur.son.data, cur.daughter.data)
}
