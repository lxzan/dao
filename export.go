package dao

type Ordering int8

const (
	Less    Ordering = -1
	Equal   Ordering = 0
	Greater Ordering = 1
)

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
	string | ~int | ~uint | float64 | float32 | byte
}

type Number[T any] interface {
	~int | ~uint | float32 | float64 | byte
}

type Integer[T any] interface {
	~int | ~uint
}

type Hashable[T any] interface{ string | ~int | ~uint | byte }

type Hasher32[T any] interface {
	GetHashCode() uint32
	Equal(x T) bool
}

type Container interface {
	Begin() (iterator interface{})
	Next(iterator interface{}) interface{}
	End(iterator interface{}) bool
}
