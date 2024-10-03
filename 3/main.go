package main

import "fmt"

// tipos de um dado especifico
type ID int

// No Go, não existe um conceito de "objeto" como em linguagens orientadas a objetos
// tradicionais, mas você pode criar tipos que encapsulam dados
// e comportamentos usando structs. Para definir um tipo que representa um "objeto",
// você pode usar uma struct
type Pessoa struct {
	Nome  string
	Idade int
}

const a ID = 1

// No Go, você não pode usar diretamente fmt.Println(p)
// para imprimir uma struct de forma detalhada,
// mas pode usar fmt.Printf com o verbo %+v para imprimir
// os campos da struct com seus nomes.
func main() {
	println(a)

	//Criando uma instancia de Pessoa
	p := Pessoa{
		Nome:  "Maria",
		Idade: 45,
	}

	println(p.Nome)
	println(p.Idade)

	//
	fmt.Printf("%+v\n", p)

}
