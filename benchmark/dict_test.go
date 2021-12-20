package benchmark

import (
	"github.com/lxzan/dao/dict"
	"testing"
)

func BenchmarkDict_Insert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var d = dict.New[int]()
		for j := 0; j < bench_count; j++ {
			d.Insert(testkeys[j], testvals[j])
		}
	}
}

func BenchmarkDict_Delete(b *testing.B) {
	var d = dict.New[int]()
	for j := 0; j < bench_count; j++ {
		d.Insert(testkeys[j], testvals[j])
	}

	for i := 0; i < b.N; i++ {
		for j := 0; j < bench_count; j++ {
			d.Delete(testkeys[j])
		}
	}
}
