package dict

import (
	"github.com/lxzan/dao/internal/utils"
	"testing"
)

const test_count = 100000

func TestNew(t *testing.T) {
	var d = New[int]()
	var m = make(map[string]map[string]uint8)
	for i := 0; i < test_count; i++ {
		var key1 = utils.Numeric.Generate(16)
		d.Insert(key1, 1)

		var key2 = key1[:4]
		if m[key2] == nil {
			m[key2] = make(map[string]uint8)
		}
		m[key2][key1] = 1
	}

	for k, v := range m {
		var arr1 = make([]string, 0, len(v))
		for k1, _ := range v {
			arr1 = append(arr1, k1)
		}
		var arr2 = make([]string, 0)
		for _, item := range d.Match(k) {
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
	var d = New[int]()
	var m = make(map[string]map[string]uint8)
	for i := 0; i < test_count; i++ {
		var key1 = utils.Numeric.Generate(16)
		d.Insert(key1, 1)

		var key2 = key1[:4]
		if m[key2] == nil {
			m[key2] = make(map[string]uint8)
		}
		m[key2][key1] = 1
	}

	for _, v := range m {
		var i = 0
		var n = len(v)
		for k1, _ := range v {
			if i >= n/2 {
				break
			}
			delete(v, k1)
			d.Delete(k1)
		}
	}

	for i := 0; i < test_count; i++ {
		var key1 = utils.Alphabet.Generate(16)
		d.Insert(key1, 1)

		var key2 = key1[:4]
		if m[key2] == nil {
			m[key2] = make(map[string]uint8)
		}
		m[key2][key1] = 1
	}

	for k, v := range m {
		var arr1 = make([]string, 0, len(v))
		for k1, _ := range v {
			arr1 = append(arr1, k1)
		}
		var arr2 = make([]string, 0)
		for _, item := range d.Match(k) {
			arr2 = append(arr2, item.Key)
		}

		arr1 = utils.UniqueString(arr1)
		arr2 = utils.UniqueString(arr2)
		if !utils.SameStrings(arr1, arr2) {
			t.Fatal("error!")
		}
	}
}
