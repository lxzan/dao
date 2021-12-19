package benchmark

import (
	"github.com/lxzan/dao/dict"
	"github.com/lxzan/dao/internal/utils"
	"testing"
)

var trie *dict.Dict[int]

func init() {
	trie = dict.New[int]()
	for _, item := range testkeys {
		trie.Insert(item, 1)
	}
}

func BenchmarkDict_Insert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var d = dict.New[int]()
		for j := 0; j < 1000000; j++ {
			d.Insert(testkeys[j], testvals[j])
		}
	}
}

func BenchmarkDict_Delete(b *testing.B) {
	var d = dict.New[int]()
	for i := 0; i < b.N; i++ {
		var j = utils.Rand.Int() % len(testkeys)
		d.Delete(testkeys[j])
	}
}
