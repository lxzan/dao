package dict

import (
	"github.com/lxzan/dao/internal/mlist"
	"github.com/lxzan/dao/internal/utils"
	"strings"
)

type element struct {
	EntryPoint mlist.EntryPoint
	Children   []*element
}

type Dict[T any] struct {
	indexes     []uint8                 // 索引
	binaryIndex bool                    // 是否使用2进制索引
	root        *element                // 根节点
	storage     *mlist.MList[string, T] // 存储
}

// New 新建字典树
// 注意: key不能重复
func New[T any]() *Dict[T] {
	return &Dict[T]{
		indexes:     defaultIndexes,
		binaryIndex: true,
		root:        &element{Children: make([]*element, defaultIndexes[0])},
		storage:     mlist.NewMList[string, T](8),
	}
}

// WithIndexes 设置索引
// 索引长度至少为2; 如果每个数字都满足y=pow(2,x), 索引效率更高.
func (c *Dict[T]) WithIndexes(indexes []uint8) *Dict[T] {
	if len(indexes) < 2 {
		panic("indexes length at least 2")
	}
	for _, item := range indexes {
		if !utils.IsBinaryNumber(item) {
			c.binaryIndex = false
			break
		}
	}
	c.indexes = indexes
	c.root.Children = make([]*element, indexes[0])
	return c
}

func (c *Dict[T]) Len() int {
	return c.storage.Len()
}

func (c *Dict[T]) Reset() {
	c.doReset(c.root)
	c.storage.Reset()
}

func (c *Dict[T]) doReset(ele *element) {
	if ele == nil {
		return
	}
	ele.EntryPoint.Head, ele.EntryPoint.Tail = 0, 0
	for _, item := range ele.Children {
		c.doReset(item)
	}
}

// Set 插入或替换元素
func (c *Dict[T]) Set(key string, val T) {
	if key == "" {
		c.storage.Push(&c.root.EntryPoint, key, val)
		return
	}

	for i := c.begin(key, true); i != nil; i = c.next(i) {
		if i.hit() {
			c.storage.Push(&i.Node.EntryPoint, key, val)
			return
		}
	}
}

// Get 根据key查询数据
func (c *Dict[T]) Get(key string) (value T, exist bool) {
	if key == "" {
		return c.storage.Find(&c.root.EntryPoint, key)
	}

	for i := c.begin(key, false); i != nil; i = c.next(i) {
		if i.hit() {
			if i.Node.EntryPoint.Head == 0 {
				return value, false
			}
			value, exist = c.storage.Find(&i.Node.EntryPoint, key)
			break
		}
	}
	return value, exist
}

// Delete 删除元素
func (c *Dict[T]) Delete(key string) {
	if key == "" {
		c.storage.Delete(&c.root.EntryPoint, key)
		return
	}

	for i := c.begin(key, false); i != nil; i = c.next(i) {
		if i.hit() {
			c.storage.Delete(&i.Node.EntryPoint, key)
			return
		}
	}
}

// Match 前缀匹配
func (c *Dict[T]) Match(prefix string, f func(key string, value T) bool) {
	var next = true
	if prefix == "" {
		c.doMatch(c.root, prefix, &next, f)
		return
	}

	for i := c.begin(prefix, true); i != nil; i = c.next(i) {
		if i.hit() {
			c.doMatch(i.Node, prefix, &next, f)
			return
		}
	}
}

func (c *Dict[T]) doMatch(node *element, prefix string, next *bool, f func(key string, value T) bool) {
	if node == nil || !*next {
		return
	}
	c.storage.Range(&node.EntryPoint, func(iter *mlist.Element[string, T]) bool {
		if strings.HasPrefix(iter.Key, prefix) {
			if ok := f(iter.Key, iter.Value); !ok {
				*next = ok
			}
		}
		return *next
	})
	for _, item := range node.Children {
		c.doMatch(item, prefix, next, f)
	}
}

// Range 遍历字典树
func (c *Dict[T]) Range(f func(key string, value T) bool) {
	for _, item := range c.storage.Buckets {
		if item.Addr == mlist.Nil {
			continue
		}
		if !f(item.Key, item.Value) {
			return
		}
	}
}
