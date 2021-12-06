package internal

import "github.com/lxzan/dao/types"

func getMedium[T any](arr []T, begin int, end int, cmp func(a, b T) types.Ordering) int {
	var mid = (begin + end) / 2
	if cmp(arr[end], arr[mid])+cmp(arr[mid], arr[begin]) != 0 {
		return mid
	}
	if cmp(arr[end], arr[begin])+cmp(arr[begin], arr[mid]) != 0 {
		return begin
	}
	return end
}

func QuickSort[T any](arr []T, begin int, end int, cmp func(a, b T) types.Ordering) {
	if begin >= end {
		return
	}
	var index = begin
	var mid = getMedium(arr, begin, end, cmp)
	arr[mid], arr[begin] = arr[begin], arr[mid]
	for i := begin + 1; i <= end; i++ {
		var order = cmp(arr[i], arr[begin])
		if order == Less || (order == Equal && i%2 == 0) {
			index++
			arr[index], arr[i] = arr[i], arr[index]
		}
	}
	arr[index], arr[begin] = arr[begin], arr[index]
	QuickSort(arr, begin, index-1, cmp)
	QuickSort(arr, index+1, end, cmp)
}
