package vector

import (
	"github.com/lxzan/dao/hashmap"
	"github.com/lxzan/dao/internal/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_GetID(t *testing.T) {
	var docs Vector[user]
	docs = append(docs, user{ID: "a"})
	docs = append(docs, user{ID: "c"})
	docs = append(docs, user{ID: "c"})
	docs = append(docs, user{ID: "b"})
	docs.UniqueByString(func(v user) string { return v.ID })
	docs.SortByString(func(v user) string {
		return v.ID
	})
	docs.Filter(func(i int, v user) bool {
		return v.ID == "b"
	})
}

type user struct {
	ID        string
	Name      string
	Age       int
	Timestamp int64
}

func (u user) GetID() string {
	return u.ID
}

func TestVector_Keys(t *testing.T) {
	t.Run("", func(t *testing.T) {
		var a = NewFromDocs("a", "b", "c")
		var b = a.MapString(func(i int, v string) string { return v })
		assert.ElementsMatch(t, b.Elem(), []string{"a", "b", "c"})
		assert.Equal(t, a.Get(0), "a")

		var values = a.Elem()
		assert.ElementsMatch(t, values, []string{"a", "b", "c"})
	})
}

func TestVector_PushBack(t *testing.T) {
	var v = New[int](8)
	v.PushBack(1)
	v.PushBack(3)
	v.PushBack(5)

	var arr []int
	for v.Len() > 0 {
		arr = append(arr, v.PopBack())
	}
	assert.True(t, utils.IsSameSlice(arr, []int{5, 3, 1}))
	assert.Equal(t, v.PopBack(), 0)
}

func TestVector_PopFront(t *testing.T) {
	var v = New[int](8)
	v.PushBack(1)
	v.PushBack(3)
	v.PushBack(5)

	var arr []int
	for v.Len() > 0 {
		arr = append(arr, v.PopFront())
	}
	assert.True(t, utils.IsSameSlice(arr, []int{1, 3, 5}))
	assert.Equal(t, v.PopFront(), 0)
}

func TestVector_Range(t *testing.T) {
	t.Run("", func(t *testing.T) {
		var a = NewFromDocs[int64](1, 3, 5)
		var v = a.Clone()
		var arr []int64
		v.Range(func(i int, value int64) bool {
			arr = append(arr, value)
			return true
		})
		assert.True(t, utils.IsSameSlice(arr, []int64{1, 3, 5}))
	})

	t.Run("", func(t *testing.T) {
		var v = NewFromDocs[int64](1, 3, 5)
		var arr []int64
		v.Range(func(i int, value int64) bool {
			arr = append(arr, value)
			return len(arr) < 2
		})
		assert.True(t, utils.IsSameSlice(arr, []int64{1, 3}))
	})
}

func TestVector_ToMap(t *testing.T) {
	t.Run("", func(t *testing.T) {
		var a = NewFromDocs[user](
			user{ID: "a"},
			user{ID: "b"},
			user{ID: "c"},
		)
		var values = hashmap.HashMap[string, user](a.ToStringMap(func(v user) string {
			return v.ID
		})).Keys()
		assert.ElementsMatch(t, values, []string{"a", "b", "c"})
	})

	t.Run("", func(t *testing.T) {
		var a = NewFromDocs[user](
			user{ID: "a", Age: 1},
			user{ID: "b", Age: 2},
			user{ID: "c", Age: 3},
		)
		var values = hashmap.HashMap[int, user](a.ToIntMap(func(v user) int {
			return v.Age
		})).Keys()
		assert.ElementsMatch(t, values, []int{1, 2, 3})
	})

	t.Run("", func(t *testing.T) {
		var a = NewFromDocs[user](
			user{ID: "a", Timestamp: 1},
			user{ID: "b", Timestamp: 2},
			user{ID: "c", Timestamp: 3},
		)
		var values = hashmap.HashMap[int64, user](a.ToInt64Map(func(v user) int64 {
			return v.Timestamp
		})).Keys()
		assert.ElementsMatch(t, values, []int64{1, 2, 3})
	})
}

func TestVector_Slice(t *testing.T) {
	var a = NewFromDocs("a", "b", "c", "d")
	var b = a.Slice(1, 3)
	var values = b.Elem()
	assert.ElementsMatch(t, values, []string{"b", "c"})

	assert.Equal(t, a.Len(), 4)
	a.Reset()
	assert.Equal(t, a.Len(), 0)
}

func TestVector_Sort(t *testing.T) {
	t.Run("", func(t *testing.T) {
		var a = NewFromDocs[int](1, 3, 5, 2, 4, 6).
			SortByInt(func(v int) int {
				return v
			}).
			Elem()
		assert.True(t, utils.IsSameSlice(a, []int{1, 2, 3, 4, 5, 6}))
	})

	t.Run("", func(t *testing.T) {
		var a = NewFromDocs[int64](1, 3, 5, 2, 4, 6).
			SortByInt64(func(v int64) int64 { return v }).
			Elem()
		assert.True(t, utils.IsSameSlice(a, []int64{1, 2, 3, 4, 5, 6}))
	})
}

func TestVector_Update(t *testing.T) {
	var v = NewFromDocs(1, 3, 5)
	assert.True(t, utils.IsSameSlice(v.Elem(), []int{1, 3, 5}))
	v.Update(0, 2)
	v.Update(1, 4)
	v.Update(2, 6)
	assert.True(t, utils.IsSameSlice(v.Elem(), []int{2, 4, 6}))
}

func TestVector_Reverse(t *testing.T) {
	var v = NewFromDocs(1, 2, 3)
	v.Reverse()
	assert.True(t, utils.IsSameSlice(v.Elem(), []int{3, 2, 1}))
}

