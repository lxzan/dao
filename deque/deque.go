package deque

import (
	"github.com/lxzan/dao/stack"
)

const Nil = 0

type (
	Pointer uint32

	Element[T any] struct {
		prev, addr, next Pointer
		value            T
	}

	Deque[T any] struct {
		head, tail Pointer              // 头尾指针
		length     int                  // 长度
		stack      stack.Stack[Pointer] // 回收站
		elements   []Element[T]         // 元素列表
		template   Element[T]           // 空值模板
	}
)

func (c Pointer) IsNil() bool {
	return c == Nil
}

func (c *Element[T]) Addr() Pointer {
	return c.addr
}

func (c *Element[T]) Next() Pointer {
	return c.next
}

func (c *Element[T]) Prev() Pointer {
	return c.prev
}

func (c *Element[T]) Value() T {
	return c.value
}

func New[T any](capacity uint32) *Deque[T] {
	return &Deque[T]{elements: make([]Element[T], 1, 1+capacity)}
}

func (c *Deque[T]) Get(addr Pointer) *Element[T] {
	if addr > 0 {
		return &(c.elements[addr])
	}
	return nil
}

func (c *Deque[T]) getElement() *Element[T] {
	if c.stack.Len() > 0 {
		addr := c.stack.Pop()
		v := c.Get(addr)
		v.addr = addr
		return v
	}

	addr := Pointer(len(c.elements))
	c.elements = append(c.elements, c.template)
	v := c.Get(addr)
	v.addr = addr
	return v
}

func (c *Deque[T]) putElement(ele *Element[T]) {
	c.stack.Push(ele.addr)
	*ele = c.template
}

func (c *Deque[T]) Reset() {
	c.head, c.tail, c.length = Nil, Nil, 0
	c.stack = c.stack[:0]
	c.elements = c.elements[:1]
}

func (c *Deque[T]) Len() int {
	return c.length
}

func (c *Deque[T]) Front() *Element[T] {
	return c.Get(c.head)
}

func (c *Deque[T]) Back() *Element[T] {
	return c.Get(c.tail)
}

func (c *Deque[T]) PushFront(value T) *Element[T] {
	ele := c.getElement()
	ele.value = value
	c.doPushFront(ele)
	return ele
}

func (c *Deque[T]) doPushFront(ele *Element[T]) {
	c.length++

	if !c.head.IsNil() {
		head := c.Get(c.head)
		head.prev = ele.addr
		ele.next = head.addr
		c.head = ele.addr
		return
	}

	c.head = ele.addr
	c.tail = ele.addr
}

func (c *Deque[T]) PushBack(value T) *Element[T] {
	ele := c.getElement()
	ele.value = value
	c.doPushBack(ele)
	return ele
}

func (c *Deque[T]) doPushBack(ele *Element[T]) {
	c.length++

	if !c.tail.IsNil() {
		tail := c.Get(c.tail)
		tail.next = ele.addr
		ele.prev = tail.addr
		c.tail = ele.addr
		return
	}

	c.head = ele.addr
	c.tail = ele.addr
}

func (c *Deque[T]) PopFront() (value T) {
	ele := c.Get(c.head)
	if ele == nil {
		return value
	}

	c.head = ele.next
	if head := c.Get(c.head); head != nil {
		head.prev = Nil
	}

	c.length--
	value = ele.value
	c.putElement(ele)
	if c.length == 0 {
		c.tail = Nil
	}

	return value
}

func (c *Deque[T]) PopBack() (value T) {
	ele := c.Get(c.tail)
	if ele == nil {
		return value
	}

	c.tail = ele.prev
	if tail := c.Get(c.tail); tail != nil {
		tail.next = Nil
	}

	c.length--
	value = ele.value
	c.putElement(ele)
	if c.length == 0 {
		c.head = Nil
	}

	return value
}

func (c *Deque[T]) InsertAfter(value T, mark Pointer) *Element[T] {
	prevEle := c.Get(mark)
	if prevEle == nil {
		return nil
	}

	c.length++
	ele := c.getElement()
	ele.prev, ele.next, ele.value = mark, prevEle.next, value
	prevEle.next = ele.addr

	if nextEle := c.Get(ele.next); nextEle != nil {
		nextEle.prev = ele.addr
		return ele
	}

	c.tail = ele.addr
	return ele
}

func (c *Deque[T]) InsertBefore(value T, mark Pointer) *Element[T] {
	nextEle := c.Get(mark)
	if nextEle == nil {
		return nil
	}

	c.length++
	ele := c.getElement()
	ele.prev, ele.next, ele.value = nextEle.prev, mark, value
	nextEle.prev = ele.addr

	if prevEle := c.Get(ele.prev); prevEle != nil {
		prevEle.next = ele.addr
		return ele
	}

	c.head = ele.addr
	return ele
}

func (c *Deque[T]) MoveToBack(addr Pointer) {
	if ele := c.Get(addr); ele != nil {
		c.doRemove(ele)
		ele.prev, ele.next = Nil, Nil
		c.doPushBack(ele)
	}
}

func (c *Deque[T]) MoveToFront(addr Pointer) {
	if ele := c.Get(addr); ele != nil {
		c.doRemove(ele)
		ele.prev, ele.next = Nil, Nil
		c.doPushFront(ele)
	}
}

func (c *Deque[T]) Update(addr Pointer, value T) {
	if ele := c.Get(addr); ele != nil {
		ele.value = value
	}
}

func (c *Deque[T]) Remove(addr Pointer) {
	if ele := c.Get(addr); ele != nil {
		c.doRemove(ele)
		c.putElement(ele)
	}
}

func (c *Deque[T]) doRemove(ele *Element[T]) {
	var prev, next *Element[T] = nil, nil
	var state = 0
	if !ele.prev.IsNil() {
		prev = c.Get(ele.prev)
		state += 1
	}
	if !ele.next.IsNil() {
		next = c.Get(ele.next)
		state += 2
	}

	switch state {
	case 3:
		prev.next = next.addr
		next.prev = prev.addr
	case 2:
		next.prev = Nil
		c.head = next.addr
	case 1:
		prev.next = Nil
		c.tail = prev.addr
	default:
		c.head = Nil
		c.tail = Nil
	}

	c.length--
}

func (c *Deque[T]) Range(f func(ele *Element[T]) bool) {
	for i := c.Get(c.head); i != nil; i = c.Get(i.next) {
		if !f(i) {
			break
		}
	}
}
