package rbtree

import (
	"github.com/lxzan/dao"
	"github.com/lxzan/dao/vector"
)

func (this *RBTree[K, V]) Find(key K) (result V, exist bool) {
	var data = &Iterator[K, V]{Key: key}
	for i := this.begin(); !this.end(i); i = this.next(i, data) {
		if key == i.data.Key {
			return i.data.Val, true
		}
	}
	return result, false
}

func (this *RBTree[K, V]) ForEach(fn func(iter *Iterator[K, V])) {
	var iter = &Iterator[K, V]{next: true}
	this.do_foreach(this.root, iter, fn)
}

func (this *RBTree[K, V]) do_foreach(node *rbtree_node[K, V], iter *Iterator[K, V], fn func(*Iterator[K, V])) {
	if this.end(node) || !iter.next {
		return
	}

	iter.Key = node.data.Key
	iter.Val = node.data.Val
	fn(iter)
	this.do_foreach(node.left, iter, fn)
	this.do_foreach(node.right, iter, fn)
}

func (this *RBTree[K, V]) GetMinKey(filter func(key K) bool) (result *Iterator[K, V], exist bool) {
	result = &Iterator[K, V]{}
	var stack = vector.New[*rbtree_node[K, V]]()
	stack.Push(this.root)
	for stack.Len() > 0 {
		var node = stack.RPop()
		if this.end(node) {
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

func (this *RBTree[K, V]) GetMaxKey(filter func(key K) bool) (result *Iterator[K, V], exist bool) {
	result = &Iterator[K, V]{}
	var stack = vector.New[*rbtree_node[K, V]](0, 0)
	stack.Push(this.root)
	for stack.Len() > 0 {
		var node = stack.RPop()
		if this.end(node) {
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

func (this *QueryBuilder[K]) init() *QueryBuilder[K] {
	if this.LeftFilter == nil {
		this.LeftFilter = AlwaysTrue[K]
	}
	if this.RightFilter == nil {
		this.RightFilter = AlwaysTrue[K]
	}
	return this
}

func (this *RBTree[K, V]) Query(q *QueryBuilder[K]) []*Iterator[K, V] {
	q.init()
	var results = make([]*Iterator[K, V], 0)
	if q.Order == DESC {
		res, exist := this.GetMaxKey(q.RightFilter)
		if exist && q.LeftFilter(res.Key) {
			results = append(results, res)
		} else {
			return results
		}

		for i := 0; i < q.Limit-1; i++ {
			res, exist = this.GetMaxKey(func(key K) bool {
				return key < res.Key
			})
			if exist && q.LeftFilter(res.Key) {
				results = append(results, res)
			} else {
				break
			}
		}
	} else {
		res, exist := this.GetMinKey(q.LeftFilter)
		if exist && q.RightFilter(res.Key) {
			results = append(results, res)
		} else {
			return results
		}

		for i := 0; i < q.Limit-1; i++ {
			res, exist = this.GetMinKey(func(key K) bool {
				return key > res.Key
			})
			if exist && q.RightFilter(res.Key) {
				results = append(results, res)
			} else {
				break
			}
		}
	}
	return results
}
