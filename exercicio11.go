package main

import (
	"fmt"
)

func main() {
	numero := 7
	fmt.Print(primo(numero))
}
func primo(numero int) (bool, error) {
	primo := true
	if numero < 0 {
		return false, fmt.Errorf("Erro")
	}

	if numero < 2 {
		primo = false
	}

	for i := 2; i < numero; i++ {
		if numero%i == 0 {
			primo = false
		}
	}
	return primo, nil

}
