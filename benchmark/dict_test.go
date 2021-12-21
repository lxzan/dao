package benchmark

import (
	"github.com/lxzan/dao/dict"
	"github.com/lxzan/dao/internal/utils"
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

func BenchmarkDict_Match(b *testing.B) {
	var d = dict.New[int]()
	for j := 0; j < bench_count; j++ {
		var length = utils.Rand.Intn(16) + 1
		d.Insert(utils.Numeric.Generate(length), testvals[j])
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		var prefix = utils.Numeric.Generate(4)
		for j := 0; j < bench_count; j++ {
			d.Match(prefix, 10)
		}
	}
}

func BenchmarkDict_Delete(b *testing.B) {
	var d = dict.New[int]()
	for j := 0; j < bench_count; j++ {
		d.Insert(testkeys[j], testvals[j])
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < bench_count; j++ {
			d.Delete(testkeys[j])
		}
	}
}
