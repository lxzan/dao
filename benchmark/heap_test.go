package benchmark

import (
	"github.com/lxzan/dao/heap"
	"github.com/lxzan/dao/types/cmp"
	"math/rand"
	"testing"
)

func BenchmarkHeap_Push_Binary(b *testing.B) {
	var tpl = heap.NewWithWays(heap.Binary, cmp.Less[int])
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
	var tpl = heap.NewWithWays(heap.Quadratic, cmp.Less[int])
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
	var tpl = heap.NewWithWays(heap.Octal, cmp.Less[int])
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
	var tpl = heap.NewWithWays(heap.Binary, cmp.Less[int])
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
	var tpl = heap.NewWithWays(heap.Quadratic, cmp.Less[int])
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
	var tpl = heap.NewWithWays(heap.Octal, cmp.Less[int])
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
