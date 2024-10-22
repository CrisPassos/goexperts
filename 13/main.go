package main

import "fmt"

type Address struct {
	Name   string
	Number int
	City   string
	State  string
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

func main() {
	cristiana := Customer{
		Name:   "Cristiana",
		Age:    31,
		Active: true,
	}

	cristiana.Disable()
}
