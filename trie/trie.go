package trie

import (
	"encoding/binary"
	"unsafe"
)

const BIT uint64 = 4

type Element struct {
	List     *Queue
	Children [BIT]*Element
}

type Trie struct {
	length int
	root   *Element
}

func NewTrie() *Trie {
	return &Trie{
		length: 0,
		root:   &Element{},
	}
}

func (c *Trie) Len() int {
	return c.length
}

func (c *Trie) format(key *string) (int, uint64) {
	var b = *(*[]byte)(unsafe.Pointer(key))
	var n = len(b)
	var a = make([]byte, 8)
	if n > 8 {
		copy(a, b[:8])
	} else {
		copy(a, b)
	}
	return n, binary.LittleEndian.Uint64(a)
}

func (c *Trie) Set(key string, val int) {
	length, x := c.format(&key)
	var node = &QueryNode{
		Node:  c.root,
		X:     x,
		Times: length * 4,
		Key:   key,
		Val:   val,
	}
	c.set(node)
}

type QueryNode struct {
	Node  *Element
	X     uint64
	Times int
	Key   string
	Val   int
}

//var deep = 0
func (c *Trie) set(cur *QueryNode) {
	cur.Times--
	var idx = cur.X % BIT
	if cur.Times == 0 {
		if cur.Node.List == nil {
			cur.Node.List = NewQueue()
		}
		cur.Node.List.Push(Pair{
			Key: cur.Key,
			Val: cur.Val,
		})
		return
	}

	if cur.Node.Children[idx] == nil {
		cur.Node.Children[idx] = &Element{
			Children: [BIT]*Element{},
		}
	}
	cur.X /= BIT
	cur.Node = cur.Node.Children[idx]
	c.set(cur)
}

//func (c *Trie) Get(key string) (int, bool) {
//	var x = c.FormatKey(key)
//	var val = 0
//	var exist = false
//	c.get(c.nodes, x, key, &val, &exist)
//	return val, exist
//}
//
//func (c *Trie) get(node *TrieNode, x uint32, key string, val *int, exist *bool) {
//	if x == 0 {
//		if node.List != nil {
//			node.List.ForEach(func(ele *Element) bool {
//				var p = ele.Value
//				if p.Key == key {
//					*exist = true
//					*val = p.Val
//					return false
//				}
//				return true
//			})
//		}
//		return
//	}
//
//	var idx = x % BIT
//	if node.Children[idx] == nil {
//		return
//	}
//	c.get(node.Children[idx], x/BIT, key, val, exist)
//}
