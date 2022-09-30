# The Great Way is Simple
Simple and high-performance data structures and algorithms library

[![Build Status](https://github.com/lxzan/dao/workflows/Go%20Test/badge.svg?branch=main)](https://github.com/lxzan/dao/actions?query=branch%3Amain)

### Benchmark
- 1,000 elements
```
goos: windows
goarch: amd64
pkg: github.com/lxzan/dao/benchmark
cpu: AMD Ryzen 5 PRO 4650G with Radeon Graphics
BenchmarkMyMap_Set/string-12               38535             32238 ns/op           53248 B/op          2 allocs/op
BenchmarkMyMap_Set/int-12                  55552             21591 ns/op           45056 B/op          2 allocs/op
BenchmarkGoMap_Set/string-12               32197             36445 ns/op           57368 B/op          2 allocs/op
BenchmarkGoMap_Set/int-12                  41067             28898 ns/op           41097 B/op          6 allocs/op
BenchmarkMyMap_Get/string-12               82111             14339 ns/op               0 B/op          0 allocs/op
BenchmarkMyMap_Get/int-12                 151774              7747 ns/op               0 B/op          0 allocs/op
BenchmarkGoMap_Get/string-12               95738             12707 ns/op               0 B/op          0 allocs/op
BenchmarkGoMap_Get/int-12                 146209              8044 ns/op               0 B/op          0 allocs/op
BenchmarkMyMap_Delete/string-12            51271             23247 ns/op           12929 B/op         11 allocs/op
BenchmarkMyMap_Delete/int-12               76872             15740 ns/op           12928 B/op         11 allocs/op
BenchmarkGoMap_Delete/string-12            38737             31319 ns/op               1 B/op          0 allocs/op
BenchmarkGoMap_Delete/int-12               51272             23187 ns/op               0 B/op          0 allocs/op
```

- 10,000 elements
```
goos: windows
goarch: amd64
pkg: github.com/lxzan/dao/benchmark
cpu: AMD Ryzen 5 PRO 4650G with Radeon Graphics
BenchmarkMyMap_Set/string-12                3187            369867 ns/op          729088 B/op          2 allocs/op
BenchmarkMyMap_Set/int-12                   5449            233284 ns/op          598016 B/op          2 allocs/op
BenchmarkGoMap_Set/string-12                3202            390564 ns/op          458776 B/op          2 allocs/op
BenchmarkGoMap_Set/int-12                   4135            296884 ns/op          322238 B/op         11 allocs/op
BenchmarkMyMap_Get/string-12                8564            149001 ns/op              85 B/op          0 allocs/op
BenchmarkMyMap_Get/int-12                  13832             85587 ns/op              43 B/op          0 allocs/op
BenchmarkGoMap_Get/string-12                5448            214066 ns/op              84 B/op          0 allocs/op
BenchmarkGoMap_Get/int-12                   6656            173191 ns/op              48 B/op          0 allocs/op
BenchmarkMyMap_Delete/string-12             4890            308094 ns/op          141334 B/op         17 allocs/op
BenchmarkMyMap_Delete/int-12                6648            160257 ns/op          141274 B/op         17 allocs/op
BenchmarkGoMap_Delete/string-12             3375            324144 ns/op             135 B/op          0 allocs/op
BenchmarkGoMap_Delete/int-12                4882            245148 ns/op              66 B/op          0 allocs/op
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
