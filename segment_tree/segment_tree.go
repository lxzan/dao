package segment_tree

type Operate uint8

const (
	OperateCreate Operate = 0
	OperateQuery  Operate = 1
	OperateUpdate Operate = 2
)

type element[T any, S any] struct {
	left     int
	right    int
	son      *element[T, S]
	daughter *element[T, S]
	data     S
}

type SegmentTree[T any, S any] struct {
	root         *element[T, S]
	arr          []T
	newSummary   NewSummary[T, S]
	mergeSummary MergeSummary[S]
}

func New[T any, S any](arr []T, newSummary NewSummary[T, S], mergeSummary MergeSummary[S]) *SegmentTree[T, S] {
	var obj = &SegmentTree[T, S]{
		root: &element[T, S]{
			left:  0,
			right: len(arr) - 1,
		},
		arr:          arr,
		newSummary:   newSummary,
		mergeSummary: mergeSummary,
	}
	obj.build(obj.root)
	return obj
}

func (c *SegmentTree[T, S]) build(cur *element[T, S]) {
	if cur.left == cur.right {
		cur.data = c.newSummary(c.arr[cur.left], OperateCreate)
		return
	}
	var mid = (cur.left + cur.right) / 2
	cur.son = &element[T, S]{
		left:  cur.left,
		right: mid,
	}
	cur.daughter = &element[T, S]{
		left:  mid + 1,
		right: cur.right,
	}
	c.build(cur.son)
	c.build(cur.daughter)
	cur.data = c.mergeSummary(cur.son.data, cur.daughter.data)
}

// Query 查询 begin <= index < end 区间
func (c *SegmentTree[T, S]) Query(begin int, end int) S {
	result := c.newSummary(c.arr[begin], OperateQuery)
	c.doQuery(c.root, begin, end-1, &result)
	return result
}

func (c *SegmentTree[T, S]) doQuery(cur *element[T, S], left int, right int, result *S) {
	if cur.left >= left && cur.right <= right {
		*result = c.mergeSummary(*result, cur.data)
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

func (c *SegmentTree[T, S]) rebuild(cur *element[T, S], i int) {
	if !(i >= cur.left && i <= cur.right) {
		return
	}
	if cur.left == cur.right && cur.left == i {
		cur.data = c.newSummary(c.arr[cur.left], OperateUpdate)
		return
	}
	c.rebuild(cur.son, i)
	c.rebuild(cur.daughter, i)
	cur.data = c.mergeSummary(cur.son.data, cur.daughter.data)
}
