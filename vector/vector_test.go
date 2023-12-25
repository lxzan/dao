package vector

import (
	"github.com/lxzan/dao/internal/utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"unsafe"
)

func TestUser_GetID(t *testing.T) {
	var docs Vector[string, user]
	docs = append(docs, user{ID: "a"})
	docs = append(docs, user{ID: "c"})
	docs = append(docs, user{ID: "c"})
	docs = append(docs, user{ID: "b"})
	docs.Unique()
	docs.Sort()
	docs.Filter(func(i int, v user) bool {
		return v.ID == "b"
	})
}

type user struct {
	ID string
}

func (u user) GetID() string {
	return u.ID
}

func TestNewFromInts(t *testing.T) {
	var a = NewFromInts(1, 3, 5)
	var b = a.GetIdList()
	assert.ElementsMatch(t, b, []int{1, 3, 5})
}

func TestNewFromInt64s(t *testing.T) {
	var a = NewFromInt64s(1, 3, 5)
	var b = a.GetIdList()
	assert.ElementsMatch(t, b, []int64{1, 3, 5})
}

func TestVector_Keys(t *testing.T) {
	t.Run("", func(t *testing.T) {
		var a = NewFromStrings("a", "b", "c")
		var b = a.GetIdList()
		assert.ElementsMatch(t, b, []string{"a", "b", "c"})
		assert.Equal(t, a.Get(0).GetID(), "a")

		var addr0 = (uintptr)(unsafe.Pointer(&(*a)[0]))
		var addr1 = (uintptr)(unsafe.Pointer(&b[0]))
		assert.Equal(t, addr0, addr1)

		var values = a.Elem()
		assert.ElementsMatch(t, values, []String{"a", "b", "c"})
	})

	t.Run("", func(t *testing.T) {
		var docs = NewFromDocs[string, user](
			user{ID: "a"},
			user{ID: "b"},
			user{ID: "c"},
		)
		assert.ElementsMatch(t, docs.GetIdList(), []string{"a", "b", "c"})
	})
}

func TestVector_Exists(t *testing.T) {
	var v = New[int, Int](8)
	v.PushBack(1)
	v.PushBack(3)
	v.PushBack(5)

	{
		_, ok := v.Exists(1)
		assert.True(t, ok)
	}

	{
		_, ok := v.Exists(3)
		assert.True(t, ok)
	}

	{
		_, ok := v.Exists(2)
		assert.False(t, ok)
	}
}

func TestVector_PushBack(t *testing.T) {
	var v = New[int, Int](8)
	v.PushBack(1)
	v.PushBack(3)
	v.PushBack(5)

	var arr []int
	for v.Len() > 0 {
		arr = append(arr, v.PopBack().GetID())
	}
	assert.True(t, utils.IsSameSlice(arr, []int{5, 3, 1}))
	assert.Equal(t, v.PopBack().GetID(), 0)
}

func TestVector_PopFront(t *testing.T) {
	var v = New[int, Int](8)
	v.PushBack(1)
	v.PushBack(3)
	v.PushBack(5)

	var arr []int
	for v.Len() > 0 {
		arr = append(arr, v.PopFront().GetID())
	}
	assert.True(t, utils.IsSameSlice(arr, []int{1, 3, 5}))
	assert.Equal(t, v.PopFront().GetID(), 0)
}

func TestVector_Range(t *testing.T) {
	t.Run("", func(t *testing.T) {
		var a = NewFromInt64s(1, 3, 5)
		var v = a.Clone()
		var arr []int64
		v.Range(func(i int, value Int64) bool {
			arr = append(arr, value.GetID())
			return true
		})
		assert.True(t, utils.IsSameSlice(arr, []int64{1, 3, 5}))
	})

	t.Run("", func(t *testing.T) {
		var v = NewFromInt64s(1, 3, 5)
		var arr []int64
		v.Range(func(i int, value Int64) bool {
			arr = append(arr, value.GetID())
			return len(arr) < 2
		})
		assert.True(t, utils.IsSameSlice(arr, []int64{1, 3}))
	})
}

func TestVector_ToMap(t *testing.T) {
	var a = NewFromDocs[string, user](
		user{ID: "a"},
		user{ID: "b"},
		user{ID: "c"},
	)
	var values = a.ToMap().Keys()
	assert.ElementsMatch(t, values, []string{"a", "b", "c"})
}

func TestVector_Slice(t *testing.T) {
	var a = NewFromStrings("a", "b", "c", "d")
	var b = a.Slice(1, 3)
	var values = b.GetIdList()
	assert.ElementsMatch(t, values, []string{"b", "c"})

	assert.Equal(t, a.Len(), 4)
	a.Reset()
	assert.Equal(t, a.Len(), 0)
}

func TestVector_Sort(t *testing.T) {
	var a = NewFromInts(1, 3, 5, 2, 4, 6).Sort().GetIdList()
	assert.True(t, utils.IsSameSlice(a, []int{1, 2, 3, 4, 5, 6}))
}

func TestVector_Update(t *testing.T) {
	var v = NewFromInts(1, 3, 5)
	assert.True(t, utils.IsSameSlice(v.Elem(), []Int{1, 3, 5}))
	v.Update(0, 2)
	v.Update(1, 4)
	v.Update(2, 6)
	assert.True(t, utils.IsSameSlice(v.Elem(), []Int{2, 4, 6}))
}

func TestVector_Reverse(t *testing.T) {
	var v = NewFromInts(1, 2, 3)
	v.Reverse()
	assert.True(t, utils.IsSameSlice(v.Elem(), []Int{3, 2, 1}))
}
