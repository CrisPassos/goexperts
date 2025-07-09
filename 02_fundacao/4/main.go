package main

import "fmt"

type ID int

var (
	a bool    = true
	c int     = 42
	d string  = "Cristiana Passos"
	e float64 = 3.14
	f ID      = 1
)

func main() {
	// PrintF formatted values
	fmt.Printf("The type of E is: %T", e)
	fmt.Printf("The value of E is: %v", e)
}
