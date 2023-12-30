package rbtree

import (
	"fmt"
	"github.com/lxzan/dao/algo"
	"github.com/lxzan/dao/hashmap"
	"github.com/lxzan/dao/internal/utils"
	"github.com/lxzan/dao/internal/validator"
	"github.com/lxzan/dao/types/cmp"
	"github.com/stretchr/testify/assert"
	"sort"
	"strconv"
	"testing"
)

func validate[K cmp.Ordered, V any](t *testing.T, node *rbtree_node[K, V]) {
	if node == nil {
		return
	}
	if node.left != nil {
		if node.left.data != nil && node.data.Key < node.left.data.Key {
			t.Error("left node error!")
		}
		validate(t, node.left)
	}

	if node.right != nil {
		if node.right.data != nil && node.data.Key > node.right.data.Key {
			t.Error("right node error!")
		}
		validate(t, node.right)
	}
}

func TestRBTree_Delete(t *testing.T) {
	var tree = New[string, uint8]()
	var keys []string
	for i := 0; i < 1000; i++ {
		key := utils.Alphabet.Generate(16)
		keys = append(keys, key)
		tree.Set(key, 1)
	}
	for _, key := range keys {
		tree.Delete(key)
	}
	assert.Equal(t, tree.Len(), 0)
}

func TestNew(t *testing.T) {
	var tree = New[string, int]()
	var m = make(map[string]int)

	for i := 0; i < 1000; i++ {
		var length = utils.Rand.Intn(16) + 1
		var key = utils.Numeric.Generate(length)
		tree.Set(key, length)
		m[key] = length
	}

	var idx = 0
	for k := range m {
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
		tree.Set(key, length)
		m[key] = length
	}

	for k, v := range m {
		result, exist := tree.Get(k)
		if !exist || result != v {
			t.Fatal("error!")
		}
	}

	if len(m) != tree.Len() {
		t.Fatal("error!")
	}

	validate(t, tree.root)
}

func TestRBTree_Get(t *testing.T) {
	var tree = New[int, string]()
	var m = make(map[int]string)

	var test_count = 1000
	for i := 0; i < test_count; i++ {
		var key = utils.Rand.Intn(test_count)
		var val = utils.Alphabet.Generate(8)
		tree.Set(key, val)
		m[key] = val
	}

	for i := 0; i < test_count; i++ {
		var key = utils.Rand.Intn(test_count)
		result, exist := tree.Get(key)
		v, ok := m[key]
		if exist != ok || (ok && result != v) {
			t.Fatal("error!")
		}
	}
}

func TestRBTree_ForEach(t *testing.T) {
	var tree = New[string, int]()
	var arr = make([]string, 0)

	for i := 0; i < 100; i++ {
		var key = utils.Alphabet.Generate(16)
		arr = append(arr, key)
		tree.Set(key, utils.Rand.Intn(1000))
	}

	var arr1 = make([]string, 0)
	tree.Range(func(key string, value int) bool {
		arr1 = append(arr1, key)
		return len(arr1) < 50
	})
	if len(arr1) != 50 {
		t.Fatal("error!")
	}

	var arr2 = make([]string, 0)
	tree.Range(func(key string, value int) bool {
		arr2 = append(arr2, key)
		return true
	})

	assert.Equal(t, len(arr2), tree.Len())
	assert.ElementsMatch(t, arr, arr2)
}

