package benchmark

import (
	"github.com/lxzan/dao/internal/utils"
	"github.com/lxzan/dao/rapid"
	"testing"
)

func BenchmarkRapid_New(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rapid.New[string, int](bench_count)
	}
}

func BenchmarkRapid_Append(b *testing.B) {
	var arr = make([]string, 0, bench_count)
	var val = 1
	for i := 0; i < bench_count; i++ {
		arr = append(arr, utils.Alphabet.Generate(8))
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		var r = rapid.New[string, int](bench_count)
		var id1 = r.NextID()
		var q1 = rapid.EntryPoint{Head: id1, Tail: id1}
		var id2 = r.NextID()
		var q2 = rapid.EntryPoint{Head: id2, Tail: id2}

		for j := 0; j < bench_count/2; j++ {
			r.Append(&q1, arr[j], val)
		}
		for j := 0; j < bench_count/2; j++ {
			r.Append(&q2, arr[bench_count/2+j], val)
		}
	}
	b.StopTimer()
}

func BenchmarkRapid_Push(b *testing.B) {
	var arr = make([]string, 0, bench_count)
	var val = 1
	for i := 0; i < bench_count; i++ {
		arr = append(arr, utils.Alphabet.Generate(8))
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		var r = rapid.New[string, int](bench_count)
		var id1 = r.NextID()
		var q1 = rapid.EntryPoint{Head: id1, Tail: id1}
		var id2 = r.NextID()
		var q2 = rapid.EntryPoint{Head: id2, Tail: id2}

		for j := 0; j < bench_count/2; j++ {
			r.Push(&q1, arr[j], val)
		}
		for j := 0; j < bench_count/2; j++ {
			r.Push(&q2, arr[bench_count/2+j], val)
		}
	}
	b.StopTimer()
}
