package dict

const bit = 16 // bit=2^n, 2/4/16

var sizes = []uint8{16, 16, 16, 16, 8, 8, 8, 8, 4, 4, 4, 4, 4, 4, 4, 4}

func (c *Dict[T]) new_iterator(key string) *iterator {
	var iter = &iterator{Node: c.root}
	var n = len(key)
	if n > c.index_length {
		n = c.index_length
	}

	iter.Bytes = make([]byte, n, n)
	for i := 0; i < n; i++ {
		iter.Bytes[i] = key[i] & (sizes[i] - 1)
	}
	iter.End = len(iter.Bytes)
	return iter
}

func (c *Dict[T]) begin(iter *iterator, initialize bool) *iterator {
	var idx = iter.Bytes[iter.Cursor]
	iter.Cursor++
	if initialize && iter.Node.Children[idx] == nil && iter.Cursor < 16 {
		iter.Node.Children[idx] = &Element{
			Children: make([]*Element, sizes[iter.Cursor], sizes[iter.Cursor]),
		}
	}
	iter.Node = iter.Node.Children[idx]
	return iter
}

func (c *Dict[T]) next(iter *iterator, initialize bool) *iterator {
	var idx = iter.Bytes[iter.Cursor]
	iter.Cursor++
	if initialize && iter.Node.Children[idx] == nil && iter.Cursor < 16 {
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
