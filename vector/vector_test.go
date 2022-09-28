package vector

import (
	"github.com/lxzan/dao"
	"github.com/lxzan/dao/algorithm"
	"github.com/lxzan/dao/internal/utils"
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

func TestSlice_Unique(t *testing.T) {
	var n = 1000
	var arr1 = New[int](0, 0)
	for i := 0; i < n; i++ {
		arr1.Push(utils.Rand.Intn(100))
	}
	var arr2 = arr1.Unique(dao.ASC[int])

	var length = arr2.Len()
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if arr2.Get(i) == arr2.Get(j) {
				t.Fatal("error!")
			}
		}
	}
}
