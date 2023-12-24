package heap

import (
	"fmt"
	"github.com/lxzan/dao/algorithm"
	"github.com/lxzan/dao/internal/utils"
	"github.com/lxzan/dao/types/cmp"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"unsafe"
)

func TestIndexedHeap_Sort(t *testing.T) {
	var h = NewIndexedHeap[int, struct{}](Quadratic, nil)
	for i := 0; i < 1000; i++ {
		h.Push(rand.Int(), struct{}{})
	}
	var arr []int
	for h.Len() > 0 {
		arr = append(arr, h.Pop().Key())
	}
	assert.True(t, algorithm.IsSorted(arr, func(a, b int) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		} else {
			return 0
		}
	}))
}

func TestHeap_Random(t *testing.T) {
	t.Run("asc", func(t *testing.T) {
		const count = 10000
		var h = NewIndexedHeap[int, struct{}](Quadratic, cmp.Less[int])
		h.SetCap(count)
		for i := 0; i < count; i++ {
			flag := rand.Intn(5)
			key := rand.Intn(count)
			switch flag {
			case 0, 1:
				h.Push(key, struct{}{})
			case 2:
				h.Pop()
			case 3:
				n := h.Len()
				if n > 0 {
					index := rand.Intn(n)
					h.UpdateKeyByIndex(index, key)
				}
			case 4:
				n := h.Len()
				if n > 0 {
					index := rand.Intn(n)
					h.DeleteByIndex(index)
				}
			}
		}

		for i, item := range h.data {
			assert.Equal(t, item.Index(), i)

			if item.Index() == 0 {
				item = h.Top()
			}
			var n = h.Len()
			var base = i << h.bits
			var end = algorithm.Min(base+h.forks, n-1)
			for j := base + 1; j <= end; j++ {
				assert.True(t, h.lessFunc(item.Key(), h.GetByIndex(j).Key()))
			}
		}
	})

	t.Run("desc", func(t *testing.T) {
		const count = 10000
		var h = NewIndexedHeap[int, struct{}](Quadratic, func(a, b int) bool {
			return a > b
		})
		h.SetCap(count)
		for i := 0; i < count; i++ {
			flag := rand.Intn(5)
			key := rand.Intn(count)
			switch flag {
			case 0, 1:
				h.Push(key, struct{}{})
			case 2:
				h.Pop()
			case 3:
				n := h.Len()
				if n > 0 {
					index := rand.Intn(n)
					h.UpdateKeyByIndex(index, key)
				}
			case 4:
				n := h.Len()
				if n > 0 {
					index := rand.Intn(n)
					h.DeleteByIndex(index)
				}
			}
		}

		for i, item := range h.data {
			assert.Equal(t, item.Index(), i)

			if item.Index() == 0 {
				item = h.Top()
			}
			var n = h.Len()
			var base = i << h.bits
			var end = algorithm.Min(base+h.forks, n-1)
			for j := base + 1; j <= end; j++ {
				assert.True(t, h.lessFunc(item.Key(), h.GetByIndex(j).Key()))
			}
		}
	})
}

func TestIndexedHeap_Range(t *testing.T) {
	var h = NewIndexedHeap[int, struct{}](Quadratic, cmp.Less[int])
	h.SetCap(8)
	h.Push(1, struct{}{})
	h.Push(3, struct{}{})
	h.Push(2, struct{}{})
	h.Push(5, struct{}{})
	h.Push(4, struct{}{})

	{
		var arr []int
		h.Range(func(ele *Element[int, struct{}]) bool {
			arr = append(arr, ele.Key())
			return true
		})
		assert.ElementsMatch(t, arr, []int{1, 2, 3, 4, 5})
	}

	{
		var arr []int
		h.Range(func(ele *Element[int, struct{}]) bool {
			arr = append(arr, ele.Key())
			return len(arr) < 2
		})
		assert.Equal(t, len(arr), 2)
	}
}

func TestIndexedHeap_SetForkNumber(t *testing.T) {
	var catch = func(f func()) (err error) {
		defer func() {
			if excp := recover(); excp != nil {
				err = fmt.Errorf("%v", excp)
			}
		}()
		f()
		return err
	}

	var err1 = catch(func() {
		NewIndexedHeap[int, struct{}](3, cmp.Less[int])
	})
	assert.Error(t, err1)

	var err2 = catch(func() {
		NewWithForks(4, cmp.Less[int])
	})
	assert.Nil(t, err2)
}

func TestIndexedHeap_Clone(t *testing.T) {
	var h = NewIndexedHeap[int, struct{}](4, nil)
	h.Push(1, struct{}{})
	h.Push(3, struct{}{})
	h.Push(4, struct{}{})
	h.Push(4, struct{}{})

	var h1 = h.Clone()
	var h2 = h
	assert.True(t, utils.IsSameSlice(h.data, h1.data))
	var addr = (uintptr)(unsafe.Pointer(&h.data[0]))
	var addr1 = (uintptr)(unsafe.Pointer(&h1.data[0]))
	var addr2 = (uintptr)(unsafe.Pointer(&h2.data[0]))
	assert.NotEqual(t, addr, addr1)
	assert.Equal(t, addr, addr2)

	h1.Reset()
	assert.Equal(t, h1.Len(), 0)
	assert.NotEqual(t, h2.Len(), 0)
}