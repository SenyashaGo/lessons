package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(tst(s))
}

func tst(f string) (string, error) {
	nF := strings.Builder{}
	if f == "" {
		return nF.String(), nil
	}
	if unicode.IsDigit(rune(f[0])) {
		return "", ErrInvalidString
	}
	for i := 0; i < len(f)-1; i++ {
		if unicode.IsDigit(rune(f[i])) && unicode.IsDigit(rune(f[i+1])) {
			return "", ErrInvalidString
		} else if !unicode.IsDigit(rune(f[i])) && unicode.IsDigit(rune(f[i+1])) {
			if s, err := strconv.Atoi(string(f[i+1])); err == nil {
				nF.WriteString(strings.Repeat(string(f[i]), s))
				continue
			}
			//else {
			//	nF.WriteRune(rune(f[i]))
			//}
		} else if !unicode.IsDigit(rune(f[i])) {
			nF.WriteRune(rune(f[i]))
		}
	}
	if _, err := strconv.Atoi(string(f[len(f)-1])); err != nil {
		nF.WriteRune(rune(f[len(f)-1]))
	}
	return nF.String(), nil
}
