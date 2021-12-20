package benchmark

import (
	"github.com/lxzan/dao/segment_tree"
	"math/rand"
	"testing"
)

const stree_count = 1000

var stree *segment_tree.SegmentTree[int, segment_tree.Schema[int]]

func init() {
	var arr = make([]int, 0)
	for i := 0; i < stree_count; i++ {
		arr = append(arr, i)
	}
	stree = segment_tree.New[int, segment_tree.Schema[int]](arr)
}

func BenchmarkSegmentTree_Query(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var j = rand.Intn(stree_count)
		var k = rand.Intn(stree_count)
		if j > k {
			j, k = k, j
		}
		stree.Query(j, k)
	}
}
