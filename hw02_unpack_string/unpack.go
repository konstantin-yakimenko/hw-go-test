package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	reverseStr := reverse(input)
	countSymbol := 1
	isPrevDigit := false
	var b strings.Builder
	for _, sym := range reverseStr {
		if unicode.IsDigit(sym) {
			if isPrevDigit {
				return "", ErrInvalidString
			}
			countSymbol = int(sym - '0')
			isPrevDigit = true
		} else {
			b.WriteString(strings.Repeat(string(sym), countSymbol))
			countSymbol = 1
			isPrevDigit = false
		}
	}
	if isPrevDigit {
		return "", ErrInvalidString
	}
	return reverse(b.String()), nil
}

func reverse(input string) string {
	r := []rune(input)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
