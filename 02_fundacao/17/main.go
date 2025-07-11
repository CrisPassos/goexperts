package main

import "fmt"

type Cliente struct {
	nome string
}

type Conta struct {
	saldo float64
}

func (c Cliente) andou() {
	// cópia não é o valor da memória
	c.nome = "Cristiana Passos"
	fmt.Printf("O cliente %v ando\n", c.nome)
}

// qualquer mudança que fizer, estou alterando diretamente o ponteiro da memoria
func (c *Conta) simular(valor float64) float64 {
	c.saldo += valor
	return c.saldo
}

// estou criando uma função que retorna o endereçamento da memmória da conta que esta criando
// a conta com o saldo zero
// qualquer local onde eu passar essa conta isso vai ser refletido em todos lugar e não em uma cópia
// muito usado
func NewConta() *Conta {
	return &Conta{saldo: 0}
}

func main() {
	cristiana := Cliente{
		nome: "Cristiana",
	}

	cristiana.andou()
	fmt.Printf("O valor da struct com nome %v\n", cristiana.nome)

	conta := Conta{
		saldo: 100.00,
	}

	conta.simular(200)
	fmt.Println(conta.saldo)
}
