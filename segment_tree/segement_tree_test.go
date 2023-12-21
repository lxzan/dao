package segment_tree

import (
	"github.com/lxzan/dao/algorithm"
	"github.com/lxzan/dao/internal/utils"
	"testing"
)

func TestSegmentTree_Query(t *testing.T) {
	var n = 10000
	var arr = make([]Int64, 0)
	for i := 0; i < n; i++ {
		arr = append(arr, Int64(utils.Rand.Intn(n)))
	}

	var tree = New[Int64Schema, Int64](arr)

	for i := 0; i < 1000; i++ {
		var left = utils.Rand.Intn(n)
		var right = utils.Rand.Intn(n)
		if left > right {
			left, right = right, left
		}
		var result1 = tree.Query(left, right+1)

		var result2 = Int64Schema{
			MaxValue: arr[left].Value(),
			MinValue: arr[left].Value(),
			Sum:      0,
		}
		for j := left; j <= right; j++ {
			result2.Sum += arr[j].Value()
			result2.MaxValue = algorithm.Max(result2.MaxValue, arr[j].Value())
			result2.MinValue = algorithm.Min(result2.MinValue, arr[j].Value())
		}

		if result1.Sum != result2.Sum || result1.MinValue != result2.MinValue || result1.MaxValue != result2.MaxValue {
			t.Fatal("error!")
		}
	}

	for i := 0; i < 1000; i++ {
		var index = utils.Rand.Intn(n)
		var value = Int64(utils.Rand.Intn(n))
		tree.Update(index, value)
	}

	for i := 0; i < 1000; i++ {
		var left = utils.Rand.Intn(n)
		var right = utils.Rand.Intn(n)
		if left > right {
			left, right = right, left
		}
		var result1 = tree.Query(left, right+1)

		var result2 = Int64Schema{
			MaxValue: arr[left].Value(),
			MinValue: arr[left].Value(),
			Sum:      0,
		}
		for j := left; j <= right; j++ {
			result2.Sum += arr[j].Value()
			result2.MaxValue = algorithm.Max(result2.MaxValue, arr[j].Value())
			result2.MinValue = algorithm.Min(result2.MinValue, arr[j].Value())
		}

		if result1.Sum != result2.Sum || result1.MinValue != result2.MinValue || result1.MaxValue != result2.MaxValue {
			t.Fatal("error!")
		}
	}
}
