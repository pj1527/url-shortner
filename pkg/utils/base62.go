package utils

import (
	"math"
	"strings"
)

const (
	base62Chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	base        = 62
)

func ToBase62(num uint64) string {
	if num == 0 {
		return string(base62Chars[0])
	}

	var result strings.Builder

	for num > 0 {
		remainder := num % base
		result.WriteByte(base62Chars[remainder])
		num = num / base
	}

	return reverse(result.String())
}

func ToInteger(str string) uint64 {
	var num uint64
	for i, char := range str {
		pos := strings.IndexRune(base62Chars, char)
		if pos == -1 {
			return 0
		}
		power := len(str) - 1 - i
		num += uint64(pos) * uint64(math.Pow(base, float64(power)))
	}
	return num
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
