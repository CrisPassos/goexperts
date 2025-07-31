package main

import "fmt"

// Thread 1
func main() {
	hello := make(chan string)
	go receiver("Olá Mundo", hello)
	read(hello)
}

// <- apenas recebe informações, publica infomações no canal
// send only channel
func receiver(name string, hello chan<- string) {
	hello <- name
}

// <- do lado esquerdo significa que o canal é de esvaziar
// send only channel
func read(data <-chan string) {
	fmt.Println(<-data)
}
