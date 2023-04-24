package main

import (
	"fmt"
	"unicode"
)

func main() {
	palavras := []string{"Olá", "Teste", "macacos"}
	fmt.Print(primeiramaiscula(palavras))
}
func primeiramaiscula(palavras []string) ([]string, error) {
	if palavras == nil {
		return palavras, fmt.Errorf("Erro ")
	}
	palavrasmaisculas := []string{}

	for i := 0; i < len(palavras); i++ {
		if unicode.IsUpper(rune(palavras[i][0])) {
			palavrasmaisculas = append(palavrasmaisculas, palavras[i])
		}
	}
	return palavrasmaisculas, nil
}
