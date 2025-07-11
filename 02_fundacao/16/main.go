package main

func sum(a, b int) int {
	a = 50
	return a + b
}

// como estou trabalhando com ponteiros, o valor é mudado independete do lugar que eu estiver
func sumPointer(a, b *int) int {
	*a = 50
	return *a + *b
}

func main() {
	var1 := 10
	var2 := 20
	// o valor aqui foi passado por referência por isso que a dentro do metodo é ignorado
	sum(var1, var2)
	println(var1)

	sumPointer(&var1, &var2)
	println(var1)

}
