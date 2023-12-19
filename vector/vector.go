package vector

import (
	"cmp"
	"github.com/lxzan/dao/algorithm"
	"github.com/lxzan/dao/hashmap"
	"reflect"
	"slices"
	"unsafe"
)

type Document[T cmp.Ordered] interface {
	GetID() T
}

type Int int

func (c Int) GetID() int {
	return int(c)
}

func New[D Document[K], K cmp.Ordered](capacity uint32) *Vector[D, K] {
	c := Vector[D, K](make([]D, 0, capacity))
	return &c
}

func NewFromDocs[D Document[K], K cmp.Ordered](arr []D) *Vector[D, K] {
	c := Vector[D, K](arr)
	return &c
}

func NewFromInts(arr []int) *Vector[Int, int] {
	var b = *(*[]Int)(unsafe.Pointer(&arr))
	v := Vector[Int, int](b)
	return &v
}

type Vector[D Document[K], K cmp.Ordered] []D

func (c *Vector[D, K]) Len() int {
	return len(*c)
}

func (c *Vector[D, K]) Elem() []D {
	return *c
}

func (c *Vector[D, K]) Exists(k K) (v D, exist bool) {
	for _, item := range *c {
		if item.GetID() == k {
			return item, true
		}
	}
	return v, exist
}

func (c *Vector[D, K]) Uniq() *Vector[D, K] {
	*c = algorithm.UniqueBy(*c, func(item D) K {
		return item.GetID()
	})
	return c
}

func (c *Vector[D, K]) Filter(f func(i int, v D) bool) *Vector[D, K] {
	*c = algorithm.Filter(*c, f)
	return c
}

func (c *Vector[D, K]) Sort() *Vector[D, K] {
	slices.SortFunc(*c, func(a, b D) int {
		return cmp.Compare(a.GetID(), b.GetID())
	})
	return c
}

func (c *Vector[D, K]) Keys() []K {
	var k K
	var d D
	if unsafe.Sizeof(k) == unsafe.Sizeof(d) && reflect.TypeOf(d).Kind() != reflect.Struct {
		var keys = *(*[]K)(unsafe.Pointer(c))
		return keys
	}

	var keys = make([]K, 0, c.Len())
	for _, item := range *c {
		keys = append(keys, item.GetID())
	}
	return keys
}

func (c *Vector[D, K]) ToMap() hashmap.HashMap[K, D] {
	var m = hashmap.New[K, D](uint32(c.Len()))
	for _, item := range *c {
		m.Set(item.GetID(), item)
	}
	return m
}

func (c *Vector[D, K]) PushBack(v D) {
	*c = append(*c, v)
}

func (c *Vector[D, K]) PopFront() (value D) {
	switch c.Len() {
	case 0:
		return value
	default:
		value = (*c)[0]
		*c = (*c)[1:]
		return value
	}
}

func (c *Vector[D, K]) PopBack() (value D) {
	n := c.Len()
	switch n {
	case 0:
		return value
	default:
		value = (*c)[n-1]
		*c = (*c)[:n-1]
		return value
	}
}

func (c *Vector[D, K]) Range(f func(i int, v D) bool) {
	for index, value := range *c {
		if !f(index, value) {
			return
		}
	}
}
