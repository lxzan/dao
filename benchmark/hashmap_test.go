package benchmark

import (
	"github.com/lxzan/dao/hashmap"
	"github.com/lxzan/dao/internal/hash"
	"testing"
	"unsafe"
)

/**
16byte
hash.MapHash(b) // 15ns/op
hash.Fnv64(b) // 13 ns/op
hash.Fnv32(b) // 12 ns/op
hash.Base36Encode(b) // 18 ns/op

8 byte
hash.MapHash(b) // 14 ns/op
hash.Fnv64(b) // 8.2 ns/op
hash.Fnv32(b) // 6.8 ns/op
hash.Base36Encode(b) // 7.5 ns/op

4byte
hash.Fnv32(b) // 5 ns/op
hash.Base36Encode(b) // 4.5ns/op
*/

func BenchmarkHashMap_Hash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < bench_count; j++ {
			var b = *(*[]byte)(unsafe.Pointer(&testkeys[j]))
			hash.NewFnv32(b)
		}
	}
}

func BenchmarkHashMap_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := hashmap.New[string, int](bench_count)
		for j := 0; j < bench_count; j++ {
			m.Set(testkeys[j], testvals[j])
		}
	}
}

func BenchmarkGolang_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := make(map[string]int, bench_count)
		for j := 0; j < bench_count; j++ {
			m[testkeys[j]] = testvals[j]
		}
	}
}

func BenchmarkHashMap_Get(b *testing.B) {
	m := hashmap.New[string, int](bench_count)
	for j := 0; j < bench_count; j++ {
		m.Set(testkeys[j], testvals[j])
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < bench_count; j++ {
			m.Get(testkeys[j])
		}
	}
	b.StopTimer()
}

func BenchmarkGolang_Get(b *testing.B) {
	m := make(map[string]int, bench_count)
	for j := 0; j < bench_count; j++ {
		m[testkeys[j]] = testvals[j]
	}
	for i := 0; i < b.N; i++ {
		for j := 0; j < bench_count; j++ {
			_ = m[testkeys[j]]
		}
	}
}

func BenchmarkHashMap_Delete(b *testing.B) {
	m := hashmap.New[string, int](bench_count)
	for j := 0; j < bench_count; j++ {
		m.Set(testkeys[j], testvals[j])
	}
	for i := 0; i < b.N; i++ {
		for j := 0; j < bench_count/2; j++ {
			m.Delete(testkeys[j])
		}
	}
}

func BenchmarkGolang_Delete(b *testing.B) {
	m := make(map[string]int, bench_count)
	for j := 0; j < bench_count; j++ {
		m[testkeys[j]] = testvals[j]
	}

	for i := 0; i < b.N; i++ {
		for j := 0; j < bench_count/2; j++ {
			delete(m, testkeys[j])
		}
	}
}
