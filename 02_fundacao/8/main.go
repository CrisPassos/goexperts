package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(sum(10, 40))
	fmt.Println(sum2(50, 2))
	fmt.Println(sum2(10, 2))

	valor, err := sumError(50, 3)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(valor)
}

// se tiver o mesmo tipo, pode omitir o tipo do segundo parâmetro
func sum(a, b int) int {
	return a + b
}

// é possivel retornar mais de um valor
func sum2(a, b int) (int, bool) {
	if a+b >= 50 {
		return a + b, true
	}

	return a + b, false
}

func sumError(a, b int) (int, error) {
	if a+b >= 50 {
		return 0, errors.New("A soma é maior que 50")
	}

	return a + b, nil
}
