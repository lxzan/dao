package linkedlist

type (
	Element[T any] struct {
		prev, next *Element[T]
		Value      T
	}

	// LinkedList 可以不使用New函数, 声明为值类型自动初始化
	LinkedList[T any] struct {
		head, tail *Element[T] // 头尾指针
		length     int         // 长度
	}
)

func (c *Element[T]) Next() *Element[T] {
	return c.next
}

func (c *Element[T]) Prev() *Element[T] {
	return c.prev
}

// New 创建双向链表
func New[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// Reset 重置链表
func (c *LinkedList[T]) Reset() {
	c.head, c.tail, c.length = nil, nil, 0
}

// Len 获取链表长度
func (c *LinkedList[T]) Len() int {
	return c.length
}

// Front 获取头部元素
func (c *LinkedList[T]) Front() *Element[T] {
	return c.head
}

// Back 获取尾部元素
func (c *LinkedList[T]) Back() *Element[T] {
	return c.tail
}

// PushFront 向头部追加元素
func (c *LinkedList[T]) PushFront(value T) *Element[T] {
	ele := &Element[T]{Value: value}
	c.doPushFront(ele)
	return ele
}

func (c *LinkedList[T]) doPushFront(ele *Element[T]) {
	c.length++

	if c.head != nil {
		head := c.head
		head.prev = ele
		ele.next = head
		c.head = ele
		return
	}

	c.head = ele
	c.tail = ele
}

// PushBack 向尾部追加元素
func (c *LinkedList[T]) PushBack(value T) *Element[T] {
	ele := &Element[T]{Value: value}
	c.doPushBack(ele)
	return ele
}

func (c *LinkedList[T]) doPushBack(ele *Element[T]) {
	c.length++

	if c.tail != nil {
		tail := c.tail
		tail.next = ele
		ele.prev = tail
		c.tail = ele
		return
	}

	c.head = ele
	c.tail = ele
}

// PopFront 从头部弹出元素
func (c *LinkedList[T]) PopFront() (value T) {
	ele := c.head
	if ele == nil {
		return value
	}

	c.head = ele.next
	if head := c.head; head != nil {
		head.prev = nil
	}

	c.length--
	ele.prev, ele.next = nil, nil
	if c.length == 0 {
		c.tail = nil
	}

	return ele.Value
}

// PopBack 从尾部弹出元素
func (c *LinkedList[T]) PopBack() (value T) {
	ele := c.tail
	if ele == nil {
		return value
	}

	c.tail = ele.prev
	if tail := c.tail; tail != nil {
		tail.next = nil
	}

	c.length--
	ele.prev, ele.next = nil, nil
	if c.length == 0 {
		c.head = nil
	}

	return ele.Value
}

// InsertAfter 在标记位置后面插入元素
func (c *LinkedList[T]) InsertAfter(value T, mark *Element[T]) *Element[T] {
	prevEle := mark
	if prevEle == nil {
		return nil
	}

	c.length++
	ele := &Element[T]{}
	ele.prev, ele.next, ele.Value = mark, prevEle.next, value
	prevEle.next = ele

	if nextEle := ele.next; nextEle != nil {
		nextEle.prev = ele
		return ele
	}

	c.tail = ele
	return ele
}

// InsertBefore 在标记位置前面插入元素
func (c *LinkedList[T]) InsertBefore(value T, mark *Element[T]) *Element[T] {
	nextEle := mark
	if nextEle == nil {
		return nil
	}

	c.length++
	ele := &Element[T]{}
	ele.prev, ele.next, ele.Value = nextEle.prev, mark, value
	nextEle.prev = ele

	if prevEle := ele.prev; prevEle != nil {
		prevEle.next = ele
		return ele
	}

	c.head = ele
	return ele
}

// MoveToBack 将指定元素移至尾部
func (c *LinkedList[T]) MoveToBack(ele *Element[T]) {
	if ele != nil {
		c.doRemove(ele)
		ele.prev, ele.next = nil, nil
		c.doPushBack(ele)
	}
}

// MoveToFront 将指定元素移至头部
func (c *LinkedList[T]) MoveToFront(ele *Element[T]) {
	if ele != nil {
		c.doRemove(ele)
		ele.prev, ele.next = nil, nil
		c.doPushFront(ele)
	}
}

// Remove 移除元素
func (c *LinkedList[T]) Remove(ele *Element[T]) {
	if ele != nil {
		c.doRemove(ele)
	}
}

func (c *LinkedList[T]) doRemove(ele *Element[T]) {
	var prev, next *Element[T] = nil, nil
	var state = 0
	if ele.prev != nil {
		prev = ele.prev
		state += 1
	}
	if ele.next != nil {
		next = ele.next
		state += 2
	}

	switch state {
	case 3:
		prev.next = next
		next.prev = prev
	case 2:
		next.prev = nil
		c.head = next
	case 1:
		prev.next = nil
		c.tail = prev
	default:
		c.head = nil
		c.tail = nil
	}

	c.length--
}

// Range 遍历链表
func (c *LinkedList[T]) Range(f func(ele *Element[T]) bool) {
	for i := c.Front(); i != nil; i = i.Next() {
		if !f(i) {
			break
		}
	}
}
