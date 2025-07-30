package main

import "fmt"

// Thread 1
func main() {
	forever := make(chan bool) // Canal vazio

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("Olá Mundo %d\n", i)
		}
		// So resolvemos o deadlock pq preenchemos o canal
		// só funciona dentro da goroutine
		forever <- true
	}()

	<-forever // Aqui o Canal fica aberto esperando o canal ser cheio, teremos um deadlock

}
