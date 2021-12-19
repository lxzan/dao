package rbtree

import (
	"fmt"
	"github.com/lxzan/dao/internal/utils"
	"sort"
	"strconv"
	"testing"
)

func (c *RBTree[K, V]) validate(t *testing.T, node *rbtree_node[K, V]) {
	if node == nil {
		return
	}
	if node.left != nil {
		if !c.is_key_empty(node.left.key) && !(node.key > node.left.key) {
			t.Error("left node error!")
		}
		c.validate(t, node.left)
	}

	if node.right != nil {
		if !c.is_key_empty(node.right.key) && !(node.key < node.right.key) {
			t.Error("right node error!")
		}
		c.validate(t, node.right)
	}
}

func TestNew(t *testing.T) {
	var tree = New[string, int]()
	var m = make(map[string]int)

	for i := 0; i < 1000; i++ {
		var length = utils.Rand.Intn(16) + 1
		var key = utils.Alphabet.Generate(length)
		tree.Insert(key, length)
		m[key] = length
	}

	var idx = 0
	for k, _ := range m {
		if idx >= 500 {
			break
		}
		delete(m, k)
		tree.Delete(k)
		idx++
	}

	for i := 0; i < 10000; i++ {
		var length = utils.Rand.Intn(16) + 1
		var key = utils.Alphabet.Generate(length)
		tree.Insert(key, length)
		m[key] = length
	}

	for k, v := range m {
		result, exist := tree.Find(k)
		if !exist || *result != v {
			t.Fatal("error!")
		}
	}

	if len(m) != tree.Len() {
		t.Fatal("error!")
	}

	tree.validate(t, tree.root)
}

func TestRBTree_ForEach(t *testing.T) {
	var tree = New[string, int]()
	for i := 0; i < 100; i++ {
		tree.Insert(utils.Alphabet.Generate(16), utils.Rand.Intn(1000))
	}

	var arr1 = make([]string, 0)
	tree.ForEach(func(key string, val *int) (continued bool) {
		arr1 = append(arr1, key)
		return len(arr1) < 50
	})
	if len(arr1) != 50 {
		t.Fatal("error!")
	}

	var arr2 = make([]string, 0)
	tree.ForEach(func(key string, val *int) (continued bool) {
		arr2 = append(arr2, key)
		return true
	})
	if len(arr2) != tree.Len() {
		t.Fatal("error!")
	}
}

func TestRBTree_Between(t *testing.T) {
	var tree = New[string, int]()
	var m = make(map[string]int)
	for i := 0; i < 10000; i++ {
		var length = utils.Rand.Intn(16) + 1
		var key = utils.Numeric.Generate(4)
		m[key] = length
		tree.Insert(key, length)
	}

	var limit = 100
	for i := 0; i < 100; i++ {
		var left = utils.Numeric.Generate(4)
		x, _ := strconv.Atoi(left)
		var right = fmt.Sprintf("%04d", x+limit)
		if left > right {
			right, left = left, right
		}
		var keys1 = tree.Query(&QueryBuilder[string]{
			LeftFilter:  func(key string) bool { return key >= left },
			RightFilter: func(key string) bool { return key <= right },
			Limit:       limit,
			Order:       ASC,
		})
		var keys2 = make([]string, 0)
		for k, _ := range m {
			if k >= left && k <= right {
				keys2 = append(keys2, k)
			}
		}
		sort.Strings(keys2)
		if len(keys2) > limit {
			keys2 = keys2[:limit]
		}

		if !utils.SameStrings(keys1, keys2) {
			t.Fatal("error!")
		}
	}
}

func TestRBTree_GreaterEqual(t *testing.T) {
	var tree = New[string, int]()
	var m = make(map[string]int)
	for i := 0; i < 10000; i++ {
		var length = utils.Rand.Intn(16) + 1
		var key = utils.Numeric.Generate(4)
		m[key] = length
		tree.Insert(key, length)
	}

	var limit = 100
	for i := 0; i < 100; i++ {
		var left = utils.Numeric.Generate(4)
		var keys1 = tree.Query(&QueryBuilder[string]{
			LeftFilter: func(key string) bool { return key >= left },
			Limit:      limit,
		})
		var keys2 = make([]string, 0)
		for k, _ := range m {
			if k >= left {
				keys2 = append(keys2, k)
			}
		}
		sort.Strings(keys2)
		if len(keys2) > limit {
			keys2 = keys2[:limit]
		}

		if !utils.SameStrings(keys1, keys2) {
			t.Fatal("error!")
		}
	}
}

func TestRBTree_LessEqual(t *testing.T) {
	var tree = New[string, int]()
	var m = make(map[string]int)
	for i := 0; i < 10000; i++ {
		var length = utils.Rand.Intn(16) + 1
		var key = utils.Numeric.Generate(4)
		m[key] = length
		tree.Insert(key, length)
	}

	var limit = 100
	for i := 0; i < 100; i++ {
		var target = utils.Numeric.Generate(4)
		var keys1 = tree.Query(&QueryBuilder[string]{
			RightFilter: func(key string) bool { return key <= target },
			Limit:       limit,
			Order:       DESC,
		})
		var keys2 = make([]string, 0)
		for k, _ := range m {
			if k <= target {
				keys2 = append(keys2, k)
			}
		}
		sort.Strings(keys2)
		utils.ReverseStrings(keys2)
		if len(keys2) > limit {
			keys2 = keys2[:limit]
		}

		if !utils.SameStrings(keys1, keys2) {
			t.Fatal("error!")
		}
	}
}
