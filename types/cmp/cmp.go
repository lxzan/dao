package cmp

type (
	Ordered interface {
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
			~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
			~float32 | ~float64 |
			~string
	}

	Number interface {
		Integer | ~float32 | ~float64
	}

	Integer interface {
		~int64 | ~int | ~int32 | ~int16 | ~int8 | ~uint64 | ~uint | ~uint32 | ~uint16 | ~uint8
	}
)

const (
	LT = -1 // 小于
	EQ = 0  // 等于
	GT = 1  // 大于
)

type (
	// LessFunc 比大小
	LessFunc[T any] func(a, b T) bool

	// CompareFunc 比较函数
	// a>b, 返回1; a<b, 返回-1; a==b, 返回0
	CompareFunc[T any] func(a, b T) int
)

// Less 比大小函数(升序)
func Less[T Ordered](x, y T) bool { return x < y }

// Compare 比较函数(升序)
func Compare[T Ordered](x, y T) int {
	if x < y {
		return -1
	}
	if x > y {
		return +1
	}
	return 0
}
