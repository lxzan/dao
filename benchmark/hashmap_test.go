package benchmark

import (
	"github.com/lxzan/dao/hashmap"
	"github.com/lxzan/dao/internal/hash"
	"github.com/lxzan/dao/internal/utils"
	"testing"
)

func BenchmarkFnv64(b *testing.B) {
	var s = utils.S2B(utils.Alphabet.Generate(16))
	for i := 0; i < b.N; i++ {
		for j := 0; j < bench_count; j++ {
			hash.HashBytes64(s)
		}
	}
}

func BenchmarkMyMap_Set(b *testing.B) {
	b.Run("string", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m := hashmap.New[string, int](bench_count)
			for j := 0; j < bench_count; j++ {
				m.Set(testkeys[j], 1)
			}
		}
	})

	b.Run("int", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m := hashmap.New[int, int](bench_count)
			for j := 0; j < bench_count; j++ {
				m.Set(testvals[j], 1)
			}
		}
	})
}

func BenchmarkGoMap_Set(b *testing.B) {
	b.Run("string", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m := make(map[string]int, bench_count)
			for j := 0; j < bench_count; j++ {
				m[testkeys[j]] = 1
			}
		}
	})

	b.Run("int", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m := make(map[int]int, bench_count)
			for j := 0; j < bench_count; j++ {
				m[testvals[j]] = 1
			}
		}
	})
}

func BenchmarkMyMap_Get(b *testing.B) {
	b.Run("string", func(b *testing.B) {
		m := hashmap.New[string, uint32](bench_count)
		for j := 0; j < bench_count; j++ {
			m.Set(testkeys[j], 1)
		}
		for i := 0; i < b.N; i++ {
			for j := 0; j < bench_count; j++ {
				m.Get(testkeys[j])
			}
		}
	})

	b.Run("int", func(b *testing.B) {
		m := hashmap.New[int, string](bench_count)
		for j := 0; j < bench_count; j++ {
			m.Set(testvals[j], testkeys[j])
		}
		for i := 0; i < b.N; i++ {
			for j := 0; j < bench_count; j++ {
				m.Get(testvals[j])
			}
		}
	})
}

func BenchmarkGoMap_Get(b *testing.B) {
	b.Run("string", func(b *testing.B) {
		m := make(map[string]int, bench_count)
		for j := 0; j < bench_count; j++ {
			m[testkeys[j]] = 1
		}
		for i := 0; i < b.N; i++ {
			for j := 0; j < bench_count; j++ {
				_ = m[testkeys[j]]
			}
		}
	})

	b.Run("int", func(b *testing.B) {
		m := make(map[int]string, bench_count)
		for j := 0; j < bench_count; j++ {
			m[testvals[j]] = testkeys[j]
		}
		for i := 0; i < b.N; i++ {
			for j := 0; j < bench_count; j++ {
				_ = m[testvals[j]]
			}
		}
	})
}

func BenchmarkMyMap_Delete(b *testing.B) {
	b.Run("string", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m := hashmap.New[string, int](bench_count)
			for j := 0; j < bench_count; j++ {
				m.Set(testkeys[j], 1)
			}

			b.StartTimer()
			for j := 0; j < bench_count; j++ {
				m.Delete(testkeys[j])
			}
			b.StopTimer()
		}
	})
	b.Run("int", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m := hashmap.New[int, int](bench_count)
			for j := 0; j < bench_count; j++ {
				m.Set(testvals[j], 1)
			}

			b.StartTimer()
			for j := 0; j < bench_count; j++ {
				m.Delete(testvals[j])
			}
			b.StopTimer()
		}
	})
}

func BenchmarkGoMap_Delete(b *testing.B) {
	b.Run("string", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m := make(map[string]int, bench_count)
			for j := 0; j < bench_count; j++ {
				m[testkeys[j]] = 1
			}

			b.StartTimer()
			for j := 0; j < bench_count; j++ {
				delete(m, testkeys[j])
			}
			b.StopTimer()
		}
	})

	b.Run("int", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m := make(map[int]int, bench_count)
			for j := 0; j < bench_count; j++ {
				m[testvals[j]] = 1
			}

			b.StartTimer()
			for j := 0; j < bench_count; j++ {
				delete(m, testvals[j])
			}
			b.StopTimer()
		}
	})
}
