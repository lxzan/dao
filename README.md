<div align="center">
    <h1>DAO</h1>
    <img src="assets/logo.png" alt="logo" width="300px">
    <h5>道生一, 一生二, 二生三, 三生万物; 万物负阴而抱阳, 冲气以为和.</h5>
</div>


[![Build Status](https://github.com/lxzan/dao/workflows/Go%20Test/badge.svg?branch=main)](https://github.com/lxzan/dao/actions?query=branch%3Amain) [![go-version](https://img.shields.io/badge/go-%3E%3D1.21-30dff3?style=flat-square&logo=go)](https://github.com/lxzan/dao)

### 简介

Go 数据结构与算法库

### 目录

- [简介](#简介)
- [目录](#目录)
- [堆](#堆)
	- [二叉堆](#二叉堆)
	- [四叉堆](#四叉堆)
- [栈](#栈)
- [双端队列](#双端队列)
- [双向链表](#双向链表)
- [红黑树](#红黑树)
	- [区间查询](#区间查询)
	- [极值查询](#极值查询)
- [前缀树](#前缀树)
- [哈希表](#哈希表)
- [线段树](#线段树)
- [基准测试](#基准测试)

### 堆

#### 二叉堆

```go
package main

import (
	"github.com/lxzan/dao/heap"
)

func main() {
	var h = heap.New[int]().SetForkNumber(heap.Binary)
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

#### 四叉堆

```go
package main

import (
	"github.com/lxzan/dao/heap"
)

func main() {
	var h = heap.New[int]().SetForkNumber(heap.Quadratic)
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

### 栈

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

### 双端队列

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

### 双向链表

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

### 红黑树

#### 区间查询

```go
package main

import (
	"github.com/lxzan/dao"
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
		Order(dao.DESC).
		Do()
	for _, item := range results {
		println(item.Key)
	}
}
```

#### 极值查询

```go
package main

import (
	"fmt"
	"github.com/lxzan/dao/rbtree"
)

func main() {
	var tree = rbtree.New[int, struct{}]()
	for i := 0; i < 10; i++ {
		tree.Set(i, struct{}{})
	}

	minimum, _ := tree.GetMinKey(rbtree.TrueFunc[int])
	maximum, _ := tree.GetMaxKey(rbtree.TrueFunc[int])
	fmt.Printf("%v %v", minimum.Key, maximum.Key)
}
```

### 前缀树

可以动态配置槽位宽度的前缀树

注意: 合理设置索引, 超出索引长度的字符不能被索引优化

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

### 哈希表

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

### 线段树

```go
package main

import (
	tree "github.com/lxzan/dao/segment_tree"
)

func main() {
	var data = []tree.Int64{1, 3, 5, 7, 9, 2, 4, 6, 8, 10}
	var lines = tree.New[tree.Int64Schema, tree.Int64](data)
	var result = lines.Query(0, 9)
	println(result.MinValue, result.MaxValue, result.Sum)
}

```

### 基准测试

- 10,000 elements

```
go test -benchmem -bench '^Benchmark' ./benchmark/
goos: windows
goarch: amd64
pkg: github.com/lxzan/dao/benchmark
cpu: AMD Ryzen 5 PRO 4650G with Radeon Graphics
BenchmarkDict_Set-12                         423           2449370 ns/op          517811 B/op      10645 allocs/op
BenchmarkDict_Get-12                         499           2419686 ns/op          480001 B/op      10000 allocs/op
BenchmarkDict_Match-12                       265           4561133 ns/op          480000 B/op      10000 allocs/op
BenchmarkHeap_Push_Binary-12                4455            287363 ns/op          507905 B/op          4 allocs/op
BenchmarkHeap_Push_Quadratic-12             5960            264085 ns/op          507906 B/op          4 allocs/op
BenchmarkHeap_Push_Octal-12                 5793            225979 ns/op          507907 B/op          4 allocs/op
BenchmarkHeap_Pop_Binary-12                  808           1497473 ns/op          163840 B/op          1 allocs/op
BenchmarkHeap_Pop_Quadratic-12               846           1454926 ns/op          163840 B/op          1 allocs/op
BenchmarkHeap_Pop_Octal-12                   673           1788673 ns/op          163840 B/op          1 allocs/op
BenchmarkStdList_Push-12                    1958            596811 ns/op          558002 B/op      19745 allocs/op
BenchmarkStdList_PushAndPop-12              1729            652220 ns/op          558001 B/op      19745 allocs/op
BenchmarkLinkedList_Push-12                 3770            319706 ns/op          240001 B/op      10000 allocs/op
BenchmarkLinkedList_PushAndPop-12           2539            465386 ns/op          240002 B/op      10000 allocs/op
BenchmarkDeque_Push-12                      8560            122805 ns/op          245761 B/op          1 allocs/op
BenchmarkDeque_PushAndPop-12                5599            378346 ns/op          386937 B/op         18 allocs/op
BenchmarkRBTree_Set-12                       540           2194732 ns/op          720051 B/op      20001 allocs/op
BenchmarkRBTree_Get-12                      3272            365139 ns/op               0 B/op          0 allocs/op
BenchmarkRBTree_Query-12                      60          18096230 ns/op         3680048 B/op      60000 allocs/op
BenchmarkSegmentTree_Query-12                418           2734015 ns/op            3917 B/op         47 allocs/op
BenchmarkSegmentTree_Update-12               686           1745124 ns/op            2387 B/op         29 allocs/op
BenchmarkSort_Quick-12                      1588            758722 ns/op           81920 B/op          1 allocs/op
BenchmarkSort_Std-12                        1377            862120 ns/op           81944 B/op          2 allocs/op
PASS
ok      github.com/lxzan/dao/benchmark  32.279s
```