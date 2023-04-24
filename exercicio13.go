package main

import (
	"fmt"
	"strconv"
)

func main() {
	numero := 5490
	fmt.Print(somaalgarismos(numero))
}

func somaalgarismos(numero int) (int, error) {

	if numero < 0 {
		return 0, fmt.Errorf("Erro ")
	}

	numString := strconv.Itoa(numero)
	soma := 0

	StringNum := 0

	for i := 0; i < len(numString); i++ {
		StringNum, _ = (strconv.Atoi(string(numString[i])))
		soma += StringNum
	}
	return soma, nil
}
