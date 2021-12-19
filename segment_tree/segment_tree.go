package segment_tree

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type Element struct {
	L     int
	R     int
	Sum   int
	LNode *Element
	RNode *Element
}

type SegmentTree struct {
	root *Element
	arr  []int
}

func New(arr []int) *SegmentTree {
	var obj = &SegmentTree{
		root: &Element{
			L: 0,
			R: len(arr) - 1,
		},
		arr: arr,
	}
	obj.build(obj.root)
	return obj
}

func (c *SegmentTree) build(cur *Element) {
	if cur.L == cur.R {
		cur.Sum = c.arr[cur.L]
		return
	}

	var mid = (cur.L + cur.R) / 2
	cur.LNode = &Element{
		L: cur.L,
		R: mid,
	}
	cur.RNode = &Element{
		L: mid + 1,
		R: cur.R,
	}
	c.build(cur.LNode)
	c.build(cur.RNode)
	cur.Sum = cur.LNode.Sum + cur.RNode.Sum
}

func (c *SegmentTree) Query(left int, right int) int {
	var result = 0
	c.doQuery(c.root, left, right, &result)
	return result
}

func (c *SegmentTree) doQuery(cur *Element, left int, right int, result *int) {
	if cur.L >= left && cur.R <= right {
		*result += cur.Sum
	} else if !(cur.L > right || cur.R < left) {
		c.doQuery(cur.LNode, left, right, result)
		c.doQuery(cur.RNode, left, right, result)
	}
}

func (c *SegmentTree) Update(i, v int) {
	c.doUpdate(c.root, c.root.L, c.root.R, i, v-c.arr[i])
	c.arr[i] = v
}

func (c *SegmentTree) doUpdate(cur *Element, left int, right int, i int, delta int) {
	if cur.L >= left && cur.R <= right {
		cur.Sum += delta
	} else if !(cur.L > right || cur.R < left) {
		c.doUpdate(cur.LNode, left, right, i, delta)
		c.doUpdate(cur.RNode, left, right, i, delta)
	}
}
