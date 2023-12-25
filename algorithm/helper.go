package algorithm

import (
	"github.com/lxzan/dao/types/cmp"
	"strconv"
)

// ToString 数字转字符串
func ToString[T cmp.Integer](x T) string {
	return strconv.Itoa(int(x))
}

// Max 获取最大值
func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Min 获取最小值
func Min[T cmp.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Swap 交换数据
func Swap[T any](a, b *T) {
	temp := *a
	*a = *b
	*b = temp
}

func Unique[T cmp.Ordered, A ~[]T](arr A) A {
	if len(arr) == 0 {
		return arr
	}

	Sort(arr)

	var n = len(arr)
	var j = 1
	for i := 1; i < n; i++ {
		if arr[i] != arr[i-1] {
			arr[j] = arr[i]
			j++
		}
	}
	arr = arr[:j]
	return arr
}

func UniqueBy[T any, K cmp.Ordered, A ~[]T](arr A, getKey func(item T) K) A {
	if len(arr) == 0 {
		return arr
	}

	SortBy(arr, func(a, b T) int {
		return cmp.Compare(getKey(a), getKey(b))
	})

	var n = len(arr)
	var j = 1
	for i := 1; i < n; i++ {
		if getKey(arr[i]) != getKey(arr[i-1]) {
			arr[j] = arr[i]
			j++
		}
	}
	arr = arr[:j]
	return arr
}

// Reduce 对数组中的每个元素按序执行一个提供的 reducer 函数，每一次运行 reducer 会将先前元素的计算结果作为参数传入，
// 最后将其结果汇总为单个返回值。
func Reduce[T any, S any](initialValue S, arr []T, reducer func(s S, item T) S) S {
	for _, item := range arr {
		initialValue = reducer(initialValue, item)
	}
	return initialValue
}

// Reverse 反转数组
func Reverse[T any, A ~[]T](arr A) A {
	var n = len(arr)
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
	}
	return arr
}

// SelectValue 选择一个值 三元操作符替代品
func SelectValue[T any](flag bool, a T, b T) T {
	if flag {
		return a
	}
	return b
}

// Contains 是否包含
func Contains[T comparable](arr []T, target T) bool {
	for i := range arr {
		if arr[i] == target {
			return true
		}
	}
	return false
}

// Map 转换器 将A数组转换为B数组
func Map[A any, B any](arr []A, transfer func(i int, v A) B) []B {
	var results = make([]B, 0, len(arr))
	for index, value := range arr {
		results = append(results, transfer(index, value))
	}
	return results
}

// Filter 过滤器
func Filter[T any, A ~[]T](arr A, check func(i int, v T) bool) A {
	var results = make([]T, 0, len(arr))
	for i, v := range arr {
		if check(i, v) {
			results = append(results, arr[i])
		}
	}
	return results
}

// IsZero 零值判断
func IsZero[T comparable](v T) bool {
	var zero T
	return v == zero
}
