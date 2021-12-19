package benchmark

import (
	"github.com/lxzan/dao"
	"github.com/lxzan/dao/algorithm"
	"github.com/lxzan/dao/internal/utils"
	"sort"
	"testing"
)

var testdata = []int{}

var asc = func(a, b int) dao.Ordering {
	if a > b {
		return dao.Greater
	} else if a == b {
		return dao.Equal
	} else {
		return dao.Less
	}
}

func init() {
	for i := 0; i < 1024; i++ {
		testdata = append(testdata, utils.Rand.Intn(100000))
	}
}

func BenchmarkSort_Quick(b *testing.B) {
	var n = len(testdata)
	for i := 0; i < b.N; i++ {
		var arr = make([]int, n)
		copy(arr, testdata)
		algorithm.QuickSort(arr, asc)
	}
}

func BenchmarkSort_Golang(b *testing.B) {
	var n = len(testdata)
	for i := 0; i < b.N; i++ {
		var arr = make([]int, n)
		copy(arr, testdata)
		sort.Ints(arr)
	}
}
