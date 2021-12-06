package dao

import (
	"github.com/lxzan/dao/internal"
	"github.com/lxzan/dao/types"
	"strconv"
)

var QuickSort = internal.QuickSort

func incr_order[T comparable[T]](a, b T) types.Ordering {
	if a > b {
		return types.Greater
	} else if a == b {
		return types.Equal
	} else {
		return types.Less
	}
}

func decr_order[T comparable[T]](a, b T) types.Ordering {
	if a > b {
		return types.Less
	} else if a == b {
		return types.Equal
	} else {
		return types.Greater
	}
}

func ToString[T inteage[T]](x T) string {
	return strconv.Itoa(int(x))
}

func LowerBound[T comparer[T]](arr []T, target T) int {
	var n = len(arr)
	if n == 0 {
		return -1
	}
	var left = 0
	var right = n - 1
	for right-left > 1 {
		var mid = (left + right) / 2
		if arr[mid] >= target {
			right = mid
		} else {
			left = mid
		}
	}

	if arr[left] == target {
		return left
	} else if arr[right] == target {
		return right
	}
	return -1
}

func UpperBound[T comparer[T]](arr []T, target T) int {
	var n = len(arr)
	if n == 0 {
		return -1
	}
	var left = 0
	var right = n - 1
	for right-left > 1 {
		var mid = (left + right) / 2
		if arr[mid] > target {
			right = mid
		} else {
			left = mid
		}
	}
	if arr[right] == target {
		return right
	} else if arr[left] == target {
		return left
	}
	return -1
}

func Max[T comparer[T]](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T comparer[T]](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Unique[T any](arr []T, cmp func(a, b T) types.Ordering) []T {
	var n = len(arr)
	var b = make([]T, 0, n)
	quickSort(arr, 0, n-1, cmp)
	b = append(b, arr[0])
	for i := 1; i < n; i++ {
		if cmp(arr[i], arr[i-1]) != types.Equal {
			b = append(b, arr[i])
		}
	}
	return b
}

func Swap[T any](a, b *T) {
	*a, *b = *b, *a
}

func Fill[T any](arr []T, v T) {
	for i, _ := range arr {
		arr[i] = v
	}
}

func Foreach[T any](arr []T, fn func(index int, value T)) {
	for i, v := range arr {
		fn(i, v)
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
