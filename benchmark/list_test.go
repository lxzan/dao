package benchmark

import (
	list2 "container/list"
	"github.com/lxzan/dao/deque"
	"github.com/lxzan/dao/linkedlist"
	"testing"
)

func BenchmarkStdList_Push(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var list = list2.New()
		for j := 0; j < bench_count; j++ {
			list.PushBack(j)
		}
	}
}

func BenchmarkStdList_PushAndPop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var list = list2.New()
		for j := 0; j < bench_count; j++ {
			list.PushBack(j)
		}
		for list.Len() > 0 {
			list.Remove(list.Front())
		}
	}
}

func BenchmarkLinkedList_Push(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var list = linkedlist.New[int]()
		for j := 0; j < bench_count; j++ {
			list.PushBack(j)
		}
	}
}

func BenchmarkLinkedList_PushAndPop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var list = linkedlist.New[int]()
		for j := 0; j < bench_count; j++ {
			list.PushBack(j)
		}
		for list.Len() > 0 {
			list.PopFront()
		}
	}
}

func BenchmarkDeque_Push(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var list = deque.New[int](bench_count)
		for j := 0; j < bench_count; j++ {
			list.PushBack(j)
		}
	}
}

func BenchmarkDeque_PushAndPop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var list = deque.New[int](bench_count)
		for j := 0; j < bench_count; j++ {
			list.PushBack(j)
		}
		for list.Len() > 0 {
			list.PopFront()
		}
	}
}
