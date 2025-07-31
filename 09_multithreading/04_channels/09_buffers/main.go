package main

func main() {
	// Indica que vou ter 2 buffers no canal
	ch := make(chan string, 2)
	ch <- "Hello"
	ch <- "World"

	println(<-ch)
	println(<-ch)
}
