package algorithm

import (
	"github.com/lxzan/dao/hashmap"
	"github.com/lxzan/dao/internal/utils"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"net/http"
	"testing"
)

func TestToString(t *testing.T) {
	if ToString(123) != "123" {
		t.Fatal("error!")
	}
}

func TestUnique(t *testing.T) {
	Unique[int, []int](nil)

	t.Run("", func(t *testing.T) {
		arr := Unique([]int{})
		assert.ElementsMatch(t, arr, []int{})
	})

	t.Run("", func(t *testing.T) {
		arr := Unique([]int{1, 3, 5, 3})
		assert.ElementsMatch(t, arr, []int{1, 3, 5})
	})

	t.Run("", func(t *testing.T) {
		arr := Unique([]int{1})
		assert.ElementsMatch(t, arr, []int{1})
	})

	t.Run("", func(t *testing.T) {
		arr := Unique([]int{1, 3, 3, 5, 5, 3, 2})
		assert.ElementsMatch(t, arr, []int{1, 2, 3, 5})
	})

	t.Run("", func(t *testing.T) {
		var m = make(map[int]uint8)
		var arr1 []int
		var arr2 []int
		var arr3 []int
		for i := 0; i < 1000; i++ {
			v := rand.Intn(1000)
			arr1 = append(arr1, v)
			m[v] = 1
		}

		for k, _ := range m {
			arr2 = append(arr2, k)
		}
		arr3 = Unique(arr1)
		assert.ElementsMatch(t, arr2, arr3)
	})
}

func TestUniqueBy(t *testing.T) {
	t.Run("", func(t *testing.T) {
		var reqs = []http.Request{}
		arr := UniqueBy[http.Request, int](reqs, func(item http.Request) int {
			return item.ProtoMajor
		})
		children := Map[http.Request, int](arr, func(i int, x http.Request) int {
			return x.ProtoMajor
		})
		assert.ElementsMatch(t, children, []int{})
	})

	t.Run("", func(t *testing.T) {
		var reqs = []http.Request{
			{ProtoMajor: 1},
			{ProtoMajor: 3},
			{ProtoMajor: 5},
			{ProtoMajor: 3},
		}
		arr := UniqueBy[http.Request, int](reqs, func(item http.Request) int {
			return item.ProtoMajor
		})
		children := Map[http.Request, int](arr, func(i int, x http.Request) int {
			return x.ProtoMajor
		})
		assert.ElementsMatch(t, children, []int{1, 3, 5})
	})

	t.Run("", func(t *testing.T) {
		var reqs = []http.Request{
			{ProtoMajor: 1},
		}
		arr := UniqueBy[http.Request, int](reqs, func(item http.Request) int {
			return item.ProtoMajor
		})
		children := Map[http.Request, int](arr, func(i int, x http.Request) int {
			return x.ProtoMajor
		})
		assert.ElementsMatch(t, children, []int{1})
	})

	t.Run("", func(t *testing.T) {
		var reqs = []http.Request{
			{ProtoMajor: 1},
			{ProtoMajor: 3},
			{ProtoMajor: 3},
			{ProtoMajor: 5},
			{ProtoMajor: 5},
			{ProtoMajor: 3},
			{ProtoMajor: 2},
		}
		arr := UniqueBy[http.Request, int](reqs, func(item http.Request) int {
			return item.ProtoMajor
		})
		children := Map[http.Request, int](arr, func(i int, x http.Request) int {
			return x.ProtoMajor
		})
		assert.ElementsMatch(t, children, []int{1, 2, 3, 5})
	})
}

func TestMin(t *testing.T) {
	assert.Equal(t, Min(1, 2), 1)
	assert.Equal(t, Min(2, 1), 1)
}

func TestMax(t *testing.T) {
	assert.Equal(t, Max(1, 2), 2)
	assert.Equal(t, Max(2, 1), 2)
}

func TestSwap(t *testing.T) {
	var a, b = 1, 2
	Swap(&a, &b)
	assert.Equal(t, a, 2)
	assert.Equal(t, b, 1)
}

func TestReverse(t *testing.T) {
	Reverse[int, []int](nil)

	t.Run("", func(t *testing.T) {
		var list = []int{1, 2, 3, 4}
		Reverse(list)
		assert.True(t, utils.IsSameSlice(list, []int{4, 3, 2, 1}))
	})

	t.Run("", func(t *testing.T) {
		var list = []int{1}
		Reverse(list)
		assert.True(t, utils.IsSameSlice(list, []int{1}))
	})

	t.Run("", func(t *testing.T) {
		var list = []int{}
		Reverse(list)
		assert.True(t, utils.IsSameSlice(list, []int{}))
	})
}

func TestSelectValue(t *testing.T) {
	assert.Equal(t, SelectValue(true, 1, 2), 1)
	assert.Equal(t, SelectValue(false, 1, 2), 2)
}

func TestContains(t *testing.T) {
	assert.True(t, Contains([]int{1, 2}, 1))
	assert.True(t, Contains([]int{1, 2}, 2))
	assert.False(t, Contains([]int{1, 2}, 3))
}

func TestFilter(t *testing.T) {
	var arr = []int{1, 2, 3, 4}
	arr = Filter(arr, func(i int, item int) bool {
		return item%2 == 0
	})
	assert.ElementsMatch(t, arr, []int{2, 4})
}

func TestIsZero(t *testing.T) {
	assert.True(t, IsZero(0))
	assert.True(t, IsZero(""))
	assert.True(t, IsZero(struct{}{}))
	assert.False(t, IsZero(1))
	assert.False(t, IsZero(" "))
}

func TestReduce(t *testing.T) {
	t.Run("", func(t *testing.T) {
		var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		var sum = Reduce(0, arr, func(summarize int, item int) int {
			return summarize + item
		})
		assert.Equal(t, sum, 55)
	})

	t.Run("", func(t *testing.T) {
		var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		var m = hashmap.New[int, struct{}](10)
		Reduce(m, arr, func(s hashmap.HashMap[int, struct{}], item int) hashmap.HashMap[int, struct{}] {
			s.Set(item, struct{}{})
			return s
		})
		assert.ElementsMatch(t, m.Keys(), []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	})
}