func TestVector_Map(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		var a = NewFromDocs[user](
			user{ID: "a", Name: "ming"},
			user{ID: "b", Name: "hong"},
			user{ID: "c", Name: "hong"},
		)
		var b = a.
			MapString(func(i int, v user) string { return v.Name }).
			UniqueByString(func(v string) string { return v }).
			Elem()
		assert.ElementsMatch(t, b, []string{"ming", "hong"})
	})

	t.Run("int", func(t *testing.T) {
		var a = NewFromDocs[user](
			user{ID: "a", Name: "ming", Age: 1},
			user{ID: "b", Name: "hong", Age: 2},
			user{ID: "c", Name: "hong", Age: 3},
			user{ID: "d", Name: "mei", Age: 2},
		)
		var b = a.
			MapInt(func(i int, v user) int { return v.Age }).
			UniqueByInt(func(v int) int { return v }).
			Elem()
		assert.ElementsMatch(t, b, []int{1, 2, 3})
	})

	t.Run("int64", func(t *testing.T) {
		var a = NewFromDocs[user](
			user{ID: "a", Name: "ming", Timestamp: 1},
			user{ID: "b", Name: "hong", Timestamp: 2},
			user{ID: "c", Name: "hong", Timestamp: 3},
			user{ID: "d", Name: "mei", Timestamp: 2},
		)
		var b = a.
			MapInt64(func(i int, v user) int64 { return v.Timestamp }).
			UniqueByInt64(func(v int64) int64 { return v }).
			Elem()
		assert.ElementsMatch(t, b, []int64{1, 2, 3})
	})
}

func TestVector_PushFront(t *testing.T) {
	var v Vector[int]
	v.PushFront(1)
	assert.ElementsMatch(t, v.Elem(), []int{1})
	v.PushBack(3)
	v.PushBack(5)
	v.PushFront(7)
	assert.True(t, utils.IsSameSlice(v.Elem(), []int{7, 1, 3, 5}))
}

func TestVector_Delete(t *testing.T) {
	t.Run("", func(t *testing.T) {
		var v = NewFromDocs(1, 3, 5, 7)
		v.Delete(0)
		assert.True(t, utils.IsSameSlice(v.Elem(), []int{3, 5, 7}))
	})

	t.Run("", func(t *testing.T) {
		var v = NewFromDocs(1, 3, 5, 7)
		v.Delete(3)
		assert.True(t, utils.IsSameSlice(v.Elem(), []int{1, 3, 5}))
	})

	t.Run("", func(t *testing.T) {
		var v = NewFromDocs(1, 3, 5, 7)
		v.Delete(1)
		assert.True(t, utils.IsSameSlice(v.Elem(), []int{1, 5, 7}))
	})
}

func TestVector_Get(t *testing.T) {
	var v = NewFromDocs(1, 3, 5)
	assert.Equal(t, v.Front(), int(1))
	assert.Equal(t, v.Back(), int(5))
}

func TestVector_GroupByInt(t *testing.T) {
	t.Run("", func(t *testing.T) {
		var arr = NewFromDocs(1, 3, 5, 7, 2, 4, 6, 8)
		var m = arr.GroupByInt(func(i int, v int) int {
			return v % 2
		})
		assert.ElementsMatch(t, m[0], Vector[int]{2, 4, 6, 8})
		assert.ElementsMatch(t, m[1], Vector[int]{1, 3, 5, 7})
	})

	t.Run("", func(t *testing.T) {
		var arr = NewFromDocs(1, 3, 5, 7, 2, 4, 6, 8)
		var m = arr.GroupByInt(func(i int, v int) int {
			return v % 3
		})
		assert.ElementsMatch(t, m[0], Vector[int]{3, 6})
		assert.ElementsMatch(t, m[1], Vector[int]{1, 4, 7})
		assert.ElementsMatch(t, m[2], Vector[int]{2, 5, 8})
	})
}

func TestVector_GroupByInt64(t *testing.T) {
	t.Run("", func(t *testing.T) {
		var arr = NewFromDocs[int64](1, 3, 5, 7, 2, 4, 6, 8)
		var m = arr.GroupByInt64(func(i int, v int64) int64 {
			return v % 2
		})
		assert.ElementsMatch(t, m[0], Vector[int64]{2, 4, 6, 8})
		assert.ElementsMatch(t, m[1], Vector[int64]{1, 3, 5, 7})
	})

	t.Run("", func(t *testing.T) {
		var arr = NewFromDocs[int64](1, 3, 5, 7, 2, 4, 6, 8)
		var m = arr.GroupByInt64(func(i int, v int64) int64 {
			return v % 3
		})
		assert.ElementsMatch(t, m[0], Vector[int64]{3, 6})
		assert.ElementsMatch(t, m[1], Vector[int64]{1, 4, 7})
		assert.ElementsMatch(t, m[2], Vector[int64]{2, 5, 8})
	})
}

func TestVector_GroupByString(t *testing.T) {
	var arr = NewFromDocs("abc", "abnormal", "oh", "oho", "bank", "bark")
	var m = arr.GroupByString(func(i int, v string) string {
		return v[:2]
	})
	assert.ElementsMatch(t, m["ab"], Vector[string]{"abc", "abnormal"})
	assert.ElementsMatch(t, m["oh"], Vector[string]{"oh", "oho"})
	assert.ElementsMatch(t, m["ba"], Vector[string]{"bank", "bark"})
}

func TestVector_Cap(t *testing.T) {
	var arr = make([]int, 0, 3)
	var v = Vector[int](arr)
	assert.Equal(t, v.Cap(), 3)
}

func TestVector_Len(t *testing.T) {
	var v *Vector[string]
	assert.Equal(t, v.Len(), 0)
}
