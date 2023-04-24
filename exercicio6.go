package main

import (
	"fmt"
	"strings"
)

func main() {
	palavras := []string{"Olá", ", ", "tenha um bom-dia"}
	fmt.Print(valorigual(palavras))
}

func valorigual(palavras []string) (string, error) {
	var juntador string
	if palavras == nil {
		return "", fmt.Errorf("Erro ")
	}
	for _, passe := range palavras {
		juntador += string(passe)
	}
	return strings.ReplaceAll(juntador, " ", ","), nil
}
