package heap

import (
	"github.com/lxzan/dao/internal/utils"
	"sort"
	"testing"
)

func TestNew(t *testing.T) {
	const count = 1000
	{
		var h = New[string](8, MinHeap[string])
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
		if !utils.SameStrings(arr1, arr2) {
			t.Fatal("error!")
		}
	}

	{
		var h = New[string](8, MaxHeap[string])
		var arr1 = make([]string, 0)
		var arr2 = make([]string, 0)
		for i := 0; i < count; i++ {
			var s = utils.Alphabet.Generate(8)
			h.Push(s)
			arr1 = append(arr1, s)
		}
		for h.Len() > 0 {
			arr2 = append(arr2, h.Pop())
		}
		sort.Strings(arr1)
		for i := 0; i < count/2; i++ {
			arr1[i], arr1[count-i-1] = arr1[count-i-1], arr1[i]
		}
		if !utils.SameStrings(arr1, arr2) {
			t.Fatal("error!")
		}
	}
}

func TestHeap_Sort(t *testing.T) {
	const count = 100
	var arr1 = make([]string, 0)
	var arr2 = make([]string, 0)
	for i := 0; i < count; i++ {
		var s = utils.Numeric.Generate(8)
		arr1 = append(arr1, s)
		arr2 = append(arr2, s)
	}
	arr2 = Init(arr2, MaxHeap[string]).Sort()
	sort.Strings(arr1)
	if !utils.SameStrings(arr1, arr2) {
		t.Fatal("error!")
	}
}
