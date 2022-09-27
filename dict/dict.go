package dict

import (
	"github.com/lxzan/dao/rapid"
	"github.com/lxzan/dao/vector"
	"math"
)

type Pair[T any] struct {
	Key string
	Val T
}

type Element struct {
	EntryPoint rapid.Pointer
	Children   []*Element
}

type Dict[T any] struct {
	index_length int // 8 Byte
	root         *Element
	storage      *rapid.Rapid[string, T]
}

func New[T any]() *Dict[T] {
	return &Dict[T]{
		index_length: 8,
		root:         &Element{Children: make([]*Element, sizes[0], sizes[0])},
		storage:      rapid.New[string, T](8),
	}
}

func (this *Dict[T]) Len() int {
	return this.storage.Length
}

// length<=32
func (this *Dict[T]) SetIndexLength(length int) {
	if length <= 0 {
		length = 8
	}
	this.index_length = length
}

// insert with unique check
func (this *Dict[T]) Insert(key string, val T) {
	for i := this.begin(key, true); !this.end(i); i = this.next(i, true) {
		if i.Cursor == i.End {
			var entrypoint = &i.Node.EntryPoint
			this.storage.Push(entrypoint, key, val)
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

func (this *Dict[T]) Find(key string) (value T, exist bool) {
	var entrypoint rapid.Pointer
	for i := this.begin(key, false); !this.end(i); i = this.next(i, false) {
		if i.Node == nil {
			return value, false
		}
		if i.Cursor == i.End {
			if i.Node == nil || i.Node.EntryPoint == 0 {
				return value, false
			}
			entrypoint = i.Node.EntryPoint
		}
	}

	for i := this.storage.Begin(entrypoint); !this.storage.End(i); i = this.storage.Next(i) {
		if i.Key == key {
			return i.Value, true
		}
	}
	return value, false
}

// limit: -1 as unlimited
func (this *Dict[T]) Match(prefix string, limit ...int) vector.Vector[Pair[T]] {
	if len(limit) == 0 {
		limit = []int{math.MaxInt}
	}

	for i := this.begin(prefix, false); !this.end(i); i = this.next(i, false) {
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
			this.doMatch(i.Node, &params)
			return params.results
		}
	}
	return nil
}

func (this *Dict[T]) doMatch(node *Element, params *match_params[T]) {
	if node == nil || len(params.results) >= params.limit {
		return
	}
	for i := this.storage.Begin(node.EntryPoint); !this.storage.End(i); i = this.storage.Next(i) {
		if len(i.Key) >= params.length && i.Key[:params.length] == params.prefix {
			params.results = append(params.results, Pair[T]{Key: i.Key, Val: i.Value})
		}
	}
	for _, item := range node.Children {
		this.doMatch(item, params)
	}
}

func (this *Dict[T]) Delete(key string) bool {
	for i := this.begin(key, false); !this.end(i); i = this.next(i, false) {
		if i.Node == nil {
			return false
		}
		if i.Cursor == i.End {
			for j := this.storage.Begin(i.Node.EntryPoint); !this.storage.End(j); j = this.storage.Next(j) {
				if j.Key == key {
					return this.storage.Delete(&i.Node.EntryPoint, j)
				}
			}
		}
	}
	return false
}

func (this *Dict[T]) ForEach(fn func(key string, val T) (continued bool)) {
	var n = len(this.storage.Buckets)
	for i := 0; i < n; i++ {
		if this.storage.Buckets[i].Ptr != 0 {
			var item = &this.storage.Buckets[i]
			if !fn(item.Key, item.Value) {
				break
			}
		}
	}
}
