package mlist

const Nil = 0

type (
	Pointer uint32

	EntryPoint struct{ Head, Tail Pointer }

	// Element 链表元素
	// 不能在外部引用*Element, 必须使用Pointer.
	Element[K comparable, V any] struct {
		Prev, Addr, Next Pointer
		Key              K
		Value            V
	}
)

// MList 多维链表
// 用于解决哈希冲突
// 在一个数组里面维护多条链表, 同个链表里面数据不能重复.
type MList[K comparable, V any] struct {
	length     int
	serial     uint32
	recyclable arrayStack
	template   Element[K, V]
	Buckets    []Element[K, V]
}

func NewMList[K comparable, V any](size uint32) *MList[K, V] {
	return &MList[K, V]{
		Buckets: make([]Element[K, V], 1, size+1),
	}
}

func (c *MList[K, V]) Reset() {
	c.length, c.serial = 0, 0
	c.recyclable = c.recyclable[:0]
	c.Buckets = c.Buckets[:1]
}

func (c *MList[K, V]) getElement() Pointer {
	if c.recyclable.Len() > 0 {
		return c.recyclable.Pop()
	}

	c.serial++
	var addr = c.serial
	c.Buckets = append(c.Buckets, c.template)
	return Pointer(addr)
}

func (c *MList[K, V]) putElement(addr Pointer) {
	c.Buckets[addr] = c.template
	c.recyclable.Push(addr)
}

func (c *MList[K, V]) newElement(key K, value V) *Element[K, V] {
	addr := c.getElement()
	ele := c.Get(addr)
	ele.Addr, ele.Key, ele.Value = addr, key, value
	c.length++
	return ele
}

func (c *MList[K, V]) Len() int {
	return c.length
}

func (c *MList[K, V]) Get(addr Pointer) *Element[K, V] {
	if addr > 0 {
		return &c.Buckets[addr]
	}
	return nil
}

func (c *MList[K, V]) Range(entrypoint *EntryPoint, f func(iter *Element[K, V]) bool) {
	for i := c.Get(entrypoint.Head); i != nil; i = c.Get(i.Next) {
		if !f(i) {
			return
		}
	}
}

// Push append an Element[] with unique check
func (c *MList[K, V]) Push(entrypoint *EntryPoint, key K, value V) (addr Pointer, exist bool) {
	// 链表为空
	if entrypoint.Head == Nil {
		ele := c.newElement(key, value)
		entrypoint.Head, entrypoint.Tail = ele.Addr, ele.Addr
		return ele.Addr, false
	}

	// key存在, 替换value
	for i := c.Get(entrypoint.Head); i != nil; i = c.Get(i.Next) {
		if i.Key == key {
			i.Value = value
			return i.Addr, true
		}
	}

	// key不存在, 插入新数据
	ele := c.newElement(key, value)
	tail := c.Get(entrypoint.Tail)
	tail.Next = ele.Addr
	ele.Prev = tail.Addr
	entrypoint.Tail = ele.Addr
	return ele.Addr, false
}

func (c *MList[K, V]) Delete(entrypoint *EntryPoint, key K) (exists bool) {
	for i := c.Get(entrypoint.Head); i != nil; i = c.Get(i.Next) {
		if i.Key == key {
			c.doDelete(entrypoint, i)
			return true
		}
	}
	return false
}

// Delete do not delete in loop if no break
func (c *MList[K, V]) doDelete(entrypoint *EntryPoint, ele *Element[K, V]) {
	var prev, next *Element[K, V] = nil, nil
	var state = 0
	if ele.Prev > 0 {
		prev = c.Get(ele.Prev)
		state += 1
	}
	if ele.Next > 0 {
		next = c.Get(ele.Next)
		state += 2
	}

	switch state {
	case 3:
		prev.Next = next.Addr
		next.Prev = prev.Addr
	case 2:
		next.Prev = 0
		entrypoint.Head = next.Addr
	case 1:
		prev.Next = 0
		entrypoint.Tail = prev.Addr
	default:
		entrypoint.Head = 0
		entrypoint.Tail = 0
	}

	c.putElement(ele.Addr)
	c.length--
}

func (c *MList[K, V]) Find(entrypoint *EntryPoint, key K) (value V, exist bool) {
	for i := c.Get(entrypoint.Head); i != nil; i = c.Get(i.Next) {
		if i.Key == key {
			return i.Value, true
		}
	}
	return value, false
}

type arrayStack []Pointer

func (c *arrayStack) Len() int {
	return len(*c)
}

func (c *arrayStack) Push(v Pointer) {
	*c = append(*c, v)
}

func (c *arrayStack) Pop() Pointer {
	var n = len(*c)
	var v = (*c)[n-1]
	*c = (*c)[:n-1]
	return v
}
