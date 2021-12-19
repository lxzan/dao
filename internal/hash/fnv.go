package hash

const (
	offset32 = 2166136261
	prime32  = 16777619
)

// NewFnv32 returns a new 32-bit FNV-1 hash.Hash.
// Its Sum method will lay the value out in big-endian byte order.
func NewFnv32(data []byte) uint32 {
	var hash uint32 = offset32
	for _, c := range data {
		hash *= prime32
		hash ^= uint32(c)
	}
	return hash
}
