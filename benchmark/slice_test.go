package benchmark

import (
	"github.com/lxzan/dao/slice"
	"testing"
)

func BenchmarkSlice_Push(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var a = slice.New[int](0, 1000)
		for i := 0; i < 1000; i++ {
			a.Push(i)
		}
	}
}
