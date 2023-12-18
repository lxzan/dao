package benchmark

import (
	"github.com/lxzan/dao"
	"github.com/lxzan/dao/algorithm"
	"sort"
	"testing"
)

func BenchmarkSort_Quick(b *testing.B) {
	var cmp = func(a, b int) dao.CompareResult {
		if a > b {
			return 1
		} else if a < b {
			return -1
		} else {
			return 0
		}
	}

	for i := 0; i < b.N; i++ {
		var arr = make([]int, bench_count, bench_count)
		copy(arr, testvals[:bench_count])
		algorithm.Sort(arr, cmp)
	}
}

func BenchmarkSort_Std(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var arr = make([]int, bench_count, bench_count)
		copy(arr, testvals[:bench_count])
		sort.Ints(arr)
	}
}
