package linkedlist

type Iterator[T any] struct {
	next *Iterator[T]
	Data T
}

func NewQueue[T any]() *Queue[T] {
	return new(Queue[T])
}

type Queue[T any] struct {
	length int
	head   *Iterator[T]
	tail   *Iterator[T]
}

func (this *Queue[T]) Clear() {
	this.head = nil
	this.tail = nil
	this.length = 0
}

func (this *Queue[T]) Begin() *Iterator[T] {
	return this.head
}

func (this *Queue[T]) Next(iter *Iterator[T]) *Iterator[T] {
	return iter.next
}

func (this *Queue[T]) End(iter *Iterator[T]) bool {
	return iter == nil
}

func (this *Queue[T]) Len() int {
	return this.length
}

func (this *Queue[T]) Push(values ...T) {
	for _, v := range values {
		var ele = &Iterator[T]{Data: v}
		if this.length > 0 {
			this.tail.next = ele
			this.tail = ele
		} else {
			this.head = ele
			this.tail = ele
		}
		this.length++
	}
}

func (this *Queue[T]) Front() *Iterator[T] {
	return this.head
}

func (this *Queue[T]) Pop() *Iterator[T] {
	if this.length == 0 {
		return nil
	}
	var result = this.head
	this.head = this.head.next
	this.length--
	if this.length == 0 {
		this.tail = nil
	}
	return result
}
