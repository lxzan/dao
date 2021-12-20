package benchmark

import (
	"github.com/lxzan/dao"
	"github.com/lxzan/dao/algorithm"
	"sort"
	"testing"
)

func BenchmarkSort_Quick(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < bench_count; j++ {
			var arr = make([]int, bench_count, bench_count)
			copy(arr, testvals[:bench_count])
			algorithm.Sort(arr, dao.ASC[int])
		}
	}
}

func BenchmarkSort_Golang(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < bench_count; j++ {
			var arr = make([]int, bench_count, bench_count)
			copy(arr, testvals[:bench_count])
			sort.Ints(arr)
		}
	}
}
