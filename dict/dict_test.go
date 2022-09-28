package dict

import (
	"github.com/lxzan/dao/internal/utils"
	"math/rand"
	"sort"
	"testing"
)

const test_count = 1000000

func TestNew(t *testing.T) {
	var d = New[int](8)
	var m = make(map[string]map[string]uint8)
	for i := 0; i < test_count; i++ {
		var key1 = utils.Numeric.Generate(16)
		d.Set(key1, 1)

		var key2 = key1[:4]
		if m[key2] == nil {
			m[key2] = make(map[string]uint8)
		}
		m[key2][key1] = 1
	}

	for k, v := range m {
		var arr1 = make([]string, 0, len(v))
		for k1 := range v {
			arr1 = append(arr1, k1)
		}
		var arr2 = make([]string, 0)
		for _, item := range d.Match(k).Elem() {
			arr2 = append(arr2, item.Key)
		}

		arr1 = utils.UniqueString(arr1)
		arr2 = utils.UniqueString(arr2)
		if !utils.SameStrings(arr1, arr2) {
			t.Fatal("error!")
		}
	}
}

func TestDict_Delete(t *testing.T) {
	var d = New[int](8)
	var m = map[string]int{}

	for i := 0; i < test_count; i++ {
		var key = utils.Numeric.Generate(16)
		var val = rand.Int()
		d.Set(key, val)
		m[key] = val
	}

	var n = d.Len()
	var keys = make([]string, 0)
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for i := 0; i < n/2; i++ {
		var key = keys[i]
		d.Delete(key)
		delete(m, key)
	}

	for i := 0; i < test_count; i++ {
		var key = utils.Numeric.Generate(16)
		var val = rand.Int()
		d.Set(key, val)
		m[key] = val
	}

	if d.Len() != len(m) {
		t.Fail()
		return
	}

	for k, v := range m {
		result, ok := d.Find(k)
		if !ok || result != v {
			t.Fail()
		}
	}
}
