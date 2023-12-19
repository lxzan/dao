package dao

import "cmp"

const (
	Less    = -1
	Equal   = 0
	Greater = 1
)

type (
	// LessFunc 比大小
	LessFunc[T any] func(a, b T) bool

	// CompareFunc 比较函数
	// a>b, 返回1; a<b, 返回-1; a==b, 返回0
	CompareFunc[T any] func(a, b T) int
)

// AscFunc 升序函数
func AscFunc[T cmp.Ordered](a, b T) bool { return a < b }

// DescFunc 降序函数
func DescFunc[T cmp.Ordered](a, b T) bool { return a > b }

type Order uint8

const (
	ASC  Order = 0 // 升序
	DESC Order = 1 // 降序
)

type (
	Number interface {
		Integer | ~float32 | ~float64
	}

	Integer interface {
		~int64 | ~int | ~int32 | ~int16 | ~int8 | ~uint64 | ~uint | ~uint32 | ~uint16 | ~uint8
	}

	// Map 键不可重复
	Map[K comparable, V any] interface {
		Len() int
		Get(key K) (V, bool)
		Set(key K, value V)
		Delete(key K)
		Range(f func(K, V) bool)
	}

	Resetter interface {
		Reset()
	}

	Cloner[T any] interface {
		// Clone 深拷贝
		Clone() T
	}
)
