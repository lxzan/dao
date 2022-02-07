package rbtree

import (
	"fmt"
	"github.com/lxzan/dao"
	"github.com/lxzan/dao/algorithm"
	"github.com/lxzan/dao/internal/utils"
	"sort"
	"strconv"
	"testing"
)

type entry struct {
	Key string
	Val int
}

func (c entry) Compare(a, b entry) dao.Ordering {
	if a.Key > b.Key {
		return dao.Greater
	} else if a.Key == b.Key {
		return dao.Equal
	} else {
		return dao.Less
	}
}

func (c *RBTree[T]) validate(t *testing.T, node *rbtree_node[T]) {
	if node == nil {
		return
	}
	if node.left != nil {
		if !c.is_key_empty(node.left.data) && c.cmp(node.data, node.left.data) != dao.Greater {
			t.Error("left node error!")
		}
		c.validate(t, node.left)
	}

	if node.right != nil {
		if !c.is_key_empty(node.right.data) && c.cmp(node.data, node.right.data) != dao.Less {
			t.Error("right node error!")
		}
		c.validate(t, node.right)
	}
}

func TestNew(t *testing.T) {
	var tree = New(func(a, b *entry) dao.Ordering {
		if a.Key > b.Key {
			return dao.Greater
		} else if a.Key == b.Key {
			return dao.Equal
		} else {
			return dao.Less
		}
	})
	var m = make(map[string]int)

	for i := 0; i < 1000; i++ {
		var length = utils.Rand.Intn(16) + 1
		var key = utils.Numeric.Generate(length)
		tree.Insert(&entry{Key: key, Val: length})
		m[key] = length
	}

	var idx = 0
	for k := range m {
		if idx >= 500 {
			break
		}
		delete(m, k)
		tree.Delete(&entry{Key: k})
		idx++
	}

	for i := 0; i < 10000; i++ {
		var length = utils.Rand.Intn(16) + 1
		var key = utils.Alphabet.Generate(length)
		tree.Insert(&entry{Key: key, Val: length})
		m[key] = length
	}

	for k, v := range m {
		result, exist := tree.Find(&entry{Key: k})
		if !exist || result.Val != v {
			t.Fatal("error!")
		}
	}

	if len(m) != tree.Len() {
		t.Fatal("error!")
	}

	tree.validate(t, tree.root)
}

func TestRBTree_ForEach(t *testing.T) {
	var tree = New(func(a, b *entry) dao.Ordering {
		if a.Key > b.Key {
			return dao.Greater
		} else if a.Key == b.Key {
			return dao.Equal
		} else {
			return dao.Less
		}
	})

	for i := 0; i < 100; i++ {
		tree.Insert(&entry{Key: utils.Alphabet.Generate(16), Val: utils.Rand.Intn(1000)})
	}

	var arr1 = make([]string, 0)
	tree.ForEach(func(item *entry) (continued bool) {
		arr1 = append(arr1, item.Key)
		return len(arr1) < 50
	})
	if len(arr1) != 50 {
		t.Fatal("error!")
	}

	var arr2 = make([]string, 0)
	tree.ForEach(func(item *entry) (continued bool) {
		arr2 = append(arr2, item.Key)
		return true
	})
	if len(arr2) != tree.Len() {
		t.Fatal("error!")
	}
}

func TestRBTree_Between(t *testing.T) {
	var tree = New(func(a, b *entry) dao.Ordering {
		if a.Key > b.Key {
			return dao.Greater
		} else if a.Key == b.Key {
			return dao.Equal
		} else {
			return dao.Less
		}
	})
	var m = make(map[string]int)
	for i := 0; i < 10000; i++ {
		var length = utils.Rand.Intn(16) + 1
		var key = utils.Numeric.Generate(4)
		m[key] = length
		tree.Insert(&entry{Key: key, Val: length})
	}

	var limit = 100
	for i := 0; i < 100; i++ {
		var left = utils.Numeric.Generate(4)
		x, _ := strconv.Atoi(left)
		var right = fmt.Sprintf("%04d", x+limit)
		if left > right {
			right, left = left, right
		}
		var keys1 = tree.Query(&QueryBuilder[entry]{
			LeftFilter:  func(d *entry) bool { return d.Key >= left },
			RightFilter: func(d *entry) bool { return d.Key <= right },
			Limit:       limit,
			Order:       DESC,
		})
		var keys2 = make([]string, 0)
		for k := range m {
			if k >= left && k <= right {
				keys2 = append(keys2, k)
			}
		}
		sort.Strings(keys2)
		algorithm.Reverse(keys2)
		if len(keys2) > limit {
			keys2 = keys2[:limit]
		}

		if !utils.SameStrings(keys2, algorithm.GetFields(keys1, func(x *entry) string {
			return x.Key
		})) {
			t.Fatal("error!")
		}
	}
}

func TestRBTree_GreaterEqual(t *testing.T) {
	var tree = New(func(a, b *entry) dao.Ordering {
		if a.Key > b.Key {
			return dao.Greater
		} else if a.Key == b.Key {
			return dao.Equal
		} else {
			return dao.Less
		}
	})
	var m = make(map[string]int)
	for i := 0; i < 10000; i++ {
		var length = utils.Rand.Intn(16) + 1
		var key = utils.Numeric.Generate(4)
		m[key] = length
		tree.Insert(&entry{Key: key, Val: length})
	}

	var limit = 100
	for i := 0; i < 100; i++ {
		var left = utils.Numeric.Generate(4)
		var keys1 = tree.Query(&QueryBuilder[entry]{
			LeftFilter: func(d *entry) bool { return d.Key >= left },
			Limit:      limit,
			Order:      ASC,
		})
		var keys2 = make([]string, 0)
		for k := range m {
			if k >= left {
				keys2 = append(keys2, k)
			}
		}
		sort.Strings(keys2)
		if len(keys2) > limit {
			keys2 = keys2[:limit]
		}

		if !utils.SameStrings(keys2, algorithm.GetFields(keys1, func(x *entry) string {
			return x.Key
		})) {
			t.Fatal("error!")
		}
	}
}

func TestRBTree_LessEqual(t *testing.T) {
	var tree = New(func(a, b *entry) dao.Ordering {
		if a.Key > b.Key {
			return dao.Greater
		} else if a.Key == b.Key {
			return dao.Equal
		} else {
			return dao.Less
		}
	})
	var m = make(map[string]int)
	for i := 0; i < 10000; i++ {
		var length = utils.Rand.Intn(16) + 1
		var key = utils.Numeric.Generate(4)
		m[key] = length
		tree.Insert(&entry{Key: key, Val: length})
	}

	var limit = 100
	for i := 0; i < 100; i++ {
		var target = utils.Numeric.Generate(4)
		var keys1 = tree.Query(&QueryBuilder[entry]{
			RightFilter: func(d *entry) bool { return d.Key <= target },
			Limit:       limit,
			Order:       DESC,
		})
		var keys2 = make([]string, 0)
		for k := range m {
			if k <= target {
				keys2 = append(keys2, k)
			}
		}
		sort.Strings(keys2)
		utils.ReverseStrings(keys2)
		if len(keys2) > limit {
			keys2 = keys2[:limit]
		}

		if !utils.SameStrings(keys2, algorithm.GetFields(keys1, func(x *entry) string {
			return x.Key
		})) {
			t.Fatal("error!")
		}
	}
}
