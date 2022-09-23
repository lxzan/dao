package segment_tree

import (
	"github.com/lxzan/dao"
	"github.com/lxzan/dao/algorithm"
)

type Schema[T dao.Number] struct {
	MaxValue T
	MinValue T
	Sum      T
}

func Init[T dao.Number](op Operate, x T) Schema[T] {
	var result = Schema[T]{
		MaxValue: x,
		MinValue: x,
		Sum:      x,
	}
	if op == Operate_Query {
		result.Sum = 0
	}
	return result
}

func Merge[T dao.Number](a, b Schema[T]) Schema[T] {
	return Schema[T]{
		MaxValue: algorithm.Max(a.MaxValue, b.MaxValue),
		MinValue: algorithm.Min(a.MinValue, b.MinValue),
		Sum:      a.Sum + b.Sum,
	}
}
