package algorithm

import (
	"cmp"
	"github.com/lxzan/dao"
	"slices"
	"strconv"
)

// ToString 数字转字符串
func ToString[T dao.Integer](x T) string {
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

func Unique[T cmp.Ordered](arr []T) []T {
	if len(arr) == 0 {
		return arr
	}

	slices.Sort(arr)

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

func UniqueBy[T any, K cmp.Ordered](arr []T, getKey func(item T) K) []T {
	if len(arr) == 0 {
		return arr
	}

	slices.SortFunc(arr, func(a, b T) int {
		x := getKey(a)
		y := getKey(b)
		if x < y {
			return -1
		} else if x > y {
			return 1
		} else {
			return 0
		}
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

// Reverse 反转数组
func Reverse[T any](arr []T) {
	var n = len(arr)
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
	}
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
func Map[A any, B any](arr []A, transfer func(item A) B) []B {
	var results = make([]B, 0, len(arr))
	for i := range arr {
		results = append(results, transfer(arr[i]))
	}
	return results
}

// Filter 过滤器
func Filter[T any](arr []T, check func(item T) bool) []T {
	var results = make([]T, 0, len(arr))
	for i, v := range arr {
		if check(v) {
			results = append(results, arr[i])
		}
	}
	return results
}
