package segment_tree

import (
	"github.com/lxzan/dao/algo"
	"github.com/lxzan/dao/types/cmp"
)

type (
	NewSummary[T any, S any] func(T, Operate) S

	MergeSummary[S any] func(a, b S) S
)

type IntSummary[T cmp.Integer] struct {
	MaxValue T
	MinValue T
	Sum      T
}

func NewIntSummary[T cmp.Integer](num T, op Operate) IntSummary[T] {
	var r = IntSummary[T]{MaxValue: num, MinValue: num, Sum: num}
	if op == OperateQuery {
		r.Sum = 0
	}
	return r
}

func MergeIntSummary[T cmp.Integer](a, b IntSummary[T]) IntSummary[T] {
	return IntSummary[T]{
		MaxValue: algo.Max(a.MaxValue, b.MaxValue),
		MinValue: algo.Min(a.MinValue, b.MinValue),
		Sum:      a.Sum + b.Sum,
	}
}