func TestRBTree_Between(t *testing.T) {
	t.Run("desc", func(t *testing.T) {
		var tree = New[string, int]()
		var m = make(map[string]int)
		for i := 0; i < 10000; i++ {
			var length = utils.Rand.Intn(16) + 1
			var key = utils.Numeric.Generate(4)
			m[key] = length
			tree.Set(key, length)
		}

		var limit = 100
		for i := 0; i < 100; i++ {
			var left = utils.Numeric.Generate(4)
			x, _ := strconv.Atoi(left)

			var right = fmt.Sprintf("%04d", x+limit)
			if left > right {
				right, left = left, right
			}
			var values = tree.
				NewQuery().
				Left(func(key string) bool { return key >= left }).
				Right(func(key string) bool { return key <= right }).
				Order(DESC).
				Limit(limit).
				FindAll()
			var keys1 = algo.Map[Pair[string, int], string](values, func(i int, v Pair[string, int]) string {
				return v.Key
			})

			var keys2 = make([]string, 0)
			for k := range m {
				if k >= left && k <= right {
					keys2 = append(keys2, k)
				}
			}
			sort.Strings(keys2)
			algo.Reverse(keys2)
			if len(keys2) > limit {
				keys2 = keys2[:limit]
			}

			assert.True(t, utils.IsSameSlice(keys1, keys2))
		}
	})

	t.Run("asc", func(t *testing.T) {
		var tree = New[string, int]()
		var m = make(map[string]int)
		for i := 0; i < 10000; i++ {
			var length = utils.Rand.Intn(16) + 1
			var key = utils.Numeric.Generate(4)
			m[key] = length
			tree.Set(key, length)
		}

		var limit = 100
		for i := 0; i < 100; i++ {
			var left = utils.Numeric.Generate(4)
			x, _ := strconv.Atoi(left)

			var right = fmt.Sprintf("%04d", x+limit)
			if left > right {
				right, left = left, right
			}
			var values = tree.
				NewQuery().
				Left(func(key string) bool { return key >= left }).
				Right(func(key string) bool { return key <= right }).
				Order(ASC).
				Limit(limit).
				Offset(10).
				FindAll()
			var keys1 = algo.Map[Pair[string, int], string](values, func(i int, v Pair[string, int]) string {
				return v.Key
			})

			var keys2 = make([]string, 0)
			for k := range m {
				if k >= left && k <= right {
					keys2 = append(keys2, k)
				}
			}
			sort.Strings(keys2)
			if len(keys2) > 10 {
				keys2 = keys2[10:]
			} else {
				keys2 = keys2[:0]
			}
			if len(keys2) > limit {
				keys2 = keys2[:limit]
			}

			assert.True(t, utils.IsSameSlice(keys1, keys2))
		}
	})

	t.Run("", func(t *testing.T) {
		var tree = New[int, uint8]()
		tree.Set(1, 1)
		tree.Set(2, 1)
		tree.Set(3, 1)
		tree.Set(4, 1)
		tree.Set(5, 1)

		var values0 = tree.
			NewQuery().
			Left(func(key int) bool { return key >= 1 }).
			Right(func(key int) bool { return key <= 3 }).
			Order(ASC).
			Limit(10).
			Offset(5).
			FindAll()
		assert.Equal(t, len(values0), 0)

		var values1 = tree.
			NewQuery().
			Left(func(key int) bool { return key >= 1 }).
			Right(func(key int) bool { return key <= 3 }).
			Order(ASC).
			Limit(10).
			FindAll()
		var keys1 = algo.Map(values1, func(i int, v Pair[int, uint8]) int { return v.Key })
		assert.True(t, utils.IsSameSlice(keys1, []int{1, 2, 3}))
	})
}

