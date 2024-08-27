package heap

import (
	"github.com/lxzan/dao/algo"
	"github.com/lxzan/dao/internal/utils"
	"github.com/lxzan/dao/types/cmp"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func validateIndexedHeap[K cmp.Ordered, V cmp.Ordered](t *testing.T, h *HashHeap[K, V], m map[K]V, compare cmp.CompareFunc[V]) {
	var n = h.Len()
	if n > 0 {
		assert.Equal(t, h.heap.data[0].Key(), h.Top().Key())
	}

	var i = 0
	h.Range(func(ele *Element[K, V]) bool {
		assert.Equal(t, ele.index, i)
		i++

		var base = ele.index << 2
		var end = algo.Min(base+4, n-1)
		for j := base + 1; j <= end; j++ {
			child := h.heap.data[j]
			assert.True(t, compare(ele.Value(), child.Value()) <= 0)
		}
		return true
	})

	assert.Equal(t, len(m), h.Len())
	for k, v := range m {
		v2, _ := h.Get(k)
		assert.Equal(t, v, v2)
	}

	var values = make([]V, 0, n)
	for h.Len() > 0 {
		values = append(values, h.Pop().Value())
	}
	assert.True(t, algo.IsSorted(values, compare))
}

func TestIndexedHeap_Random(t *testing.T) {
	const count = 10000

	var f = func(lessFunc cmp.LessFunc[int], compareFunc cmp.CompareFunc[int]) {
		var h = NewHashHeap[int, int](lessFunc)
		var m = make(map[int]int)
		for i := 0; i < count; i++ {
			flag := utils.Alphabet.Intn(6)
			key := rand.Intn(count)
			switch flag {
			case 0, 1, 2:
				h.Set(key, i)
				m[key] = i
			case 3:
				if ele := h.Pop(); ele != nil {
					delete(m, ele.Key())
				}
			case 4, 5:
				h.Delete(key)
				delete(m, key)
			}
		}

		validateIndexedHeap(t, h, m, compareFunc)
	}

	f(cmp.Less[int], cmp.Compare[int])
	f(cmp.Great[int], compareDesc[int])
}

func TestHashHeap_Range(t *testing.T) {
	var hm = NewHashHeap[string, int](func(a, b int) bool {
		return a < b
	})
	hm.Set("a", 1)
	hm.Set("b", 2)
	hm.Set("c", 3)
	var values []int
	hm.Range(func(ele *Element[string, int]) bool {
		values = append(values, ele.Value())
		return len(values) < 2
	})
	assert.Equal(t, len(values), 2)
}

func TestHashHeap_Pop(t *testing.T) {
	var hm = NewHashHeap[string, int](func(a, b int) bool {
		return a < b
	})
	assert.Nil(t, hm.Pop())
	assert.Nil(t, hm.heap.Pop())
}
