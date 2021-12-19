package benchmark

import (
	"github.com/lxzan/dao/rapid"
	"testing"
)

const rapid_count = 10000

func BenchmarkRapid_New(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rapid.New[string, int](rapid_count)
	}
}

func BenchmarkRapid_Append(b *testing.B) {
	var r = rapid.New[string, int](rapid_count)
	var id1 = r.NextID()
	var q1 = rapid.EntryPoint{Head: id1, Tail: id1}
	var id2 = r.NextID()
	var q2 = rapid.EntryPoint{Head: id2, Tail: id2}
	for i := 0; i < b.N/2; i++ {
		r.Append(&q1, &rapid.Entry[string, int]{
			HashCode: 0,
			Key:      "hello",
			Val:      1,
		})
	}
	for i := 0; i < b.N/2; i++ {
		r.Append(&q2, &rapid.Entry[string, int]{
			HashCode: 0,
			Key:      "hello",
			Val:      1,
		})
	}
}
