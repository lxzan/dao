package dao

type (
	// Map 键不可重复
	Map[K comparable, V any] interface {
		Len() int
		Get(key K) (V, bool)
		Set(key K, value V)
		Delete(key K)
		Range(f func(K, V) bool)
	}

	Resetter interface {
		Reset()
	}

	Cloner[T any] interface {
		// Clone 深拷贝
		Clone() T
	}
)
