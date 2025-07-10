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

// Desativar é um método que desativa o cliente, muda a propriedade do Structs
// usamos (c Cliente) para indicar que é um método do tipo Cliente
func (c Cliente) Desativar() {
	c.Ativo = false
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

	cristiana.Desativar()
	fmt.Println("Ativo após desativar:", cristiana.Ativo)

}
