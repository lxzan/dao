package rbtree

import (
	"cmp"
	"github.com/lxzan/dao"
	"github.com/lxzan/dao/algorithm"
	"github.com/lxzan/dao/stack"
)

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

// GetMinKey 获取最小的key, 过滤条件可为空
func (c *RBTree[K, V]) GetMinKey(filter func(key K) bool) (result Pair[K, V], exist bool) {
	return c.doGetMinKey(stack.New[*rbtree_node[K, V]](10), filter)
}

func (c *RBTree[K, V]) doGetMinKey(s *stack.Stack[*rbtree_node[K, V]], filter func(key K) bool) (result Pair[K, V], exist bool) {
	s.Reset()
	filter = algorithm.SelectValue(filter == nil, TrueFunc[K], filter)
	s.Push(c.root)
	for s.Len() > 0 {
		var node = s.Pop()
		if c.end(node) {
			continue
		}
		if filter(node.data.Key) {
			if !exist || node.data.Key < result.Key {
				exist = true
				result = *node.data
			}
			s.Push(node.left)
		} else {
			s.Push(node.right)
		}
	}
	return result, exist
}

// GetMaxKey 获取最大的key, 过滤条件可为空
func (c *RBTree[K, V]) GetMaxKey(filter func(key K) bool) (result Pair[K, V], exist bool) {
	return c.doGetMaxKey(stack.New[*rbtree_node[K, V]](10), filter)
}

func (c *RBTree[K, V]) doGetMaxKey(s *stack.Stack[*rbtree_node[K, V]], filter func(key K) bool) (result Pair[K, V], exist bool) {
	s.Reset()
	filter = algorithm.SelectValue(filter == nil, TrueFunc[K], filter)
	s.Push(c.root)
	for s.Len() > 0 {
		var node = s.Pop()
		if c.end(node) {
			continue
		}
		if filter(node.data.Key) {
			if !exist || node.data.Key > result.Key {
				exist = true
				result = *node.data
			}
			s.Push(node.right)
		} else {
			s.Push(node.left)
		}
	}
	return result, exist
}

func TrueFunc[K cmp.Ordered](d K) bool {
	return true
}

type QueryBuilder[K cmp.Ordered, V any] struct {
	tree        *RBTree[K, V]
	leftFilter  func(key K) bool
	rightFilter func(key K) bool
	limit       int
	order       dao.Order
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
func (c *QueryBuilder[K, V]) Order(o dao.Order) *QueryBuilder[K, V] {
	c.order = o
	return c
}

// Limit 设置结果数量限制, 默认10条
func (c *QueryBuilder[K, V]) Limit(n int) *QueryBuilder[K, V] {
	c.limit = n
	return c
}

// Do 执行查询
func (c *QueryBuilder[K, V]) Do() []Pair[K, V] {
	return c.tree.do_query(c)
}

// NewQuery 新建一个查询
func (c *RBTree[K, V]) NewQuery() *QueryBuilder[K, V] {
	return &QueryBuilder[K, V]{tree: c}
}

func (c *RBTree[K, V]) do_query(q *QueryBuilder[K, V]) []Pair[K, V] {
	q.init()
	var results = make([]Pair[K, V], 0, q.limit)
	var s = stack.New[*rbtree_node[K, V]](uint32(q.limit))

	if q.order == dao.DESC {
		maxEle, exist := c.doGetMaxKey(s, q.rightFilter)
		if exist && q.leftFilter(maxEle.Key) {
			results = append(results, maxEle)
		} else {
			return results
		}

		for i := 0; i < q.limit-1; i++ {
			result, exist := c.doGetMaxKey(s, func(key K) bool {
				return key < maxEle.Key
			})
			if exist && q.leftFilter(result.Key) {
				results = append(results, result)
				maxEle = result
			} else {
				break
			}
		}
	} else {
		minEle, exist := c.doGetMinKey(s, q.leftFilter)
		if exist && q.rightFilter(minEle.Key) {
			results = append(results, minEle)
		} else {
			return results
		}

		for i := 0; i < q.limit-1; i++ {
			result, exist := c.doGetMinKey(s, func(key K) bool {
				return key > minEle.Key
			})
			if exist && q.rightFilter(result.Key) {
				results = append(results, result)
				minEle = result
			} else {
				break
			}
		}
	}
	return results
}
