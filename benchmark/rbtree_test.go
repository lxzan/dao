package benchmark

import (
	"fmt"
	"github.com/lxzan/dao/internal/utils"
	"github.com/lxzan/dao/rbtree"
	"strconv"
	"testing"
)

const rbtree_count = 100000

var benchtree *rbtree.RBTree[string, int]

func init() {
	benchtree = rbtree.New[string, int]()
	for i := 0; i < rbtree_count; i++ {
		var length = utils.Rand.Intn(16) + 1
		var key = utils.Numeric.Generate(4)
		benchtree.Insert(key, length)
	}
}

func BenchmarkRBTree_Insert(b *testing.B) {
	var tree = rbtree.New[string, int]()
	for i := 0; i < b.N; i++ {
		tree.Insert(strconv.Itoa(i), 1)
	}
}

func BenchmarkRBTree_Find(b *testing.B) {
	var tree = rbtree.New[string, int]()
	for i := 0; i < b.N; i++ {
		tree.Insert(strconv.Itoa(i), 1)
	}

	for i := 0; i < b.N; i++ {
		tree.Find(strconv.Itoa(i))
	}
}

func BenchmarkRBTree_Delete(b *testing.B) {
	var tree = rbtree.New[string, int]()
	for i := 0; i < b.N; i++ {
		tree.Insert(strconv.Itoa(i), 1)
	}

	for i := 0; i < b.N; i++ {
		tree.Delete(strconv.Itoa(i))
	}
}

func BenchmarkRBTree_Between(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var left = utils.Numeric.Generate(4)
		x, _ := strconv.Atoi(left)
		var right = fmt.Sprintf("%04d", x+10)
		if left > right {
			right, left = left, right
		}
		benchtree.Query(&rbtree.QueryBuilder[string]{
			LeftFilter:  func(key string) bool { return key >= left },
			RightFilter: func(key string) bool { return key < right },
			Limit:       10,
			Order:       rbtree.ASC,
		})
	}
}

func BenchmarkRBTree_GreaterEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var k = utils.Numeric.Generate(4)
		benchtree.Query(&rbtree.QueryBuilder[string]{
			LeftFilter: func(key string) bool { return key >= k },
			Limit:      10,
			Order:      rbtree.ASC,
		})
	}
}

func BenchmarkRBTree_LessEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var k = utils.Numeric.Generate(4)
		benchtree.Query(&rbtree.QueryBuilder[string]{
			RightFilter: func(key string) bool { return key <= k },
			Limit:       10,
			Order:       rbtree.DESC,
		})
	}
}
