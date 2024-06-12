package benchmark

import (
	tree "github.com/lxzan/dao/segment_tree"
	"math/rand"
	"testing"
)

func BenchmarkSegmentTree_Query(b *testing.B) {
	var arr = make([]int, 0)
	for i := 0; i < bench_count; i++ {
		arr = append(arr, testvals[i])
	}
	var st = tree.New(arr, tree.NewIntSummary[int], tree.MergeIntSummary[int])

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < bench_count; j++ {
			var left = rand.Intn(bench_count)
			var right = rand.Intn(bench_count)
			if left > right {
				left, right = right, left
			}
			st.Query(left, right)
		}
	}
}

func BenchmarkSegmentTree_Update(b *testing.B) {
	var arr1 = make([]int, 0)
	for i := 0; i < bench_count; i++ {
		arr1 = append(arr1, testvals[i])
	}
	var st = tree.New(arr1, tree.NewIntSummary[int], tree.MergeIntSummary[int])

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < bench_count; j++ {
			var x = rand.Intn(bench_count)
			st.Update(x, x)
		}
	}
}
