package main

import "fmt"

// Closures ou funções anónimas, chamar uma func dentro de outra func
func main() {
	total := func() int {
		return sum(10, 20, 30, 40, 50, 60)
	}()
	fmt.Println(total)
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
