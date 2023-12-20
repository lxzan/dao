package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsSameSlice(t *testing.T) {
	assert.True(t, IsSameSlice(
		[]int{1, 2, 3},
		[]int{1, 2, 3},
	))

	assert.False(t, IsSameSlice(
		[]int{1, 2, 3},
		[]int{1, 2},
	))

	assert.False(t, IsSameSlice(
		[]int{1, 2, 3},
		[]int{1, 2, 4},
	))
}

func TestRandomString_Generate(t *testing.T) {
	var s = Numeric.Generate(6)
	assert.Equal(t, len(s), 6)
}

func TestReverseStrings(t *testing.T) {
	{
		var arr = []string{"a", "b", "c"}
		ReverseStrings(arr)
		assert.True(t, IsSameSlice(arr, []string{"c", "b", "a"}))
	}

	{
		var arr = []string{"a", "b"}
		ReverseStrings(arr)
		assert.True(t, IsSameSlice(arr, []string{"b", "a"}))
	}

	{
		var arr = []string{"a"}
		ReverseStrings(arr)
		assert.True(t, IsSameSlice(arr, []string{"a"}))
	}

	{
		var arr = []string{}
		ReverseStrings(arr)
		assert.True(t, IsSameSlice(arr, []string{}))
	}
}

func TestIsBinaryNumber(t *testing.T) {
	assert.True(t, IsBinaryNumber(1))
	assert.True(t, IsBinaryNumber(2))
	assert.True(t, IsBinaryNumber(16))

	assert.False(t, IsBinaryNumber(3))
	assert.False(t, IsBinaryNumber(7))
	assert.False(t, IsBinaryNumber(21))
}

func TestClone(t *testing.T) {
	{
		var a []int
		var b = Clone(a)
		assert.True(t, len(b) == 0)
	}

	{
		var a = []int{1, 2, 3}
		var b = Clone(a)
		assert.ElementsMatch(t, b, a)
	}
}
