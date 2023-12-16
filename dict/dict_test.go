package dict

import (
	"fmt"
	"github.com/lxzan/dao/internal/utils"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"testing"
)

const test_count = 1000

func TestNew(t *testing.T) {
	var d = New[int]()
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
		d.Match(k, func(key string, value int) bool {
			arr2 = append(arr2, key)
			return true
		})

		arr1 = utils.UniqueString(arr1)
		arr2 = utils.UniqueString(arr2)
		if !utils.SameStrings(arr1, arr2) {
			t.Fatal("error!")
		}
	}
}

func TestDict_Delete(t *testing.T) {
	var d = New[int]()
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
		result, ok := d.Get(k)
		if !ok || result != v {
			t.Fail()
		}
	}
}

func TestDict_Match(t *testing.T) {
	var d = New[uint8]()
	d.Set("hasaki", 1)
	d.Set("oh", 2)
	d.Set("haha", 3)
	var list []string
	d.Match("ha", func(key string, value uint8) bool {
		list = append(list, key)
		return len(list) < 1
	})
	assert.Equal(t, len(list), 1)
}

func TestDict_Range(t *testing.T) {
	var d = New[uint8]()
	d.Set("hasaki", 1)
	d.Set("oh", 2)
	d.Set("haha", 3)
	var list []string
	d.Range(func(key string, value uint8) bool {
		list = append(list, key)
		return len(list) < 1
	})
	assert.Equal(t, len(list), 1)
}

func TestDict_Get(t *testing.T) {
	indexes := []uint8{4, 4, 4, 4}

	t.Run("", func(t *testing.T) {
		var d = New[int]().WithIndexes(indexes)
		d.Set("", 1)
		d.Set("", 2)
		val, _ := d.Get("")
		assert.Equal(t, val, 2)

		d.Reset()
		assert.Equal(t, d.Len(), 0)
	})

	t.Run("", func(t *testing.T) {
		var d = New[int]().WithIndexes(indexes)
		d.Set("a", 1)
		d.Delete("a")
		_, ok := d.Get("a")
		assert.False(t, ok)
	})

	t.Run("", func(t *testing.T) {
		var d = New[int]().WithIndexes([]uint8{1, 1})
		d.Set("abc", 1)
		d.Set("abd", 2)
		_, ok := d.Get("abc")
		assert.True(t, ok)

		for i := d.begin("abd", false); i != nil; i = i.next() {
		}
	})
}

func TestDict_WithIndexes(t *testing.T) {
	var f = func(cb func()) (err error) {
		defer func() {
			if exception := recover(); exception != nil {
				err = fmt.Errorf("%v", exception)
			}
		}()
		cb()
		return
	}

	t.Run("", func(t *testing.T) {
		err := f(func() {
			New[uint8]().WithIndexes([]uint8{1, 2, 3, 4})
		})
		assert.Error(t, err)
	})

	t.Run("", func(t *testing.T) {
		err := f(func() {
			New[uint8]().WithIndexes([]uint8{1})
		})
		assert.Error(t, err)
	})
}

func TestDict_Random(t *testing.T) {
	var count = 1000000
	var d = New[int]()
	var m = make(map[string]int)
	for i := 0; i < count; i++ {
		key := strconv.Itoa(i)
		val := rand.Int()
		flag := rand.Intn(4)
		switch flag {
		case 0, 1, 2:
			d.Set(key, val)
			m[key] = val
		default:
			d.Delete(key)
			delete(m, key)
		}
	}

	assert.Equal(t, d.Len(), len(m))
	d.Range(func(key string, value int) bool {
		assert.Equal(t, m[key], value)
		return true
	})

	for i := 0; i < 10; i++ {
		var prefix = strconv.Itoa(rand.Intn(10000))
		var list1 []string
		var list2 []string
		for k, _ := range m {
			if strings.HasPrefix(k, prefix) {
				list1 = append(list1, k)
			}
		}
		d.Match(prefix, func(key string, value int) bool {
			list2 = append(list2, key)
			return true
		})
		assert.ElementsMatch(t, list1, list2)
	}
}
