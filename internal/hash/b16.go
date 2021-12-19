package hash

import "unsafe"

var (
	bases16 = [8]uint64{}
	map16   [256]uint64
)

func init() {
	var base uint64 = 1
	for i := 0; i < 8; i++ {
		bases16[i] = base
		base *= B16
	}

	for i := 0; i < 256; i++ {
		var k = uint8(i)
		var v uint8 = 0
		if k >= '0' && k <= '9' {
			v = k - '0'
		} else if k >= 'A' && k <= 'F' {
			v = k - 'A' + 10
		} else if k >= 'a' && k <= 'f' {
			v = k - 'a' + 10
		} else {
			v = k % B16
		}
		map16[k] = uint64(v)
	}
}

func Hex64(str *string) uint64 {
	b := *(*[]byte)(unsafe.Pointer(str))
	var n = len(b)
	var sum uint64 = 0
	if n <= 8 {
		for i := 0; i < n; i++ {
			sum += map16[b[i]] * bases16[i]
		}
	} else {
		var temp = make([]byte, 8)
		for i := 0; i < n; i++ {
			temp[i%8] ^= b[i]
		}
		for i, j := range temp {
			sum += map16[j] * bases16[i]
		}
	}
	return sum
}
