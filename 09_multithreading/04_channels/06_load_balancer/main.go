package main

import (
	"fmt"
	"time"
)

// faz o processamento dos jobs
func worker(workerId int, data <-chan int) {
	for x := range data {
		fmt.Printf("Worker %d processando job %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}

func main() {
	data := make(chan int)
	qtdWorkers := 1000000

	//inicializa os workers
	for i := 0; i < qtdWorkers; i++ {
		go worker(i, data)
	}

	//envio informações para o canal
	for i := 0; i < 1000000; i++ {
		data <- i
	}

}
