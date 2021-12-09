package hashmap

const B = 62

var bases = [8]uint64{1, 62, 3844, 238328, 14776336, 916132832, 56800235584, 3521614606208}

var alphabetMap [256]uint64

func init() {
	for i := 0; i < 256; i++ {
		var k = uint8(i)
		var v uint8 = 0
		if k >= '0' && k <= '9' {
			v = k - '0'
		} else if k >= 'A' && k <= 'Z' {
			v = k - 'A' + 10
		} else if k >= 'a' && k <= 'z' {
			v = k - 'a' + 36
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
		var indexes = make([]byte, 8)
		indexes[0] = 0
		indexes[1] = byte(n - 1)
		indexes[2] = (indexes[0] + indexes[1]) / 2
		indexes[3] = (indexes[0] + indexes[2]) / 2
		indexes[4] = (indexes[1] + indexes[2]) / 2
		indexes[5] = (indexes[0] + indexes[3]) / 2
		indexes[6] = (indexes[1] + indexes[4]) / 2
		indexes[7] = (indexes[2] + indexes[3]) / 2

		for i, j := range indexes {
			sum += alphabetMap[b[j]] * bases[i]
		}
	}
	return sum
}
