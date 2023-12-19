package algorithm

import (
	"cmp"
	"github.com/lxzan/dao/internal/utils"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func desc[T cmp.Ordered](a, b T) int {
	return -1 * cmp.Compare(a, b)
}

func TestGetMedium(t *testing.T) {
	var as = assert.New(t)
	{
		arr := []int{1, 2, 3}
		idx := getMedium(arr, 0, 2, cmp.Compare[int])
		as.Equal(2, arr[idx])
	}
	{
		arr := []int{1, 3, 2}
		idx := getMedium(arr, 0, 2, cmp.Compare[int])
		as.Equal(2, arr[idx])
	}
	{
		arr := []int{5, 2, 3}
		idx := getMedium(arr, 0, 2, cmp.Compare[int])
		as.Equal(3, arr[idx])
	}
	{
		arr := []int{3, 5, 1}
		idx := getMedium(arr, 0, 2, cmp.Compare[int])
		as.Equal(3, arr[idx])
	}
}

func TestIsSorted(t *testing.T) {
	var as = assert.New(t)
	as.Equal(true, IsSorted([]int{1, 2, 3}, cmp.Compare[int]))
	as.Equal(true, IsSorted([]int{1, 2, 3, 4}, cmp.Compare[int]))
	as.Equal(true, IsSorted([]int{}, cmp.Compare[int]))
	as.Equal(true, IsSorted([]int{1}, cmp.Compare[int]))
	as.Equal(true, IsSorted([]int{1, 2, 2, 2}, cmp.Compare[int]))
	as.Equal(true, IsSorted([]int{3, 2, 1}, desc[int]))

	as.Equal(false, IsSorted([]int{1, 3, 2}, cmp.Compare[int]))
	as.Equal(false, IsSorted([]int{1, 2, 3, 2}, cmp.Compare[int]))
	as.Equal(false, IsSorted([]int{1, 2, 2, 1}, cmp.Compare[int]))
	as.Equal(false, IsSorted([]int{3, 2, 1}, cmp.Compare[int]))
	as.Equal(false, IsSorted([]int{3, 2, 1, 0}, cmp.Compare[int]))
}

func TestSort(t *testing.T) {
	var arr = make([]int, 0)
	for i := 0; i < 999; i++ {
		arr = append(arr, rand.Intn(1000))
	}
	SortBy(arr, cmp.Compare[int])

	if !IsSorted(arr, cmp.Compare[int]) {
		t.Error("not sorted!")
	}

	t.Run("", func(t *testing.T) {
		var a = []int{1, 2, 3}
		SortBy(a, cmp.Compare[int])
		assert.True(t, utils.IsSameSlice(a, []int{1, 2, 3}))
	})

	t.Run("", func(t *testing.T) {
		var a = []int{1, 3, 5, 2, 4, 6}
		Sort(a)
		assert.True(t, utils.IsSameSlice(a, []int{1, 2, 3, 4, 5, 6}))
	})

	t.Run("", func(t *testing.T) {
		var a = []int{1, 2, 3, 4}
		Sort(a)
		assert.True(t, utils.IsSameSlice(a, []int{1, 2, 3, 4}))
	})
}

func TestBinarySearch(t *testing.T) {
	t.Run("", func(t *testing.T) {
		var list []int
		var index = BinarySearch(list, 1, cmp.Compare[int])
		assert.Equal(t, index, -1)
	})

	t.Run("", func(t *testing.T) {
		var list = []int{1, 2, 3}
		var index = BinarySearch(list, 3, cmp.Compare[int])
		assert.Equal(t, index, 2)

		index = BinarySearch(list, 1, cmp.Compare[int])
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

		SortBy(arr, cmp.Compare[int])
		for i := 0; i < count; i++ {
			k := rand.Intn(10000)
			_, ok := m[k]
			index := BinarySearch(arr, k, cmp.Compare[int])
			if ok {
				assert.Equal(t, arr[index], k)
			} else {
				assert.Equal(t, index, -1)
			}
		}
	})
}
