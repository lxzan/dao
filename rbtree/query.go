package rbtree

import (
	"github.com/lxzan/dao"
	"github.com/lxzan/dao/vector"
)

func (c *RBTree[K, V]) Get(key K) (result V, exist bool) {
	for i := c.begin(); !c.end(i); i = c.next(i, key) {
		if key == i.data.Key {
			return i.data.Val, true
		}
	}
	return result, false
}

func (c *RBTree[K, V]) ForEach(fn func(iter *Iterator[K, V])) {
	var iter = &Iterator[K, V]{}
	c.do_foreach(c.root, iter, fn)
}

func (c *RBTree[K, V]) do_foreach(node *rbtree_node[K, V], iter *Iterator[K, V], fn func(iterator *Iterator[K, V])) {
	if c.end(node) || iter.broken {
		return
	}

	iter.Key = node.data.Key
	iter.Val = node.data.Val
	fn(iter)
	c.do_foreach(node.left, iter, fn)
	c.do_foreach(node.right, iter, fn)
}

func (c *RBTree[K, V]) GetMinKey(filter func(key K) bool) (result *Element[K, V], exist bool) {
	result = &Element[K, V]{}
	var stack = vector.New[*rbtree_node[K, V]]()
	stack.Push(c.root)
	for stack.Len() > 0 {
		var node = stack.RPop()
		if c.end(node) {
			continue
		}
		if filter(node.data.Key) {
			if !exist || node.data.Key < result.Key {
				exist = true
				result = node.data
			}
			stack.Push(node.left)
		} else {
			stack.Push(node.right)
		}
	}
	return result, exist
}

func (c *RBTree[K, V]) GetMaxKey(filter func(key K) bool) (result *Element[K, V], exist bool) {
	result = &Element[K, V]{}
	var stack = vector.New[*rbtree_node[K, V]](0, 0)
	stack.Push(c.root)
	for stack.Len() > 0 {
		var node = stack.RPop()
		if c.end(node) {
			continue
		}
		if filter(node.data.Key) {
			if !exist || node.data.Key > result.Key {
				exist = true
				result = node.data
			}
			stack.Push(node.right)
		} else {
			stack.Push(node.left)
		}
	}
	return result, exist
}

type Order uint8

const (
	ASC  Order = 0
	DESC Order = 1
)

func AlwaysTrue[K dao.Comparable](d K) bool {
	return true
}

type QueryBuilder[K dao.Comparable] struct {
	LeftFilter  func(d K) bool
	RightFilter func(d K) bool
	Limit       int
	Order       Order
}

func (c *QueryBuilder[K]) init() *QueryBuilder[K] {
	if c.LeftFilter == nil {
		c.LeftFilter = AlwaysTrue[K]
	}
	if c.RightFilter == nil {
		c.RightFilter = AlwaysTrue[K]
	}
	return c
}

func (c *RBTree[K, V]) Query(q *QueryBuilder[K]) []*Element[K, V] {
	q.init()
	var results = make([]*Element[K, V], 0)

	if q.Order == DESC {
		maxEle, exist := c.GetMaxKey(q.RightFilter)
		if exist && q.LeftFilter(maxEle.Key) {
			results = append(results, maxEle)
		} else {
			return results
		}

		for i := 0; i < q.Limit-1; i++ {
			result, exist := c.GetMaxKey(func(key K) bool {
				return key < maxEle.Key
			})
			if exist && q.LeftFilter(result.Key) {
				results = append(results, result)
				maxEle = result
			} else {
				break
			}
		}
	} else {
		minEle, exist := c.GetMinKey(q.LeftFilter)
		if exist && q.RightFilter(minEle.Key) {
			results = append(results, minEle)
		} else {
			return results
		}

		for i := 0; i < q.Limit-1; i++ {
			result, exist := c.GetMinKey(func(key K) bool {
				return key > minEle.Key
			})
			if exist && q.RightFilter(result.Key) {
				results = append(results, result)
				minEle = result
			} else {
				break
			}
		}
	}
	return results
}
