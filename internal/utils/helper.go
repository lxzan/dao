package utils

import (
	"math/rand"
	"sort"
	"time"
)

type RandomString string

var Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

const (
	Alphabet RandomString = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Numeric  RandomString = "0123456789"
)

func (c RandomString) Generate(n int) string {
	var b = make([]byte, n)
	var length = len(c)
	for i := 0; i < n; i++ {
		var idx = Rand.Intn(length)
		b[i] = c[idx]
	}
	return string(b)
}

func SameInts(arr1, arr2 []int) bool {
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

func SameStrings(arr1, arr2 []string) bool {
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
