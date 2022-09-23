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

func (this *Rapid[T]) Collect(ptr Pointer) {
	var node = &this.Buckets[ptr]
	node.Ptr = 0
	node.NextPtr = 0
	node.PrevPtr = 0
	this.Recyclable.Push(ptr)
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

func (this Rapid[T]) Begin(entrypoint *EntryPoint) *Iterator[T] {
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
	if entrypoint.Head == 0 {
		var ptr = this.NextID()
		entrypoint.Head = ptr
		entrypoint.Tail = ptr
	}

	var head = &this.Buckets[entrypoint.Head]
	if head.Ptr == 0 {
		this.Length++
		this.Buckets[entrypoint.Head] = Iterator[T]{
			Ptr:     entrypoint.Head,
			PrevPtr: 0,
			NextPtr: 0,
			Data:    *data,
		}
		return false
	}

	for i := this.Begin(entrypoint); !this.End(i); i = this.Next(i) {
		if this.Equal(&i.Data, data) {
			i.Data = *data
			return true
		}
	}

	var cursor = this.NextID()
	var tail = &this.Buckets[entrypoint.Tail]
	tail.NextPtr = cursor
	entrypoint.Tail = cursor
	this.Buckets[cursor] = Iterator[T]{
		Ptr:     cursor,
		PrevPtr: tail.Ptr,
		NextPtr: 0,
		Data:    *data,
	}
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

	// delete last node
	if target.NextPtr == 0 && target.PrevPtr == 0 {
		entrypoint.Head = 0
		entrypoint.Tail = 0
		this.Collect(target.Ptr)
		return true
	}

	// delete head
	if target.PrevPtr == 0 {
		var next = &this.Buckets[target.NextPtr]
		entrypoint.Head = next.Ptr
		next.PrevPtr = 0
		this.Collect(target.Ptr)
		return true
	}

	// delete tail
	if target.NextPtr == 0 {
		var prev = &this.Buckets[target.PrevPtr]
		entrypoint.Tail = prev.Ptr
		prev.NextPtr = 0
		this.Collect(target.Ptr)
		return true
	}

	var prev = &this.Buckets[target.PrevPtr]
	var next = &this.Buckets[target.NextPtr]
	next.PrevPtr = prev.Ptr
	prev.NextPtr = next.Ptr
	this.Collect(target.Ptr)
	return true
}

func (this *Rapid[T]) Find(entrypoint *EntryPoint, data *T) (result *Iterator[T], exist bool) {
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
