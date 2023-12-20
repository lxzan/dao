package rbtree

import (
	"github.com/lxzan/dao/algorithm"
	"github.com/lxzan/dao/stack"
	"github.com/lxzan/dao/types/cmp"
)

type Order uint8

const (
	ASC  Order = 0 // 升序
	DESC Order = 1 // 降序
)

func TrueFunc[K cmp.Ordered](d K) bool { return true }

// NewQuery 新建一个查询
func (c *RBTree[K, V]) NewQuery() *QueryBuilder[K, V] {
	return &QueryBuilder[K, V]{tree: c}
}

type QueryBuilder[K cmp.Ordered, V any] struct {
	tree    *RBTree[K, V] // 红黑树
	results []Pair[K, V]  // 查询结果

	leftFilter  func(key K) bool // 左边界条件
	rightFilter func(key K) bool // 右边界条件
	limit       int              // 单页限制
	total       int              // 总条数
	offset      int              // 偏移量
	order       Order            // 排序
}

func (c *QueryBuilder[K, V]) init() *QueryBuilder[K, V] {
	if c.leftFilter == nil {
		c.leftFilter = TrueFunc[K]
	}
	if c.rightFilter == nil {
		c.rightFilter = TrueFunc[K]
	}
	if c.limit <= 0 {
		c.limit = 10
	}
	c.total = c.offset + c.limit
	return c
}

// Left 设置左边界过滤条件
func (c *QueryBuilder[K, V]) Left(f func(key K) bool) *QueryBuilder[K, V] {
	c.leftFilter = f
	return c
}

// Right 设置右边界过滤条件
func (c *QueryBuilder[K, V]) Right(f func(key K) bool) *QueryBuilder[K, V] {
	c.rightFilter = f
	return c
}

// Order 设置排序, 默认ASC
func (c *QueryBuilder[K, V]) Order(o Order) *QueryBuilder[K, V] {
	c.order = o
	return c
}

// Limit 设置结果数量限制, 默认10条
func (c *QueryBuilder[K, V]) Limit(n int) *QueryBuilder[K, V] {
	c.limit = n
	return c
}

func (c *QueryBuilder[K, V]) Offset(n int) *QueryBuilder[K, V] {
	c.offset = n
	return c
}

// FindAll 执行查询
func (c *QueryBuilder[K, V]) FindAll() []Pair[K, V] {
	c.init()
	c.results = make([]Pair[K, V], 0, c.total)

	switch c.order {
	case DESC:
		c.rangeDesc(c.tree.root)
	case ASC:
		c.rangeAsc(c.tree.root)
	}

	if c.offset > 0 {
		if len(c.results) > c.offset {
			c.results = c.results[c.offset:]
		} else {
			c.results = c.results[:0]
		}
	}
	return c.results
}

// 降序遍历 中序遍历是有序的
func (c *QueryBuilder[K, V]) rangeDesc(node *rbtree_node[K, V]) {
	if c.tree.end(node) || len(c.results) >= c.total {
		return
	}

	state := 0
	if c.rightFilter(node.data.Key) {
		state += 1
	}
	if c.leftFilter(node.data.Key) {
		state += 2
	}

	switch state {
	case 3:
		c.rangeDesc(node.right)
		if len(c.results) < c.total {
			c.results = append(c.results, *node.data)
		} else {
			return
		}
		c.rangeDesc(node.left)
	case 2:
		c.rangeDesc(node.left)
	case 1:
		c.rangeDesc(node.right)
	}
}

// 升序遍历 中序遍历是有序的
func (c *QueryBuilder[K, V]) rangeAsc(node *rbtree_node[K, V]) {
	if c.tree.end(node) || len(c.results) >= c.total {
		return
	}

	state := 0
	if c.rightFilter(node.data.Key) {
		state += 1
	}
	if c.leftFilter(node.data.Key) {
		state += 2
	}

	switch state {
	case 3:
		c.rangeAsc(node.left)
		if len(c.results) < c.total {
			c.results = append(c.results, *node.data)
		} else {
			return
		}
		c.rangeAsc(node.right)
	case 2:
		c.rangeAsc(node.left)
	case 1:
		c.rangeAsc(node.right)
	}
}

func (c *QueryBuilder[K, V]) FindOne() (p Pair[K, V], exist bool) {
	c.init()

	switch c.order {
	case DESC:
		if v, ok := c.getMaxPair(c.rightFilter); ok {
			if c.leftFilter(v.Key) {
				return *v, true
			}
		}
	case ASC:
		if v, ok := c.getMinPair(c.leftFilter); ok {
			if c.rightFilter(v.Key) {
				return *v, true
			}
		}
	}

	return p, exist
}

func (c *QueryBuilder[K, V]) getMaxPair(filter func(key K) bool) (result *Pair[K, V], exist bool) {
	var s = stack.Stack[*rbtree_node[K, V]]{}
	s.Push(c.tree.root)
	for s.Len() > 0 {
		var node = s.Pop()
		if c.tree.end(node) {
			continue
		}
		if filter(node.data.Key) {
			if !exist || node.data.Key > result.Key {
				exist = true
				result = node.data
			}
			s.Push(node.right)
		} else {
			s.Push(node.left)
		}
	}
	return result, exist
}

func (c *QueryBuilder[K, V]) getMinPair(filter func(key K) bool) (result *Pair[K, V], exist bool) {
	var s = stack.Stack[*rbtree_node[K, V]]{}
	filter = algorithm.SelectValue(filter == nil, TrueFunc[K], filter)
	s.Push(c.tree.root)
	for s.Len() > 0 {
		var node = s.Pop()
		if c.tree.end(node) {
			continue
		}
		if filter(node.data.Key) {
			if !exist || node.data.Key < result.Key {
				exist = true
				result = node.data
			}
			s.Push(node.left)
		} else {
			s.Push(node.right)
		}
	}
	return result, exist
}

// Get 查询一个key
func (c *RBTree[K, V]) Get(key K) (result V, exist bool) {
	for i := c.begin(); !c.end(i); i = c.next(i, key) {
		if key == i.data.Key {
			return i.data.Val, true
		}
	}
	return result, false
}

// Range 遍历树
func (c *RBTree[K, V]) Range(fn func(key K, value V) bool) {
	var next = true
	c.do_range(c.root, &next, fn)
}

func (c *RBTree[K, V]) do_range(node *rbtree_node[K, V], next *bool, fn func(K, V) bool) {
	if c.end(node) || !*next {
		return
	}

	if ok := fn(node.data.Key, node.data.Val); !ok {
		*next = ok
	}

	c.do_range(node.left, next, fn)
	c.do_range(node.right, next, fn)
}
