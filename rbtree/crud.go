package rbtree

import (
	"github.com/lxzan/dao"
	"github.com/lxzan/dao/heap"
)

func (c *RBTree[T]) Len() int {
	return c.length
}

func (c *RBTree[T]) Clear() {
	var node rbtree_node[T]
	c.root = &node
	c.sentinel = &node
	c.length = 0
}

// insert with unique check
func (c *RBTree[T]) Insert(data *T) (success bool) {
	for i := c.begin(); !c.end(i); i = c.next(i, data) {
		if c.cmp(data, i.data) == dao.Equal {
			return false
		}
	}

	c.length++
	var node = &rbtree_node[T]{data: data}
	var root = &c.root
	var temp, sentinel *rbtree_node[T]

	/* a binary tree insert */

	sentinel = c.sentinel
	if *root == sentinel {
		node.parent = nil
		node.left = sentinel
		node.right = sentinel
		node.set_black()
		*root = node

		return
	}
	c.do_insert(*root, node, sentinel)

	/* re-balance tree */

	for node != *root && node.parent.is_red() {
		if node.parent == node.parent.parent.left {
			temp = node.parent.parent.right
			if temp.is_red() {
				node.parent.set_black()
				temp.set_black()
				node.parent.parent.set_red()
				node = node.parent.parent
			} else {
				if node == node.parent.right {
					node = node.parent
					c.left_rotate(root, sentinel, node)
				}
				node.parent.set_black()
				node.parent.parent.set_red()
				c.right_rotate(root, sentinel, node.parent.parent)
			}
		} else {
			temp = node.parent.parent.left

			if temp.is_red() {
				node.parent.set_black()
				temp.set_black()
				node.parent.parent.set_red()
				node = node.parent.parent
			} else {
				if node == node.parent.left {
					node = node.parent
					c.right_rotate(root, sentinel, node)
				}
				node.parent.set_black()
				node.parent.parent.set_red()
				c.left_rotate(root, sentinel, node.parent.parent)
			}
		}
	}
	(*root).set_black()
	return true
}

func (c *RBTree[T]) Delete(ele *T) (success bool) {
	for i := c.begin(); !c.end(i); i = c.next(i, ele) {
		if c.cmp(ele, i.data) == dao.Equal {
			c.length--
			c.do_delete(i)
			return true
		}
	}
	return false
}

func (c *RBTree[T]) Find(ele *T) (result *T, exist bool) {
	for i := c.begin(); !c.end(i); i = c.next(i, ele) {
		if c.cmp(ele, i.data) == dao.Equal {
			return i.data, true
		}
	}
	return nil, false
}

func (c *RBTree[T]) Update(data *T) (success bool) {
	for i := c.begin(); !c.end(i); i = c.next(i, data) {
		if c.cmp(i.data, data) == dao.Equal {
			*(i.data) = *data
			return true
		}
	}
	return false
}

func (c *RBTree[T]) ForEach(fn func(item *T) (continued bool)) {
	var next = true
	c.do_foreach(c.root, &next, fn)
}

func (c *RBTree[T]) do_foreach(node *rbtree_node[T], next *bool, fn func(item *T) bool) {
	if c.end(node) || !(*next) {
		return
	}
	*next = fn(node.data)
	c.do_foreach(node.left, next, fn)
	c.do_foreach(node.right, next, fn)
}

func (c *RBTree[T]) GetMaxKey() *T {
	var maxKey = *(c.root.data)
	c.do_get_max_key(c.root, &maxKey)
	return &maxKey
}

func (c *RBTree[T]) do_get_max_key(node *rbtree_node[T], maxKey *T) {
	if c.end(node) {
		return
	}
	if c.cmp(node.data, maxKey) == dao.Greater {
		*maxKey = *(node.data)
	}
	c.do_get_max_key(node.right, maxKey)
}

func (c *RBTree[T]) GetMinKey() *T {
	var minKey = *(c.root.data)
	c.do_get_min_key(c.root, &minKey)
	return &minKey
}

func (c *RBTree[T]) do_get_min_key(node *rbtree_node[T], minKey *T) {
	if c.end(node) {
		return
	}
	if c.cmp(node.data, minKey) == dao.Less {
		*minKey = *(node.data)
	}
	c.do_get_min_key(node.right, minKey)
}

type Order uint8

const (
	ASC  Order = 0
	DESC Order = 1
)

func AlwaysTrue[T any](d *T) bool {
	return true
}

type QueryBuilder[T any] struct {
	LeftFilter  func(d *T) bool
	RightFilter func(d *T) bool
	Limit       int
	Order       Order
	results     *heap.Heap[*T]
}

func (c *QueryBuilder[T]) init(cmp func(a, b *T) dao.Ordering) *QueryBuilder[T] {
	var typ = func(a, b *T) bool {
		return cmp(a, b) == dao.Greater
	}
	if c.Order == DESC {
		typ = func(a, b *T) bool {
			return cmp(a, b) == dao.Less
		}
	}
	if c.LeftFilter == nil {
		c.LeftFilter = AlwaysTrue[T]
	}
	if c.RightFilter == nil {
		c.RightFilter = AlwaysTrue[T]
	}
	if c.Limit <= 0 {
		c.results = heap.New[*T](10, typ)
	} else {
		c.results = heap.New[*T](c.Limit, typ)
	}
	return c
}

func (c *RBTree[T]) Query(q *QueryBuilder[T]) []*T {
	c.do_query1(c.root, q.init(c.cmp))
	return q.results.Sort()
}

func (c *RBTree[T]) do_query1(node *rbtree_node[T], q *QueryBuilder[T]) {
	if c.end(node) {
		return
	}
	if q.LeftFilter(node.data) && q.RightFilter(node.data) {
		if q.results.Len() < q.Limit {
			q.results.Push(node.data)
			c.do_query2(node, q)
		} else if q.results.Less(q.results.Data[0], node.data) {
			q.results.Data[0] = node.data
			q.results.Down(0, q.Limit)
			c.do_query2(node, q)
		} else {
			c.do_query1(node.left, q)
		}
	} else {
		if !q.LeftFilter(node.data) {
			c.do_query1(node.right, q)
		} else if !q.RightFilter(node.data) {
			c.do_query1(node.left, q)
		}
	}
}

func (c *RBTree[T]) do_query2(node *rbtree_node[T], q *QueryBuilder[T]) {
	if q.Order == ASC {
		c.do_query1(node.left, q)
		c.do_query1(node.right, q)
	} else {
		c.do_query1(node.right, q)
		c.do_query1(node.left, q)
	}
}
