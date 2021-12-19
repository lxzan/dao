package double_linkedlist

import (
	"container/list"
	"github.com/lxzan/dao/internal/utils"
	"testing"
)

func TestList_Push(t *testing.T) {
	var list1 = New[int]()
	var list2 = list.New()

	for i := 0; i < 100; i++ {
		var x = utils.Rand.Int()
		list1.RPush(x)
		list2.PushBack(x)
	}
	for i := 0; i < 100; i++ {
		var x = utils.Rand.Int()
		list1.LPush(x)
		list2.PushFront(x)
	}
	for i := 0; i < 100; i++ {
		var x = utils.Rand.Int()
		list1.RPush(x)
		list2.PushBack(x)
	}
	for i := 0; i < 100; i++ {
		var x = utils.Rand.Int()
		list1.LPush(x)
		list2.PushFront(x)
	}

	var arr1 = make([]int, 0)
	var arr2 = make([]int, 0)
	for list1.Len() > 0 {
		arr1 = append(arr1, list1.LPop().Data)
	}
	for list2.Len() > 0 {
		var ele = list2.Front()
		arr2 = append(arr2, ele.Value.(int))
		list2.Remove(ele)
	}
	if !utils.SameInts(arr1, arr2) {
		t.Fatal("error!")
	}
}

func TestList_Delete(t *testing.T) {
	var dl = New[int]()
	var m = make(map[int]uint8)

	for i := 0; i < 1000; i++ {
		var x = utils.Rand.Int()
		m[x] = 1
	}
	for k, _ := range m {
		dl.RPush(k)
	}

	for i := dl.Begin(); !dl.End(i); i = dl.Next(i) {
		var x = utils.Rand.Int()
		if x%2 == 0 {
			delete(m, i.Data)
			dl.Delete(i)
		}
	}

	for i := dl.Begin(); !dl.End(i); i = dl.Next(i) {
		if _, exist := m[i.Data]; !exist {
			t.Fatal("error!")
		}
	}

	for i := dl.Begin(); !dl.End(i); i = dl.Next(i) {
		dl.Delete(i)
	}
	if dl.Len() != 0 {
		t.Fatal("error!")
	}
}

func TestList_InsertAfter(t *testing.T) {
	var dl = New[int]()
	dl.RPush(1)
	dl.RPush(9)
	var item1 = dl.Front()
	dl.InsertAfter(item1, 3)
	var item2 = dl.Next(item1)
	dl.InsertAfter(item2, 5)
	var item3 = dl.Next(item2)
	dl.InsertAfter(item3, 7)

	var arr = make([]int, 0)
	for i := dl.Begin(); !dl.End(i); i = dl.Next(i) {
		arr = append(arr, i.Data)
	}
	if !utils.SameInts(arr, []int{1, 3, 5, 7, 9}) {
		t.Fatal("error!")
	}
}

func TestList_InsertBefore(t *testing.T) {
	var dl = New[int]()
	dl.RPush(1)
	dl.RPush(9)
	var item1 = dl.Back()
	dl.InsertBefore(item1, 7)
	var item2 = item1.prev
	dl.InsertBefore(item2, 5)
	var item3 = item2.prev
	dl.InsertBefore(item3, 3)

	var arr = make([]int, 0)
	for i := dl.Begin(); !dl.End(i); i = dl.Next(i) {
		arr = append(arr, i.Data)
	}
	if !utils.SameInts(arr, []int{1, 3, 5, 7, 9}) {
		t.Fatal("error!")
	}
}
