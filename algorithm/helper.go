package algorithm

import (
	"github.com/lxzan/dao"
	"github.com/lxzan/dao/types"
	"strconv"
)

func ForEach[I any](c types.Iterable[I], fn func(iter I)) {
	for i := c.Begin(); !c.End(i); i = c.Next(i) {
		fn(i)
	}
}

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

func Swap[T any](a, b *T) {
	*a, *b = *b, *a
}

func Unique[T any, K comparable](arr *[]T, getKey func(x T) K) {
	var n = len(*arr)
	var m = make(map[K]T, n)
	for i := range *arr {
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
	for i := range arr {
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

func GetFields[T any, K any](arr []T, get_field func(x T) K) []K {
	var results = make([]K, 0, len(arr))
	for i := range arr {
		results = append(results, get_field(arr[i]))
	}
	return results
}

func Contains[T dao.Comparable[T]](arr []T, target T) bool {
	for i := range arr {
		if arr[i] == target {
			return true
		}
	}
	return false
}
