package rapid

import (
	"github.com/lxzan/dao/double_linkedlist"
	"github.com/lxzan/dao/internal/utils"
	"testing"
)

func TestRapid_Push(t *testing.T) {
	var queens1 = New[string, int]()
	var queens2 = make([]*double_linkedlist.List[int], 0)
	var entrypoints = make([]EntryPoint, 0)
	for i := 0; i < 10; i++ {
		var ptr = queens1.NextID()
		entrypoints = append(entrypoints, EntryPoint{Head: ptr, Tail: ptr})
		queens2 = append(queens2, double_linkedlist.New[int]())
	}

	for i := 0; i < 10000; i++ {
		var j = i % 10
		var key = utils.Alphabet.Generate(8)
		var val = utils.Rand.Int()
		queens1.Push(&entrypoints[j], &key, &val)
		queens2[j].RPush(val)
	}

	for i := 0; i < 10; i++ {
		var arr1 = make([]int, 0)
		var arr2 = make([]int, 0)
		for j := queens1.Begin(entrypoints[i]); !queens1.End(j); j = queens1.Next(j) {
			arr1 = append(arr1, j.Data)
		}
		for j := queens2[i].Begin(); !queens2[i].End(j); j = queens2[i].Next(j) {
			arr2 = append(arr2, j.Data)
		}
		if !utils.SameInts(arr1, arr2) {
			t.Fatal("error!")
		}
	}
}

func TestRapid_Delete(t *testing.T) {
	var queens1 = New[string, int]()
	var queens2 = make([]*double_linkedlist.List[int], 0)
	var entrypoints = make([]EntryPoint, 0)
	for i := 0; i < 10; i++ {
		var ptr = queens1.NextID()
		entrypoints = append(entrypoints, EntryPoint{Head: ptr, Tail: ptr})
		queens2 = append(queens2, double_linkedlist.New[int]())
	}

	for i := 0; i < 10000; i++ {
		var j = i % 10
		var key = utils.Alphabet.Generate(8)
		var val = utils.Rand.Int()
		queens1.Push(&entrypoints[j], &key, &val)
		queens2[j].RPush(val)
	}

	for i := 0; i < 10; i++ {
		var k = utils.Rand.Intn(10)
		var idx = 0
		for j := queens1.Begin(entrypoints[i]); !queens1.End(j); j = queens1.Next(j) {
			if idx == k {
				queens1.Delete(&entrypoints[i], j)
				break
			}
			idx++
		}

		idx = 0
		for j := queens2[i].Begin(); !queens2[i].End(j); j = queens2[i].Next(j) {
			if idx == k {
				queens2[i].Delete(j)
				break
			}
			idx++
		}
	}

	for i := 0; i < 10; i++ {
		var arr1 = make([]int, 0)
		var arr2 = make([]int, 0)
		for j := queens1.Begin(entrypoints[i]); !queens1.End(j); j = queens1.Next(j) {
			arr1 = append(arr1, j.Data)
		}
		for j := queens2[i].Begin(); !queens2[i].End(j); j = queens2[i].Next(j) {
			arr2 = append(arr2, j.Data)
		}
		if !utils.SameInts(arr1, arr2) {
			t.Fatal("error!")
		}
	}
}
