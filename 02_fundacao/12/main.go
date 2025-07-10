package main

import "fmt"

// composiçao de structs, tipo herança
type Endereco struct {
	Cidade string
	Estado string
	Cep    string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
	// criei uma propriedade do tipo Endereco
	Address Endereco
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

	cristiana.Endereco.Cidade = "São Paulo"
	fmt.Println("Cidade:", cristiana.Endereco.Cidade)

	cristiana.Address.Cidade = "Rio de Janeiro"
	fmt.Println("Cidade Address:", cristiana.Address.Cidade)

}
