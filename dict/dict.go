package dict

import (
	"github.com/lxzan/dao/rapid"
	"math"
)

type Pair[T any] struct {
	Key string
	Val T
}

type Element struct {
	EntryPoint rapid.EntryPoint
	Children   []*Element
}

type Dict[T any] struct {
	index_length int // 8 Byte
	root         *Element
	storage      *rapid.Rapid[Pair[T]]
}

type iterator struct {
	Node   *Element
	Bytes  []byte
	Cursor int
	End    int
}

func New[T any]() *Dict[T] {
	return &Dict[T]{
		index_length: 8,
		root:         &Element{Children: make([]*Element, sizes[0], sizes[0])},
		storage: rapid.New[Pair[T]](8, func(a, b *Pair[T]) bool {
			return a.Key == b.Key
		}),
	}
}

func (c *Dict[T]) Len() int {
	return c.storage.Length
}

// length<=16
func (c *Dict[T]) SetIndexLength(length int) {
	if length <= 0 {
		length = 8
	}
	c.index_length = length
}

// insert with unique check
func (c *Dict[T]) Insert(key string, val T) {
	var data = Pair[T]{
		Key: key,
		Val: val,
	}
	for i := c.begin(c.new_iterator(key), true); true; i = c.next(i, true) {
		if i.Cursor == i.End {
			var entrypoint = &i.Node.EntryPoint
			if entrypoint.Head == 0 {
				var ptr = c.storage.NextID()
				entrypoint.Head = ptr
				entrypoint.Tail = ptr
			}
			c.storage.Push(entrypoint, &data)
			break
		}
	}
}

type match_params[T any] struct {
	node    *Element
	results []Pair[T]
	limit   int
	prefix  string
	length  int
}

// limit: -1 as unlimited
func (c *Dict[T]) Match(prefix string, limit ...int) []Pair[T] {
	if len(limit) == 0 {
		limit = []int{math.MaxInt}
	}
	for i := c.begin(c.new_iterator(prefix), false); !c.end(i); i = c.next(i, false) {
		if i.Node == nil {
			return nil
		}
		if i.Cursor == i.End {
			var params = match_params[T]{
				node:    i.Node,
				results: make([]Pair[T], 0),
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
	if node == nil || len(params.results) >= params.limit {
		return
	}
	for i := c.storage.Begin(node.EntryPoint); !c.storage.End(i); i = c.storage.Next(i) {
		if len(i.Data.Key) >= params.length && i.Data.Key[:params.length] == params.prefix {
			params.results = append(params.results, i.Data)
		}
	}
	for _, item := range node.Children {
		c.doMatch(item, params)
	}
}

func (c *Dict[T]) Delete(key string) bool {
	for i := c.begin(c.new_iterator(key), false); !c.end(i); i = c.next(i, false) {
		if i.Node == nil {
			return false
		}
		if i.Cursor == i.End {
			for j := c.storage.Begin(i.Node.EntryPoint); !c.storage.End(j); j = c.storage.Next(j) {
				if j.Data.Key == key {
					return c.storage.Delete(&i.Node.EntryPoint, j)
				}
			}
		}
	}
	return false
}

func (c *Dict[T]) ForEach(fn func(item *Pair[T]) (continued bool)) {
	var n = len(c.storage.Buckets)
	for i := 0; i < n; i++ {
		if c.storage.Buckets[i].Ptr != 0 {
			if !fn(&c.storage.Buckets[i].Data) {
				break
			}
		}
	}
}
