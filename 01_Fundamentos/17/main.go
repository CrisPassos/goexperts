package main

import "fmt"

type Customer struct {
	Name string
}

func (c Customer) walk() {
	fmt.Printf("O cliente %v andou \n", c.Name)
}

type Account struct {
	Balance int
}

func (c *Account) simulation(value int) int {
	c.Balance += value
	println(c.Balance)
	return c.Balance
}

func main() {
	cristiana := Customer{
		Name: "Cristiana",
	}

	cristiana.walk()

	account := Account{Balance: 100}
	account.simulation(200)
	println(account.Balance)
}
