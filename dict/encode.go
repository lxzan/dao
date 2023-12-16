package dict

var defaultIndexes = []uint8{16, 16, 16, 8, 8, 8, 8, 4, 4, 4, 4, 2, 2, 2, 2, 2, 2}

type iterator struct {
	Node       *element
	Key        string
	N          int
	Cursor     int
	Initialize bool
	Indexes    []uint8
}

func (c *iterator) hit() bool {
	return c.Cursor == c.N-1
}

func (c *iterator) getIndex() int {
	return int(c.Key[c.Cursor]) & int(c.Indexes[c.Cursor]-1)
}

func (c *Dict[T]) begin(key string, initialize bool) *iterator {
	var iter = &iterator{Node: c.root, Key: key, N: min(len(key), len(c.indexes)-1), Initialize: initialize, Indexes: c.indexes}
	var idx = iter.getIndex()
	if iter.Node.Children[idx] == nil {
		if !iter.Initialize {
			return nil
		}
		iter.Node.Children[idx] = &element{Children: make([]*element, iter.Indexes[1])}
	}
	iter.Node = iter.Node.Children[idx]
	return iter
}

func (c *iterator) next() *iterator {
	c.Cursor++
	if c.Cursor >= c.N {
		return nil
	}

	var idx = c.getIndex()
	if c.Node.Children[idx] == nil {
		if !c.Initialize {
			return nil
		}
		c.Node.Children[idx] = &element{Children: make([]*element, c.Indexes[c.Cursor+1])}
	}
	c.Node = c.Node.Children[idx]

	return c
}
