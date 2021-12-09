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
	var node = &queryNode{
		Node:  c.root,
		X:     x,
		Times: length * int(BIT),
		Key:   key,
		Val:   val,
	}
	c.set(node)
}

type queryNode struct {
	Node  *Element
	X     uint64
	Times int
	Key   string
	Val   int
}

//var deep = 0
func (c *Trie) set(cur *queryNode) {
	cur.Times--
	var idx = cur.X % BIT
	if cur.Times == 0 {
		if cur.Node.List == nil {
			cur.Node.List = NewQueue()
		}

		var exist = false
		cur.Node.List.ForEach(func(ele *ListElement) (next bool) {
			if ele.Value.Key == cur.Key {
				exist = true
				ele.Value.Val = cur.Val
				return false
			}
			return true
		})

		if !exist {
			c.length++
			cur.Node.List.Push(Pair{
				Key: cur.Key,
				Val: cur.Val,
			})
		}
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

func (c *Trie) Get(key string) (int, bool) {
	length, x := c.format(&key)
	var node = &queryNode{
		Node:  c.root,
		X:     x,
		Times: length * 4,
		Key:   key,
	}
	exist := c.get(node)
	return node.Val, exist
}

func (c *Trie) get(cur *queryNode) (exist bool) {
	cur.Times--
	var idx = cur.X % BIT
	if cur.Times == 0 {
		if cur.Node.List == nil {
			return false
		}

		cur.Node.List.ForEach(func(ele *ListElement) (next bool) {
			if ele.Value.Key == cur.Key {
				cur.Val = ele.Value.Val
				exist = true
				return false
			}
			return true
		})
		return
	}

	if cur.Node.Children[idx] == nil {
		return false
	}
	cur.X /= BIT
	cur.Node = cur.Node.Children[idx]
	return c.get(cur)
}

func (c *Trie) PrefixMatch(key string) []Pair {
	length, x := c.format(&key)
	var node = &queryNode{
		Node:  c.root,
		X:     x,
		Times: length * 4,
		Key:   key,
	}

	var arr = make([]Pair, 0)
	c.matchPairs(c.matchElement(node), &arr)
	return arr
}

func (c *Trie) matchElement(cur *queryNode) *Element {
	cur.Times--
	var idx = cur.X % BIT
	if cur.Times == 0 {
		return cur.Node

	}
	if cur.Node.Children[idx] == nil {
		return nil
	}
	cur.X /= BIT
	cur.Node = cur.Node.Children[idx]
	return c.matchElement(cur)
}

func (c *Trie) matchPairs(cur *Element, arr *[]Pair) {
	if cur == nil {
		return
	}

	if cur.List != nil {
		cur.List.ForEach(func(ele *ListElement) (next bool) {
			*arr = append(*arr, ele.Value)
			return true
		})
	}
	for _, item := range cur.Children {
		c.matchPairs(item, arr)
	}
}
