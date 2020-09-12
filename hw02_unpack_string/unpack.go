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
	if len(str) == 0 {
		return "", nil
	}
	if unicode.IsDigit(rune(str[0])) {
		return "", ErrInvalidString
	}
	for i := 0; i < len(str); i++ {
		if i+1 < len(str) && unicode.IsDigit(rune(str[i+1])) {
			if !unicode.IsDigit(rune(str[i+2])) {
				dig, _ := strconv.Atoi(string(str[i+1]))
				strb.WriteString(strings.Repeat(string(str[i]), dig))
			} else {
				return "", ErrInvalidString
			}
		} else if !unicode.IsDigit(rune(str[i])) {
			strb.WriteString(string(str[i]))
		}
		//if i == len(str)-1 && !unicode.IsDigit(rune(str[i])) {
		//	strb.WriteString(string(str[i]))
		//}
	}
	return strb.String(), nil
}
