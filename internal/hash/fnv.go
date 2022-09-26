package hash

import "encoding/binary"

const (
	// FNV-1
	offset64 = uint64(14695981039346656037)
	prime64  = uint64(1099511628211)

	// Init64 is what 64 bits hash values should be initialized with.
	Init64 = offset64
)

// HashBytes64 returns the hash of u.
func HashBytes64(b []byte) uint64 {
	return AddBytes64(Init64, b)
}

// HashUint64 returns the hash of u.
func HashUint64(u uint64) uint64 {
	return AddUint64(Init64, u)
}

// AddBytes64 adds the hash of b to the precomputed hash value h.
func AddBytes64(h uint64, b []byte) uint64 {
	for len(b) >= 8 {
		h = (h * prime64) ^ binary.LittleEndian.Uint64(b[:8])
		b = b[8:]
	}

	if len(b) >= 4 {
		h = (h * prime64) ^ uint64(binary.LittleEndian.Uint32(b[:4]))
		b = b[4:]
	}

	for _, code := range b {
		h = (h * prime64) ^ uint64(code)
	}

	return h
}

// AddUint64 adds the hash value of the 8 bytes of u to h.
func AddUint64(h uint64, u uint64) uint64 {
	h = (h * prime64) ^ ((u >> 56) & 0xFF)
	h = (h * prime64) ^ ((u >> 48) & 0xFF)
	h = (h * prime64) ^ ((u >> 40) & 0xFF)
	h = (h * prime64) ^ ((u >> 32) & 0xFF)
	h = (h * prime64) ^ ((u >> 24) & 0xFF)
	h = (h * prime64) ^ ((u >> 16) & 0xFF)
	h = (h * prime64) ^ ((u >> 8) & 0xFF)
	h = (h * prime64) ^ ((u >> 0) & 0xFF)
	return h
}
