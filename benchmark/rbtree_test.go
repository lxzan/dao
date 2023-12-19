package benchmark

import (
	"github.com/lxzan/dao"
	"github.com/lxzan/dao/algorithm"
	"github.com/lxzan/dao/rbtree"
	"math/rand"
	"testing"
)

func BenchmarkRBTree_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var tree = rbtree.New[int, string]()
		for j := 0; j < bench_count; j++ {
			tree.Set(j, "")
		}
	}
}

func BenchmarkRBTree_Get(b *testing.B) {
	var tree = rbtree.New[int, string]()
	for j := 0; j < bench_count; j++ {
		tree.Set(j, "")
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < bench_count; j++ {
			tree.Get(j)
		}
	}
}

func BenchmarkRBTree_Query(b *testing.B) {
	var tree = rbtree.New[int, string]()
	for j := 0; j < bench_count; j++ {
		tree.Set(j, "")
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < bench_count; j++ {
			x, y := rand.Intn(bench_count), rand.Intn(bench_count)
			if x > y {
				algorithm.Swap(&x, &y)
			}
			tree.
				NewQuery().
				Left(func(key int) bool { return key >= x }).
				Right(func(key int) bool { return key <= y }).
				Order(dao.DESC).
				Limit(10).
				Do()
		}
	}
}
