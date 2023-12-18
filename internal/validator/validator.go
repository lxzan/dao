package validator

import (
	"github.com/lxzan/dao"
	"github.com/lxzan/dao/internal/utils"
	"math/rand"
)

func ValidateMapImpl(m dao.Map[string, int]) bool {
	const count = 10000
	var std = make(map[string]int)

	for i := 0; i < count; i++ {
		k := utils.Numeric.Generate(4)
		v := rand.Int()
		flag := rand.Intn(10)
		switch flag {
		case 0, 1, 2:
			m.Set(k, v)
			std[k] = v
		case 3:
			m.Delete(k)
			delete(std, k)
		}
	}

	if len(std) != m.Len() {
		return false
	}

	for k, v := range std {
		value, ok := m.Get(k)
		if !ok {
			return false
		}
		if v != value {
			return false
		}
	}

	return true
}
