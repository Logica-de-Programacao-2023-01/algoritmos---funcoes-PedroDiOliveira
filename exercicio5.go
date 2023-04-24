package main

import (
	"fmt"
)

func main() {
	slice := []int{0, 5, 8, 20, 7}
	num := 5
	fmt.Print(igualdade(slice, num))
}

func igualdade(slice []int, num int) int {
	valor := -1
	for i := 0; i < len(slice); i++ {
		if slice[i] == num {
			valor = i
		}
	}
	return valor
}
