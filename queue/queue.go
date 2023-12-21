package queue

import "github.com/lxzan/dao/internal/utils"

type Queue[T any] struct {
	offset int
	tpl    T
	data   []T
}

// New 创建队列
func New[T any](capacity int) *Queue[T] {
	return &Queue[T]{data: make([]T, 0, capacity)}
}

// NewFrom 从切片创建队列
func NewFrom[T any](values ...T) *Queue[T] {
	return &Queue[T]{data: values}
}

// Reset 重置
func (c *Queue[T]) Reset() {
	c.offset = 0
	c.data = c.data[:0]
}

// Len 获取队列长度
func (c *Queue[T]) Len() int {
	return len(c.data) - c.offset
}

// Push 追加元素到队列尾部
func (c *Queue[T]) Push(v T) {
	c.data = append(c.data, v)
}

// Pop 从队列头部弹出元素
func (c *Queue[T]) Pop() (value T) {
	if n := c.Len(); n > 0 {
		value = c.data[c.offset]
		c.data[c.offset] = c.tpl
		c.offset++
		if c.offset == len(c.data) {
			c.Reset()
		}
	}
	return value
}

// Range 遍历
func (c *Queue[T]) Range(f func(value T) bool) {
	for _, item := range c.data {
		if !f(item) {
			return
		}
	}
}

// UnWrap 解包, 返回底层数组
func (c *Queue[T]) UnWrap() []T {
	return c.data[c.offset:]
}

// Clone 拷贝副本
func (c *Queue[T]) Clone() *Queue[T] {
	return &Queue[T]{data: utils.Clone(c.data)}
}
