package benchmark

import (
	"github.com/lxzan/dao/double_linkedlist"
	"testing"
)

func BenchmarkList_RPush(b *testing.B) {
	for i := 0; i < b.N; i++ {
		list := double_linkedlist.New[int]()
		for j := 0; j < bench_count; j++ {
			list.RPush(1)
		}
	}
}

func BenchmarkList_LPop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		list := double_linkedlist.New[int]()
		for j := 0; j < bench_count; j++ {
			list.RPush(1)
		}

		for j := 0; j < bench_count; j++ {
			list.LPop()
		}
	}
}
