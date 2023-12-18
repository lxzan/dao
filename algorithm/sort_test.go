package algorithm

import (
	"cmp"
	"github.com/lxzan/dao"
	"github.com/lxzan/dao/internal/utils"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func asc[T cmp.Ordered](a, b T) dao.CompareResult {
	if a < b {
		return dao.Less
	} else if a > b {
		return dao.Greater
	} else {
		return dao.Equal
	}
}

func desc[T cmp.Ordered](a, b T) dao.CompareResult {
	return -1 * asc(a, b)
}

func TestGetMedium(t *testing.T) {
	var as = assert.New(t)
	{
		arr := []int{1, 2, 3}
		idx := getMedium(arr, 0, 2, asc[int])
		as.Equal(2, arr[idx])
	}
	{
		arr := []int{1, 3, 2}
		idx := getMedium(arr, 0, 2, asc[int])
		as.Equal(2, arr[idx])
	}
	{
		arr := []int{5, 2, 3}
		idx := getMedium(arr, 0, 2, asc[int])
		as.Equal(3, arr[idx])
	}
	{
		arr := []int{3, 5, 1}
		idx := getMedium(arr, 0, 2, asc[int])
		as.Equal(3, arr[idx])
	}
}

func TestIsSorted(t *testing.T) {
	var as = assert.New(t)
	as.Equal(true, IsSorted([]int{1, 2, 3}, asc[int]))
	as.Equal(true, IsSorted([]int{1, 2, 3, 4}, asc[int]))
	as.Equal(true, IsSorted([]int{}, asc[int]))
	as.Equal(true, IsSorted([]int{1}, asc[int]))
	as.Equal(true, IsSorted([]int{1, 2, 2, 2}, asc[int]))
	as.Equal(true, IsSorted([]int{3, 2, 1}, desc[int]))

	as.Equal(false, IsSorted([]int{1, 3, 2}, asc[int]))
	as.Equal(false, IsSorted([]int{1, 2, 3, 2}, asc[int]))
	as.Equal(false, IsSorted([]int{1, 2, 2, 1}, asc[int]))
	as.Equal(false, IsSorted([]int{3, 2, 1}, asc[int]))
	as.Equal(false, IsSorted([]int{3, 2, 1, 0}, asc[int]))
}

func TestSort(t *testing.T) {
	var arr = make([]int, 0)
	for i := 0; i < 999; i++ {
		arr = append(arr, rand.Intn(1000))
	}
	Sort(arr, asc[int])

	if !IsSorted(arr, asc[int]) {
		t.Error("not sorted!")
	}

	t.Run("", func(t *testing.T) {
		var a = []int{1, 2, 3}
		Sort(a, asc[int])
		assert.True(t, utils.IsSameSlice(a, []int{1, 2, 3}))
	})
}

func TestBinarySearch(t *testing.T) {
	t.Run("", func(t *testing.T) {
		var list []int
		var index = BinarySearch(list, 1, asc[int])
		assert.Equal(t, index, -1)
	})

	t.Run("", func(t *testing.T) {
		var list = []int{1, 2, 3}
		var index = BinarySearch(list, 3, asc[int])
		assert.Equal(t, index, 2)

		index = BinarySearch(list, 1, asc[int])
		assert.Equal(t, index, 0)
	})

	t.Run("", func(t *testing.T) {
		const count = 1000
		var m = make(map[int]uint8)
		var arr []int
		for i := 0; i < count; i++ {
			v := rand.Intn(10000)
			arr = append(arr, v)
			m[v] = 1
		}

		Sort(arr, asc[int])
		for i := 0; i < count; i++ {
			k := rand.Intn(10000)
			_, ok := m[k]
			index := BinarySearch(arr, k, asc[int])
			if ok {
				assert.Equal(t, arr[index], k)
			} else {
				assert.Equal(t, index, -1)
			}
		}
	})
}
