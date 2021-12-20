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

func (c Schema[T]) Init(op Operate, x T) Schema[T] {
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

func (c Schema[T]) Merge(a, b Schema[T]) Schema[T] {
	return Schema[T]{
		MaxValue: algorithm.Max[T](a.MaxValue, b.MaxValue),
		MinValue: algorithm.Min[T](a.MinValue, b.MinValue),
		Sum:      a.Sum + b.Sum,
	}
}
