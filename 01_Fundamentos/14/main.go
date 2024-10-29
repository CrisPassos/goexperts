package main

import "fmt"

type Address struct {
	Name   string
	Number int
	City   string
	State  string
}

// a interface do GO só permite assinatura de métodos e somente isso, não é possível settar propriedades
type Person interface {
	Disable()
}

type Customer struct {
	Name   string
	Age    int
	Active bool
	Address
}

func (c Customer) Disable() {
	c.Active = false
	fmt.Printf("O cliente %s foi desativado\n", c.Name)
}

func Disabled(person Person) {
	person.Disable()
}

func main() {
	cristiana := Customer{
		Name:   "Cristiana",
		Age:    31,
		Active: true,
	}

	Disabled(cristiana)
}
