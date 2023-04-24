package main

import (
	"fmt"
)

func main() {
	slice := []int{10, 5, 8, 20, 7}
	fmt.Print(ordemcrescente(slice))
}
func ordemcrescente(slice []int) ([]int, error) {

	if len(slice) == 0 {
		return slice, nil
	}

	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice); j++ {

			if slice[i] < slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	return slice, nil
}
