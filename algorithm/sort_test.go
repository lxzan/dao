package algorithm

import (
	"github.com/lxzan/dao"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestIsSorted(t *testing.T) {
	var as = assert.New(t)
	as.Equal(true, IsSorted([]int{1, 2, 3}, dao.ASC[int]))
	as.Equal(true, IsSorted([]int{1, 2, 3, 4}, dao.ASC[int]))
	as.Equal(true, IsSorted([]int{}, dao.ASC[int]))
	as.Equal(true, IsSorted([]int{1}, dao.ASC[int]))
	as.Equal(true, IsSorted([]int{1, 2, 2, 2}, dao.ASC[int]))
	as.Equal(true, IsSorted([]int{3, 2, 1}, dao.DESC[int]))

	as.Equal(false, IsSorted([]int{1, 3, 2}, dao.ASC[int]))
	as.Equal(false, IsSorted([]int{1, 2, 3, 2}, dao.ASC[int]))
	as.Equal(false, IsSorted([]int{1, 2, 2, 1}, dao.ASC[int]))
	as.Equal(false, IsSorted([]int{3, 2, 1}, dao.ASC[int]))
	as.Equal(false, IsSorted([]int{3, 2, 1, 0}, dao.ASC[int]))
}

func TestSort(t *testing.T) {
	var arr = make([]int, 0)
	for i := 0; i < 999; i++ {
		arr = append(arr, rand.Intn(1000))
	}
	Sort(arr, dao.ASC[int])

	if !IsSorted(arr, dao.ASC[int]) {
		t.Error("not sorted!")
	}
}

func TestBinarySearch(t *testing.T) {
	var count = 1000
	var arr = make([]int, 0, count)
	for i := 0; i < count; i++ {
		arr = append(arr, rand.Intn(count))
	}
	arr = Unique(arr, func(x int) int { return x })
	Sort(arr, dao.DESC[int])

	var m = make(map[int]int)
	for i, v := range arr {
		m[v] = i
	}

	for k, v := range m {
		res := BinarySearch(arr, k, dao.DESC[int])
		if res != v {
			t.Fatal("error!")
		}
	}

	if res := BinarySearch(arr, count, dao.DESC[int]); res != -1 {
		t.Fatal("error!")
	}
}
