package stack

// Stack 可以不使用New函数, 声明为值类型自动初始化
type Stack[T any] []T

// New 创建栈
func New[T any](capacity int) *Stack[T] {
	s := Stack[T](make([]T, 0, capacity))
	return &s
}

// NewFrom 从可变参数切片创建栈
func NewFrom[T any](values ...T) *Stack[T] {
	c := Stack[T](values)
	return &c
}

func (c *Stack[T]) Reset() {
	*c = (*c)[:0]
}

func (c *Stack[T]) Len() int {
	return len(*c)
}

func (c *Stack[T]) Push(v T) {
	*c = append(*c, v)
}

func (c *Stack[T]) Pop() (value T) {
	n := c.Len()
	switch n {
	case 0:
		return
	default:
		value = (*c)[n-1]
		*c = (*c)[:n-1]
		return
	}
}

func (c *Stack[T]) Range(f func(value T) bool) {
	for _, item := range *c {
		if !f(item) {
			return
		}
	}
}

// UnWrap 解包为切片
func (c *Stack[T]) UnWrap() []T {
	return *(*[]T)(c)
}
