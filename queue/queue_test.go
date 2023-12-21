package queue

import (
	"github.com/lxzan/dao/internal/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue(t *testing.T) {
	t.Run("", func(t *testing.T) {
		q := NewFrom(1, 3, 5, 7, 9)
		var a []int
		for q.Len() > 0 {
			a = append(a, q.Pop())
		}
		assert.True(t, utils.IsSameSlice(a, []int{1, 3, 5, 7, 9}))
		assert.Equal(t, q.offset, 0)
	})

	t.Run("", func(t *testing.T) {
		q := NewFrom[int](1)
		q.Push(3)
		q.Pop()
		assert.Equal(t, q.offset, 1)
		assert.Equal(t, q.data[0], 0)
		assert.Equal(t, q.Len(), 1)
	})
}

func TestQueue_UnWrap(t *testing.T) {
	t.Run("", func(t *testing.T) {
		q := NewFrom(1, 3, 5, 7, 9)
		q.Pop()
		a := q.UnWrap()
		assert.True(t, utils.IsSameSlice(a, []int{3, 5, 7, 9}))
	})

	t.Run("", func(t *testing.T) {
		q := NewFrom(1)
		q.Pop()
		a := q.UnWrap()
		assert.Equal(t, len(a), 0)
	})

	t.Run("", func(t *testing.T) {
		q := New[int](8)
		a := q.UnWrap()
		assert.Equal(t, len(a), 0)
	})
}

func TestQueue_Clone(t *testing.T) {
	q := NewFrom(1, 3, 5, 7, 9)
	b := q.Clone()
	assert.True(t, utils.IsSameSlice(b.UnWrap(), []int{1, 3, 5, 7, 9}))
}

func TestQueue_Range(t *testing.T) {
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
