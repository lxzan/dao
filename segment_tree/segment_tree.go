package segment_tree

type segment_tree_interface[T any, Element any] interface {
	New(T) Element
	Merge(Element, Element) Element
}

type segment_tree[T any, Element segment_tree_interface[T, Element]] struct {
	L     int
	R     int
	Node  N
	LNode *segment_tree[T, Element]
	RNode *segment_tree[T, Element]
}

func newLineTree[T any, Element segment_tree_interface[T, Element]](arr []T) *segment_tree[T, Element] {
	var n = len(arr)
	if n == 0 {
		panic("arr cannot be empty!")
	}
	var obj = &segment_tree[T, Element]{
		L: 0,
		R: len(arr) - 1,
	}
	var vnode Element
	obj.build(obj, arr, vnode)
	return obj
}

func (c *segment_tree[T, Element]) build(cur *segment_tree[T, Element], arr []T, vnode Element) {
	if cur.L == cur.R {
		cur.Node = vnode.New(arr[cur.L])
		return
	}

	var mid = (cur.L + cur.R) / 2
	cur.LNode = &segment_tree[T, Element]{L: cur.L, R: mid}
	cur.RNode = &segment_tree[T, Element]{L: mid + 1, R: cur.R}
	c.build(cur.LNode, arr, vnode)
	c.build(cur.RNode, arr, vnode)
	cur.Node = vnode.Merge(cur.LNode.Node, cur.RNode.Node)
}

func (c *segment_tree[T, Element]) Query(left int, right int) Element {
	var result Element
	var serial = 0
	c.doQuery(&serial, c, left, right, &result)
	return result
}

func (c *segment_tree[T, Element]) doQuery(serial *int, cur *segment_tree[T, Element], left int, right int, result *N) {
	if cur.L >= left && cur.R <= right {
		if *serial == 0 {
			*result = cur.Node
		} else {
			*result = result.Merge(*result, cur.Node)
		}
		*serial++
	} else if !(cur.L > right || cur.R < left) {
		c.doQuery(serial, cur.LNode, left, right, result)
		c.doQuery(serial, cur.RNode, left, right, result)
	}
}

