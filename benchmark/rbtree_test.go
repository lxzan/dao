package benchmark

import (
	"github.com/lxzan/dao"
	"github.com/lxzan/dao/internal/utils"
	"github.com/lxzan/dao/rbtree"
	"testing"
)

type rbtree_datanode struct {
	Key int
	Val string
}

func rbtree_datanode_compare(a, b *rbtree_datanode) dao.Ordering {
	if a.Key > b.Key {
		return dao.Greater
	} else if a.Key == b.Key {
		return dao.Equal
	} else {
		return dao.Less
	}
}

func BenchmarkRBTree_Insert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var tree = rbtree.New(rbtree_datanode_compare)
		for j := 0; j < bench_count; j++ {
			tree.Insert(&rbtree_datanode{Key: j, Val: ""})
		}
	}
}

func BenchmarkRBTree_Find(b *testing.B) {
	var tree = rbtree.New(rbtree_datanode_compare)
	for j := 0; j < bench_count; j++ {
		tree.Insert(&rbtree_datanode{Key: j, Val: ""})
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < bench_count; j++ {
			tree.Find(&rbtree_datanode{Key: j, Val: ""})
		}
	}
}

func BenchmarkRBTree_Delete(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var tree = rbtree.New(rbtree_datanode_compare)
		for j := 0; j < bench_count; j++ {
			tree.Insert(&rbtree_datanode{Key: j, Val: ""})
		}

		for j := 0; j < bench_count; j++ {
			tree.Delete(&rbtree_datanode{Key: j, Val: ""})
		}
	}
}

func BenchmarkRBTree_Between(b *testing.B) {
	var tree = rbtree.New(rbtree_datanode_compare)
	for j := 0; j < bench_count; j++ {
		tree.Insert(&rbtree_datanode{Key: j, Val: ""})
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < bench_count; j++ {
			var left = utils.Rand.Intn(bench_count)
			var right = left + 10
			tree.Query(&rbtree.QueryBuilder[rbtree_datanode]{
				LeftFilter:  func(d *rbtree_datanode) bool { return d.Key >= left },
				RightFilter: func(d *rbtree_datanode) bool { return d.Key < right },
				Limit:       10,
				Order:       rbtree.ASC,
			})
		}
	}
	b.StopTimer()
}

func BenchmarkRBTree_GreaterEqual(b *testing.B) {
	var tree = rbtree.New(rbtree_datanode_compare)
	for j := 0; j < bench_count; j++ {
		tree.Insert(&rbtree_datanode{Key: j, Val: ""})
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < bench_count; j++ {
			var k = utils.Rand.Intn(bench_count)
			tree.Query(&rbtree.QueryBuilder[rbtree_datanode]{
				LeftFilter: func(d *rbtree_datanode) bool { return d.Key >= k },
				Limit:      10,
				Order:      rbtree.ASC,
			})
		}
	}
	b.StopTimer()
}

func BenchmarkRBTree_LessEqual(b *testing.B) {
	var tree = rbtree.New(rbtree_datanode_compare)
	for j := 0; j < bench_count; j++ {
		tree.Insert(&rbtree_datanode{Key: j, Val: ""})
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < bench_count; j++ {
			var k = utils.Rand.Intn(bench_count)
			tree.Query(&rbtree.QueryBuilder[rbtree_datanode]{
				RightFilter: func(d *rbtree_datanode) bool { return d.Key <= k },
				Limit:       10,
				Order:       rbtree.DESC,
			})
		}
	}
	b.StopTimer()
}

func BenchmarkRBTree_GetMinKey(b *testing.B) {
	var tree = rbtree.New(rbtree_datanode_compare)
	for j := 0; j < bench_count; j++ {
		tree.Insert(&rbtree_datanode{Key: j, Val: ""})
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < bench_count; j++ {
			tree.GetMinKey(&rbtree_datanode{Key: j, Val: ""})
			// res := tree.GetMinKey(&rbtree_datanode{Key: j, Val: ""})
			// println(res)
		}
	}
	b.StopTimer()
}
