package rbtree

import (
	"fmt"
	"github.com/lxzan/dao"
	"github.com/lxzan/dao/algorithm"
	"github.com/lxzan/dao/internal/utils"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sort"
	"strconv"
	"testing"
)

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

	tree.validate(t, tree.root)
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
		var keys1 = tree.
			NewQuery().
			Left(func(key string) bool { return key >= left }).
			Right(func(key string) bool { return key <= right }).
			Order(dao.DESC).
			Limit(limit).
			Do()
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

		if !utils.IsSameSlice(keys2, algorithm.GetFields(keys1, func(x *Pair[string, int]) string {
			return x.Key
		})) {
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
		tree.Set(key, length)
	}

	var limit = 100
	for i := 0; i < 100; i++ {
		var left = utils.Numeric.Generate(4)
		var keys1 = tree.
			NewQuery().
			Left(func(key string) bool { return key >= left }).
			Limit(limit).
			Do()
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

		if !utils.IsSameSlice(keys2, algorithm.GetFields(keys1, func(x *Pair[string, int]) string {
			return x.Key
		})) {
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
		tree.Set(key, length)
	}

	var limit = 100
	for i := 0; i < 100; i++ {
		var target = utils.Numeric.Generate(4)
		var keys1 = tree.
			NewQuery().
			Right(func(key string) bool { return key <= target }).
			Order(dao.DESC).
			Limit(limit).
			Do()
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

		if !utils.IsSameSlice(keys2, algorithm.GetFields(keys1, func(x *Pair[string, int]) string {
			return x.Key
		})) {
			t.Fatal("error!")
		}
	}
}

func TestRBTree_GetMinKey(t *testing.T) {
	var tree = New[string, int]()

	const test_count = 100
	for i := 0; i < test_count; i++ {
		var v = rand.Intn(10000)
		tree.Set(strconv.Itoa(v), v)
	}

	for i := 0; i < test_count; i++ {
		var k = strconv.Itoa(rand.Intn(10000))
		result, exist := tree.GetMinKey(func(key string) bool {
			return key >= k
		})

		if !exist {
			tree.Range(func(key string, value int) bool {
				if key >= k {
					t.Fatal("error!")
				}
				return true
			})
		} else {
			tree.Range(func(key string, value int) bool {
				if key < result.Key && key >= k {
					t.Fatal("error!")
				}
				return true
			})
		}
	}
}

func TestRBTree_GetMaxKey(t *testing.T) {
	var tree = New[string, int]()

	const test_count = 100
	for i := 0; i < test_count; i++ {
		var v = rand.Intn(10000)
		tree.Set(strconv.Itoa(v), v)
	}

	for i := 0; i < test_count; i++ {
		var k = strconv.Itoa(rand.Intn(10000))
		result, exist := tree.GetMaxKey(func(key string) bool {
			return key <= k
		})

		if !exist {
			tree.Range(func(key string, value int) bool {
				if key <= k {
					t.Fatal("error!")
				}
				return true
			})
		} else {
			tree.Range(func(key string, value int) bool {
				if key > result.Key && key <= k {
					t.Fatal("error!")
				}
				return true
			})
		}
	}
}
