package main

// Uma constante é um valor que não pode ser alterado
const a = "Hello, World!"

// declaração de escopo global
var (
	// declarando uma variável do tipo boolean
	// por default, o valor de uma variável do tipo boolean é false
	b bool
	// por default, o valor de uma variável do tipo int é 0
	c int
	// por default, o valor de uma variável do tipo string é uma string vazia
	d string = "Cristiana Passos"
	// por default, o valor de uma variável do tipo float64 é 0.0
	e float64
)

func main() {
	b = true
	println(a)
	println(b)
	println(c)
	println(e)

	//declaração de escopo local
	var f string
	println(f)

	// short hand declaration, usa o :=
	// := só pode ser usado na primeira vez que a variável é declarada
	// se a variável já foi declarada, deve-se usar o operador de atribuição =
	g := "Criei, declarei e atribui a varíavel: Hello, Go!"
	println(g)
}

func x() {
}
