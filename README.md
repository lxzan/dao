# The Great Way is Simple
Simple and high-performance data structures and algorithms library



### Benchmark
- 1,000 elements
```
goos: darwin
goarch: arm64
pkg: github.com/lxzan/dao/benchmark
BenchmarkMyMap_Set/string-8                59996             19892 ns/op           53248 B/op          2 allocs/op
BenchmarkMyMap_Set/int-8                   91070             13291 ns/op           45056 B/op          2 allocs/op
BenchmarkGoMap_Set/string-8                36583             32849 ns/op           57368 B/op          2 allocs/op
BenchmarkGoMap_Set/int-8                   54259             22276 ns/op           41097 B/op          6 allocs/op
BenchmarkMyMap_Get/string-8               113704             10354 ns/op               0 B/op          0 allocs/op
BenchmarkMyMap_Get/int-8                  218997              5541 ns/op               0 B/op          0 allocs/op
BenchmarkGoMap_Get/string-8               168564              7316 ns/op               0 B/op          0 allocs/op
BenchmarkGoMap_Get/int-8                  167908              6520 ns/op               0 B/op          0 allocs/op
BenchmarkMyMap_Delete/string-8             49182             24258 ns/op           12929 B/op         11 allocs/op
BenchmarkMyMap_Delete/int-8                90664             13053 ns/op           12928 B/op         11 allocs/op
BenchmarkGoMap_Delete/string-8             32680             36706 ns/op               1 B/op          0 allocs/op
BenchmarkGoMap_Delete/int-8                50990             23574 ns/op               0 B/op          0 allocs/op
PASS
ok      github.com/lxzan/dao/benchmark  31.351s
```

- 10,000 elements
```
goos: darwin
goarch: arm64
pkg: github.com/lxzan/dao/benchmark
BenchmarkMyMap_Set/string-8                 3825            309373 ns/op          729092 B/op          2 allocs/op
BenchmarkMyMap_Set/int-8                    8720            134059 ns/op          598016 B/op          2 allocs/op
BenchmarkGoMap_Set/string-8                 3256            364237 ns/op          458777 B/op          2 allocs/op
BenchmarkGoMap_Set/int-8                    4876            243187 ns/op          322224 B/op         11 allocs/op
BenchmarkMyMap_Get/string-8                 5068            232323 ns/op             143 B/op          0 allocs/op
BenchmarkMyMap_Get/int-8                   15595             76749 ns/op              38 B/op          0 allocs/op
BenchmarkGoMap_Get/string-8                 3990            300194 ns/op             114 B/op          0 allocs/op
BenchmarkGoMap_Get/int-8                    6386            180545 ns/op              50 B/op          0 allocs/op
BenchmarkMyMap_Delete/string-8              3648            319692 ns/op          141386 B/op         17 allocs/op
BenchmarkMyMap_Delete/int-8                 9091            127630 ns/op          141250 B/op         17 allocs/op
BenchmarkGoMap_Delete/string-8              3242            363678 ns/op             141 B/op          0 allocs/op
BenchmarkGoMap_Delete/int-8                 4821            242847 ns/op              66 B/op          0 allocs/op
PASS
ok      github.com/lxzan/dao/benchmark  21.932s
```

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
	m.ForEach(func(iter *hashmap.Iterator[string, int]) bool {
		println(iter.Key, iter.Value)
		return iter.Continue()
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
	"github.com/lxzan/dao/rbtree"
	"strconv"
)

func main() {
	var tree = rbtree.New[int, string]()

	var rows = make([]*rbtree.Iterator[int, string], 0)
	for i := 0; i < 10; i++ {
		rows = append(rows, &rbtree.Iterator[int, string]{Key: i, Val: strconv.Itoa(i)})
	}
	for _, item := range rows {
		tree.Insert(item)
	}

	results := tree.Query(&rbtree.QueryBuilder[int]{
		LeftFilter: func(d int) bool { return d >= 5 },
		Limit:      10,
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
	var d = dict.New[int]()
	d.Set("teemo", 1)
	d.Set("tesla", 2)
	d.Set("task", 3)
	d.Set("hasaki", 4)
	d.Set("test", 5)
	d.Set("aha", 6)
	var results = d.Match("te").Elem()
	for _, item := range results {
		println(item.Key)
	}
}

```
