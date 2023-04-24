package main

import (
	"fmt"
)

func contemNoSlice(slice []int, n int) (bool, error) {
	if len(slice) == 0 {
		return false, fmt.Errorf("deu erro")
	}
	for _, rs := range slice {
		if n == rs {
			return true, nil
		}
	}
	return false, nil
}
func main() {
	slice := []int{1, 2, 3, 4, 5}
	var num int
	fmt.Println("digite um número:")
	fmt.Scan(&num)
	resultado, erro := contemNoSlice(slice, num)
	if erro != nil {
		fmt.Print(erro)

	} else {
		fmt.Print(resultado)
	}
}
