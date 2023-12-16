package benchmark

import (
	"github.com/armon/go-radix"
	"github.com/lxzan/dao/dict"
	"testing"
)

func BenchmarkDict_Set(b *testing.B) {
	var d = dict.New[int]()
	for i := 0; i < b.N; i++ {
		key := testkeys[i%bench_count]
		d.Set(key, 1)
	}
}

func BenchmarkRadix_Set(b *testing.B) {
	var d = radix.New()
	for i := 0; i < b.N; i++ {
		key := testkeys[i%bench_count]
		d.Insert(key, 1)
	}
}

func BenchmarkDict_Get(b *testing.B) {
	var d = dict.New[int]()
	for j := 0; j < bench_count; j++ {
		d.Set(testkeys[j], 1)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var key = testkeys[i%bench_count]
		d.Get(key)
	}
}

func BenchmarkRadix_Get(b *testing.B) {
	var d = radix.New()
	for j := 0; j < bench_count; j++ {
		d.Insert(testkeys[j], 1)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var key = testkeys[i%bench_count]
		d.Get(key)
	}
}

func BenchmarkDict_Match(b *testing.B) {
	var d = dict.New[int]()
	for j := 0; j < bench_count; j++ {
		d.Set(testkeys[j], testvals[j])
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var prefix = testkeys[i%bench_count][:8]
		var num = 0
		d.Match(prefix, func(key string, value int) bool {
			num++
			return num < 10
		})
	}
}

func BenchmarkRadix_Match(b *testing.B) {
	var d = radix.New()
	for j := 0; j < bench_count; j++ {
		d.Insert(testkeys[j], testvals[j])
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var prefix = testkeys[i%bench_count][:8]
		var num = 0
		d.WalkPrefix(prefix, func(s string, v interface{}) bool {
			num++
			return num < 10
		})
	}
}
