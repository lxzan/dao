package dao

type CompareResult int8

const (
	Less    CompareResult = -1
	Equal   CompareResult = 0
	Greater CompareResult = 1
)

type Order uint8

const (
	ASC  Order = 0
	DESC Order = 1
)

type Number interface {
	~int64 | ~int | ~int32 | ~int16 | ~int8 | ~uint64 | ~uint | ~uint32 | ~uint16 | ~uint8 | ~float32 | ~float64
}

type Integer interface {
	~int64 | ~int | ~int32 | ~int16 | ~int8 | ~uint64 | ~uint | ~uint32 | ~uint16 | ~uint8
}

// Map 键不可重复
type Map[K comparable, V any] interface {
	Len() int
	Get(key K) (V, bool)
	Set(key K, value V)
	Delete(key K)
	Range(f func(K, V) bool)
}

type Resetter interface {
	Reset()
}
