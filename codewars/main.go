package main

//"abcde" -> 0 # no characters repeats more than once
//"aabbcde" -> 2 # 'a' and 'b'
//"aabBcde" -> 2 # 'a' occurs twice and 'b' twice (`b` and `B`)
//"indivisibility" -> 1 # 'i' occurs six times
//"Indivisibilities" -> 2 # 'i' occurs seven times and 's' occurs twice
//"aA11" -> 2 # 'a' and '1'
//"ABBA" -> 2 # 'A' and 'B' each occur twice

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Print(duplicate_count(s))
}

func duplicate_count(s1 string) int {
	var k1 int
	s1 = strings.ToLower(s1)
	var S []string
	for _, v := range s1 {
		k := strings.Count(s1, string(v))
		if k > 1 && !Contains(S, string(v)) {
			k1++
			S = append(S, string(v))

		}

	}
	return k1
}

func Contains[S ~[]E, E comparable](s S, v E) bool {
	return Index(s, v) >= 0
}

func Index[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}
