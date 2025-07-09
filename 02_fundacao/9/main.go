package main

import "fmt"

func main() {
	fmt.Println(sum(1, 4, 5, 22, 543, 765756))
}

// posso passar quantos parametros eu quiser desde que sejam inteiors
func sum(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}
