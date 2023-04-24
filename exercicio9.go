package main

import "fmt"
import "strings"

func concatpalavras(palavra string) ([]string, error) {
	if len(palavra) == 0 {
		return nil, fmt.Errorf("deu erro")

	}
	strings := strings.Fields(palavra)
	return strings, nil

}
func main() {

	palavra, erro := concatpalavras("oi rs ola")
	if erro != nil {
		fmt.Print(erro)
	} else {
		fmt.Print(palavra)
	}
}
