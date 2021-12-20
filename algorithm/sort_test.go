package algorithm

import (
	"github.com/lxzan/dao"
	"math/rand"
	"testing"
)

func TestSort(t *testing.T) {
	var arr = make([]int, 0)
	for i := 0; i < 999; i++ {
		arr = append(arr, rand.Intn(1000))
	}
	Sort[int](arr, dao.ASC[int])

	if !IsSorted[int](arr, dao.ASC[int]) {
		t.Error("not sorted!")
	}
}
