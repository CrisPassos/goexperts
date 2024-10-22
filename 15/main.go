package main

// Ponteiros é muito comum o uso de ponteiros "funciona como um cache ou apontamento da memória, usado como escopo global"
func main() {
	//Memória -> Endereço -> Valor
	a := 10
	// variavel -> ponteiro -> ponteiro tem um valor
	var ponteiro *int = &a
	//endereço da memória Oxc000058738
	println(&a)
	println(ponteiro)

	*ponteiro = 20
	println(a)
	b := &a
	println(b)
	//direfference colocar o * na frente aponta para o valor no endereço &a
	println(*b)
}
