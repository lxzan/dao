package benchmark

import (
	"github.com/lxzan/dao"
	"github.com/lxzan/dao/heap"
	"math/rand"
	"testing"
)

func BenchmarkHeap_Push_Binary(b *testing.B) {
	var tpl = heap.New(dao.AscFunc[int]).SetForkNumber(heap.Binary)
	for j := 0; j < bench_count; j++ {
		tpl.Push(rand.Int())
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var h = tpl.Clone()
		for j := 0; j < bench_count; j++ {
			var val = testvals[j]
			h.Push(val)
		}
	}
}

func BenchmarkHeap_Push_Quadratic(b *testing.B) {
	var tpl = heap.New(dao.AscFunc[int]).SetForkNumber(heap.Quadratic)
	for j := 0; j < bench_count; j++ {
		tpl.Push(rand.Int())
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var h = tpl.Clone()
		for j := 0; j < bench_count; j++ {
			var val = testvals[j]
			h.Push(val)
		}
	}
}
func BenchmarkHeap_Push_Octal(b *testing.B) {
	var tpl = heap.New(dao.AscFunc[int]).SetForkNumber(heap.Octal)
	for j := 0; j < bench_count; j++ {
		tpl.Push(rand.Int())
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var h = tpl.Clone()
		for j := 0; j < bench_count; j++ {
			var val = testvals[j]
			h.Push(val)
		}
	}
}

func BenchmarkHeap_Pop_Binary(b *testing.B) {
	var tpl = heap.New(dao.AscFunc[int]).SetForkNumber(heap.Binary)
	for j := 0; j < bench_count*2; j++ {
		tpl.Push(rand.Int())
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := tpl.Clone()
		for j := 0; j < bench_count; j++ {
			h.Pop()
		}
	}
}

func BenchmarkHeap_Pop_Quadratic(b *testing.B) {
	var tpl = heap.New(dao.AscFunc[int]).SetForkNumber(heap.Quadratic)
	for j := 0; j < bench_count*2; j++ {
		tpl.Push(rand.Int())
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := tpl.Clone()
		for j := 0; j < bench_count; j++ {
			h.Pop()
		}
	}
}

func BenchmarkHeap_Pop_Octal(b *testing.B) {
	var tpl = heap.New(dao.AscFunc[int]).SetForkNumber(heap.Octal)
	for j := 0; j < bench_count*2; j++ {
		tpl.Push(rand.Int())
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := tpl.Clone()
		for j := 0; j < bench_count; j++ {
			h.Pop()
		}
	}
}
