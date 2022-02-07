package benchmark

import (
	"github.com/lxzan/dao/heap"
	"testing"
)

func BenchmarkHeap_Push(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var h = heap.New(b.N, heap.MinHeap[int])
		var n = len(testkeys)
		for j := 0; j < bench_count; j++ {
			h.Push(testvals[j%n])
		}
	}
}

// a little slow
//func BenchmarkHashHeap_Find(b *testing.B) {
//	type entry struct {
//		Key string
//		Val int
//	}
//
//	var max_heap = func(a, b *entry) dao.Ordering {
//		if a.Key > b.Key {
//			return dao.Less
//		} else if a.Key == b.Key {
//			return dao.Equal
//		} else {
//			return dao.Greater
//		}
//	}
//
//	m := heap.New[*entry](bench_count, max_heap)
//	for i := 0; i < bench_count; i++ {
//		m.Push(&entry{Key: testkeys[i], Val: testvals[i]})
//	}
//
//	b.StartTimer()
//	for i := 0; i < b.N; i++ {
//		for j := 0; j < bench_count; j++ {
//			m.Find(&entry{Key: testkeys[i]})
//		}
//	}
//	b.StopTimer()
//}
