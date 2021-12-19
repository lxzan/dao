package hash

import "unsafe"

const (
	B16 = 16
	B36 = 36
)

var (
	bases36 = [8]uint64{}
	map36   [256]uint64
)

func init() {
	var base uint64 = 1
	for i := 0; i < 8; i++ {
		bases36[i] = base
		base *= B36
	}

	for i := 0; i < 256; i++ {
		var k = uint8(i)
		var v uint8 = 0
		if k >= '0' && k <= '9' {
			v = k - '0'
		} else if k >= 'A' && k <= 'Z' {
			v = k - 'A' + 10
		} else if k >= 'a' && k <= 'z' {
			v = k - 'a' + 10
		} else {
			v = k % B36
		}
		map36[k] = uint64(v)
	}
}

func Alphabet64(str *string) uint64 {
	var b = *(*[]byte)(unsafe.Pointer(str))
	var n = len(b)
	var sum uint64 = 0
	if n <= 8 {
		for i := 0; i < n; i++ {
			sum += map36[b[i]] * bases36[i]
		}
	} else {
		var temp = make([]byte, 8)
		for i := 0; i < n; i++ {
			temp[i%8] ^= b[i]
		}
		for i, j := range temp {
			sum += map36[j] * bases36[i]
		}
	}
	return sum
}
