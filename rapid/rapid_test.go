package rapid

import (
	"github.com/lxzan/dao/double_linkedlist"
	"github.com/lxzan/dao/internal/utils"
	"testing"
)

func TestRapid_Push(t *testing.T) {
	var queens1 = New[string, int](8)
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
		queens1.Push(&entrypoints[j], key, val)
		queens2[j].RPush(val)
	}

	for i := 0; i < 10; i++ {
		var arr1 = make([]int, 0)
		var arr2 = make([]int, 0)
		for j := queens1.Begin(&entrypoints[i]); !queens1.End(j); j = queens1.Next(j) {
			arr1 = append(arr1, j.Value)
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
	const test_count = 10000
	var q = New[string, int](0)
	var ptr = q.NextID()
	var entrypoint = &EntryPoint{Head: ptr, Tail: ptr}
	var m = make(map[string]int)
	var keys = make([]string, 0)

	for i := 0; i < test_count; i++ {
		var key = utils.Alphabet.Generate(16)
		m[key] = 1
		q.Push(entrypoint, key, 1)
		keys = append(keys, key)
	}

	for i := 0; i < test_count/2; i++ {
		var ptr = entrypoint.Tail
		var iter = &q.Buckets[ptr]
		var key = iter.Key
		q.Delete(entrypoint, iter)
		delete(m, key)
	}

	for i := 0; i < test_count; i++ {
		var key = utils.Alphabet.Generate(16)
		m[key] = 1
		q.Push(entrypoint, key, 1)
	}

	var arr1 = make([]string, 0)
	var arr2 = make([]string, 0)
	for k, _ := range m {
		arr1 = append(arr1, k)
	}
	for i := q.Begin(entrypoint); !q.End(i); i = q.Next(i) {
		arr2 = append(arr2, i.Key)
	}
	if !utils.SameStrings(arr1, arr2) {
		t.Fail()
	}
}
