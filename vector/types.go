package vector

import (
	"github.com/lxzan/dao/types/cmp"
	"unsafe"
)

type (
	Document[T cmp.Ordered] interface {
		GetID() T
	}

	Int int

	Int64 int64

	String string
)

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

func (c Int) GetID() int {
	return int(c)
}

func (c Int64) GetID() int64 {
	return int64(c)
}

func (c String) GetID() string { return string(c) }
