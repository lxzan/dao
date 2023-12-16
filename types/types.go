package types

const Nil = 0

type (
	Pointer uint32

	Pair[K any, V any] struct {
		Key   K
		Value V
	}

	Integer interface {
		~int64 | ~int | ~int32 | ~int16 | ~int8 | ~uint64 | ~uint | ~uint32 | ~uint16 | ~uint8
	}

	Map[K comparable, V any] interface {
		Len() int
		Get(key K) (V, bool)
		Set(key K, value V) bool
		Delete(key K) bool
		Range(f func(K, V) bool)
	}
)
