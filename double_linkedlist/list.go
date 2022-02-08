package double_linkedlist

func New[T any]() *List[T] {
	return new(List[T])
}

// safe delete in loop
type Iterator[T any] struct {
	prev, next *Iterator[T]
	Data       T
}

type List[T any] struct {
	length int
	head   *Iterator[T]
	tail   *Iterator[T]
}

func (this List[T]) Begin() *Iterator[T] {
	return this.head
}

func (this *List[T]) Next(iter *Iterator[T]) *Iterator[T] {
	return iter.next
}

func (this List[T]) End(iter *Iterator[T]) bool {
	return iter == nil
}

func (this List[T]) Len() int {
	return this.length
}

func (this List[T]) Clear() {
	this.head = nil
	this.tail = nil
	this.length = 0
}

func (this *List[T]) Front() *Iterator[T] {
	return this.head
}

func (this *List[T]) Back() *Iterator[T] {
	return this.tail
}

func (this *List[T]) RPush(values ...T) {
	for _, v := range values {
		var ele = &Iterator[T]{Data: v}
		if this.length > 0 {
			this.tail.next = ele
			ele.prev = this.tail
			this.tail = ele
		} else {
			this.head = ele
			this.tail = ele
		}
		this.length++
	}
}

func (this *List[T]) LPush(values ...T) {
	for _, v := range values {
		var ele = &Iterator[T]{Data: v}
		if this.length > 0 {
			ele.next = this.head
			this.head.prev = ele
			this.head = ele
		} else {
			this.head = ele
			this.tail = ele
		}
		this.length++
	}
}

func (this *List[T]) LPop() *Iterator[T] {
	if this.length == 0 {
		return nil
	}
	var result = this.head
	this.head = this.head.next
	result.next = nil
	if this.head != nil {
		this.head.prev = nil
	}
	if this.length == 1 {
		this.tail = nil
	}
	this.length--
	return result
}

func (this *List[T]) RPop() *Iterator[T] {
	if this.length == 0 {
		return nil
	}

	var result = this.tail
	this.tail = this.tail.prev
	result.prev = nil
	if this.tail != nil {
		this.tail.next = nil
	}
	if this.length == 1 {
		this.head = nil
	}
	this.length--
	return result
}

// Delete it's safe delete in loop
func (this *List[T]) Delete(iter *Iterator[T]) {
	var prev = iter.prev
	var next = iter.next
	if prev != nil && next != nil {
		prev.next = next
		next.prev = prev
	} else if prev != nil && next == nil {
		prev.next = nil
		this.tail = prev
	} else if prev == nil && next != nil {
		next.prev = nil
		this.head = next
	} else {
		this.head = nil
		this.tail = nil
	}
	this.length--
}

func (this *List[T]) InsertAfter(iter *Iterator[T], v T) {
	var next = iter.next
	var cur = &Iterator[T]{Data: v, prev: iter, next: next}
	iter.next = cur
	if next != nil {
		next.prev = cur
	}
	this.length++
}

func (this *List[T]) InsertBefore(iter *Iterator[T], v T) {
	var prev = iter.prev
	var cur = &Iterator[T]{Data: v, prev: prev, next: iter}
	iter.prev = cur
	if prev != nil {
		prev.next = cur
	}
	this.length++
}
