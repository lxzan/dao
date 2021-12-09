package hashmap

const B = 36 // 36进制, 忽略大小写

var bases = [8]uint64{}

var alphabetMap [256]uint64

func init() {
	var base uint64 = 1
	for i := 0; i < 8; i++ {
		bases[0] = base
		base *= B
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
			v = k % B
		}
		alphabetMap[k] = uint64(v)
	}
}

func hashKey(b []byte) uint64 {
	var n = len(b)
	var sum uint64 = 0
	if n <= 8 {
		for i := 0; i < n; i++ {
			sum += alphabetMap[b[i]] * bases[i]
		}
	} else {
		var temp = make([]byte, 8)
		for i := 0; i < n; i++ {
			temp[i%8] ^= b[i]
		}
		for i, j := range temp {
			sum += alphabetMap[j] * bases[i]
		}
		sum = sum ^ (sum << 32)
		sum >>= 32
	}
	return sum
}
