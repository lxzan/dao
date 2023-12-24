package hashmap

import (
	"github.com/lxzan/dao/internal/utils"
	"github.com/lxzan/dao/internal/validator"
	"github.com/stretchr/testify/assert"
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
	var m1 = New[string, int](0)
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
	var m1 = New[string, int](0)
	var m2 = make(map[string]int)

	for _, item := range testdata {
		var val = utils.Rand.Int()
		m1.Set(item, val)
		m2[item] = val
	}

	var sum = 0
	m1.Range(func(key string, val int) bool {
		sum++
		if m2[key] != val {
			t.Error("error!")
		}
		return true
	})

	if m1.Len() != len(m2) || sum != len(m2) {
		println(m1.Len(), len(m2))
		t.Error("m1.length != m2.length")
	}
}

func TestDict_Map(t *testing.T) {
	assert.True(t, validator.ValidateMapImpl(New[string, int](8)))
}

func TestHashMap_Keys(t *testing.T) {
	var m = New[string, int](8)
	m.Set("a", 1)
	m.Set("b", 2)
	m.Set("c", 3)
	assert.ElementsMatch(t, m.Keys(), []string{"a", "b", "c"})
}

func TestHashMap_Values(t *testing.T) {
	var m = New[string, int](8)
	m.Set("a", 1)
	m.Set("b", 2)
	m.Set("c", 3)
	assert.ElementsMatch(t, m.Values(), []int{1, 2, 3})
}

func TestHashMap_Range(t *testing.T) {
	var m = New[string, int](8)
	m.Set("a", 1)
	m.Set("b", 2)
	m.Set("c", 3)
	assert.True(t, m.Exists("a"))
	assert.False(t, m.Exists("d"))

	var keys []string
	m.Range(func(key string, val int) bool {
		keys = append(keys, key)
		return true
	})
	assert.ElementsMatch(t, keys, []string{"a", "b", "c"})

	keys = keys[:0]
	m.Range(func(key string, val int) bool {
		keys = append(keys, key)
		return len(keys) < 2
	})
	assert.Equal(t, len(keys), 2)
}
