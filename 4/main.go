package main

// biblioteca padrão
import "fmt"

type ID int

var (
	a bool    = true
	c int     = 10
	d string  = "Wesley"
	e float32 = 1.2
	f ID      = 1
)

func main() {
	// concatenação de strings, o FMT é uma biblioteca de formatação da própria linguagem GO
	fmt.Printf("O tipo de E é %T", e)
	fmt.Printf("O tipo de E é %v", e)

	fmt.Printf("O tipo de E é %T", f)
}
