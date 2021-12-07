package dao

import (
	"math/rand"
	"time"
)

type RandomString string

const (
	Alphabet RandomString = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Numeric  RandomString = "0123456789"
)

var Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func (c RandomString) Generate(n int) string {
	var b = make([]byte, n)
	var length = len(c)
	for i := 0; i < n; i++ {
		var idx = Rand.Intn(length)
		b[i] = c[idx]
	}
	return string(b)
}
