package types

type Ordering int8

const (
	Less    Ordering = -1
	Equal   Ordering = 0
	Greater Ordering = 1
)

type comparable[T any] interface {
	Compare(a, b T) Ordering
}

type (
	String string
	Uint   uint
	Uint64 uint64
	Uint32 uint32
	Uint16 uint16
	Uint8  uint8
	Int    int
	Int64  int64
	Int32  int32
	Int16  int16
	Int8   int8
)

func (c String) Compare(a, b String) Ordering {
	if a > b {
		return Greater
	} else if a == b {
		return Equal
	} else {
		return Less
	}
}

func (c Uint) Compare(a, b Uint) Ordering {
	if a > b {
		return Greater
	} else if a == b {
		return Equal
	} else {
		return Less
	}
}

func (c Uint64) Compare(a, b Uint64) Ordering {
	if a > b {
		return Greater
	} else if a == b {
		return Equal
	} else {
		return Less
	}
}

func (c Uint32) Compare(a, b Uint32) Ordering {
	if a > b {
		return Greater
	} else if a == b {
		return Equal
	} else {
		return Less
	}
}

func (c Uint16) Compare(a, b Uint16) Ordering {
	if a > b {
		return Greater
	} else if a == b {
		return Equal
	} else {
		return Less
	}
}

func (c Uint8) Compare(a, b Uint8) Ordering {
	if a > b {
		return Greater
	} else if a == b {
		return Equal
	} else {
		return Less
	}
}

func (c Int) Compare(a, b Int) Ordering {
	if a > b {
		return Greater
	} else if a == b {
		return Equal
	} else {
		return Less
	}
}

func (c Int64) Compare(a, b Int64) Ordering {
	if a > b {
		return Greater
	} else if a == b {
		return Equal
	} else {
		return Less
	}
}

func (c Int32) Compare(a, b Int32) Ordering {
	if a > b {
		return Greater
	} else if a == b {
		return Equal
	} else {
		return Less
	}
}

func (c Int16) Compare(a, b Int16) Ordering {
	if a > b {
		return Greater
	} else if a == b {
		return Equal
	} else {
		return Less
	}
}

func (c Int8) Compare(a, b Int8) Ordering {
	if a > b {
		return Greater
	} else if a == b {
		return Equal
	} else {
		return Less
	}
}
