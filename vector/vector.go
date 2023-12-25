package vector

import (
	"github.com/lxzan/dao/algorithm"
	"github.com/lxzan/dao/hashmap"
	"github.com/lxzan/dao/internal/utils"
	"github.com/lxzan/dao/types/cmp"
	"unsafe"
)

// New 创建动态数组
func New[K cmp.Ordered, V Document[K]](capacity int) *Vector[K, V] {
	c := Vector[K, V](make([]V, 0, capacity))
	return &c
}

// NewFromDocs 从可变参数创建动态数组
func NewFromDocs[K cmp.Ordered, V Document[K]](values ...V) *Vector[K, V] {
	c := Vector[K, V](values)
	return &c
}

// NewFromInts 创建动态数组
func NewFromInts(values ...int) *Vector[int, Int] {
	var b = *(*[]Int)(unsafe.Pointer(&values))
	v := Vector[int, Int](b)
	return &v
}

// NewFromInt64s 创建动态数组
func NewFromInt64s(values ...int64) *Vector[int64, Int64] {
	var b = *(*[]Int64)(unsafe.Pointer(&values))
	v := Vector[int64, Int64](b)
	return &v
}

// NewFromStrings 创建动态数组
func NewFromStrings(values ...string) *Vector[string, String] {
	var b = *(*[]String)(unsafe.Pointer(&values))
	v := Vector[string, String](b)
	return &v
}

// Vector 动态数组
type Vector[K cmp.Ordered, V Document[K]] []V

// Reset 重置
func (c *Vector[K, V]) Reset() {
	*c = (*c)[:0]
}

// Len 获取元素数量
func (c *Vector[K, V]) Len() int {
	return len(*c)
}

// Get 根据下标取值
func (c *Vector[K, V]) Get(index int) V {
	return (*c)[index]
}

// Update 根据下标修改值
func (c *Vector[K, V]) Update(index int, value V) {
	(*c)[index] = value
}

// Elem 取值
func (c *Vector[K, V]) Elem() []V {
	return *c
}

// Exists 根据id判断某条数据是否存在
func (c *Vector[K, V]) Exists(id K) (v V, exist bool) {
	for _, item := range *c {
		if item.GetID() == id {
			return item, true
		}
	}
	return v, exist
}

// Unique 排序并根据id去重
func (c *Vector[K, V]) Unique() *Vector[K, V] {
	*c = algorithm.UniqueBy(*c, func(item V) K {
		return item.GetID()
	})
	return c
}

// Filter 过滤
func (c *Vector[K, V]) Filter(f func(i int, v V) bool) *Vector[K, V] {
	*c = algorithm.Filter(*c, f)
	return c
}

// Sort 排序
func (c *Vector[K, V]) Sort() *Vector[K, V] {
	algorithm.SortBy(*c, func(a, b V) int {
		return cmp.Compare(a.GetID(), b.GetID())
	})
	return c
}

// GetIdList 获取id数组
func (c *Vector[K, V]) GetIdList() []K {
	var v V
	switch any(v).(type) {
	case Int, Int64, String:
		var keys = *(*[]K)(unsafe.Pointer(c))
		return keys
	default:
		var keys = make([]K, 0, c.Len())
		for _, item := range *c {
			keys = append(keys, item.GetID())
		}
		return keys
	}
}

// ToMap 生成map[K]D
func (c *Vector[K, V]) ToMap() hashmap.HashMap[K, V] {
	var m = hashmap.New[K, V](c.Len())
	for _, item := range *c {
		m.Set(item.GetID(), item)
	}
	return m
}

// PushBack 向尾部追加元素
func (c *Vector[K, V]) PushBack(v V) {
	*c = append(*c, v)
}

// PopFront 从头部弹出元素
func (c *Vector[K, V]) PopFront() (value V) {
	switch c.Len() {
	case 0:
		return value
	default:
		value = (*c)[0]
		*c = (*c)[1:]
		return value
	}
}

// PopBack 从尾部弹出元素
func (c *Vector[K, V]) PopBack() (value V) {
	n := c.Len()
	switch n {
	case 0:
		return value
	default:
		value = (*c)[n-1]
		*c = (*c)[:n-1]
		return value
	}
}

// Range 遍历
func (c *Vector[K, V]) Range(f func(i int, v V) bool) {
	for index, value := range *c {
		if !f(index, value) {
			return
		}
	}
}

// Clone 拷贝
func (c *Vector[K, V]) Clone() *Vector[K, V] {
	var d = utils.Clone(*c)
	return &d
}

// Slice 截取子数组
func (c *Vector[K, V]) Slice(start, end int) *Vector[K, V] {
	var children = (*c)[start:end]
	return &children
}

func (c *Vector[K, V]) Reverse() *Vector[K, V] {
	*c = algorithm.Reverse(*c)
	return c
}
