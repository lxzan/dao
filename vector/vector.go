package vector

import (
	"github.com/lxzan/dao/algorithm"
	"github.com/lxzan/dao/hashmap"
	"github.com/lxzan/dao/internal/utils"
	"github.com/lxzan/dao/types/cmp"
	"unsafe"
)

// New 创建动态数组
func New[D Document[K], K cmp.Ordered](capacity int) *Vector[D, K] {
	c := Vector[D, K](make([]D, 0, capacity))
	return &c
}

// NewFromDocs 从可变参数创建动态数组
func NewFromDocs[D Document[K], K cmp.Ordered](values ...D) *Vector[D, K] {
	c := Vector[D, K](values)
	return &c
}

// NewFromInts 创建动态数组
func NewFromInts(values ...int) *Vector[Int, int] {
	var b = *(*[]Int)(unsafe.Pointer(&values))
	v := Vector[Int, int](b)
	return &v
}

// NewFromInt64s 创建动态数组
func NewFromInt64s(values ...int64) *Vector[Int64, int64] {
	var b = *(*[]Int64)(unsafe.Pointer(&values))
	v := Vector[Int64, int64](b)
	return &v
}

// NewFromStrings 创建动态数组
func NewFromStrings(values ...string) *Vector[String, string] {
	var b = *(*[]String)(unsafe.Pointer(&values))
	v := Vector[String, string](b)
	return &v
}

// Vector 动态数组
type Vector[D Document[K], K cmp.Ordered] []D

// Reset 重置
func (c *Vector[D, K]) Reset() {
	*c = (*c)[:0]
}

// Len 获取元素数量
func (c *Vector[D, K]) Len() int {
	return len(*c)
}

// Get 根据下标取值
func (c *Vector[D, K]) Get(index int) D {
	return (*c)[index]
}

// Update 根据下标修改值
func (c *Vector[D, K]) Update(index int, value D) {
	(*c)[index] = value
}

// Elem 取值
func (c *Vector[D, K]) Elem() []D {
	return *c
}

// Exists 根据id判断某条数据是否存在
func (c *Vector[D, K]) Exists(id K) (v D, exist bool) {
	for _, item := range *c {
		if item.GetID() == id {
			return item, true
		}
	}
	return v, exist
}

// Unique 排序并根据id去重
func (c *Vector[D, K]) Unique() *Vector[D, K] {
	*c = algorithm.UniqueBy(*c, func(item D) K {
		return item.GetID()
	})
	return c
}

// Filter 过滤
func (c *Vector[D, K]) Filter(f func(i int, v D) bool) *Vector[D, K] {
	*c = algorithm.Filter(*c, f)
	return c
}

// Sort 排序
func (c *Vector[D, K]) Sort() *Vector[D, K] {
	algorithm.SortBy(*c, func(a, b D) int {
		return cmp.Compare(a.GetID(), b.GetID())
	})
	return c
}

// IdList 获取id数组
func (c *Vector[D, K]) IdList() []K {
	var d D
	switch any(d).(type) {
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
func (c *Vector[D, K]) ToMap() hashmap.HashMap[K, D] {
	var m = hashmap.New[K, D](c.Len())
	for _, item := range *c {
		m.Set(item.GetID(), item)
	}
	return m
}

// PushBack 向尾部追加元素
func (c *Vector[D, K]) PushBack(v D) {
	*c = append(*c, v)
}

// PopFront 从头部弹出元素
func (c *Vector[D, K]) PopFront() (value D) {
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
func (c *Vector[D, K]) PopBack() (value D) {
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
func (c *Vector[D, K]) Range(f func(i int, v D) bool) {
	for index, value := range *c {
		if !f(index, value) {
			return
		}
	}
}

// Clone 拷贝
func (c *Vector[D, K]) Clone() *Vector[D, K] {
	var d = utils.Clone(*c)
	return &d
}

// Slice 截取子数组
func (c *Vector[D, K]) Slice(start, end int) *Vector[D, K] {
	var children = (*c)[start:end]
	return &children
}

func (c *Vector[D, K]) Reverse() *Vector[D, K] {
	*c = algorithm.Reverse(*c)
	return c
}
