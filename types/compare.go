package types

type Ordering int8

const (
	Less    Ordering = -1
	Equal   Ordering = 0
	Greater Ordering = 1
)

//type comparable[T any] interface{ type int8, int16, int32, int64, int, uint8, uint16, uint32, uint64, uint, string }

//type inteage[T any] interface{ type int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8 }

//type interator[Collection any, Iterator any] interface {
//	ForEach(func(Iterator))
//}
