package benchmark

import (
	"github.com/lxzan/dao/vector"
	"testing"
)

func BenchmarkSlice_Push(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var a = vector.New[int](0, bench_count)
		for j := 0; j < bench_count; j++ {
			a.Push(j)
		}
	}
}
