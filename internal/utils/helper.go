package utils

import (
	"math/rand"
	"reflect"
	"sort"
	"time"
	"unsafe"
)

type RandomString string

var Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

const (
	Alphabet RandomString = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Numeric  RandomString = "0123456789"
)

func (this RandomString) Generate(n int) string {
	var b = make([]byte, n)
	var length = len(this)
	for i := 0; i < n; i++ {
		var idx = rand.Intn(length)
		b[i] = this[idx]
	}
	return string(b)
}

func SameInts(arr1, arr2 []int) bool {
	sort.Ints(arr1)
	sort.Ints(arr2)
	var n = len(arr1)
	if n != len(arr2) {
		return false
	}
	for i := 0; i < n; i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
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

func UniqueString(arr []string) []string {
	sort.Strings(arr)
	var n = len(arr)
	var b = make([]string, 0, n)
	if n > 0 {
		b = append(b, arr[0])
	}
	for i := 1; i < n; i++ {
		if arr[i] != arr[i-1] {
			b = append(b, arr[i])
		}
	}
	return b
}

func ReverseStrings(arr []string) {
	var n = len(arr)
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
	}
}

func S2B(s string) (b []byte) {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data = sh.Data
	bh.Cap = sh.Len
	bh.Len = sh.Len
	return b
}

func XOR64(v uint64) uint32 {
	return uint32((v >> 32) ^ (v << 32 >> 32))
}

type Integer interface {
	~int64 | ~int | ~int32 | ~int16 | ~int8 | ~uint64 | ~uint | ~uint32 | ~uint16 | ~uint8
}

func IsBinaryNumber[T Integer](x T) bool {
	return x&(x-1) == 0
}
