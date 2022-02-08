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

func ASC[T Comparable[T]](a, b T) Ordering {
	if a > b {
		return Greater
	} else if a < b {
		return Less
	} else {
		return Equal
	}
}

func DESC[T Comparable[T]](a, b T) Ordering {
	return -1 * ASC(a, b)
}

type Comparable[T any] interface {
	~string | ~int64 | ~int | ~int32 | ~int16 | ~int8 | ~uint64 | ~uint | ~uint32 | ~uint16 | ~uint8 | ~float64 | ~float32
}

type Number[T any] interface {
	~int64 | ~int | ~int32 | ~int16 | ~int8 | ~uint64 | ~uint | ~uint32 | ~uint16 | ~uint8 | ~float32 | ~float64
}

type Integer[T any] interface {
	~int64 | ~int | ~int32 | ~int16 | ~int8 | ~uint64 | ~uint | ~uint32 | ~uint16 | ~uint8
}

type Hasher32[T any] interface {
	GetHashCode() uint32
}

type Container interface {
	Begin() (iterator interface{})
	Next(iterator interface{}) interface{}
	End(iterator interface{}) bool
}
