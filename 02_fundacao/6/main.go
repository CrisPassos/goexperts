package main

import "fmt"

func main() {
	// Slice with initial values, [] em branco significa um slice
	s := []int{2, 4, 6, 8, 10}

	//len(s) retorna o tamanho do slice
	//cap(s) retorna a capacidade de sustentar valores
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	// s[:0] tudo que estiver depois do indice 0 (:0) é considerado vazio(removido)
	// apesar da remoção a capacidade do slice continua a mesma
	// : é um ponte de corte
	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])

	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4])

	// s[2:] tudo que estiver antes do indice 2 (2:) é considerado removido
	// apesar da remoção a capacidade do slice continua a mesma
	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])

	// a capacidade do slice pode ser aumentada com append
	// toda vez que o slice é aumentado, a capacidade é dobrada
	s = append(s, 12, 14, 16)
	fmt.Printf("len=%d cap=%d %v\n", len(s[:2]), cap(s[:2]), s[:2])

	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
