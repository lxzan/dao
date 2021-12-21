package benchmark

import "github.com/lxzan/dao/internal/utils"

const bench_count = 1000000

var (
	testkeys []string
	testvals []int
)

func init() {
	testkeys = make([]string, 0, 1000000)
	testvals = make([]int, 0, 1000000)
	for i := 0; i < 1000000; i++ {
		var length = utils.Rand.Intn(16) + 1
		testkeys = append(testkeys, utils.Alphabet.Generate(length))
		testvals = append(testvals, utils.Rand.Int())
	}
}
