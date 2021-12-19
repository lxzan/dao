package benchmark

import (
	"github.com/lxzan/dao/double_linkedlist"
	"testing"
)

func BenchmarkList_RPush(b *testing.B) {
	list := double_linkedlist.New[int]()
	for i := 0; i < b.N; i++ {
		list.RPush(1)
	}
}
