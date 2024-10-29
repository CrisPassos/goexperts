package main

import (
	"errors"
	"fmt"
)

func main() {
	valor, err := sumError(53, 10)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(valor)

	fmt.Println("-------")

	fmt.Println(sum(1, 2))
	fmt.Println(sum2(3, 5))
	fmt.Println(sum2(51, 4))
}

// se os parametros são do mesmo tipo eu posso passar os parametros assim
func sum(a, b int) int {
	return a + b
}

// é possível retornar mais de um valor
func sum2(a, b int) (int, bool) {
	if a+b >= 50 {
		return a + b, true
	}

	return a + b, false
}

// é possível retornar mais de um valor exemplo error
func sumError(a, b int) (int, error) {
	if a+b >= 50 {
		return 0, errors.New("A Soma é maior que 50")
	}

	return a + b, nil
}
