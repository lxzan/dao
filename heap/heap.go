package heap

func newHeap[T any](less func(a, b T) bool) *heap[T] {
	var obj heap[T]
	obj.data = make([]T, 0)
	obj.less = less
	return &obj
}

type heap[T any] struct {
	data []T
	less func(a, b T) bool
}

func (c heap[T]) Len() int {
	return len(c.data)
}

func (c heap[T]) Get(i int) T {
	return c.data[i]
}

func (c *heap[T]) Swap(i, j int) {
	c.data[i], c.data[j] = c.data[j], c.data[i]
}

func (c *heap[T]) Push(eles ...T) {
	for _, item := range eles {
		c.data = append(c.data, item)
		c.up(c.Len() - 1)
	}
}

func (c *heap[T]) up(i int) {
	var j = (i - 1) / 2
	if j >= 0 && c.less(c.data[i], c.data[j]) {
		c.Swap(i, j)
		c.up(j)
	}
}

func (c *heap[T]) Pop() T {
	var n = c.Len()
	var result = c.data[0]
	c.data[0] = c.data[n-1]
	c.data = c.data[:n-1]
	c.down(0, n-1)
	return result
}

func (c *heap[T]) down(i, n int) {
	var j = 2*i + 1
	if j < n && c.less(c.data[j], c.data[i]) {
		c.Swap(i, j)
		c.down(j, n)
	}
	var k = 2*i + 2
	if k < n && c.less(c.data[k], c.data[i]) {
		c.Swap(i, k)
		c.down(k, n)
	}
}
