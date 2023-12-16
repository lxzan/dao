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
