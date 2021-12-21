package benchmark

import (
	"github.com/lxzan/dao"
	"github.com/lxzan/dao/hashmap"
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

func BenchmarkHashHeap_Find(b *testing.B) {
	type entry struct {
		Key string
		Val int
	}

	var max_heap = func(a, b entry) dao.Ordering {

	}

	for i := 0; i < b.N; i++ {
		m := hashmap.New[string, int](bench_count)
		for j := 0; j < bench_count; j++ {
			m.Set(testkeys[j], testvals[j])
		}
	}
}
