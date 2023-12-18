package benchmark

import (
	"github.com/lxzan/dao/internal/utils"
)

const bench_count = 10000

var (
	testkeys []string
	testvals []int
)

func init() {
	testkeys = make([]string, 0, bench_count)
	testvals = make([]int, 0, bench_count)
	for i := 0; i < bench_count; i++ {
		length := 16
		testkeys = append(testkeys, utils.Alphabet.Generate(length))
		testvals = append(testvals, utils.Rand.Int())
	}
}

func getKeys(n int) []string {
	var keys = make([]string, 0, n)
	for i := 0; i < n; i++ {
		keys = append(keys, utils.Alphabet.Generate(16))
	}
	return keys
}
