package stack

import (
	"github.com/lxzan/dao/internal/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_Pop(t *testing.T) {
	var s = Stack[int]{}
	s.Push(1)
	s.Push(3)
	s.Push(5)

	var arr []int
	for s.Len() > 0 {
		arr = append(arr, s.Pop())
	}
	assert.True(t, utils.IsSameSlice(arr, []int{5, 3, 1}))
	assert.Equal(t, s.Pop(), 0)
}

func TestStack_Range(t *testing.T) {
	t.Run("", func(t *testing.T) {
		var s = NewFrom(1, 3, 5)

		var arr []int
		s.Range(func(value int) bool {
			arr = append(arr, value)
			return true
		})
		assert.True(t, utils.IsSameSlice(arr, []int{1, 3, 5}))

		s.Reset()
		assert.Equal(t, s.Len(), 0)
	})

	t.Run("", func(t *testing.T) {
		var s = New[int](8)
		s.Push(1)
		s.Push(3)
		s.Push(5)

		var arr []int
		s.Range(func(value int) bool {
			arr = append(arr, value)
			return len(arr) < 2
		})
		assert.True(t, utils.IsSameSlice(arr, []int{1, 3}))
	})
}

func TestStack_UnWrap(t *testing.T) {
	var s = NewFrom(1, 3, 5)
	var a = s.UnWrap()
	assert.ElementsMatch(t, a, []int{1, 3, 5})
}
