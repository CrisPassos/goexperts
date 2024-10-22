package main

import "fmt"

func main() {
	//declaração de um slice (por de trás dos panos é executado um array)
	s := []int{10, 20, 30, 40, 50, 60, 100}

	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	//o : faz um corte ou um slide, para direita ou para esquerda
	//nesse caso ele fez o corte para a esquerda, por isso o resultado é vazio
	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])
	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4])

	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])

	//aumentar a capacidade do slice
	s = append(s, 100)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
