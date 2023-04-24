package main

import (
	"fmt"
)

func main() {
	palavra := []string{"macacos", "me", "mordam"}
	fmt.Println(palavrasmaisque5(palavra))
}

func palavrasmaisque5(palavra []string) ([]string, error) {
	palavrasmaisque5 := []string{}

	if len(palavra) == 0 {
		return palavrasmaisque5, fmt.Errorf("ERRO ")
	}

	for _, rangervermelho := range palavra {
		if len(rangervermelho) > 5 {
			palavrasmaisque5 = append(palavrasmaisque5, rangervermelho)
		}
	}

	return palavrasmaisque5, nil
}
