package main

import (
	"strings"
	"sync"
)

var (
	urlStore = make(map[string]string)
	mu       = &sync.RWMutex{}
	counter  = 10000000 // Start with a larger number for longer keys
)

const base62Chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// toBase62 converts an integer to a base62 string.
func toBase62(num int) string {
	if num == 0 {
		return string(base62Chars[0])
	}

	var result strings.Builder
	base := len(base62Chars)

	for num > 0 {
		remainder := num % base
		result.WriteByte(base62Chars[remainder])
		num = num / base
	}

	// The result is reversed, so we need to reverse it back.
	return reverse(result.String())
}

// reverse reverses a string.
func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
