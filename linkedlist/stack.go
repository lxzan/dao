package linkedlist

func NewStack[T any]() *Stack[T] {
	return new(Stack[T])
}

type Stack[T any] struct {
	length int
	head   *Iterator[T]
}

func (this *Stack[T]) Next(iter *Iterator[T]) *Iterator[T] {
	return iter.next
}

func (this *Stack[T]) Begin() *Iterator[T] {
	return this.head
}

func (this *Stack[T]) End(iter *Iterator[T]) bool {
	return iter == nil
}

func (this *Stack[T]) Len() int {
	return this.length
}

func (this *Stack[T]) Clear() {
	this.head = nil
	this.length = 0
}

func (this *Stack[T]) Push(values ...T) {
	for _, v := range values {
		var ele = &Iterator[T]{Data: v}
		if this.length > 0 {
			ele.next = this.head
			this.head = ele
		} else {
			this.head = ele
		}
		this.length++
	}
}

func (this *Stack[T]) Front() *Iterator[T] {
	return this.head
}

func (this *Stack[T]) Pop() *Iterator[T] {
	if this.length == 0 {
		return nil
	}
	var result = this.head
	this.head = this.head.next
	this.length--
	return result
}
