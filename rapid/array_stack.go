package rapid

type array_stack []Pointer

func (this array_stack) Len() int {
	return len(this)
}

func (this *array_stack) Push(v Pointer) {
	*this = append(*this, v)
}

func (this *array_stack) Pop() Pointer {
	var n = this.Len()
	if n >= 1 {
		var result = (*this)[n-1]
		*this = (*this)[:n-1]
		return result
	}
	return 0
}
