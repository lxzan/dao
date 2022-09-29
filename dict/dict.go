package dict

import (
	"github.com/lxzan/dao/internal/mlist"
	"github.com/lxzan/dao/vector"
	"math"
)

type Iterator[T any] struct {
	Key    string
	Value  T
	broken bool
}

func (c *Iterator[T]) Break() {
	c.broken = true
}

type Element struct {
	EntryPoint mlist.Pointer
	Children   []*Element
}

type Dict[T any] struct {
	indexLength int      // 8byte index
	root        *Element // root node
	storage     *mlist.MList[string, T]
}

// New 4<=indexLength<=32
func New[T any](indexLength ...int) *Dict[T] {
	if len(indexLength) == 0 {
		indexLength = []int{8}
	}
	return &Dict[T]{
		indexLength: indexLength[0],
		root:        &Element{Children: make([]*Element, sizes[0], sizes[0])},
		storage:     mlist.NewMList[string, T](8),
	}
}

func (c *Dict[T]) Len() int {
	return c.storage.Length
}

// Set insert a element into the hashmap
// if key exists, value will be replaced
func (c *Dict[T]) Set(key string, val T) {
	for i := c.begin(key, true); !c.end(i); i = c.next(i, true) {
		if i.Cursor == i.End {
			c.storage.Push(&i.Node.EntryPoint, key, val)
			break
		}
	}
}

func (c *Dict[T]) Get(key string) (value T, exist bool) {
	var entrypoint mlist.Pointer
	for i := c.begin(key, false); !c.end(i); i = c.next(i, false) {
		if i.Cursor == i.End {
			if i.Node == nil || i.Node.EntryPoint == 0 {
				return value, false
			}
			entrypoint = i.Node.EntryPoint
		}
	}

	for i := c.storage.Begin(entrypoint); !c.storage.End(i); i = c.storage.Next(i) {
		if i.Key == key {
			return i.Value, true
		}
	}
	return value, false
}

type match_params[T any] struct {
	node    *Element
	results *vector.Vector[Iterator[T]]
	limit   int
	prefix  string
	length  int
}

// Match limit: -1 as unlimited
func (c *Dict[T]) Match(prefix string, limit ...int) *vector.Vector[Iterator[T]] {
	if len(limit) == 0 {
		limit = []int{math.MaxInt}
	}

	for i := c.begin(prefix, false); !c.end(i); i = c.next(i, false) {
		if i.Node == nil {
			return nil
		}
		if i.Cursor == i.End {
			var params = match_params[T]{
				node:    i.Node,
				results: vector.New[Iterator[T]](),
				limit:   limit[0],
				prefix:  prefix,
				length:  len(prefix),
			}
			c.doMatch(i.Node, &params)
			return params.results
		}
	}
	return nil
}

func (c *Dict[T]) doMatch(node *Element, params *match_params[T]) {
	if node == nil || params.results.Len() >= params.limit {
		return
	}
	for i := c.storage.Begin(node.EntryPoint); !c.storage.End(i); i = c.storage.Next(i) {
		if len(i.Key) >= params.length && i.Key[:params.length] == params.prefix {
			params.results.Push(Iterator[T]{Key: i.Key, Value: i.Value})
		}
	}
	for _, item := range node.Children {
		c.doMatch(item, params)
	}
}

func (c *Dict[T]) Delete(key string) bool {
	for i := c.begin(key, false); !c.end(i); i = c.next(i, false) {
		if i.Node == nil {
			return false
		}
		if i.Cursor == i.End {
			for j := c.storage.Begin(i.Node.EntryPoint); !c.storage.End(j); j = c.storage.Next(j) {
				if j.Key == key {
					return c.storage.Delete(&i.Node.EntryPoint, j)
				}
			}
		}
	}
	return false
}

func (c *Dict[T]) ForEach(fn func(iter *Iterator[T])) {
	var iter = &Iterator[T]{}
	for i := 1; i < int(c.storage.Serial); i++ {
		var item = &c.storage.Buckets[i]
		if item.Ptr > 0 {
			iter.Key = item.Key
			iter.Value = item.Value
			fn(iter)
			if iter.broken {
				return
			}
		}
	}
}
