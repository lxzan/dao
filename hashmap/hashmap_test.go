package hashmap

import (
	"github.com/lxzan/dao"
	"github.com/lxzan/dao/internal/utils"
	"testing"
)

const test_count = 1000

var testdata []string

func init() {
	for i := 0; i < test_count; i++ {
		length := utils.Rand.Intn(16) + 1
		testdata = append(testdata, utils.Alphabet.Generate(length))
	}
}

func TestHashMap(t *testing.T) {
	var m1 = New[string, int]()
	var m2 = make(map[string]int)

	for _, item := range testdata {
		var val = utils.Rand.Int()
		m1.Set(item, val)
		m2[item] = val
	}

	for i := 0; i < test_count/2; i++ {
		m1.Delete(testdata[i])
		delete(m2, testdata[i])
	}

	for i := 0; i < test_count/2; i++ {
		var key = utils.Alphabet.Generate(8)
		var val = utils.Rand.Int()
		m1.Set(key, val)
		m2[key] = val
	}

	if m1.Len() != len(m2) {
		println(m1.Len(), len(m2))
		t.Error("m1.length != m2.length")
	}

	for k, v := range m2 {
		v1, ok := m1.Get(k)
		if !ok || v1 != v {
			t.Error("error!")
		}
	}
}

func TestHashMap_ForEach(t *testing.T) {
	var m1 = New[string, int]()
	var m2 = make(map[string]int)

	for _, item := range testdata {
		var val = utils.Rand.Int()
		m1.Set(item, val)
		m2[item] = val
	}

	var sum = 0
	m1.ForEach(func(key string, value int) bool {
		sum++
		if m2[key] != value {
			t.Error("error!")
		}
		return dao.Continue
	})

	if m1.Len() != len(m2) || sum != len(m2) {
		println(m1.Len(), len(m2))
		t.Error("m1.length != m2.length")
	}
}
