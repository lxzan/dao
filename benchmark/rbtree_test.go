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

func (c rbtree_datanode) Compare(a, b entry) dao.Ordering {
	if a.Key > b.Key {
		return dao.Greater
	} else if a.Key == b.Key {
		return dao.Equal
	} else {
		return dao.Less
	}
}

const rbtree_count = 100000

var benchtree *rbtree.RBTree[rbtree_datanode]

func init() {
	benchtree = rbtree.New[rbtree_datanode](func(a, b *rbtree_datanode) dao.Ordering {
		if a.Key > b.Key {
			return dao.Greater
		} else if a.Key == b.Key {
			return dao.Equal
		} else {
			return dao.Less
		}
	})
	for i := 0; i < rbtree_count; i++ {
		var length = utils.Rand.Intn(16) + 1
		var key = utils.Numeric.Generate(4)
		benchtree.Insert(&rbtree_datanode{Key: key, Val: length})
	}
}

func BenchmarkRBTree_Insert(b *testing.B) {
	var tree = rbtree.New[rbtree_datanode](func(a, b *rbtree_datanode) dao.Ordering {
		if a.Key > b.Key {
			return dao.Greater
		} else if a.Key == b.Key {
			return dao.Equal
		} else {
			return dao.Less
		}
	})
	for i := 0; i < b.N; i++ {
		tree.Insert(&rbtree_datanode{Key: strconv.Itoa(i), Val: 1})
	}
}

func BenchmarkRBTree_Find(b *testing.B) {
	var tree = rbtree.New[rbtree_datanode](func(a, b *rbtree_datanode) dao.Ordering {
		if a.Key > b.Key {
			return dao.Greater
		} else if a.Key == b.Key {
			return dao.Equal
		} else {
			return dao.Less
		}
	})
	for i := 0; i < b.N; i++ {
		tree.Insert(&rbtree_datanode{Key: strconv.Itoa(i), Val: 1})
	}

	for i := 0; i < b.N; i++ {
		tree.Find(&rbtree_datanode{Key: strconv.Itoa(i)})
	}
}

func BenchmarkRBTree_Delete(b *testing.B) {
	var tree = rbtree.New[rbtree_datanode](func(a, b *rbtree_datanode) dao.Ordering {
		if a.Key > b.Key {
			return dao.Greater
		} else if a.Key == b.Key {
			return dao.Equal
		} else {
			return dao.Less
		}
	})
	for i := 0; i < b.N; i++ {
		tree.Insert(&rbtree_datanode{Key: strconv.Itoa(i), Val: 1})
	}

	for i := 0; i < b.N; i++ {
		tree.Delete(&rbtree_datanode{Key: strconv.Itoa(i)})
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
		benchtree.Query(&rbtree.QueryBuilder[rbtree_datanode]{
			LeftFilter:  func(d *rbtree_datanode) bool { return d.Key >= left },
			RightFilter: func(d *rbtree_datanode) bool { return d.Key < right },
			Limit:       10,
			Order:       rbtree.ASC,
		})
	}
}

func BenchmarkRBTree_GreaterEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var k = utils.Numeric.Generate(4)
		benchtree.Query(&rbtree.QueryBuilder[rbtree_datanode]{
			LeftFilter: func(d *rbtree_datanode) bool { return d.Key >= k },
			Limit:      10,
			Order:      rbtree.ASC,
		})
	}
}

func BenchmarkRBTree_LessEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var k = utils.Numeric.Generate(4)
		benchtree.Query(&rbtree.QueryBuilder[rbtree_datanode]{
			RightFilter: func(d *rbtree_datanode) bool { return d.Key <= k },
			Limit:       10,
			Order:       rbtree.DESC,
		})
	}
}
