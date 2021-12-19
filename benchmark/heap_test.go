package benchmark

import (
	"github.com/lxzan/dao/heap"
	"testing"
)

func BenchmarkMinHeap_Push(b *testing.B) {
	var h = heap.New[int](b.N, heap.MinHeap[int])
	var n = len(testkeys)
	for i := 0; i < b.N; i++ {
		h.Push(testvals[i%n])
	}
}

func BenchmarkMaxHeap_Push(b *testing.B) {
	var h = heap.New[int](b.N, heap.MaxHeap[int])
	var n = len(testkeys)
	for i := 0; i < b.N; i++ {
		h.Push(testvals[i%n])
	}
}
