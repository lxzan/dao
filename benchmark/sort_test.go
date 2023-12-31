package benchmark

import (
	"github.com/lxzan/dao/algo"
	"sort"
	"testing"
)

func BenchmarkSort_Quick(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var arr = make([]int, bench_count)
		copy(arr, testvals[:bench_count])
		algo.Sort(arr)
	}
}

func BenchmarkSort_Std(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var arr = make([]int, bench_count)
		copy(arr, testvals[:bench_count])
		sort.Ints(arr)
	}
}
