package array_list

import (
	"github.com/lxzan/dao"
	"github.com/lxzan/dao/algorithm"
	"testing"
)

func TestSlice_Sort(t *testing.T) {
	var arr = New[int](0, 0)
	arr.Push(1, 3, 5, 7, 9, 2, 4, 6, 8, 0)
	arr.Sort(dao.ASC[int])
	if !algorithm.IsSorted(arr.Elem(), dao.ASC[int]) {
		t.Fatal("error!")
	}
}
