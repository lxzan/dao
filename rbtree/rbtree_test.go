package rbtree

import (
	"fmt"
	"github.com/lxzan/dao/algorithm"
	"github.com/lxzan/dao/internal/utils"
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
		tree.Insert(&Iterator[string, int]{Key: key, Val: length})
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
		tree.Insert(&Iterator[string, int]{Key: key, Val: length})
		m[key] = length
	}

	for k, v := range m {
		result, exist := tree.Find(k)
		if !exist || result != v {
			t.Fatal("error!")
		}
	}

	if len(m) != tree.Len() {
		t.Fatal("error!")
	}

	tree.validate(t, tree.root)
}

func TestRBTree_Find(t *testing.T) {
	var tree = New[int, string]()
	var m = make(map[int]string)

	var test_count = 1000
	for i := 0; i < test_count; i++ {
		var key = utils.Rand.Intn(test_count)
		var val = utils.Alphabet.Generate(8)
		tree.Insert(&Iterator[int, string]{Key: key, Val: val})
		if _, ok := m[key]; !ok {
			m[key] = val
		}
	}

	for i := 0; i < test_count; i++ {
		var key = utils.Rand.Intn(test_count)
		result, exist := tree.Find(key)
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
		tree.Insert(&Iterator[string, int]{Key: key, Val: utils.Rand.Intn(1000)})
	}

	var arr1 = make([]string, 0)
	tree.ForEach(func(iter *Iterator[string, int]) {
		arr1 = append(arr1, iter.Key)
		if len(arr1) >= 50 {
			iter.Break()
		}
	})
	if len(arr1) != 50 {
		t.Fatal("error!")
	}

	var arr2 = make([]string, 0)
	tree.ForEach(func(iter *Iterator[string, int]) {
		arr2 = append(arr2, iter.Key)
	})

	if len(arr2) != tree.Len() || !utils.SameStrings(arr, arr2) {
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
		tree.Insert(&Iterator[string, int]{Key: key, Val: length})
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
			LeftFilter:  func(d string) bool { return d >= left },
			RightFilter: func(d string) bool { return d <= right },
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

		if !utils.SameStrings(keys2, algorithm.GetFields(keys1, func(x *Iterator[string, int]) string {
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
		tree.Insert(&Iterator[string, int]{Key: key, Val: length})
	}

	var limit = 100
	for i := 0; i < 100; i++ {
		var left = utils.Numeric.Generate(4)
		var keys1 = tree.Query(&QueryBuilder[string]{
			LeftFilter: func(d string) bool { return d >= left },
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

		if !utils.SameStrings(keys2, algorithm.GetFields(keys1, func(x *Iterator[string, int]) string {
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
		tree.Insert(&Iterator[string, int]{Key: key, Val: length})
	}

	var limit = 100
	for i := 0; i < 100; i++ {
		var target = utils.Numeric.Generate(4)
		var keys1 = tree.Query(&QueryBuilder[string]{
			RightFilter: func(d string) bool { return d <= target },
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

		if !utils.SameStrings(keys2, algorithm.GetFields(keys1, func(x *Iterator[string, int]) string {
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
		tree.Insert(&Iterator[string, int]{Key: strconv.Itoa(v), Val: v})
	}

	for i := 0; i < test_count; i++ {
		var k = strconv.Itoa(rand.Intn(10000))
		result, exist := tree.GetMinKey(func(key string) bool {
			return key >= k
		})

		if !exist {
			tree.ForEach(func(iter *Iterator[string, int]) {
				if iter.Key >= k {
					t.Fatal("error!")
				}
			})
		} else {
			tree.ForEach(func(iter *Iterator[string, int]) {
				if iter.Key < result.Key && iter.Key >= k {
					t.Fatal("error!")
				}
			})
		}
	}
}

func TestRBTree_GetMaxKey(t *testing.T) {
	var tree = New[string, int]()

	const test_count = 100
	for i := 0; i < test_count; i++ {
		var v = rand.Intn(10000)
		tree.Insert(&Iterator[string, int]{Key: strconv.Itoa(v), Val: v})
	}

	for i := 0; i < test_count; i++ {
		var k = strconv.Itoa(rand.Intn(10000))
		result, exist := tree.GetMaxKey(func(key string) bool {
			return key <= k
		})

		if !exist {
			tree.ForEach(func(iter *Iterator[string, int]) {
				if iter.Key <= k {
					t.Fatal("error!")
				}
			})
		} else {
			tree.ForEach(func(iter *Iterator[string, int]) {
				if iter.Key > result.Key && iter.Key <= k {
					t.Fatal("error!")
				}
			})
		}
	}
}
