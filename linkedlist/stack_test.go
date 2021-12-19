package linkedlist

import (
	"github.com/lxzan/dao/internal/utils"
	"testing"
)

var testdata1 = []int{1, 3, 5, 7, 9}
var testdata2 = []int{9, 7, 5, 3, 1}

func TestNewStack(t *testing.T) {
	var stack = NewStack[int]()
	stack.Push(testdata1...)

	var results = make([]int, 0)
	for i := stack.Begin(); !stack.End(i); i = stack.Next(i) {
		results = append(results, i.Data)
	}
	if !utils.SameInts(results, testdata2) {
		t.Error("error!")
	}
}

func TestStack_Pop(t *testing.T) {
	var stack = NewStack[int]()
	stack.Push(testdata1...)

	var results = make([]int, 0)
	for stack.Len() > 0 {
		results = append(results, stack.Pop().Data)
	}
	if !utils.SameInts(results, testdata2) {
		t.Error("error!")
	}
}

func TestStack_Push(t *testing.T) {
	var stack = NewStack[int]()
	stack.Push(testdata1...)
	if stack.Len() != len(testdata1) {
		t.Error("error!")
	}
}
