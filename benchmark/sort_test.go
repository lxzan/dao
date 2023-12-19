package benchmark

import (
	"github.com/lxzan/dao/algorithm"
	"sort"
	"testing"
)

func BenchmarkSort_Quick(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var arr = make([]int, bench_count)
		copy(arr, testvals[:bench_count])
		algorithm.Sort(arr)
	}
}

func BenchmarkSort_Std(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var arr = make([]int, bench_count)
		copy(arr, testvals[:bench_count])
		sort.Ints(arr)
	}
}
