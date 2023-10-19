package main

import "fmt"

func itoa(i int) (s string) {

	return s
}

func main() {
	type pair struct {
		i int
		s string
	}
	test := []pair{
		{0, "0"},
		{22, "22"},
		{03245, "03245"},
		{-3, "-3"},
	}
	for _, t := range test {
		s := itoa(t.i)
		if t.s == s {
			fmt.Printf("%-10d - %s\n", t.i, "OK")
		} else {
			fmt.Printf("%-10d - %s (got %q)\n", t.i, "FAIL", s)

		}
	}
}
