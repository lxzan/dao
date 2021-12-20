package benchmark

import (
	"github.com/lxzan/dao/heap"
	"testing"
)

func BenchmarkHeap_Push(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var h = heap.New[int](b.N, heap.MinHeap[int])
		var n = len(testkeys)
		for j := 0; j < bench_count; j++ {
			h.Push(testvals[j%n])
		}
	}
}
