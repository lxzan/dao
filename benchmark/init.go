package benchmark

import "github.com/lxzan/dao/internal/utils"

const bench_count = 100000

var (
	testkeys []string
	testvals []int
)

func init() {
	testkeys = make([]string, 0, bench_count)
	testvals = make([]int, 0, bench_count)
	for i := 0; i < bench_count; i++ {
		var length = utils.Rand.Intn(16) + 1
		testkeys = append(testkeys, utils.Alphabet.Generate(length))
		testvals = append(testvals, utils.Rand.Int())
	}
}
