package algorithm

import (
	"github.com/lxzan/dao"
	"math/rand"
	"testing"
)

func TestQuickSort(t *testing.T) {
	var arr = make([]int, 0)
	for i := 0; i < 999; i++ {
		arr = append(arr, rand.Intn(1000))
	}
	QuickSort[int](arr, dao.ASC[int])

	if !IsSorted[int](arr, dao.ASC[int]) {
		t.Error("not sorted!")
	}
}
