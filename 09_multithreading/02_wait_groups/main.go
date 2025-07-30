package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("Task is running %s: %d\n", name, i)
		time.Sleep(1 * time.Second)
		wg.Done()
	}
}

// Thread 1
func main() {
	waitGroup := &sync.WaitGroup{}
	//sabemos que vamos executar 25 operações, então adicionamos 25 ao waitGroup
	waitGroup.Add(25)
	//Thread 2
	//trabalhamos simultaneamente com goroutines
	go task("A", waitGroup)

	//Thread 3
	go task("B", waitGroup)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("Task is ANONIMO: %d\n", i)
			time.Sleep(1 * time.Second)
			waitGroup.Done()
		}
	}()

	// ao todo sabemos que vamos executar 25 operações para o sistema sair
	waitGroup.Wait()
}
