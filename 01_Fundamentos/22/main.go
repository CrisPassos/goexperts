package main

// tipos de for
func main() {
	for i := 0; i < 10; i++ {
		println(i)
	}

	slices := []string{"um", "dois", "três"}
	for indice, value := range slices {
		println(indice, value)
	}

	i := 0
	for i < 3 {
		println(i)
		i++
	}

	//loop infinito
	for {
		println("Infinity")
	}
}
