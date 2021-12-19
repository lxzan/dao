package algorithm

import (
	"github.com/lxzan/dao"
	"strconv"
)

func ToString[T dao.Integer[T]](x T) string {
	return strconv.Itoa(int(x))
}

func Max[T dao.Comparable[T]](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T dao.Comparable[T]](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Unique[T any, K dao.Hashable[K]](arr []T, fn func(x T) K) []T {
	var n = len(arr)
	var results = make([]T, 0, n)
	var m = make(map[K]T, n)
	for i, _ := range arr {
		var key = fn(arr[i])
		m[key] = arr[i]
	}
	for _, v := range m {
		results = append(results, v)
	}
	return results
}

func Fill[T any](arr []T, v T) {
	for i, _ := range arr {
		arr[i] = v
	}
}

func Reverse[T any](arr []T) {
	var n = len(arr)
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
	}
}

func GetValue[T any](flag bool, a T, b T) T {
	if flag {
		return a
	}
	return b
}
