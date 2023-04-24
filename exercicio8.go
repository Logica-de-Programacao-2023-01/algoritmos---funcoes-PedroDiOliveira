package main

import (
	"fmt"
	"strings"
)

func main() {
	palavra := "meu nome e pedro"
	fmt.Println(palavrasslice(palavra))
}

func palavrasslice(palavra string) ([]string, error) {
	palavras := []string{}
	if palavra == "" {
		return palavras, fmt.Errorf("ERRO ")
	}

	separador := strings.Split(palavra, " ")

	return separador, nil
}
