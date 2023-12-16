package heap

import (
	"github.com/lxzan/dao/internal/utils"
	"sort"
	"testing"
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
