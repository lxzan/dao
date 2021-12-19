package linkedlist

import (
	"github.com/lxzan/dao/internal/utils"
	"testing"
)

func TestNewQueue(t *testing.T) {
	var q = NewQueue[int]()
	q.Push(testdata1...)

	var results = make([]int, 0)
	for i := q.Begin(); !q.End(i); i = q.Next(i) {
		results = append(results, i.Data)
	}
	if !utils.SameInts(results, testdata1) {
		t.Error("error!")
	}
}

func TestQueue_Pop(t *testing.T) {
	var q = NewQueue[int]()
	q.Push(testdata1...)

	var results = make([]int, 0)
	for q.Len() > 0 {
		results = append(results, q.Pop().Data)
	}
	if !utils.SameInts(results, testdata1) {
		t.Error("error!")
	}
}

func TestQueue_Push(t *testing.T) {
	var q = NewQueue[int]()
	q.Push(testdata1...)
	if q.Len() != len(testdata1) {
		t.Error("error!")
	}
}
