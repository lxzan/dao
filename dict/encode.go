package dict

import "github.com/lxzan/dao/algorithm"

var defaultIndexes = []uint8{32, 32, 16, 16, 16, 16, 16, 8, 8, 8, 8, 8, 8, 4, 4, 4, 4}

type iterator struct {
	Node       *element
	Key        string
	N          int
	Cursor     int
	Initialize bool
}

func (c *iterator) hit() bool {
	return c.Cursor == c.N-1
}

func (c *Dict[T]) getIndex(iter *iterator) int {
	return int(iter.Key[iter.Cursor]) & int(c.indexes[iter.Cursor]-1)
}

func (c *Dict[T]) begin(key string, initialize bool) *iterator {
	var iter = &iterator{Node: c.root, Key: key, N: algorithm.Min(len(key), len(c.indexes)-1), Initialize: initialize}
	var idx = c.getIndex(iter)
	if iter.Node.Children[idx] == nil {
		if !iter.Initialize {
			return nil
		}
		iter.Node.Children[idx] = &element{Children: make([]*element, c.indexes[1])}
	}
	iter.Node = iter.Node.Children[idx]
	return iter
}

func (c *Dict[T]) next(iter *iterator) *iterator {
	iter.Cursor++
	if iter.Cursor >= iter.N {
		return nil
	}

	var idx = c.getIndex(iter)
	if iter.Node.Children[idx] == nil {
		if !iter.Initialize {
			return nil
		}
		iter.Node.Children[idx] = &element{Children: make([]*element, c.indexes[iter.Cursor+1])}
	}
	iter.Node = iter.Node.Children[idx]

	return iter
}
