package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	cristiana := Cliente{
		Nome:  "Cristiana Passos",
		Idade: 35,
		Ativo: true,
	}

	fmt.Println("Nome:", cristiana.Nome)
	fmt.Println("Idade:", cristiana.Idade)
	fmt.Println("Ativo:", cristiana.Ativo)

	cristiana.Ativo = false
	fmt.Println("Ativo:", cristiana.Ativo)

}
