package hashmap

type comparable[T any] interface{ type string, int8, int16, int32, int64, int, uint8, uint16, uint32, uint64, uint }

type element[K comparable[K], V any] struct {
	Ptr     uint32
	NextPtr uint32
	Data    entry[K, V]
}

type entry[K comparable[K], V any] struct {
	HashCode uint32
	Key      K
	Val      V
}
