package main

// O Go só trabalha com o FOR como loop
func main() {

	// for mais simples
	for i := 0; i < 10; i++ {
		println(i)
	}

	//range
	numeros := []string{"um", "dois", "tres"}

	for i, v := range numeros {
		println(i, v)
	}

	// percorrer somente
	i := 10

	for i < 20 {
		println(i)
		i++
	}

	// loop infinito
	for {
		println("Hello World")
	}

}
