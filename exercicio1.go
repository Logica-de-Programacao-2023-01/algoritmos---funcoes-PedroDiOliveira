package main

import "fmt"

func media(slice []int) int {
	resultado := 0
	for _, numeros := range slice {
		resultado += numeros
	}
	return resultado / 3

}
func main() {
	slice := []int{4, 5, 6}
	resultado := media(slice)
	fmt.Println(resultado)
}
