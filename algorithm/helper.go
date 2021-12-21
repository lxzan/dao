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

func Unique[T any, K dao.Hashable[K]](arr *[]T, getKey func(x T) K) {
	var n = len(*arr)
	var m = make(map[K]T, n)
	for i, _ := range *arr {
		var key = getKey((*arr)[i])
		m[key] = (*arr)[i]
	}

	var i = 0
	for _, v := range m {
		(*arr)[i] = v
		i++
	}
	*arr = (*arr)[:i]
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

func GetKeys[T any, K any](arr []T, getKey func(x T) K) []K {
	var results = make([]K, 0, len(arr))
	for i, _ := range arr {
		results = append(results, getKey(arr[i]))
	}
	return results
}
