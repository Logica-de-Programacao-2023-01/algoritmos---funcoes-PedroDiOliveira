package main

import "fmt"

func somaResultado(slice []int, fn func(int) int) ([]int, error) {
	if len(slice) == 0 {
		return nil, fmt.Errorf("deu erro")
	}
	var soma int
	resultado := make([]int, len(slice))
	for i, rs := range slice {
		valor := fn(rs)
		soma += valor
		resultado[i] = valor
	}
	return resultado, nil
}
func main() {
	slice := []int{1, 2, 3, 4, 5}
	soma := func(num int) int { return num * 2 }
	resultado, erro := somaResultado(slice, soma)
	if erro != nil {
		fmt.Print(erro)
	} else {
		fmt.Print(resultado)
	}
}
