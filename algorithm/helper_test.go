package algorithm

import (
	"github.com/lxzan/dao/internal/utils"
	"testing"
)

func TestToString(t *testing.T) {
	if ToString(123) != "123" {
		t.Fatal("error!")
	}
}

func TestUnique(t *testing.T) {
	var n = 1000
	var arr1 = make([]int, 0)
	for i := 0; i < n; i++ {
		arr1 = append(arr1, utils.Rand.Intn(100))
	}
	Unique(&arr1, func(x int) int { return x })

	var length = len(arr1)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if arr1[i] == arr1[j] {
				t.Fatal("error!")
			}
		}
	}
}
