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

// Reset 重置
func (c *Stack[T]) Reset() {
	*c = (*c)[:0]
}

// Len 获取元素数量
func (c *Stack[T]) Len() int {
	return len(*c)
}

// Push 追加元素
func (c *Stack[T]) Push(v T) {
	*c = append(*c, v)
}

// Pop 弹出元素
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

// Range 遍历
func (c *Stack[T]) Range(f func(value T) bool) {
	for _, item := range *c {
		if !f(item) {
			return
		}
	}
}

// UnWrap 解包, 返回底层数组
func (c *Stack[T]) UnWrap() []T {
	return *(*[]T)(c)
}
