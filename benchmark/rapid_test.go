package benchmark

import (
	"github.com/lxzan/dao/rapid"
	"testing"
)

const bench_count = 1000

type entry struct {
	Key string
	Val int
}

func (c entry) Equal(x *entry) bool {
	return c.Key == x.Key
}

func BenchmarkRapid_New(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rapid.New[entry](bench_count, func(a, b *entry) bool {
			return a.Key == b.Key
		})
	}
}

func BenchmarkRapid_Append(b *testing.B) {
	var arr = make([]*entry, 0, bench_count)
	for i := 0; i < bench_count; i++ {
		arr = append(arr, &entry{
			Key: "hello",
			Val: 1,
		})
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		var r = rapid.New[entry](1000, func(a, b *entry) bool {
			return a.Key == b.Key
		})
		var id1 = r.NextID()
		var q1 = rapid.EntryPoint{Head: id1, Tail: id1}
		var id2 = r.NextID()
		var q2 = rapid.EntryPoint{Head: id2, Tail: id2}

		for i := 0; i < bench_count/2; i++ {
			r.Append(&q1, arr[i])
		}
		for i := 0; i < bench_count/2; i++ {
			r.Append(&q2, arr[bench_count/2+i])
		}
	}
	b.StopTimer()
}
