package main

func main() {
	a := 10
	// mostra o valor
	println(a)

	// mostra a alocação da memória
	println(&a)

	// variavel -> ponteiro que tem um endereco em memoria, que tem um valor
	var ponteiro *int = &a
	println(ponteiro)

	// a esta apontado para o mesmo endereço de memoria que o ponteiro
	*ponteiro = 20
	println(a)

	b := &a
	println(b)
	//dereference, mostra o valor que está guardado na memória
	println(*b)

	*b = 30
	println(*b)
	println(a)
}
