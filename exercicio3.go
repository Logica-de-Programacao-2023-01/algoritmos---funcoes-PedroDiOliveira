package main

import "fmt"

func main() {
	palavras := []string{"Olá", ", ", "tenha um bom-dia"}
	fmt.Print(concatenador(palavras))
}
func concatenador(palavras []string) string {
	var juntador string
	for _, passe := range palavras {
		juntador += string(passe)
	}
	return juntador
}
