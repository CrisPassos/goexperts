package main

import "fmt"

func main() {
	// esse array vai ter 5 posições fixas
	var myArray [5]int
	myArray[0] = 10
	myArray[1] = 20
	myArray[2] = 30
	myArray[3] = 40
	myArray[4] = 50

	fmt.Println(len(myArray))     //tamanho do array
	fmt.Println(len(myArray) - 1) //ultima posição

	fmt.Println(myArray[len(myArray)-1]) //acessando o último elemento do array

	for i, v := range myArray {
		fmt.Printf("The value of index %d é %d\n", i, v)
	}
}
