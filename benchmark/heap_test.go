package benchmark

import (
	"github.com/lxzan/dao"
	"github.com/lxzan/dao/heap"
	"testing"
)

func BenchmarkHeap_Push_Binary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var h = heap.New(dao.AscFunc[int]).SetForkNumber(heap.Binary)
		for j := 0; j < bench_count; j++ {
			var val = testvals[j]
			h.Push(val)
		}
	}
}

func BenchmarkHeap_Push_Quadratic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var h = heap.New(dao.AscFunc[int]).SetForkNumber(heap.Quadratic)
		for j := 0; j < bench_count; j++ {
			var val = testvals[j]
			h.Push(val)
		}
	}
}
func BenchmarkHeap_Push_Octal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var h = heap.New(dao.AscFunc[int]).SetForkNumber(heap.Octal)
		for j := 0; j < bench_count; j++ {
			var val = testvals[j]
			h.Push(val)
		}
	}
}

func BenchmarkHeap_PushAndPop_Binary(b *testing.B) {
	var h = heap.New(dao.AscFunc[int]).SetForkNumber(heap.Binary)
	for j := 0; j < bench_count; j++ {
		h.Push(testvals[j])
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < bench_count; j++ {
			var val = testvals[j]
			h.Push(val)
			h.Pop()
		}
	}
}

func BenchmarkHeap_PushAndPop_Quadratic(b *testing.B) {
	var h = heap.New(dao.AscFunc[int]).SetForkNumber(heap.Quadratic)
	for j := 0; j < bench_count; j++ {
		h.Push(testvals[j])
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < bench_count; j++ {
			var val = testvals[j]
			h.Push(val)
			h.Pop()
		}
	}
}

func BenchmarkHeap_PushAndPop_Octal(b *testing.B) {
	var h = heap.New(dao.AscFunc[int]).SetForkNumber(heap.Octal)
	for j := 0; j < bench_count; j++ {
		h.Push(testvals[j])
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < bench_count; j++ {
			var val = testvals[j]
			h.Push(val)
			h.Pop()
		}
	}
}
