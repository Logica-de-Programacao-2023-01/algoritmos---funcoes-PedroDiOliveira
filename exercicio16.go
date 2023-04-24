package main

import (
	"fmt"
	"strings"
)

func vogais(s string) (string, error) {
	if len(s) == 0 {
		return "", fmt.Errorf("deu erro")
	}
	vogais := "aeiouAEIOU"
	var a string
	for _, rs := range s {
		if strings.ContainsRune(vogais, rs) {
			a += "*"
		} else {
			a += string(rs)
		}
	}
	return a, nil

}
func main() {
	nome := "arara"
	resultado, erro := vogais(nome)
	if erro != nil {
		fmt.Print(erro)
	} else {
		fmt.Print(resultado)
	}
}
