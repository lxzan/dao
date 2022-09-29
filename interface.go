package dao

type Ordering int8

const (
	Less    Ordering = -1
	Equal   Ordering = 0
	Greater Ordering = 1
)

type Comparer[T any] interface {
	Compare(a, b *T) Ordering
}

type Comparable interface {
	~string | ~int64 | ~int | ~int32 | ~int16 | ~int8 | ~uint64 | ~uint | ~uint32 | ~uint16 | ~uint8 | ~float64 | ~float32
}

func ASC[T Comparable](a, b T) Ordering {
	if a > b {
		return Greater
	} else if a < b {
		return Less
	} else {
		return Equal
	}
}

func DESC[T Comparable](a, b T) Ordering {
	return -1 * ASC(a, b)
}

type Number interface {
	~int64 | ~int | ~int32 | ~int16 | ~int8 | ~uint64 | ~uint | ~uint32 | ~uint16 | ~uint8 | ~float32 | ~float64
}

type Integer interface {
	~int64 | ~int | ~int32 | ~int16 | ~int8 | ~uint64 | ~uint | ~uint32 | ~uint16 | ~uint8
}

type Hashable interface {
	Integer | ~string
}

type Iterable[I any] interface {
	Begin() I
	Next(I) I
	End(I) bool
}

type Mapper[K Hashable, V any] interface {
	Set(key K, val V)
	Get(key K) (val V, exist bool)
	Delete(key K) bool
}
