package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

//func contains(f2 string) bool {
//
//}

func UnpackTEST(f string) (string, error) {
	nF := strings.Builder{}
	//if f == "" {
	//	return
	//}
	if unicode.IsDigit(rune(f[0])) {
		//nF += string(f[0])
		//f = f[1:]
		//nF = "некорректная строка"
		return "", errors.New("некорректная строка")
	}
	for i := 0; i < len(f)-1; i++ {
		if !unicode.IsDigit(rune(f[i])) && unicode.IsDigit(rune(f[i+1])) {
			if s, err := strconv.Atoi(string(f[i+1])); err == nil {
				//if s != 0 {
				//	for j := 0; j < s; j++ {
				//		//nF += string(f[i-1])
				//		nF.WriteRune(rune(f[i]))
				//	}
				//}
				nF.WriteString(strings.Repeat(string(f[i]), s))

			} else {
				//nF += string(f[i])
				nF.WriteRune(rune(f[i]))
			}
			i++
		} else {
			return "", errors.New("некорректная строка")
		}

	}
	if _, err := strconv.Atoi(string(f[len(f)-1])); err != nil {
		//fmt.Println(s2, err)
		nF.WriteRune(rune(f[len(f)-1]))
	}
	return nF.String(), nil
	//for i := 0; i < len(f)-1; i++ {
	//	if s, err := strconv.Atoi(string(f[i+1])); err == nil {
	//		for j := 0; j < s; j++ {
	//			nF.WriteRune(rune(f[i]))
	//			continue
	//		}
	//	} else {
	//		nF.WriteRune(rune(f[i]))
	//	}
	//}
	//return nF.String(), nil
}

func main() {
	//A := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var stroka string
	fmt.Println("Введите строку")
	fmt.Scan(&stroka)
	fmt.Println(UnpackTEST(stroka))
}
