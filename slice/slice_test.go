package slice

import (
	"github.com/lxzan/dao"
	"testing"
)

func TestNew(t *testing.T) {
	var a = Slice[int]([]int{1, 3, 5, 7, 9})
	for i := a.Begin(); !a.End(i); i = a.Next(i) {
		println(i.Index, i.Value)
	}
}

func TestSlice_Sort(t *testing.T) {
	var slice Slice[int] = New[int]()
	slice.Push(1, 3, 5, 7, 9, 2, 4, 6, 8, 0)
	slice.Sort(dao.ASC[int])
	println(&slice)
}
