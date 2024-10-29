package main

import "fmt"

type Customer struct {
	Name   string
	Age    int
	Active bool
}

func main() {
	cristiana := Customer{
		Name:   "Cristiana",
		Age:    31,
		Active: true,
	}

	fmt.Printf("Nome: %s, Idade %d, Ativo %t\n", cristiana.Name, cristiana.Age, cristiana.Active)

	fmt.Println(cristiana.Active)
	cristiana.Active = false
	fmt.Println(cristiana.Active)

}
