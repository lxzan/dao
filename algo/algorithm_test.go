package algo

import (
	"errors"
	"github.com/lxzan/dao/hashmap"
	"github.com/lxzan/dao/internal/utils"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"net"
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
	t.Run("", func(t *testing.T) {
		var arr = []int{1, 2, 3, 4}
		arr = Filter(arr, func(i int, item int) bool { return item%2 == 0 })
		assert.ElementsMatch(t, arr, []int{2, 4})
	})

	t.Run("", func(t *testing.T) {
		var arr = []int{1, 2, 3, 4}
		arr = Filter(arr, func(i int, item int) bool { return item%2 == 1 })
		assert.ElementsMatch(t, arr, []int{1, 3})
	})

	t.Run("", func(t *testing.T) {
		var arr = []int{}
		arr = Filter(arr, func(i int, item int) bool { return item%2 == 0 })
		assert.ElementsMatch(t, arr, []int{})
	})
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
		var sum = Reduce(arr, 0, func(summarize int, i int, item int) int {
			return summarize + item
		})
		assert.Equal(t, sum, 55)
	})

	t.Run("", func(t *testing.T) {
		var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		var m = hashmap.New[int, struct{}](10)
		Reduce(arr, m, func(s hashmap.HashMap[int, struct{}], i int, item int) hashmap.HashMap[int, struct{}] {
			s.Set(item, struct{}{})
			return s
		})
		assert.ElementsMatch(t, m.Keys(), []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	})
}

func TestSum(t *testing.T) {
	assert.Equal(t, Sum([]int{}), 0)
	assert.Equal(t, Sum([]int{1, 3, 5, 7, 9}), 25)
	assert.Equal(t, Sum([]uint32{1, 3, 5, 7, 9}), uint32(25))
}

func TestGroupBy(t *testing.T) {
	t.Run("", func(t *testing.T) {
		var arr = []int{1, 3, 5, 7, 2, 4, 6, 8}
		var m = GroupBy(arr, func(i int, v int) int {
			return v % 2
		})
		assert.ElementsMatch(t, m[0], []int{2, 4, 6, 8})
		assert.ElementsMatch(t, m[1], []int{1, 3, 5, 7})
	})

	t.Run("", func(t *testing.T) {
		var arr = []int{1, 3, 5, 7, 2, 4, 6, 8}
		var m = GroupBy(arr, func(i int, v int) int {
			return v % 3
		})
		assert.ElementsMatch(t, m[0], []int{3, 6})
		assert.ElementsMatch(t, m[1], []int{1, 4, 7})
		assert.ElementsMatch(t, m[2], []int{2, 5, 8})
	})
}

func TestIsNil(t *testing.T) {
	var conn1 *net.TCPConn
	var conn2 net.Conn = conn1
	var conn3 = &net.TCPConn{}
	assert.True(t, IsNil(conn1))
	assert.True(t, IsNil(conn2))
	assert.False(t, IsNil(conn3))

	assert.False(t, NotNil(conn1))
	assert.False(t, NotNil(conn2))
	assert.True(t, NotNil(conn3))
}

func TestNotNil(t *testing.T) {
	assert.True(t, NotNil(errors.New("1")))
	assert.False(t, NotNil(nil))
}

func TestWithDefault(t *testing.T) {
	assert.Equal(t, WithDefault(0, 1), 1)
	assert.Equal(t, WithDefault(2, 1), 2)
	assert.Equal(t, WithDefault("", "1"), "1")
	assert.Equal(t, WithDefault("2", "1"), "2")
}

func TestSplitWithCallback(t *testing.T) {
	t.Run("", func(t *testing.T) {
		var path = "/api/v1/greet"
		var values []string
		SplitWithCallback(path, "/", func(index int, item string) bool {
			values = append(values, item)
			return true
		})
		assert.ElementsMatch(t, values, []string{"", "api", "v1", "greet"})
	})

	t.Run("", func(t *testing.T) {
		var path = "/api/v1/greet/"
		var values []string
		SplitWithCallback(path, "/", func(index int, item string) bool {
			values = append(values, item)
			return true
		})
		assert.ElementsMatch(t, values, []string{"", "api", "v1", "greet", ""})
	})

	t.Run("", func(t *testing.T) {
		var path = "/api/v1/greet"
		var values []string
		SplitWithCallback(path, "/", func(index int, item string) bool {
			values = append(values, item)
			return len(values) < 2
		})
		assert.ElementsMatch(t, values, []string{"", "api"})
	})
}

func TestSliceToMap(t *testing.T) {
	var m = SliceToMap([]string{"a", "b", "c"}, func(i int, v string) (string, int) {
		return v, i
	})
	assert.Equal(t, m["a"], 0)
	assert.Equal(t, m["c"], 2)
}
