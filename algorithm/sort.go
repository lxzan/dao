package algorithm

import (
	"github.com/lxzan/dao/types/cmp"
)

func IsSorted[T any](arr []T, compare cmp.CompareFunc[T]) bool {
	var n = len(arr)
	if n <= 1 {
		return true
	}
	for i := 1; i < n; i++ {
		if compare(arr[i], arr[i-1]) < 0 {
			return false
		}
	}
	return true
}

func Sort[T cmp.Ordered](arr []T) {
	var f = cmp.Compare[T]
	if IsSorted(arr, f) {
		return
	}
	QuickSort(arr, 0, len(arr)-1, f)
}

func SortBy[T any](arr []T, compare cmp.CompareFunc[T]) {
	if IsSorted(arr, compare) {
		return
	}
	QuickSort(arr, 0, len(arr)-1, compare)
}

func getMedium[T any](arr []T, begin int, end int, compare cmp.CompareFunc[T]) int {
	var mid = (begin + end) / 2
	var x = compare(arr[begin], arr[mid])
	var y = compare(arr[mid], arr[end])
	if x+y != 0 {
		return mid
	}

	var z = compare(arr[begin], arr[end])
	y *= -1
	if y+z != 0 {
		return end
	}
	return begin
}

func insertionSort[T any](arr []T, a, b int, compare cmp.CompareFunc[T]) {
	for i := a + 1; i <= b; i++ {
		for j := i; j > a && compare(arr[j], arr[j-1]) == cmp.LT; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}

// QuickSort 快速排序 begin <= x <= end 区间
// 对于随机数据, 此算法比标准库稍快; 对于本身比较有序的数据, 标准库表现更佳.
func QuickSort[T any](arr []T, begin int, end int, compare cmp.CompareFunc[T]) {
	if begin >= end {
		return
	}
	if end-begin <= 15 {
		insertionSort(arr, begin, end, compare)
		return
	}

	var index = begin
	var mid = getMedium(arr, begin, end, compare)
	arr[mid], arr[begin] = arr[begin], arr[mid]
	for i := begin + 1; i <= end; i++ {
		var flag = compare(arr[i], arr[begin])
		if flag == cmp.LT || (flag == cmp.EQ && i%2 == 0) {
			index++
			arr[index], arr[i] = arr[i], arr[index]
		}
	}
	arr[index], arr[begin] = arr[begin], arr[index]

	QuickSort(arr, begin, index-1, compare)
	QuickSort(arr, index+1, end, compare)
}

// BinarySearch 二分搜索
// @return 数组下标 如果不存在, 返回-1
func BinarySearch[T any](arr []T, target T, compare cmp.CompareFunc[T]) int {
	var n = len(arr)
	if n == 0 {
		return -1
	}

	var left = 0
	var right = n - 1
	for right-left > 1 {
		var mid = (left + right) / 2
		switch compare(arr[mid], target) {
		case cmp.EQ:
			return mid
		case cmp.GT:
			right = mid
		default:
			left = mid
		}
	}

	if compare(arr[left], target) == cmp.EQ {
		return left
	} else if compare(arr[right], target) == cmp.EQ {
		return right
	} else {
		return -1
	}
}
