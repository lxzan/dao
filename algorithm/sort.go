package algorithm

import (
	"cmp"
	"github.com/lxzan/dao"
)

func IsSorted[T any](arr []T, cmp dao.CompareFunc[T]) bool {
	var n = len(arr)
	if n <= 1 {
		return true
	}
	for i := 1; i < n; i++ {
		if cmp(arr[i], arr[i-1]) < 0 {
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

func SortBy[T any](arr []T, cmp dao.CompareFunc[T]) {
	if IsSorted(arr, cmp) {
		return
	}
	QuickSort(arr, 0, len(arr)-1, cmp)
}

func getMedium[T any](arr []T, begin int, end int, cmp dao.CompareFunc[T]) int {
	var mid = (begin + end) / 2
	var x = cmp(arr[begin], arr[mid])
	var y = cmp(arr[mid], arr[end])
	if x+y != 0 {
		return mid
	}

	var z = cmp(arr[begin], arr[end])
	y *= -1
	if y+z != 0 {
		return end
	}
	return begin
}

func insertionSort[T any](arr []T, a, b int, cmp dao.CompareFunc[T]) {
	for i := a + 1; i <= b; i++ {
		for j := i; j > a && cmp(arr[j], arr[j-1]) == dao.Less; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}

// QuickSort 快速排序 begin <= x <= end 区间
// 对于随机数据, 此算法比标准库稍快; 对于本身比较有序的数据, 标准库表现更佳.
func QuickSort[T any](arr []T, begin int, end int, cmp dao.CompareFunc[T]) {
	if begin >= end {
		return
	}
	if end-begin <= 15 {
		insertionSort(arr, begin, end, cmp)
		return
	}

	var index = begin
	var mid = getMedium(arr, begin, end, cmp)
	arr[mid], arr[begin] = arr[begin], arr[mid]
	for i := begin + 1; i <= end; i++ {
		var flag = cmp(arr[i], arr[begin])
		if flag == dao.Less || (flag == dao.Equal && i%2 == 0) {
			index++
			arr[index], arr[i] = arr[i], arr[index]
		}
	}
	arr[index], arr[begin] = arr[begin], arr[index]

	QuickSort(arr, begin, index-1, cmp)
	QuickSort(arr, index+1, end, cmp)
}

// BinarySearch 二分搜索
// @return 数组下标 如果不存在, 返回-1
func BinarySearch[T any](arr []T, target T, cmp dao.CompareFunc[T]) int {
	var n = len(arr)
	if n == 0 {
		return -1
	}

	var left = 0
	var right = n - 1
	for right-left > 1 {
		var mid = (left + right) / 2
		switch cmp(arr[mid], target) {
		case dao.Equal:
			return mid
		case dao.Greater:
			right = mid
		default:
			left = mid
		}
	}

	if cmp(arr[left], target) == dao.Equal {
		return left
	} else if cmp(arr[right], target) == dao.Equal {
		return right
	} else {
		return -1
	}
}
