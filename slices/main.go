package main

import "fmt"

func main() {
	fmt.Println(concat(create()))
}

func concat(n1 [][]int) []int {
	arr := make([]int, 0)
	for _, v := range n1 {
		arr = append(arr, v...)
	}
	return arr
}

func create() [][]int {
	var el1 int
	var el2 int
	fmt.Print("Первое значение: ")
	fmt.Scan(&el1)
	fmt.Print("Второе значение: ")
	fmt.Scan(&el2)
	arr := make([][]int, 0)
	for i := 0; i < el1; i++ {
		arrIn := make([]int, 0)
		for j := 0; j < el2; j++ {
			var elIn int
			fmt.Scan(&elIn)
			arrIn = append(arrIn, elIn)
		}
		arr = append(arr, arrIn)
	}
	return arr
}
