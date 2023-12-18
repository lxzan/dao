package dao

import "cmp"

type CompareResult int8

const (
	Less    CompareResult = -1
	Equal   CompareResult = 0
	Greater CompareResult = 1
)

// LessFunc 比大小
type LessFunc[T any] func(a, b T) bool

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
)
