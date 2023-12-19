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
func (c *RBTree[K, V]) GetMinKey(filter func(key K) bool) (pair Pair[K, V], exist bool) {
	if v, ok := c.getMinNode(stack.New[*rbtree_node[K, V]](10), filter); ok {
		return *v.data, true

	}
	return pair, false
}

func (c *RBTree[K, V]) getMinNode(s *stack.Stack[*rbtree_node[K, V]], filter func(key K) bool) (result *rbtree_node[K, V], exist bool) {
	s.Reset()
	filter = algorithm.SelectValue(filter == nil, TrueFunc[K], filter)
	s.Push(c.root)
	for s.Len() > 0 {
		var node = s.Pop()
		if c.end(node) {
			continue
		}
		if filter(node.data.Key) {
			if !exist || node.data.Key < result.data.Key {
				exist = true
				result = node
			}
			s.Push(node.left)
		} else {
			s.Push(node.right)
		}
	}
	return result, exist
}

// GetMaxKey 获取最大的key, 过滤条件可为空
func (c *RBTree[K, V]) GetMaxKey(filter func(key K) bool) (pair Pair[K, V], exist bool) {
	if v, ok := c.getMaxNode(stack.New[*rbtree_node[K, V]](10), filter); ok {
		return *v.data, true
	}
	return pair, false
}

func (c *RBTree[K, V]) getMaxNode(s *stack.Stack[*rbtree_node[K, V]], filter func(key K) bool) (result *rbtree_node[K, V], exist bool) {
	s.Reset()
	filter = algorithm.SelectValue(filter == nil, TrueFunc[K], filter)
	s.Push(c.root)
	for s.Len() > 0 {
		var node = s.Pop()
		if c.end(node) {
			continue
		}
		if filter(node.data.Key) {
			if !exist || node.data.Key > result.data.Key {
				exist = true
				result = node
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
	tree    *RBTree[K, V] // 红黑树
	results []Pair[K, V]  // 查询结果
	history map[K]uint8   // 历史记录

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
	c.results = make([]Pair[K, V], 0, c.limit)
	c.history = make(map[K]uint8, c.limit)
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
	c.tree.do_query(c)
	return c.results
}

// NewQuery 新建一个查询
func (c *RBTree[K, V]) NewQuery() *QueryBuilder[K, V] {
	return &QueryBuilder[K, V]{tree: c}
}

func (c *RBTree[K, V]) do_query(qb *QueryBuilder[K, V]) {
	qb.init()
	var s = stack.New[*rbtree_node[K, V]](uint32(qb.limit))

	if qb.order == dao.DESC {
		var filter = qb.rightFilter
		for {
			maxEle, exist := c.getMaxNode(s, filter)
			if !exist {
				break
			}

			var count = 0
			c.rangeInOrder(maxEle, qb, func(node *rbtree_node[K, V]) bool {
				if len(qb.results) < qb.limit {
					qb.results = append(qb.results, *node.data)
					maxEle = node
					count++
					return true
				}
				return false
			})

			if count == 0 {
				break
			}

			filter = func(key K) bool { return key < maxEle.data.Key }
		}
		return
	}

	var filter = qb.leftFilter
	for {
		minEle, exist := c.getMinNode(s, filter)
		if !exist {
			break
		}

		var count = 0
		c.rangeInOrder(minEle, qb, func(node *rbtree_node[K, V]) bool {
			if len(qb.results) < qb.limit {
				qb.results = append(qb.results, *node.data)
				minEle = node
				count++
				return true
			}
			return false
		})

		if count == 0 {
			break
		}

		filter = func(key K) bool { return key > minEle.data.Key }
	}
}

// 有序遍历 中序遍历是有序的
func (c *RBTree[K, V]) rangeInOrder(curNode *rbtree_node[K, V], qb *QueryBuilder[K, V], visit func(*rbtree_node[K, V]) bool) {
	if curNode == nil || c.end(curNode) || qb.history[curNode.data.Key] == 1 || len(qb.results) >= qb.limit {
		return
	}

	qb.history[curNode.data.Key] = 1
	state := 0
	if qb.rightFilter(curNode.data.Key) {
		state += 1
	}
	if qb.leftFilter(curNode.data.Key) {
		state += 2
	}

	if qb.order == dao.DESC {
		switch state {
		case 3:
			c.rangeInOrder(curNode.right, qb, visit)
			if !visit(curNode) {
				return
			}
			c.rangeInOrder(curNode.left, qb, visit)
		case 2:
			c.rangeInOrder(curNode.left, qb, visit)
		case 1:
			c.rangeInOrder(curNode.right, qb, visit)
		default:
		}
		return
	}

	switch state {
	case 3:
		c.rangeInOrder(curNode.left, qb, visit)
		if !visit(curNode) {
			return
		}
		c.rangeInOrder(curNode.right, qb, visit)
	case 2:
		c.rangeInOrder(curNode.right, qb, visit)
	case 1:
		c.rangeInOrder(curNode.left, qb, visit)
	default:
	}
}
