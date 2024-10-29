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

func main() {
	cristiana := Customer{
		Name:   "Cristiana",
		Age:    31,
		Active: true,
	}

	cristiana.City = "São Paulo"
	cristiana.Address.State = "São Paulo"

	fmt.Println(cristiana.City)
	fmt.Println(cristiana.State)
	fmt.Println(cristiana.Address.State)
}
