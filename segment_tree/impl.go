package segment_tree

import (
	"github.com/lxzan/dao"
	"github.com/lxzan/dao/algorithm"
)

type Schema[T dao.Number[T]] struct {
	MaxValue T
	MinValue T
	Sum      T
}

func Init[T dao.Number[T]](op Operate, x T) Schema[T] {
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

func Merge[T dao.Number[T]](a, b Schema[T]) Schema[T] {
	return Schema[T]{
		MaxValue: algorithm.Max[T](a.MaxValue, b.MaxValue),
		MinValue: algorithm.Min[T](a.MinValue, b.MinValue),
		Sum:      a.Sum + b.Sum,
	}
}
