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

// Delete 根据下标删除某个元素
// 性能不好, 少用
func (c *Vector[K, V]) Delete(index int) {
	var n = c.Len()
	for i := index; i < n-1; i++ {
		(*c)[i] = (*c)[i+1]
	}
	*c = (*c)[:n-1]
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

// MapString 转换为字符串数组
func (c *Vector[K, V]) MapString(transfer func(i int, v V) string) *Vector[string, String] {
	return NewFromStrings(algorithm.Map(*c, transfer)...)
}

// MapInt 转换为int数组
func (c *Vector[K, V]) MapInt(transfer func(i int, v V) int) *Vector[int, Int] {
	return NewFromInts(algorithm.Map(*c, transfer)...)
}

// MapInt64 转换为int64数组
func (c *Vector[K, V]) MapInt64(transfer func(i int, v V) int64) *Vector[int64, Int64] {
	return NewFromInt64s(algorithm.Map(*c, transfer)...)
}

// Unique 排序并根据id去重
func (c *Vector[K, V]) Unique() *Vector[K, V] {
	*c = algorithm.UniqueBy(*c, func(item V) K { return item.GetID() })
	return c
}

func (c *Vector[K, V]) UniqueByString(transfer func(v V) string) *Vector[K, V] {
	*c = algorithm.UniqueBy(*c, transfer)
	return c
}

func (c *Vector[K, V]) UniqueByInt(transfer func(v V) int) *Vector[K, V] {
	*c = algorithm.UniqueBy(*c, transfer)
	return c
}

func (c *Vector[K, V]) UniqueByInt64(transfer func(v V) int64) *Vector[K, V] {
	*c = algorithm.UniqueBy(*c, transfer)
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

// ToMap 生成hashmap.HashMap[K, V]
func (c *Vector[K, V]) ToMap() hashmap.HashMap[K, V] {
	var m = hashmap.New[K, V](c.Len())
	for _, item := range *c {
		m.Set(item.GetID(), item)
	}
	return m
}

// PushFront 向头部追加元素
// 性能不好, 少用
func (c *Vector[K, V]) PushFront(v V) {
	var d = New[K, V](c.Len() + 1)
	d.PushBack(v)
	*d = append(*d, *c...)
	*c = *d
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
