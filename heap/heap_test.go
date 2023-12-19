package heap

import (
	"fmt"
	"github.com/lxzan/dao"
	"github.com/lxzan/dao/internal/utils"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
	"unsafe"
)

func TestNew(t *testing.T) {
	const count = 1000
	{
		var h = New[string](func(a, b string) bool {
			return a < b
		})
		var arr1 = make([]string, 0)
		var arr2 = make([]string, 0)
		for i := 0; i < count; i++ {
			var s = utils.Numeric.Generate(8)
			h.Push(s)
			arr1 = append(arr1, s)
		}
		for h.Len() > 0 {
			arr2 = append(arr2, h.Pop())
		}
		sort.Strings(arr1)
		if !utils.IsSameSlice(arr1, arr2) {
			t.Fatal("error!")
		}
	}
}

func TestDesc(t *testing.T) {
	var h = New(dao.DescFunc[int])
	h.SetCap(8)
	h.SetForkNumber(Octal)
	h.Push(1)
	assert.Equal(t, h.Top(), 1)
	h.Push(3)
	h.Push(2)
	h.Push(5)
	h.Push(4)

	var arr []int
	for h.Len() > 0 {
		arr = append(arr, h.Pop())
	}
	assert.True(t, utils.IsSameSlice(arr, []int{5, 4, 3, 2, 1}))
}

func TestAsc(t *testing.T) {
	var h = New(dao.AscFunc[int])
	h.SetCap(8)
	h.SetForkNumber(Binary)
	h.Push(1)
	h.Push(3)
	h.Push(2)
	h.Push(5)
	h.Push(4)

	var arr []int
	for h.Len() > 0 {
		arr = append(arr, h.Pop())
	}
	assert.True(t, utils.IsSameSlice(arr, []int{1, 2, 3, 4, 5}))
}

func TestHeap_Range(t *testing.T) {
	var h = New(dao.AscFunc[int])
	h.SetCap(8)
	h.Push(1)
	h.Push(3)
	h.Push(2)
	h.Push(5)
	h.Push(4)

	{
		var arr []int
		h.Range(func(index int, value int) bool {
			arr = append(arr, value)
			return true
		})
		assert.ElementsMatch(t, arr, []int{1, 2, 3, 4, 5})
	}

	{
		var arr []int
		h.Range(func(index int, value int) bool {
			arr = append(arr, value)
			return len(arr) < 2
		})
		assert.Equal(t, len(arr), 2)
	}
}

func TestHeap_Reset(t *testing.T) {
	var h = New(dao.AscFunc[int])
	h.Push(1)
	h.Push(3)
	h.Push(2)
	h.Push(5)
	h.Push(4)
	h.Reset()
	assert.Equal(t, h.Len(), 0)
}

func TestHeap_Pop(t *testing.T) {
	var h = New(dao.AscFunc[int])
	assert.Equal(t, h.Pop(), 0)
	h.Push(1)
	assert.Equal(t, h.Pop(), 1)
}

func TestHeap_SetForkNumber(t *testing.T) {
	var h = New(dao.AscFunc[int])
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
		h.SetForkNumber(3)
	})
	assert.Error(t, err1)

	var err2 = catch(func() {
		h.SetForkNumber(4)
	})
	assert.Nil(t, err2)
}

func TestHeap_Clone(t *testing.T) {
	var h = New(dao.AscFunc[int])
	h.SetForkNumber(4)
	h.Push(1)
	h.Push(3)
	h.Push(2)
	h.Push(4)

	var h1 = h.Clone()
	var h2 = h
	assert.True(t, utils.IsSameSlice(h.data, h1.data))
	var addr = (uintptr)(unsafe.Pointer(&h.data[0]))
	var addr1 = (uintptr)(unsafe.Pointer(&h1.data[0]))
	var addr2 = (uintptr)(unsafe.Pointer(&h2.data[0]))
	assert.NotEqual(t, addr, addr1)
	assert.Equal(t, addr, addr2)
}
