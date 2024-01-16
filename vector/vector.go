package vector

import (
	"github.com/lxzan/dao/algo"
	"github.com/lxzan/dao/internal/utils"
	"github.com/lxzan/dao/types/cmp"
)

// New 创建动态数组
func New[T any](capacity int) *Vector[T] {
	c := Vector[T](make([]T, 0, capacity))
	return &c
}

// NewFromDocs 从可变参数创建动态数组
func NewFromDocs[T any](values ...T) *Vector[T] {
	c := Vector[T](values)
	return &c
}

// Vector 动态数组
type Vector[T any] []T

// Reset 重置
func (c *Vector[T]) Reset() {
	*c = (*c)[:0]
}

// Len 获取元素数量
func (c *Vector[T]) Len() int {
	if c == nil {
		return 0
	}
	return len(*c)
}

func (c *Vector[T]) Cap() int {
	if c == nil {
		return 0
	}
	return cap(*c)
}

// Get 根据下标取值
func (c *Vector[T]) Get(index int) T {
	return (*c)[index]
}

// Front 获取头部元素
// 注意: 未作越界检查
func (c *Vector[T]) Front() T {
	return c.Get(0)
}

// Back 获取尾部元素
// 注意: 未作越界检查
func (c *Vector[T]) Back() T {
	return c.Get(c.Len() - 1)
}

// Delete 根据下标删除某个元素
// 性能不好, 少用
func (c *Vector[T]) Delete(index int) {
	var n = c.Len()
	for i := index; i < n-1; i++ {
		(*c)[i] = (*c)[i+1]
	}
	*c = (*c)[:n-1]
}

// Update 根据下标修改值
func (c *Vector[T]) Update(index int, value T) {
	(*c)[index] = value
}

// Elem 取值
func (c *Vector[T]) Elem() []T {
	return *c
}

// MapString 转换为字符串数组
func (c *Vector[T]) MapString(transfer func(i int, v T) string) *Vector[string] {
	v := Vector[string](algo.Map(*c, transfer))
	return &v
}

// MapInt 转换为int数组
func (c *Vector[T]) MapInt(transfer func(i int, v T) int) *Vector[int] {
	v := Vector[int](algo.Map(*c, transfer))
	return &v
}

// MapInt64 转换为int64数组
func (c *Vector[T]) MapInt64(transfer func(i int, v T) int64) *Vector[int64] {
	v := Vector[int64](algo.Map(*c, transfer))
	return &v
}

// UniqueByString 通过string类型字段去重
func (c *Vector[T]) UniqueByString(transfer func(v T) string) *Vector[T] {
	*c = algo.UniqueBy(*c, transfer)
	return c
}

// UniqueByInt 通过int类型字段去重
func (c *Vector[T]) UniqueByInt(transfer func(v T) int) *Vector[T] {
	*c = algo.UniqueBy(*c, transfer)
	return c
}

// UniqueByInt64 通过int64类型字段去重
func (c *Vector[T]) UniqueByInt64(transfer func(v T) int64) *Vector[T] {
	*c = algo.UniqueBy(*c, transfer)
	return c
}

// GroupByString 通过string类型字段分组
func (c *Vector[T]) GroupByString(transfer func(i int, v T) string) map[string][]T {
	return algo.GroupBy(c.Elem(), transfer)
}

// GroupByInt 通过int类型字段分组
func (c *Vector[T]) GroupByInt(transfer func(i int, v T) int) map[int][]T {
	return algo.GroupBy(c.Elem(), transfer)
}

// GroupByInt64 通过int64类型字段分组
func (c *Vector[T]) GroupByInt64(transfer func(i int, v T) int64) map[int64][]T {
	return algo.GroupBy(c.Elem(), transfer)
}

// Filter 过滤
func (c *Vector[T]) Filter(f func(i int, v T) bool) *Vector[T] {
	*c = algo.Filter(*c, f)
	return c
}

// ToStringMap 转换为map
func (c *Vector[T]) ToStringMap(transfer func(v T) string) map[string]T {
	var m = make(map[string]T, c.Len())
	for _, item := range *c {
		key := transfer(item)
		m[key] = item
	}
	return m
}

// ToIntMap 转换为map
func (c *Vector[T]) ToIntMap(transfer func(v T) int) map[int]T {
	var m = make(map[int]T, c.Len())
	for _, item := range *c {
		key := transfer(item)
		m[key] = item
	}
	return m
}

// ToInt64Map 转换为map
func (c *Vector[T]) ToInt64Map(transfer func(v T) int64) map[int64]T {
	var m = make(map[int64]T, c.Len())
	for _, item := range *c {
		key := transfer(item)
		m[key] = item
	}
	return m
}

// PushFront 向头部追加元素
// 性能不好, 少用
func (c *Vector[T]) PushFront(v T) {
	var d = New[T](c.Len() + 1)
	d.PushBack(v)
	*d = append(*d, *c...)
	*c = *d
}

// PushBack 向尾部追加元素
func (c *Vector[T]) PushBack(v T) {
	*c = append(*c, v)
}

// PopFront 从头部弹出元素
func (c *Vector[T]) PopFront() (value T) {
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
func (c *Vector[T]) PopBack() (value T) {
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

func (c *Vector[T]) SortByString(transfer func(v T) string) *Vector[T] {
	algo.SortBy(c.Elem(), func(a, b T) int {
		return cmp.Compare(transfer(a), transfer(b))
	})
	return c
}

func (c *Vector[T]) SortByInt(transfer func(v T) int) *Vector[T] {
	algo.SortBy(c.Elem(), func(a, b T) int {
		return cmp.Compare(transfer(a), transfer(b))
	})
	return c
}

func (c *Vector[T]) SortByInt64(transfer func(v T) int64) *Vector[T] {
	algo.SortBy(c.Elem(), func(a, b T) int {
		return cmp.Compare(transfer(a), transfer(b))
	})
	return c
}

// Range 遍历
func (c *Vector[T]) Range(f func(i int, v T) bool) {
	for index, value := range *c {
		if !f(index, value) {
			return
		}
	}
}

// Clone 拷贝
func (c *Vector[T]) Clone() *Vector[T] {
	var d = utils.Clone(*c)
	return &d
}

// Slice 截取子数组
func (c *Vector[T]) Slice(start, end int) *Vector[T] {
	var children = (*c)[start:end]
	return &children
}

func (c *Vector[T]) Reverse() *Vector[T] {
	*c = algo.Reverse(*c)
	return c
}
