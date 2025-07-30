package main

import "fmt"

// Thread 1
func main() {
	ch := make(chan int)
	go publish(ch)
	//vai segurar o canal aberto
	reader(ch)
}

func reader(ch chan int) {
	for x := range ch {
		fmt.Printf("Recebido: %d\n", x)
	}
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	//evita os deadlocks
	close(ch) // Fecha o canal após enviar todos os valores
}
