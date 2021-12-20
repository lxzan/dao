package benchmark

import (
	"fmt"
	"github.com/lxzan/dao"
	"github.com/lxzan/dao/internal/utils"
	"github.com/lxzan/dao/rbtree"
	"strconv"
	"testing"
)

type rbtree_datanode struct {
	Key string
	Val int
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
		var tree = rbtree.New[rbtree_datanode](rbtree_datanode_compare)
		for j := 0; j < bench_count; j++ {
			tree.Insert(&rbtree_datanode{Key: testkeys[j], Val: testvals[j]})
		}
	}
}

func BenchmarkRBTree_Find(b *testing.B) {
	var tree = rbtree.New[rbtree_datanode](rbtree_datanode_compare)
	for j := 0; j < bench_count; j++ {
		tree.Insert(&rbtree_datanode{Key: testkeys[j], Val: testvals[j]})
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < bench_count; j++ {
			tree.Find(&rbtree_datanode{Key: testkeys[j]})
		}
	}
}

func BenchmarkRBTree_Delete(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var tree = rbtree.New[rbtree_datanode](rbtree_datanode_compare)
		for j := 0; j < bench_count; j++ {
			tree.Insert(&rbtree_datanode{Key: testkeys[j], Val: testvals[j]})
		}

		for j := 0; j < bench_count; j++ {
			tree.Delete(&rbtree_datanode{Key: testkeys[j]})
		}
	}
}

func BenchmarkRBTree_Between(b *testing.B) {
	var tree = rbtree.New[rbtree_datanode](rbtree_datanode_compare)
	for j := 0; j < bench_count; j++ {
		tree.Insert(&rbtree_datanode{Key: testkeys[j], Val: testvals[j]})
	}

	for i := 0; i < b.N; i++ {
		for j := 0; j < bench_count; j++ {
			var left = utils.Numeric.Generate(4)
			x, _ := strconv.Atoi(left)
			var right = fmt.Sprintf("%04d", x+10)
			if left > right {
				right, left = left, right
			}
			tree.Query(&rbtree.QueryBuilder[rbtree_datanode]{
				LeftFilter:  func(d *rbtree_datanode) bool { return d.Key >= left },
				RightFilter: func(d *rbtree_datanode) bool { return d.Key < right },
				Limit:       10,
				Order:       rbtree.ASC,
			})
		}
	}
}

func BenchmarkRBTree_GreaterEqual(b *testing.B) {
	var tree = rbtree.New[rbtree_datanode](rbtree_datanode_compare)
	for j := 0; j < bench_count; j++ {
		tree.Insert(&rbtree_datanode{Key: testkeys[j], Val: testvals[j]})
	}

	for i := 0; i < b.N; i++ {
		for j := 0; j < bench_count; j++ {
			var k = utils.Numeric.Generate(4)
			tree.Query(&rbtree.QueryBuilder[rbtree_datanode]{
				LeftFilter: func(d *rbtree_datanode) bool { return d.Key >= k },
				Limit:      10,
				Order:      rbtree.ASC,
			})
		}
	}
}

func BenchmarkRBTree_LessEqual(b *testing.B) {
	var tree = rbtree.New[rbtree_datanode](rbtree_datanode_compare)
	for j := 0; j < bench_count; j++ {
		tree.Insert(&rbtree_datanode{Key: testkeys[j], Val: testvals[j]})
	}

	for i := 0; i < b.N; i++ {
		for j := 0; j < bench_count; j++ {
			var k = utils.Numeric.Generate(4)
			tree.Query(&rbtree.QueryBuilder[rbtree_datanode]{
				RightFilter: func(d *rbtree_datanode) bool { return d.Key <= k },
				Limit:       10,
				Order:       rbtree.DESC,
			})
		}
	}
}
