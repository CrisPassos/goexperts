package main

const a = "Hello World"

// declando a variavel do tipo boleano
// por default ela inferiu o valor false para essa bool
// ao declarar uma varíavel no GO ele infere automaticamente um valor para cada tipo
// essa declaração é do escopo global
var (
	b bool
	c int
	d string
	e float64
	g string = "Cristiana"
)

func main() {
	//declaração escopo local
	var f string = "Escopo Local"
	println(f)

	println("--------")
	b = true
	println(b)
	println(c)
	println(d)
	println(e)

	println("--------")
	println(g)

	println("--------")
	//short declaration
	h := "tipo string"
	i := false
	println(h)
	println(i)
}
