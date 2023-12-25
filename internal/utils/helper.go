package utils

import (
	"math/rand"
	"sync"
	"time"
)

type RandomString struct {
	mu   sync.Mutex
	rand *rand.Rand
	dict string
}

var Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

var (
	Alphabet = (&RandomString{dict: "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"}).init()
	Numeric  = (&RandomString{dict: "0123456789"}).init()
)

func (c *RandomString) init() *RandomString {
	c.rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	return c
}

func (c *RandomString) Generate(n int) string {
	c.mu.Lock()
	defer c.mu.Unlock()

	var b = make([]byte, n)
	var length = len(c.dict)
	for i := 0; i < n; i++ {
		var idx = c.rand.Intn(length)
		b[i] = c.dict[idx]
	}
	return string(b)
}

func (c *RandomString) Intn(n int) int {
	c.mu.Lock()
	v := c.rand.Intn(n)
	c.mu.Unlock()
	return v
}

func IsSameSlice[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func ReverseStrings(arr []string) {
	var n = len(arr)
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
	}
}

type Integer interface {
	~int64 | ~int | ~int32 | ~int16 | ~int8 | ~uint64 | ~uint | ~uint32 | ~uint16 | ~uint8
}

func IsBinaryNumber[T Integer](x T) bool {
	return x&(x-1) == 0
}

// GetBinaryExponential 获取指数
func GetBinaryExponential(n int) int {
	sum := 0
	for n > 1 {
		n >>= 1
		sum++
	}
	return sum
}

func Clone[S ~[]E, E any](s S) S {
	if s == nil {
		return nil
	}
	return append(S([]E{}), s...)
}
