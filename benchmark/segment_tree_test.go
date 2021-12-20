package benchmark

import (
	"github.com/lxzan/dao/segment_tree"
	"math/rand"
	"testing"
)

func BenchmarkSegmentTree_Query(b *testing.B) {
	var arr = make([]int, 0)
	for i := 0; i < bench_count; i++ {
		arr = append(arr, testvals[i])
	}
	var tree = segment_tree.New[int, segment_tree.Schema[int]](arr, segment_tree.Init[int], segment_tree.Merge[int])

	for i := 0; i < b.N; i++ {
		for j := 0; j < bench_count; j++ {
			var j = rand.Intn(bench_count)
			var k = rand.Intn(bench_count)
			if j > k {
				j, k = k, j
			}
			tree.Query(j, k)
		}
	}
}
