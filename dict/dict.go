package dict

import (
	"github.com/lxzan/dao/internal/mlist"
	"github.com/lxzan/dao/vector"
	"math"
)

type Pair[T any] struct {
	Key string
	Val T
}

type Element struct {
	EntryPoint mlist.Pointer
	Children   []*Element
}

type Dict[T any] struct {
	indexLength int // 8 Byte
	root        *Element
	storage     *mlist.MList[string, T]
}

// New 4<=indexLength<=32
func New[T any](indexLength int) *Dict[T] {
	return &Dict[T]{
		indexLength: indexLength,
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

func (c *Dict[T]) Find(key string) (value T, exist bool) {
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
	results *vector.Vector[Pair[T]]
	limit   int
	prefix  string
	length  int
}

// Match limit: -1 as unlimited
func (c *Dict[T]) Match(prefix string, limit ...int) *vector.Vector[Pair[T]] {
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
				results: vector.New[Pair[T]](),
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
			params.results.Push(Pair[T]{Key: i.Key, Val: i.Value})
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

func (c *Dict[T]) ForEach(fn func(key string, val T) bool) {
	var n = len(c.storage.Buckets)
	for i := 0; i < n; i++ {
		if c.storage.Buckets[i].Ptr != 0 {
			var item = &c.storage.Buckets[i]
			if !fn(item.Key, item.Value) {
				break
			}
		}
	}
}
