package heap

import (
	"fmt"
	"github.com/lxzan/dao/algo"
	"github.com/lxzan/dao/internal/utils"
	"github.com/lxzan/dao/types/cmp"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"unsafe"
)

func validateIndexedHeap[K cmp.Ordered, V any](t *testing.T, h *IndexedHeap[K, V], compare cmp.CompareFunc[K]) {
	var n = h.Len()
	if n > 0 {
		assert.Equal(t, h.data[0].Key(), h.Top().Key())
	}

	var i = 0
	h.Range(func(ele *Element[K, V]) bool {
		assert.Equal(t, ele.Index(), i)
		i++

		var base = ele.Index() << h.bits
		var end = algo.Min(base+h.ways, n-1)
		for j := base + 1; j <= end; j++ {
			child := h.GetByIndex(j)
			assert.True(t, compare(ele.Key(), child.Key()) <= 0)
		}
		return true
	})

	var keys = make([]K, 0, n)
	for h.Len() > 0 {
		keys = append(keys, h.Pop().Key())
	}
	assert.True(t, algo.IsSorted(keys, compare))
}

func TestIndexedHeap_Random(t *testing.T) {
	const count = 10000

	var f = func(ways uint32, lessFunc cmp.LessFunc[int], compareFunc cmp.CompareFunc[int]) {
		var h = NewIndexedHeap[int, struct{}](ways, lessFunc)
		h.SetCap(count)
		for i := 0; i < count; i++ {
			flag := utils.Alphabet.Intn(6)
			key := rand.Intn(count)
			switch flag {
			case 0, 1, 2:
				h.Push(key, struct{}{})
			case 3:
				h.Pop()
			case 4:
				n := h.Len()
				if n > 0 {
					index := rand.Intn(n)
					h.UpdateKeyByIndex(index, key)
				}
			case 5:
				n := h.Len()
				if n > 0 {
					index := rand.Intn(n)
					h.DeleteByIndex(index)
				}
			}
		}

		validateIndexedHeap(t, h, compareFunc)
	}

	f(2, cmp.Less[int], cmp.Compare[int])
	f(2, cmp.Great[int], compareDesc[int])
	f(4, cmp.Less[int], cmp.Compare[int])
	f(4, cmp.Great[int], compareDesc[int])
	f(8, cmp.Less[int], cmp.Compare[int])
	f(8, cmp.Great[int], compareDesc[int])
	f(16, cmp.Less[int], cmp.Compare[int])
	f(16, cmp.Great[int], compareDesc[int])
}

func TestIndexedHeap_Sort(t *testing.T) {
	var h = NewIndexedHeap[int, struct{}](Quadratic, nil)
	for i := 0; i < 1000; i++ {
		h.Push(rand.Int(), struct{}{})
	}
	var arr []int
	for h.Len() > 0 {
		arr = append(arr, h.Pop().Key())
	}
	assert.True(t, algo.IsSorted(arr, func(a, b int) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		} else {
			return 0
		}
	}))
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
		NewWithWays(4, cmp.Less[int])
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

func TestIndexedHeap_DeleteByIndex(t *testing.T) {
	t.Run("", func(t *testing.T) {
		var h = NewIndexedHeap[int, string](Quadratic, cmp.Less[int])
		h.Push(1, "")
		h.Push(2, "")
		var ele = h.Push(3, "")
		h.Push(4, "")
		h.DeleteByIndex(ele.Index())
		assert.Equal(t, ele.Index(), -1)
		var arr []int
		h.Range(func(ele *Element[int, string]) bool {
			arr = append(arr, ele.Key())
			return true
		})
		assert.ElementsMatch(t, arr, []int{1, 2, 4})
	})

	t.Run("", func(t *testing.T) {
		var h = NewIndexedHeap[int, string](Quadratic, cmp.Less[int])
		h.Push(1, "")
		h.Push(2, "")
		h.Push(3, "")
		h.Push(4, "")
		var ele = h.Pop()
		assert.Equal(t, ele.Index(), -1)

		var arr []int
		h.Range(func(ele *Element[int, string]) bool {
			arr = append(arr, ele.Key())
			return true
		})
		assert.ElementsMatch(t, arr, []int{2, 3, 4})
	})
}
