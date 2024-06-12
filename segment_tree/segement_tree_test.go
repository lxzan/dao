package segment_tree

import (
	"github.com/lxzan/dao/algo"
	"github.com/lxzan/dao/internal/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSegmentTree_Query(t *testing.T) {
	var n = 10000
	var arr = make([]int, 0)
	for i := 0; i < n; i++ {
		arr = append(arr, utils.Rand.Intn(n))
	}
	var stree = New(arr, NewIntSummary[int], MergeIntSummary[int])
	for i := 0; i < 100; i++ {
		var x, y = utils.Alphabet.Intn(n), utils.Alphabet.Intn(n)
		if x == y {
			continue
		}
		if x > y {
			x, y = y, x
		}

		var flag = utils.Alphabet.Intn(4)
		switch flag {
		case 0:
			stree.Update(x, y)
		default:
			r0 := stree.Query(x, y)
			r1 := NewIntSummary(arr[x], OperateQuery)
			for j := x; j < y; j++ {
				r1.MaxValue = algo.Max(r1.MaxValue, arr[j])
				r1.MinValue = algo.Min(r1.MinValue, arr[j])
				r1.Sum += arr[j]
			}
			assert.Equal(t, r0.MaxValue, r1.MaxValue)
			assert.Equal(t, r0.MinValue, r1.MinValue)
			assert.Equal(t, r0.Sum, r1.Sum)
		}
	}
}
