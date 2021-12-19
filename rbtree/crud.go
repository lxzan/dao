package rbtree

import (
	"github.com/lxzan/dao"
	"github.com/lxzan/dao/heap"
)

func (c *RBTree[K, V]) Len() int {
	return c.length
}

func New[K dao.Comparable[K], V any]() *RBTree[K, V] {
	var node rbtree_node[K, V]
	return &RBTree[K, V]{root: &node, sentinel: &node}
}

func (c *RBTree[K, V]) Clear() {
	var node rbtree_node[K, V]
	c.root = &node
	c.sentinel = &node
	c.length = 0
}

// insert with unique check
func (c *RBTree[K, V]) Insert(key K, val V) (success bool) {
	for i := c.begin(); !c.end(i); i = c.next(i, key) {
		if i.key == key {
			return false
		}
	}

	c.length++
	var node = &rbtree_node[K, V]{key: key, data: &val}
	var root = &c.root
	var temp, sentinel *rbtree_node[K, V]

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

func (c *RBTree[K, V]) Delete(key K) (success bool) {
	for i := c.begin(); !c.end(i); i = c.next(i, key) {
		if i.key == key {
			c.length--
			c.do_delete(i)
			return true
		}
	}
	return false
}

func (c *RBTree[K, V]) Find(key K) (result *V, exist bool) {
	for i := c.begin(); !c.end(i); i = c.next(i, key) {
		if i.key == key {
			return i.data, true
		}
	}
	return nil, false
}

func (c *RBTree[K, V]) Update(key K, val *V) (success bool) {
	for i := c.begin(); !c.end(i); i = c.next(i, key) {
		if i.key == key {
			i.data = val
			return true
		}
	}
	return false
}

func (c *RBTree[K, V]) ForEach(fn func(key K, val *V) (continued bool)) {
	var next = true
	c.do_foreach(c.root, &next, fn)
}

func (c *RBTree[K, V]) do_foreach(node *rbtree_node[K, V], next *bool, fn func(key K, val *V) bool) {
	if c.end(node) || !(*next) {
		return
	}
	*next = fn(node.key, node.data)
	c.do_foreach(node.left, next, fn)
	c.do_foreach(node.right, next, fn)
}

func (c *RBTree[K, V]) GetMaxKey() K {
	var maxKey = c.root.key
	c.do_get_max_key(c.root, &maxKey)
	return maxKey
}

func (c *RBTree[K, V]) do_get_max_key(node *rbtree_node[K, V], maxKey *K) {
	if c.end(node) {
		return
	}
	if node.key > *maxKey {
		*maxKey = node.key
	}
	c.do_get_max_key(node.right, maxKey)
}

func (c *RBTree[K, V]) GetMinKey() K {
	var minKey = c.root.key
	c.do_get_min_key(c.root, &minKey)
	return minKey
}

func (c *RBTree[K, V]) do_get_min_key(node *rbtree_node[K, V], minKey *K) {
	if c.end(node) {
		return
	}
	if node.key < *minKey {
		*minKey = node.key
	}
	c.do_get_min_key(node.left, minKey)
}

type Order uint8

const (
	ASC  Order = 0
	DESC Order = 1
)

func AlwaysTrue[K dao.Comparable[K]](key K) bool {
	return true
}

type QueryBuilder[K dao.Comparable[K]] struct {
	LeftFilter  func(key K) bool
	RightFilter func(key K) bool
	Limit       int
	Order       Order
	results     *heap.Heap[K]
}

func (c *QueryBuilder[K]) init() *QueryBuilder[K] {
	var typ = heap.MaxHeap[K]
	if c.Order == DESC {
		typ = heap.MinHeap[K]
	}
	if c.LeftFilter == nil {
		c.LeftFilter = AlwaysTrue[K]
	}
	if c.RightFilter == nil {
		c.RightFilter = AlwaysTrue[K]
	}
	if c.Limit <= 0 {
		c.results = heap.New[K](10, typ)
	} else {
		c.results = heap.New[K](c.Limit, typ)
	}
	return c
}

func (c *RBTree[K, V]) Query(q *QueryBuilder[K]) []K {
	c.do_query1(c.root, q.init())
	return q.results.Sort()
}

func (c *RBTree[K, V]) do_query1(node *rbtree_node[K, V], q *QueryBuilder[K]) {
	if c.end(node) {
		return
	}
	if q.LeftFilter(node.key) && q.RightFilter(node.key) {
		if q.results.Len() < q.Limit {
			q.results.Push(node.key)
			c.do_query2(node, q)
		} else if q.results.Less(q.results.Data[0], node.key) {
			q.results.Data[0] = node.key
			q.results.Down(0, q.Limit)
			c.do_query2(node, q)
		} else {
			c.do_query1(node.left, q)
		}
	} else {
		if !q.LeftFilter(node.key) {
			c.do_query1(node.right, q)
		} else if !q.RightFilter(node.key) {
			c.do_query1(node.left, q)
		}
	}
}

func (c *RBTree[K, V]) do_query2(node *rbtree_node[K, V], q *QueryBuilder[K]) {
	if q.Order == ASC {
		c.do_query1(node.left, q)
		c.do_query1(node.right, q)
	} else {
		c.do_query1(node.right, q)
		c.do_query1(node.left, q)
	}
}
