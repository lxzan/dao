package rbtree

import (
	"github.com/lxzan/dao"
	"github.com/lxzan/dao/heap"
)

func (this *RBTree[T]) Len() int {
	return this.length
}

func (this *RBTree[T]) Clear() {
	var node rbtree_node[T]
	this.root = &node
	this.sentinel = &node
	this.length = 0
}

// insert with unique check
func (this *RBTree[T]) Insert(data *T) (success bool) {
	for i := this.begin(); !this.end(i); i = this.next(i, data) {
		if this.cmp(data, i.data) == dao.Equal {
			return false
		}
	}

	this.length++
	var node = &rbtree_node[T]{data: data}
	var root = &this.root
	var temp, sentinel *rbtree_node[T]

	/* a binary tree insert */

	sentinel = this.sentinel
	if *root == sentinel {
		node.parent = nil
		node.left = sentinel
		node.right = sentinel
		node.set_black()
		*root = node

		return
	}
	this.do_insert(*root, node, sentinel)

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
					this.left_rotate(root, sentinel, node)
				}
				node.parent.set_black()
				node.parent.parent.set_red()
				this.right_rotate(root, sentinel, node.parent.parent)
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
					this.right_rotate(root, sentinel, node)
				}
				node.parent.set_black()
				node.parent.parent.set_red()
				this.left_rotate(root, sentinel, node.parent.parent)
			}
		}
	}
	(*root).set_black()
	return true
}

func (this *RBTree[T]) Delete(ele *T) (success bool) {
	for i := this.begin(); !this.end(i); i = this.next(i, ele) {
		if this.cmp(ele, i.data) == dao.Equal {
			this.length--
			this.do_delete(i)
			return true
		}
	}
	return false
}

func (this *RBTree[T]) Find(ele *T) (result *T, exist bool) {
	for i := this.begin(); !this.end(i); i = this.next(i, ele) {
		if this.cmp(ele, i.data) == dao.Equal {
			return i.data, true
		}
	}
	return nil, false
}

func (this *RBTree[T]) Update(data *T) (success bool) {
	for i := this.begin(); !this.end(i); i = this.next(i, data) {
		if this.cmp(i.data, data) == dao.Equal {
			*(i.data) = *data
			return true
		}
	}
	return false
}

func (this *RBTree[T]) ForEach(fn func(item *T) (continued bool)) {
	var next = true
	this.do_foreach(this.root, &next, fn)
}

func (this *RBTree[T]) do_foreach(node *rbtree_node[T], next *bool, fn func(item *T) bool) {
	if this.end(node) || !(*next) {
		return
	}
	*next = fn(node.data)
	this.do_foreach(node.left, next, fn)
	this.do_foreach(node.right, next, fn)
}

func (this *RBTree[T]) GetMaxKey() *T {
	var maxKey = *(this.root.data)
	this.do_get_max_key(this.root, &maxKey)
	return &maxKey
}

func (this *RBTree[T]) do_get_max_key(node *rbtree_node[T], maxKey *T) {
	if this.end(node) {
		return
	}
	if this.cmp(node.data, maxKey) == dao.Greater {
		*maxKey = *(node.data)
	}
	this.do_get_max_key(node.right, maxKey)
}

func (this *RBTree[T]) GetMinKey(base *T) *T {
	if this.root.data == nil {
		return nil
	}

	var minKey T
	if this.cmp(this.root.data, base) == dao.Greater {
		minKey = *(this.root.data)
		this.do_get_min_key(this.root.left, base, &minKey)
	} else {
		this.do_get_min_key(this.root.right, base, &minKey)
	}
	return &minKey
}

func (this *RBTree[T]) do_get_min_key(node *rbtree_node[T], base *T, minKey *T) {
	if this.end(node) {
		return
	}
	if this.cmp(node.data, minKey) == dao.Less && this.cmp(node.data, base) == dao.Greater {
		*minKey = *(node.data)
	}
	this.do_get_min_key(node.left, base, minKey)
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

func (this *QueryBuilder[T]) init(cmp func(a, b *T) dao.Ordering) *QueryBuilder[T] {
	var typ = func(a, b *T) dao.Ordering {
		return -1 * cmp(a, b)
	}
	if this.Order == DESC {
		typ = cmp
	}

	if this.LeftFilter == nil {
		this.LeftFilter = AlwaysTrue[T]
	}
	if this.RightFilter == nil {
		this.RightFilter = AlwaysTrue[T]
	}
	if this.Limit <= 0 {
		this.results = heap.New(10, typ)
	} else {
		this.results = heap.New(this.Limit, typ)
	}
	return this
}

func (this *RBTree[T]) Query(q *QueryBuilder[T]) []*T {
	q.init(this.cmp)
	if q.Order == ASC {
		this.do_query_asc(this.root, q)
	} else {
		this.do_query_desc(this.root, q)
	}
	return q.results.Sort()
}

func (this *RBTree[T]) do_query_asc(node *rbtree_node[T], q *QueryBuilder[T]) {
	if this.end(node) {
		return
	}

	var flag1 = q.LeftFilter(node.data)
	var flag2 = q.RightFilter(node.data)
	if flag1 && flag2 {
		if q.results.Len() < q.Limit {
			q.results.Push(node.data)
			this.do_query_asc(node.left, q)
			this.do_query_asc(node.right, q)
		} else if q.results.Cmp(q.results.Data[0], node.data) == dao.Less {
			q.results.Data[0] = node.data
			q.results.Down(0, q.Limit)
			this.do_query_asc(node.left, q)
			this.do_query_asc(node.right, q)
		} else {
			this.do_query_asc(node.left, q)
		}
	} else {
		if !flag1 {
			this.do_query_asc(node.right, q)
		} else if !flag2 {
			this.do_query_asc(node.left, q)
		}
	}
}

func (this *RBTree[T]) do_query_desc(node *rbtree_node[T], q *QueryBuilder[T]) {
	if this.end(node) {
		return
	}

	if q.LeftFilter(node.data) && q.RightFilter(node.data) {
		if q.results.Len() < q.Limit {
			q.results.Push(node.data)
			this.do_query_desc(node.right, q)
			this.do_query_desc(node.left, q)
		} else if q.results.Cmp(q.results.Data[0], node.data) == dao.Less {
			q.results.Data[0] = node.data
			q.results.Down(0, q.Limit)
			this.do_query_desc(node.right, q)
			this.do_query_desc(node.left, q)
		} else {
			this.do_query_desc(node.left, q)
		}
	} else {
		if !q.LeftFilter(node.data) {
			this.do_query_desc(node.right, q)
		} else if !q.RightFilter(node.data) {
			this.do_query_desc(node.left, q)
		}
	}
}
