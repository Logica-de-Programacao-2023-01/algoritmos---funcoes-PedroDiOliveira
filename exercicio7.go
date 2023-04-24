package main

import "fmt"

func main() {
	conj := []int{2, 10, 9, 8, 5490}
	fmt.Println(conta(conj, multiplique1000))

}

func multiplique1000(x int) int {

	return 1000 * x
}

func conta(conj []int, multiplique1000 func(int) int) ([]int, error) {
	for i := 0; i < len(conj); i++ {
		conj[i] = multiplique1000(conj[i])
	}
	if len(conj) == 0 {
		return conj, fmt.Errorf("Erro")
	}
	return conj, nil
}
