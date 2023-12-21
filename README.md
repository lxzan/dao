<div align="center">
    <h1>DAO</h1>
    <img src="assets/logo.png" alt="logo" width="300px">
    <h5>道生一, 一生二, 二生三, 三生万物; 万物负阴而抱阳, 冲气以为和</h5>
</div>


[![Build Status](https://github.com/lxzan/dao/workflows/Go%20Test/badge.svg?branch=main)](https://github.com/lxzan/dao/actions?query=branch%3Amain) [![codecov](https://codecov.io/gh/lxzan/dao/graph/badge.svg?token=BQM1JHCDEE)](https://codecov.io/gh/lxzan/dao) [![go-version](https://img.shields.io/badge/go-%3E%3D1.18-30dff3?style=flat-square&logo=go)](https://github.com/lxzan/dao) [![license](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
### 简介

`DAO` 是一个基于泛型的数据结构与算法库

### 目录

- [简介](#简介)
- [目录](#目录)
- [动态数组](#动态数组)
	- [去重](#去重)
	- [排序](#排序)
	- [过滤](#过滤)
- [堆](#堆)
	- [二叉堆](#二叉堆)
	- [四叉堆](#四叉堆)
	- [八叉堆](#八叉堆)
- [栈](#栈)
- [队列](#队列)
- [双端队列](#双端队列)
- [双向链表](#双向链表)
- [红黑树](#红黑树)
- [字典树](#字典树)
- [哈希表](#哈希表)
- [线段树](#线段树)
- [基准测试](#基准测试)

### 动态数组

#### 去重

```go
package main

import (
	"fmt"
	"github.com/lxzan/dao/vector"
)

func main() {
	var v = vector.NewFromInts(1, 3, 5, 3)
	v.Uniq()
	fmt.Printf("%v", v.Elem())
}

```

#### 排序

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

#### 过滤

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

### 堆

**堆** 又称之为优先队列, 堆顶元素总是最大或最小的. 常用的是四叉堆, `Push/Pop` 性能较为均衡.

#### 二叉堆

```go
package main

import (
	"github.com/lxzan/dao/heap"
	"github.com/lxzan/dao/types/cmp"
)

func main() {
	var h = heap.NewWithForks(heap.Binary, cmp.Less[int])
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
	"github.com/lxzan/dao/types/cmp"
)

func main() {
	var h = heap.NewWithForks(heap.Quadratic, cmp.Less[int])
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

#### 八叉堆

```go
package main

import (
	"github.com/lxzan/dao/heap"
	"github.com/lxzan/dao/types/cmp"
)

func main() {
	var h = heap.NewWithForks(heap.Octal, cmp.Less[int])
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

**栈** 先进后出 (`LIFO`) 的数据结构

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

### 队列

**队列** 先进先出 (`FIFO`) 的数据结构. `dao/queue` 在全部元素弹出后会自动重置, 复用内存空间

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

### 双端队列

**双端队列** 类似于双向链表, 两端均可高效执行插入删除操作.

`dao/deque` 基于数组下标模拟指针实现, 删除后的空间后续仍可复用, 且不依赖 `sync.Pool`

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

高性能红黑树实现, 可作为内存数据库使用.

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

### 字典树

**字典树** 又叫前缀树, 可以高效匹配字符串前缀. `dao/dict` 可以动态配置槽位宽度(由索引控制).

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

**线段树** 是一种二叉树，它的每个节点都表示一个区间。 线段树的特点是可以在`O(logn)`的时间内进行区间查询和区间更新。

```go
package main

import (
	tree "github.com/lxzan/dao/segment_tree"
)

func main() {
	var data = []tree.Int64{1, 3, 5, 7, 9, 2, 4, 6, 8, 10}
	var lines = tree.New[tree.Int64Schema, tree.Int64](data)
	var result = lines.Query(0, 10)
	println(result.MinValue, result.MaxValue, result.Sum)
}

```

### 基准测试

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
