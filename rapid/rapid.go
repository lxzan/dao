package rapid

type (
	Pointer uint32

	EntryPoint struct {
		Head Pointer
		Tail Pointer
	}

	Iterator[T any] struct {
		Ptr     Pointer
		PrevPtr Pointer
		NextPtr Pointer
		Data    T
	}
)

func (this *Iterator[T]) Reset() {
	this.Ptr = 0
	this.NextPtr = 0
}

type Rapid[T any] struct {
	Length     int
	Serial     uint32
	Recyclable array_stack // do not recycle head
	Buckets    []Iterator[T]
	Equal      func(a, b *T) bool
}

func New[T any](size uint32, eq func(a, b *T) bool) *Rapid[T] {
	return &Rapid[T]{
		Serial:     1,
		Recyclable: []Pointer{},
		Buckets:    make([]Iterator[T], size+1),
		Length:     0,
		Equal:      eq,
	}
}

func (this Rapid[T]) Begin(entrypoint EntryPoint) *Iterator[T] {
	return &this.Buckets[entrypoint.Head]
}

func (this Rapid[T]) Next(iter *Iterator[T]) *Iterator[T] {
	return &this.Buckets[iter.NextPtr]
}

func (this Rapid[T]) End(iter *Iterator[T]) bool {
	return iter.Ptr == 0
}

// NextID apply a pointer
func (this *Rapid[T]) NextID() Pointer {
	if this.Recyclable.Len() > 0 {
		return this.Recyclable.Pop()
	}

	var result = this.Serial
	if result >= uint32(len(this.Buckets)) {
		var ele Iterator[T]
		this.Buckets = append(this.Buckets, ele)
	}
	this.Serial++
	return Pointer(result)
}

// Push append an element with unique check
func (this *Rapid[T]) Push(entrypoint *EntryPoint, data *T) (replaced bool) {
	var head = &this.Buckets[entrypoint.Head]
	if head.Ptr == 0 {
		head.Ptr = entrypoint.Head
		head.Data = *data
		this.Length++
		return false
	}

	for i := head; !this.End(i); i = this.Next(i) {
		if this.Equal(&i.Data, data) {
			i.Data = *data
			return true
		}
	}

	var cursor = this.NextID()
	var tail = &this.Buckets[entrypoint.Tail]
	tail.NextPtr = cursor
	entrypoint.Tail = cursor
	var target = &this.Buckets[cursor]
	target.Ptr = cursor
	target.Data = *data
	target.PrevPtr = tail.Ptr
	this.Length++
	return false
}

// Append append an element without unique check
func (this *Rapid[T]) Append(entrypoint *EntryPoint, data *T) {
	var head = &this.Buckets[entrypoint.Head]
	if head.Ptr == 0 {
		head.Ptr = entrypoint.Head
		head.Data = *data
		this.Length++
		return
	}

	var cursor = this.NextID()
	var tail = &this.Buckets[entrypoint.Tail]
	tail.NextPtr = cursor
	entrypoint.Tail = cursor
	var target = &this.Buckets[cursor]
	target.Ptr = cursor
	target.Data = *data
	target.PrevPtr = tail.Ptr
	this.Length++
}

// Delete do not delete in loop if no break
func (this *Rapid[T]) Delete(entrypoint *EntryPoint, target *Iterator[T]) (deleted bool) {
	var head = this.Buckets[entrypoint.Head]
	if head.Ptr == 0 || target == nil || target.Ptr == 0 {
		return false
	}

	this.Length--
	if target.NextPtr == 0 {
		if target.PrevPtr != 0 {
			var prev = &this.Buckets[target.PrevPtr]
			prev.NextPtr = 0
			entrypoint.Tail = prev.Ptr
			this.Recyclable.Push(target.Ptr)
		}
		target.Reset()
		return true
	}

	var next = &this.Buckets[target.NextPtr]
	this.Recyclable.Push(next.Ptr)
	next.Ptr = target.Ptr
	next.PrevPtr = target.PrevPtr
	*target = *next
	next.Reset()
	if target.NextPtr == 0 {
		entrypoint.Tail = target.Ptr
	}
	return true
}

func (this *Rapid[T]) Find(entrypoint EntryPoint, data *T) (result *Iterator[T], exist bool) {
	if entrypoint.Head == 0 {
		return nil, false
	}
	for i := this.Begin(entrypoint); !this.End(i); i = this.Next(i) {
		if this.Equal(&i.Data, data) {
			return i, true
		}
	}
	return nil, false
}
