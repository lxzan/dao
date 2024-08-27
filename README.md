[中文](README_CN.md)

<div align="center">
    <h1>DAO</h1>
    <img src="assets/logo.png" alt="logo" width="300px">
    <h5>道生一, 一生二, 二生三, 三生万物; 万物负阴而抱阳, 冲气以为和</h5>
</div>

<div align="center">

[![Build Status](https://github.com/lxzan/dao/workflows/Go%20Test/badge.svg?branch=main)](https://github.com/lxzan/dao/actions?query=branch%3Amain) [![codecov](https://codecov.io/gh/lxzan/dao/graph/badge.svg?token=BQM1JHCDEE)](https://codecov.io/gh/lxzan/dao) [![go-version](https://img.shields.io/badge/go-%3E%3D1.18-30dff3?style=flat-square&logo=go)](https://github.com/lxzan/dao) [![license](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

</div>

### Description

`DAO` is a library of generic-based data structures and algorithms that complements the standard library in terms of
data containers and algorithms to simplify business development.

### Index

- [Description](#description)
- [Index](#index)
- [Vector](#vector)
  - [Unique](#unique)
  - [Sort](#sort)
  - [Filter](#filter)
- [Heap](#heap)
  - [N-Way Heap](#n-way-heap)
  - [HashHeap](#hashheap)
- [Stack](#stack)
- [Queue](#queue)
- [Deque](#deque)
- [LinkedList](#linkedlist)
- [RBTree](#rbtree)
- [Dict](#dict)
- [HashMap](#hashmap)
- [Segment Tree](#segment-tree)
- [Benchmark](#benchmark)


### Vector

#### Unique

```go
package main

import (
    "fmt"
    "github.com/lxzan/dao/vector"
)

func main() {
    var v = vector.NewFromInts(1, 3, 5, 3)
    v.Unique()
    fmt.Printf("%v", v.Elem())
}

```

#### Sort

```go
package main

import (
    "fmt"
    "github.com/lxzan/dao/vector"
)

func main() {
    var v = vector.NewFromInts(1, 3, 5, 2, 4, 6)
    v.Sort()
    fmt.Printf("%v", v.Elem())
}

```

#### Filter

```go
package main

import (
    "fmt"
    "github.com/lxzan/dao/vector"
)

func main() {
    var v = vector.NewFromInts(1, 3, 5, 2, 4, 6)
    v.Filter(func(i int, v vector.Int) bool {
        return v.GetID()%2 == 0
    })
    fmt.Printf("%v", v.Elem())
}

```

### Heap

**Heap**, also known as a priority queue, where the top element of the heap is always the largest or smallest. Commonly
used is the quadruple heap, `Push/Pop` is more balanced. Using `y=pow(2,x)` as the number of forks, to speed up
parent-child computation.

#### N-Way Heap

```go
package main

import (
    "github.com/lxzan/dao/heap"
    "github.com/lxzan/dao/types/cmp"
)

func main() {
    var h = heap.NewWithWays(heap.Binary, cmp.Less[int])
    h.Push(1)
    h.Push(3)
    h.Push(5)
    h.Push(2)
    h.Push(4)
    h.Push(6)
    for h.Len() > 0 {
        println(h.Pop())
    }
}

```

#### HashHeap

Heap structure with hash index.

```go
package main

import (
    "github.com/lxzan/dao/heap"
)

func main() {
    var h = heap.NewHashHeap[string, int](func(a, b int) bool { return a < b })
    h.Set("a", 1)
    h.Set("b", 2)
    h.Set("c", 3)
    h.Set("d", 4)
    h.Set("e", 5)
    h.Set("f", 6)

    h.Delete("c")
    h.Set("d", 0)
    h.Set("g", 3)
    for h.Len() > 0 {
        ele := h.Pop()
        println(ele.Key(), ele.Value())
    }
}

```

### Stack

**Stack** Last in first out (`LIFO`) data structure

```go
package main

import (
    "github.com/lxzan/dao/stack"
)

func main() {
    var s stack.Stack[int]
    s.Push(1)
    s.Push(3)
    s.Push(5)
    for s.Len() > 0 {
        println(s.Pop())
    }
}
```

### Queue

**Queue** First in first out (`FIFO`) data structure. The `dao/queue` is automatically reset when all elements are
ejected, reusing memory space.

```go
package main

import (
    "github.com/lxzan/dao/queue"
)

func main() {
    var s = queue.New[int](0)
    s.Push(1)
    s.Push(3)
    s.Push(5)
    for s.Len() > 0 {
        println(s.Pop())
    }
}

```

### Deque

**Deque** are similar to doubly linked list, where insertion and deletion operations can be
performed efficiently at both ends.

`dao/deque` is based on array subscripts to emulate pointers, the deleted space can still be reused later, and does not
depend on `sync.Pool`.

```go
package main

import (
    "fmt"
    "github.com/lxzan/dao/deque"
)

func main() {
    var list = deque.New[int](8)
    list.PushBack(1)
    list.PushBack(3)
    list.PushBack(5)
    list.PushBack(7)
    list.PushBack(9)
    for i := list.Front(); i != nil; i = list.Get(i.Next()) {
        fmt.Printf("%v ", i.Value())
    }

    println()
    for i := list.Back(); i != nil; i = list.Get(i.Prev()) {
        fmt.Printf("%v ", i.Value())
    }
}
```

### LinkedList

```go
package main

import (
    "fmt"
    "github.com/lxzan/dao/linkedlist"
)

func main() {
    var list = linkedlist.New[int]()
    list.PushBack(1)
    list.PushBack(3)
    list.PushBack(5)
    list.PushBack(7)
    list.PushBack(9)
    for i := list.Front(); i != nil; i = i.Next() {
        fmt.Printf("%v ", i.Value)
    }

    println()
    for i := list.Back(); i != nil; i = i.Prev() {
        fmt.Printf("%v ", i.Value)
    }
}
```

### RBTree

A high-performance red-black tree implementation that can be used as an in-memory database.

```go
package main

import (
    "github.com/lxzan/dao/rbtree"
)

func main() {
    var tree = rbtree.New[int, struct{}]()
    for i := 0; i < 10; i++ {
        tree.Set(i, struct{}{})
    }

    var results = tree.
        NewQuery().
        Left(func(key int) bool { return key >= 3 }).
        Right(func(key int) bool { return key <= 5 }).
        Order(rbtree.ASC).
        FindAll()
    for _, item := range results {
        println(item.Key)
    }
}

```

### Dict

**Dict** aka prefix tree, efficiently matches string prefixes. `dao/dict` can dynamically configure the width
of slots (controlled by the index).

Note: Set the index reasonably, characters beyond the index length can not be optimized for indexing.

```go
package main

import (
    "github.com/lxzan/dao/dict"
)

func main() {
    var tree = dict.New[int]()
    tree.Set("listen", 1)
    tree.Set("list", 2)
    tree.Set("often", 3)
    tree.Set("oh!", 4)
    tree.Set("haha", 5)
    tree.Set("", 6)

    tree.Match("list", func(key string, value int) bool {
        println(key, value)
        return true
    })
}
```

### HashMap

An alias for `Runtime Map`, extending `Keys`, `Values`, `Range` and other utility methods.

```go
package main

import (
    "github.com/lxzan/dao/hashmap"
)

func main() {
    var m = hashmap.New[string, int](8)
    m.Set("a", 1)
    m.Set("b", 2)
    m.Set("c", 3)
    m.Range(func(key string, val int) bool {
        println(key, val)
        return true
    })
}
```

### Segment Tree

**Segment Tree** is a binary tree in which each node represents an interval. Line segment trees are characterized by the
ability to perform interval queries and interval updates in `O(logn)` time.

```go
package main

import (
    "fmt"
    st "github.com/lxzan/dao/segment_tree"
)

func main() {
    var a = []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10}
    var t = st.New(a, st.NewIntSummary[int], st.MergeIntSummary[int])
    var r = t.Query(3, 6)
    fmt.Printf("%v\n", r)
}

```

### Benchmark

- 1,000 elements

```
go test -benchmem -bench '^Benchmark' ./benchmark/
goos: windows
goarch: amd64
pkg: github.com/lxzan/dao/benchmark
cpu: AMD Ryzen 5 PRO 4650G with Radeon Graphics
BenchmarkDict_Set-12                        8647            124.1 ns/op           48190 B/op       1003 allocs/op
BenchmarkDict_Get-12                        9609            122.0 ns/op           48000 B/op       1000 allocs/op
BenchmarkDict_Match-12                      3270            349.3 ns/op           48000 B/op       1000 allocs/op
BenchmarkHeap_Push_Binary-12               55809             21.0 ns/op           38912 B/op          3 allocs/op
BenchmarkHeap_Push_Quadratic-12            65932             18.0 ns/op           38912 B/op          3 allocs/op
BenchmarkHeap_Push_Octal-12                71005             16.1 ns/op           38912 B/op          3 allocs/op
BenchmarkHeap_Pop_Binary-12                10000            100.1 ns/op           16384 B/op          1 allocs/op
BenchmarkHeap_Pop_Quadratic-12             10000            100.7 ns/op           16384 B/op          1 allocs/op
BenchmarkHeap_Pop_Octal-12                  9681            124.9 ns/op           16384 B/op          1 allocs/op
BenchmarkStdList_Push-12                   24715             48.7 ns/op           54000 B/op       1745 allocs/op
BenchmarkStdList_PushAndPop-12             22006             54.7 ns/op           54000 B/op       1745 allocs/op
BenchmarkLinkedList_Push-12                38464             31.7 ns/op           24000 B/op       1000 allocs/op
BenchmarkLinkedList_PushAndPop-12          36898             32.6 ns/op           24000 B/op       1000 allocs/op
BenchmarkDeque_Push-12                    100468             11.7 ns/op           24576 B/op          1 allocs/op
BenchmarkDeque_PushAndPop-12               51649             21.7 ns/op           37496 B/op         12 allocs/op
BenchmarkRBTree_Set-12                      9999            113.9 ns/op           72048 B/op       2001 allocs/op
BenchmarkRBTree_Get-12                     51806             22.7 ns/op               0 B/op          0 allocs/op
BenchmarkRBTree_FindAll-12                  2808            421.3 ns/op          288001 B/op       5000 allocs/op
BenchmarkRBTree_FindAOne-12                 4722            252.2 ns/op           56000 B/op       5000 allocs/op
BenchmarkSegmentTree_Query-12               7498            164.4 ns/op              20 B/op          0 allocs/op
BenchmarkSegmentTree_Update-12             10000            108.6 ns/op              15 B/op          0 allocs/op
BenchmarkSort_Quick-12                     24488             48.5 ns/op               0 B/op          0 allocs/op
BenchmarkSort_Std-12                       21703             54.9 ns/op            8216 B/op          2 allocs/op
PASS
ok      github.com/lxzan/dao/benchmark  31.100s
```

- 1,000,000 elements

```
go test -benchmem -bench '^Benchmark' ./benchmark/
goos: windows
goarch: amd64
pkg: github.com/lxzan/dao/benchmark
cpu: AMD Ryzen 5 PRO 4650G with Radeon Graphics
BenchmarkDict_Set-12                           1        2295.2 ns/op        1405087408 B/op 24868109 allocs/op
BenchmarkDict_Get-12                           2         784.0 ns/op        48000000 B/op    1000000 allocs/op
BenchmarkDict_Match-12                         2         961.0 ns/op        48000000 B/op    1000000 allocs/op
BenchmarkHeap_Push_Binary-12                  48          24.8 ns/op        65708034 B/op          5 allocs/op
BenchmarkHeap_Push_Quadratic-12               58          19.4 ns/op        65708033 B/op          5 allocs/op
BenchmarkHeap_Push_Octal-12                   69          17.1 ns/op        65708033 B/op          5 allocs/op
BenchmarkHeap_Pop_Binary-12                    3         376.3 ns/op        16007168 B/op          1 allocs/op
BenchmarkHeap_Pop_Quadratic-12                 3         342.8 ns/op        16007168 B/op          1 allocs/op
BenchmarkHeap_Pop_Octal-12                     3         374.8 ns/op        16007168 B/op          1 allocs/op
BenchmarkStdList_Push-12                      21          55.0 ns/op        55998007 B/op    1999745 allocs/op
BenchmarkStdList_PushAndPop-12                15          67.5 ns/op        55998008 B/op    1999745 allocs/op
BenchmarkLinkedList_Push-12                   43          29.5 ns/op        24000000 B/op    1000000 allocs/op
BenchmarkLinkedList_PushAndPop-12             39          34.7 ns/op        24000002 B/op    1000000 allocs/op
BenchmarkDeque_Push-12                       123           9.4 ns/op        24002560 B/op          1 allocs/op
BenchmarkDeque_PushAndPop-12                  60          18.7 ns/op        45098876 B/op         37 allocs/op
BenchmarkRBTree_Set-12                         6         171.9 ns/op        72000064 B/op    2000001 allocs/op
BenchmarkRBTree_Get-12                        22          50.1 ns/op               0 B/op          0 allocs/op
BenchmarkRBTree_FindAll-12                     1        1936.8 ns/op        288000128 B/op   5000001 allocs/op
BenchmarkRBTree_FindAOne-12                    1        1793.4 ns/op        56000000 B/op    5000000 allocs/op
BenchmarkSegmentTree_Query-12                  1        1630.0 ns/op        169678048 B/op   2000038 allocs/op
BenchmarkSegmentTree_Update-12                 1        1025.0 ns/op        169678048 B/op   2000038 allocs/op
BenchmarkSort_Quick-12                        10         109.5 ns/op         8003584 B/op          1 allocs/op
BenchmarkSort_Std-12                           9         123.0 ns/op         8003608 B/op          2 allocs/op
PASS
ok      github.com/lxzan/dao/benchmark  47.376s
```
