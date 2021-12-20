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
	if a > b {
		return Less
	} else if a < b {
		return Greater
	} else {
		return Equal
	}
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

type Hashable[T any] interface {
	~string | ~int64 | ~int | ~int32 | ~int16 | ~int8 | ~uint64 | ~uint | ~uint32 | ~uint16 | ~uint8 | ~float64 | ~float32
}

type Equaler[T any] interface {
	Equal(x *T) bool
}

type Hasher32[T any] interface {
	GetHashCode() uint32
	Equal(x T) bool
}

type Container interface {
	Begin() (iterator interface{})
	Next(iterator interface{}) interface{}
	End(iterator interface{}) bool
}
