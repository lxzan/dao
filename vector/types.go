package vector

import (
	"github.com/lxzan/dao/types/cmp"
)

type Document[T cmp.Ordered] interface {
	GetID() T
}

type (
	Int int

	Int64 int64

	String string
)

func (c Int) GetID() int {
	return int(c)
}

func (c Int64) GetID() int64 {
	return int64(c)
}

func (c String) GetID() string { return string(c) }
