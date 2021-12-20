# The Great Way is Simple
> Simple and high-performance data structures and algorithms library



### Benchmark

| Container        | Operate | Elements | Cost(ns/op) |
| ---------------- | ------- | -------- | ----------- |
| DoubleLinkedList | RPush   | 1,000    | 29,704      |
|                  |         |          |             |
|                  |         |          |             |



### HashMap

```go
package main

import (
	"github.com/lxzan/dao/hashmap"
)

func main() {
	var m = hashmap.New[string, int]()
	m.Set("hello", 1)
	m.Set("world", 2)
	m.Set("!", 3)
	m.ForEach(func(item *hashmap.Pair[string, int]) (continued bool) {
		println(item.Key)
		return true
	})
}

```

### DoubleLinkedList

```go
package main

import (
	"github.com/lxzan/dao/double_linkedlist"
)

func main() {
	var list = double_linkedlist.New[int]()
	list.RPush(1, 3)
	list.LPush(5, 7)
	for i := list.Begin(); !list.End(i); i = list.Next(i) {
		println(i.Data)
	}
}

```

### Heap

```go
package main

import (
	"github.com/lxzan/dao/heap"
)

func main() {
	var h = heap.New(10, heap.MaxHeap[int])
	h.Push(1, 3, 5, 7, 9, 2, 4, 6, 8, 0)
	for h.Len() > 0 {
		println(h.Pop())
	}
}

```

### SegmentTree

```go
package main

import (
	"github.com/lxzan/dao/segment_tree"
)

func main() {
	var arr = []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 0}
	var tree *segment_tree.SegmentTree[int, segment_tree.Schema[int]] = segment_tree.New(arr, segment_tree.Init[int], segment_tree.Merge[int])

	var query1 = tree.Query(0, 9)
	println(query1.Sum, query1.MaxValue, query1.MinValue)

	tree.Update(1, 18)
	tree.Update(3, -1)
	var query2 = tree.Query(0, 9)
	println(query2.Sum, query2.MaxValue, query2.MinValue)
}
```

### RedBlackTree

```go
package main

import (
	"github.com/lxzan/dao"
	"github.com/lxzan/dao/rbtree"
	"strconv"
)

type entry struct {
	Key int
	Val string
}

func main() {
	var tree = rbtree.New(func(a, b *entry) dao.Ordering {
		if a.Key > b.Key {
			return dao.Greater
		} else if a.Key == b.Key {
			return dao.Equal
		} else {
			return dao.Less
		}
	})

	var rows = make([]*entry, 0)
	for i := 0; i < 10; i++ {
		rows = append(rows, &entry{Key: i, Val: strconv.Itoa(i)})
	}
	for _, item := range rows {
		tree.Insert(item)
	}

	results := tree.Query(&rbtree.QueryBuilder[entry]{
		LeftFilter: func(d *entry) bool { return d.Key > 5 },
		Limit:      3,
		Order:      rbtree.ASC,
	})
	for _, item := range results {
		println(item.Key)
	}
}
```

### Trie

```go
package main

import (
	"github.com/lxzan/dao/dict"
)

func main() {
	var tree = dict.New[int]()
	tree.Insert("teemo", 1)
	tree.Insert("tesla", 2)
	tree.Insert("task", 3)
	tree.Insert("hasaki", 4)
	tree.Insert("oh", 5)
	tree.Insert("aha", 6)
	var results = tree.Match("te")
	for _, item := range results {
		println(item.Key)
	}
}

```