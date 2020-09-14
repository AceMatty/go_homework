package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var strb strings.Builder
	runes := []rune(str)
	if len(runes) == 0 {
		return "", nil
	}
	if unicode.IsDigit(runes[0]) {
		return "", ErrInvalidString
	}
	for i := 0; i < len(runes); i++ {
		if i+1 < len(runes) && !unicode.IsDigit(runes[i]) {
			if unicode.IsDigit(runes[i+1]) {
				dig, _ := strconv.Atoi(string(runes[i+1]))
				strb.WriteString(strings.Repeat(string(runes[i]), dig))
			} else {
				strb.WriteString(string(runes[i]))
			}
		}
		if i == len(runes)-1 && !unicode.IsDigit(runes[i]) {
			strb.WriteString(string(runes[i]))
		}
		if i+1 < len(runes) && unicode.IsDigit(runes[i]) {
			if unicode.IsDigit(runes[i+1]) {
				return "", ErrInvalidString
			}
		}
	}
	return strb.String(), nil
}
