package dict

const max_length = 32 // max_length=2^n, 2/4/16

var sizes = [max_length]uint8{16, 16, 16, 16, 8, 8, 8, 8, 4, 4, 4, 4, 4, 4, 4, 4, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}

type iterator struct {
	Node   *Element
	Bytes  []byte
	Cursor int
	End    int
}

func (c *Dict[T]) begin(key string, initialize bool) *iterator {
	var iter = &iterator{Node: c.root}
	var n = len(key)
	if n > c.indexLength {
		n = c.indexLength
	}
	iter.Bytes = make([]byte, n, n)
	for i := 0; i < n; i++ {
		iter.Bytes[i] = key[i] & (sizes[i] - 1)
	}
	iter.End = len(iter.Bytes)
	if key == "" {
		return iter
	}

	var idx = iter.Bytes[iter.Cursor]
	iter.Cursor++
	if initialize && iter.Node.Children[idx] == nil && iter.Cursor < max_length {
		iter.Node.Children[idx] = &Element{
			Children: make([]*Element, sizes[iter.Cursor], sizes[iter.Cursor]),
		}
	}
	iter.Node = iter.Node.Children[idx]
	return iter
}

func (c *Dict[T]) next(iter *iterator, initialize bool) *iterator {
	if iter.Cursor >= iter.End {
		return nil
	}

	var idx = iter.Bytes[iter.Cursor]
	iter.Cursor++
	if initialize && iter.Node.Children[idx] == nil && iter.Cursor < max_length {
		iter.Node.Children[idx] = &Element{
			Children: make([]*Element, sizes[iter.Cursor], sizes[iter.Cursor]),
		}
	}
	iter.Node = iter.Node.Children[idx]
	return iter
}

func (c *Dict[T]) end(iter *iterator) bool {
	return iter == nil || iter.Cursor > iter.End
}
