# The Great Way is Simple
Simple and high-performance data structures and algorithms library



### Benchmark

```
goos: windows
goarch: amd64
pkg: github.com/lxzan/dao/benchmark
cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
```

| Container        | Operate            | Elements | ns/op | allocs/op |
|------------------| ------------------ | -------- |-------| --------- |
| DoubleLinkedList | RPush              | 1,000    | 29.7  | 1000      |
| Dict             | Insert             | 1,000    | 651.6 | 11006     |
| Dict             | Delete             | 1,000    | 59.3  | 1000      |
| Dict             | Match (limit 10)   | 1,000    | 41.3  | 5697      |
| HashMap          | Set                | 1,000    | 36.7  | 5         |
| Go Map           | Set                | 1,000    | 38.1  | 2         |
| HashMap          | Get                | 1,000    | 17.0  | 0         |
| Go Map           | Get                | 1,000    | 19.3  | 0         |
| RBTree           | Insert             | 1,000    | 341.7 | 2001      |
| RBTree           | Find               | 1,000    | 138.4 | 2000      |
| RBTree           | Delete             | 1,000    | 377.4 | 1000      |
| RBTree           | Between (limit 10) | 1,000    | 777.7 | 4001      |
| SegmentTree      | Query              | 1,000    | 190.8 | 0         |
| SegmentTree      | Update             | 1,000    | 144.3 | 0         |
| Heap             | Push               | 1,000    | 63.8  | 1         |



| Container        | Operate            | Elements     | ns/op  | allocs/op  |
| ---------------- | ------------------ | --------     |--------|------------|
| DoubleLinkedList | RPush              | 1,000,000    | 42.5   | 1,000,000  |
| Dict             | Insert             | 1,000,000    | 1196.2 | 6,370,600  |
| Dict             | Delete             | 1,000,000    | 1673.7 | 7,370,623  |
| Dict             | Match (limit 10)   | 1,000,000    | 2247.6 | 14,849,896 |
| HashMap          | Set                | 1,000,000    | 127.4  | 5          |
| Go Map           | Set                | 1,000,000    | 112.3  | 2          |
| HashMap          | Get                | 1,000,000    | 116.0  | 0          |
| Go Map           | Get                | 1,000,000    | 99.2   | 0          |
| RBTree           | Insert             | 1,000,000    | 449.2  | 2,000,001  |
| RBTree           | Find               | 1,000,000    | 251.9  | 1,400,000  |
| RBTree           | Delete             | 1,000,000    | 574.6  | 4,000,001  |
| RBTree           | Between (limit 10) | 1,000,000    | 3330.4 | 5,000,001  |
| SegmentTree      | Query              | 1,000,000    | 1834.9 | 2,000,038  |
| SegmentTree      | Update             | 1,000,000    | 1051.6 | 2,000,038  |
| Heap             | Push               | 1,000,000    | 37.8   | 33         |

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
	m.ForEach(func(key string, val int) (continued bool) {
		println(key)
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
