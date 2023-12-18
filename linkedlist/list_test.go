package linkedlist

import (
	"container/list"
	"github.com/lxzan/dao/internal/utils"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func validate[T comparable](q *LinkedList[T]) bool {
	var sum = 0
	for i := q.Front(); i != nil; i = i.Next() {
		sum++
		next := i.next
		if next == nil {
			continue
		}
		if i.next.Value != next.Value {
			return false
		}
		if next.prev.Value != i.Value {
			return false
		}
	}

	if q.Len() != sum {
		return false
	}

	if head := q.Front(); head != nil {
		if head.prev != nil {
			return false
		}
	}

	if tail := q.Back(); tail != nil {
		if tail.next != nil {
			return false
		}
	}

	if q.Len() == 1 && q.Front().Value != q.Back().Value {
		return false
	}

	return true
}

func TestLinkedList_Reset(t *testing.T) {
	var q = New[int]()
	q.PushBack(1)
	q.PushBack(2)
	q.PushBack(3)
	q.Reset()
	assert.True(t, validate(q))
	assert.Equal(t, q.Len(), 0)
}

func TestLinkedList_PopBack(t *testing.T) {
	var q = New[int]()
	assert.Equal(t, q.PopBack(), 0)

	q.PushBack(1)
	assert.Equal(t, q.PopBack(), 1)
}

func TestQueue_Range(t *testing.T) {
	const count = 1000

	t.Run("", func(t *testing.T) {
		var q = New[int]()
		var a []int
		for i := 0; i < count; i++ {
			v := rand.Intn(count)
			q.PushBack(v)
			a = append(a, v)
		}

		assert.Equal(t, q.Len(), count)

		var b []int
		q.Range(func(ele *Element[int]) bool {
			b = append(b, ele.Value)
			return len(b) < 100
		})
		assert.Equal(t, len(b), 100)

		var i = 0
		for q.Len() > 0 {
			v := q.PopFront()
			assert.Equal(t, a[i], v)
			i++
		}
	})

	t.Run("", func(t *testing.T) {
		var q = New[int]()
		for i := 0; i < count; i++ {
			v := rand.Intn(count)
			q.PushBack(v)
		}

		var a1 []int
		var a2 []int
		for i := q.Front(); i != nil; i = i.Next() {
			a1 = append(a1, i.Value)
		}
		for i := q.Back(); i != nil; i = i.Prev() {
			a2 = append(a2, i.Value)
		}

		assert.ElementsMatch(t, a1, a2)
	})
}

func TestQueue_Addr(t *testing.T) {
	const count = 1000
	var q = New[int]()
	for i := 0; i < count; i++ {
		v := rand.Intn(count)
		if v&7 == 0 {
			q.PopFront()
		} else {
			q.PushBack(v)
		}
	}

	var sum = 0
	for i := q.Front(); i != nil; i = i.Next() {
		sum++

		prev := i.prev
		next := i.next
		if prev != nil {
			assert.Equal(t, prev.next.Value, i.Value)
		}
		if next != nil {
			assert.Equal(t, i.Value, next.prev.Value)
		}
	}

	assert.Equal(t, q.Len(), sum)
	if head := q.head; head != nil {
		assert.Zero(t, head.prev)
	}
	if tail := q.tail; tail != nil {
		assert.Zero(t, tail.next)
	}
}

func TestQueue_Pop(t *testing.T) {
	var q = New[int]()
	assert.Zero(t, q.Front())
	assert.Zero(t, q.PopFront())

	q.PushBack(1)
	q.PushBack(2)
	q.PushBack(3)
	q.PopFront()
	q.PushBack(4)
	q.PushBack(5)
	q.PopFront()

	var arr []int
	q.Range(func(ele *Element[int]) bool {
		arr = append(arr, ele.Value)
		return true
	})
	assert.Equal(t, q.Front().Value, 3)
	assert.True(t, utils.IsSameSlice(arr, []int{3, 4, 5}))
}

func TestLinkedList_InsertAfter(t *testing.T) {
	t.Run("", func(t *testing.T) {
		var q = New[int]()
		assert.Nil(t, q.InsertAfter(1, q.Front()))

	})

	t.Run("", func(t *testing.T) {
		var q = New[int]()
		q.PushBack(1)
		var node = q.PushBack(2)
		q.PushBack(4)
		q.InsertAfter(3, node)

		var arr []int
		q.Range(func(ele *Element[int]) bool {
			arr = append(arr, ele.Value)
			return true
		})

		assert.True(t, utils.IsSameSlice(arr, []int{1, 2, 3, 4}))
		assert.True(t, validate(q))
	})

	t.Run("", func(t *testing.T) {
		var q = New[int]()
		q.PushBack(1)
		q.PushBack(2)
		var node = q.PushBack(4)
		q.InsertAfter(3, node)

		var arr []int
		q.Range(func(ele *Element[int]) bool {
			arr = append(arr, ele.Value)
			return true
		})
		assert.True(t, utils.IsSameSlice(arr, []int{1, 2, 4, 3}))
		assert.True(t, validate(q))
	})
}

func TestLinkedList_InsertBefore(t *testing.T) {
	t.Run("", func(t *testing.T) {
		var q = New[int]()
		assert.Nil(t, q.InsertBefore(1, q.Front()))
	})

	t.Run("", func(t *testing.T) {
		var q = New[int]()
		q.PushBack(1)
		var node = q.PushBack(2)
		q.PushBack(4)
		q.InsertBefore(3, node)

		var arr []int
		q.Range(func(ele *Element[int]) bool {
			arr = append(arr, ele.Value)
			return true
		})

		assert.True(t, utils.IsSameSlice(arr, []int{1, 3, 2, 4}))
		assert.True(t, validate(q))
	})

	t.Run("", func(t *testing.T) {
		var q = New[int]()
		var node = q.PushBack(1)
		q.PushBack(2)
		q.PushBack(4)
		q.InsertBefore(3, node)

		var arr []int
		q.Range(func(ele *Element[int]) bool {
			arr = append(arr, ele.Value)
			return true
		})
		assert.True(t, utils.IsSameSlice(arr, []int{3, 1, 2, 4}))
		assert.True(t, validate(q))
	})
}

func TestLinkedList_Delete(t *testing.T) {
	t.Run("", func(t *testing.T) {
		var q = New[int]()
		var node = q.PushBack(1)
		q.PushBack(2)
		q.PushBack(3)
		q.Remove(node)

		var arr []int
		q.Range(func(ele *Element[int]) bool {
			arr = append(arr, ele.Value)
			return true
		})
		assert.True(t, utils.IsSameSlice(arr, []int{2, 3}))
		assert.True(t, validate(q))
	})

	t.Run("", func(t *testing.T) {
		var q = New[int]()
		q.PushBack(1)
		var node = q.PushBack(2)
		q.PushBack(3)
		q.Remove(node)

		var arr []int
		q.Range(func(ele *Element[int]) bool {
			arr = append(arr, ele.Value)
			return true
		})
		assert.True(t, utils.IsSameSlice(arr, []int{1, 3}))
		assert.True(t, validate(q))
	})

	t.Run("", func(t *testing.T) {
		var q = New[int]()
		q.PushBack(1)
		q.PushBack(2)
		var node = q.PushBack(3)
		q.Remove(node)

		var arr []int
		q.Range(func(ele *Element[int]) bool {
			arr = append(arr, ele.Value)
			return true
		})
		assert.True(t, utils.IsSameSlice(arr, []int{1, 2}))
		assert.True(t, validate(q))
	})

	t.Run("", func(t *testing.T) {
		var q = New[int]()
		var node = q.PushBack(3)
		q.Remove(node)
		assert.Equal(t, q.Len(), 0)
		assert.True(t, validate(q))
	})
}

func TestLinkedList_PushFront(t *testing.T) {
	var q = New[int]()
	q.PushFront(1)
	assert.True(t, validate(q))
}

func TestQueue_Random(t *testing.T) {
	var count = 10000
	var q = LinkedList[int]{}
	var linkedlist = list.New()
	for i := 0; i < count; i++ {
		var flag = rand.Intn(13)
		var val = rand.Int()
		switch flag {
		case 0, 1:
			q.PushBack(val)
			linkedlist.PushBack(val)
		case 2, 3:
			q.PushFront(val)
			linkedlist.PushFront(val)
		case 4:
			if q.Len() > 0 {
				q.PopFront()
				linkedlist.Remove(linkedlist.Front())
			}
		case 5:
			if q.Len() > 0 {
				q.PopBack()
				linkedlist.Remove(linkedlist.Back())
			}
		case 6:
			if node := q.Front(); node != nil {
				q.MoveToBack(node)
				linkedlist.MoveToBack(linkedlist.Front())
			}
		case 7:
			if node := q.Back(); node != nil {
				q.MoveToFront(node)
				linkedlist.MoveToFront(linkedlist.Back())
			}
		case 8:
			if node := q.Back(); node != nil {
				q.MoveToFront(node)
				linkedlist.MoveToFront(linkedlist.Back())
			}
		case 9:
			var n = rand.Intn(10)
			var index = 0
			for iter := q.Front(); iter != nil; iter = iter.Next() {
				index++
				if index >= n {
					q.InsertAfter(val, iter)
					break
				}
			}

			index = 0
			for iter := linkedlist.Front(); iter != nil; iter = iter.Next() {
				index++
				if index >= n {
					linkedlist.InsertAfter(val, iter)
					break
				}
			}
		case 10:
			var n = rand.Intn(10)
			var index = 0
			for iter := q.Front(); iter != nil; iter = iter.Next() {
				index++
				if index >= n {
					q.InsertBefore(val, iter)
					break
				}
			}

			index = 0
			for iter := linkedlist.Front(); iter != nil; iter = iter.Next() {
				index++
				if index >= n {
					linkedlist.InsertBefore(val, iter)
					break
				}
			}
		case 11, 12:
			var n = rand.Intn(10)
			var index = 0
			for iter := q.Front(); iter != nil; iter = iter.Next() {
				index++
				if index >= n {
					q.Remove(iter)
					break
				}
			}

			index = 0
			for iter := linkedlist.Front(); iter != nil; iter = iter.Next() {
				index++
				if index >= n {
					linkedlist.Remove(iter)
					break
				}
			}
		default:

		}
	}

	assert.True(t, validate(&q))
	for i := linkedlist.Front(); i != nil; i = i.Next() {
		var val = q.PopFront()
		assert.Equal(t, i.Value, val)
	}
}
