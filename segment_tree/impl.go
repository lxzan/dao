package segment_tree

import (
	"github.com/lxzan/dao/algorithm"
)

type Int64 int64

// Init 初始化摘要结构
func (c Int64) Init(op Operate) Int64Schema {
	var val = int64(c)
	var result = Int64Schema{
		MaxValue: val,
		MinValue: val,
		Sum:      val,
	}
	if op == OperateQuery {
		result.Sum = 0
	}
	return result
}

func (c Int64) Value() int64 {
	return int64(c)
}

type Int64Schema struct {
	MaxValue int64
	MinValue int64
	Sum      int64
}

// Merge 合并摘要信息
func (c Int64Schema) Merge(d Int64Schema) Int64Schema {
	return Int64Schema{
		MaxValue: algorithm.Max(c.MaxValue, d.MaxValue),
		MinValue: algorithm.Min(c.MinValue, d.MinValue),
		Sum:      c.Sum + d.Sum,
	}
}
