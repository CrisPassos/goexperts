package main

import "fmt"

func main() {
	//quando declaramos um array em Go não podemos acrescentar mais posições após a sua declaração
	var meuArray [3]int
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30

	fmt.Println(meuArray)
	fmt.Println(meuArray[1])
	fmt.Println(meuArray[len(meuArray)-1])

	fmt.Println("-----FOR-----")
	for i, v := range meuArray {
		//%d é usado para dígito
		fmt.Printf("O indice é %d e o valor é %d\n", i, v)
	}
}
