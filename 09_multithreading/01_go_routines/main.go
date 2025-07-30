package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("Task is running %s: %d\n", name, i)
		time.Sleep(1 * time.Second)
	}
}

// Thread 1
func main() {
	//Thread 2
	//trabalhamos simultaneamente com goroutines
	go task("A")

	//Thread 3
	go task("B")

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("Task is ANONIMO: %d\n", i)
			time.Sleep(1 * time.Second)
		}
	}()

	//Thread 4
	// Wait for a while to let goroutines finish(solução paleativa)
	time.Sleep(12 * time.Second)
	fmt.Println("All tasks completed.")

	//fim
}
