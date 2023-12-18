package benchmark

import (
	tree "github.com/lxzan/dao/segment_tree"
	"math/rand"
	"testing"
)

func BenchmarkSegmentTree_Query(b *testing.B) {
	var arr = make([]tree.Int64, 0)
	for i := 0; i < bench_count; i++ {
		arr = append(arr, tree.Int64(testvals[i]))
	}
	var st = tree.New[tree.Int64Schema, tree.Int64](arr)

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
	var arr1 = make([]tree.Int64, 0)
	for i := 0; i < bench_count; i++ {
		arr1 = append(arr1, tree.Int64(testvals[i]))
	}
	var st = tree.New[tree.Int64Schema, tree.Int64](arr1)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < bench_count; j++ {
			var x = rand.Intn(bench_count)
			st.Update(x, tree.Int64(x))
		}
	}
}
