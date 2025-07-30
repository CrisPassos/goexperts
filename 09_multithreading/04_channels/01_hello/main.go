package main

import "fmt"

// Thread 1
func main() {
	canal := make(chan string) // Canal vazio

	//Thread 2
	//eu pego a string ola mundo e preencho o canal
	go func() {
		canal <- "Olá Mundo" // Canal Cheio
	}()

	//Thread 1
	msg := <-canal // eu espero o canal me retornar uma string -- Canal esvaziou
	fmt.Println(msg)

	// Eu peguei o valor do canal 2 e coloquei no canal 1

}
