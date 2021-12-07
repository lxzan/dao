package hashmap

const MaxU32 uint64 = 4294967295

func hashKey(b []byte) uint32 {
	var n = len(b)
	if n <= 4 {
		var sum uint32 = 0
		var base uint32 = 1
		for i := 0; i < n; i++ {
			sum += uint32(b[i]) * base
			base <<= 7
		}
		return sum
	} else if n <= 8 {
		var sum uint64 = 0
		var base uint64 = 1
		for i := 0; i < n; i++ {
			sum += uint64(b[i]) * base
			base <<= 7
		}
		if sum > MaxU32 {
			return uint32((sum ^ (sum >> 32)) & MaxU32)
		}
		var x = uint32(sum & MaxU32)
		return x ^ (x >> 16)
	}

	var indexes = make([]byte, 9)
	indexes[1] = 0
	indexes[2] = byte(n - 1)
	indexes[3] = (indexes[1] + indexes[2]) / 2
	indexes[4] = (indexes[1] + indexes[3]) / 2
	indexes[5] = (indexes[2] + indexes[3]) / 2
	indexes[6] = (indexes[1] + indexes[4]) / 2
	indexes[7] = (indexes[2] + indexes[5]) / 2
	indexes[8] = (indexes[3] + indexes[4]) / 2
	var sum uint64 = 0
	var base uint64 = 1
	for i := 1; i <= 8; i++ {
		sum += uint64(b[indexes[i]]) * base
		base <<= 7
	}
	if sum > MaxU32 {
		return uint32((sum ^ (sum >> 32)) & MaxU32)
	}
	var x = uint32(sum & MaxU32)
	return x ^ (x >> 16)
}
