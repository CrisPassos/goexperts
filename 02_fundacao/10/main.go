package main

import "fmt"

func main() {
	total := func() int {
		return sum(1, 2, 4, 5345, 34534) * 2
	}()

	fmt.Println(total)
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