func TestRBTree_GreaterEqual(t *testing.T) {
	var tree = New[string, int]()
	var m = make(map[string]int)
	for i := 0; i < 10000; i++ {
		var length = utils.Rand.Intn(16) + 1
		var key = utils.Numeric.Generate(4)
		m[key] = length
		tree.Set(key, length)
	}

	var limit = 100
	for i := 0; i < 100; i++ {
		var left = utils.Numeric.Generate(4)
		var values = tree.
			NewQuery().
			Left(func(key string) bool { return key >= left }).
			Limit(limit).
			FindAll()
		var keys1 = algo.Map[Pair[string, int], string](values, func(i int, v Pair[string, int]) string {
			return v.Key
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

		assert.True(t, utils.IsSameSlice(keys1, keys2))
	}
}

func TestRBTree_LessEqual(t *testing.T) {
	var tree = New[string, int]()
	var m = make(map[string]int)
	for i := 0; i < 10000; i++ {
		var length = utils.Rand.Intn(16) + 1
		var key = utils.Numeric.Generate(4)
		m[key] = length
		tree.Set(key, length)
	}

	var limit = 10
	for i := 0; i < 100; i++ {
		var target = utils.Numeric.Generate(4)
		var results = tree.
			NewQuery().
			Right(func(key string) bool { return key <= target }).
			Order(DESC).
			Limit(limit).
			FindAll()
		var keys1 = algo.Map[Pair[string, int], string](results, func(i int, v Pair[string, int]) string {
			return v.Key
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

		assert.True(t, utils.IsSameSlice(keys1, keys2))
	}
}

func TestRBTree_FindOne(t *testing.T) {
	t.Run("desc", func(t *testing.T) {
		var tree = New[string, int]()
		var m = hashmap.New[string, int](0)
		for i := 0; i < 10000; i++ {
			var length = utils.Rand.Intn(16) + 1
			var key = utils.Numeric.Generate(4)
			m.Set(key, length)
			tree.Set(key, length)
		}

		for i := 0; i < 100; i++ {
			var target = utils.Numeric.Generate(4)
			v0, ok0 := tree.
				NewQuery().
				Right(func(key string) bool { return key <= target }).
				Order(DESC).
				FindOne()

			var v1, ok1 = "", false
			m.Range(func(key string, val int) bool {
				if key <= target && (v1 == "" || key > v1) {
					v1 = key
					ok1 = true
				}
				return true
			})

			assert.Equal(t, ok0, ok1)
			if ok0 {
				assert.Equal(t, v0.Key, v1)
			}
		}
	})

	t.Run("asc", func(t *testing.T) {
		var tree = New[string, int]()
		var m = hashmap.New[string, int](0)
		for i := 0; i < 10000; i++ {
			var length = utils.Rand.Intn(16) + 1
			var key = utils.Numeric.Generate(4)
			m.Set(key, length)
			tree.Set(key, length)
		}

		for i := 0; i < 100; i++ {
			var target = utils.Numeric.Generate(4)
			v0, ok0 := tree.
				NewQuery().
				Left(func(key string) bool { return key >= target }).
				Order(ASC).
				FindOne()

			var v1, ok1 = "", false
			m.Range(func(key string, val int) bool {
				if key >= target && (v1 == "" || key < v1) {
					v1 = key
					ok1 = true
				}
				return true
			})

			assert.Equal(t, ok0, ok1)
			if ok0 {
				assert.Equal(t, v0.Key, v1)
			}
		}
	})

	t.Run("", func(t *testing.T) {
		var tree = New[string, int]()
		var qb = QueryBuilder[string, int]{tree: tree}
		_, ok := qb.FindOne()
		assert.False(t, ok)
	})
}

func TestDict_Map(t *testing.T) {
	assert.True(t, validator.ValidateMapImpl(New[string, int]()))
}

func TestRBTree_NewQuery(t *testing.T) {
	var tree = New[int, uint8]()
	tree.Set(1, 1)
	tree.Set(2, 1)
	tree.Set(5, 1)
	tree.Set(2, 1)
	tree.Set(4, 1)
	tree.Set(6, 1)

	t.Run("", func(t *testing.T) {
		var results = tree.
			NewQuery().
			Left(func(key int) bool { return key > 10 }).
			FindAll()
		assert.Equal(t, len(results), 0)
	})

	t.Run("", func(t *testing.T) {
		var results = tree.
			NewQuery().
			Left(func(key int) bool { return key > 10 }).
			Order(DESC).
			FindAll()
		assert.Equal(t, len(results), 0)
	})
}
