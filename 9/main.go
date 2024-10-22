package main

import "fmt"

// Funções variadicas
func main() {
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7))
	fmt.Println(sum(10, 20, 30, 40, 50, 60))
}

// essa função permite a passagem de parametros sem saber quantos serão BIGn
// por de trás dos panos é um hashmap
func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}
