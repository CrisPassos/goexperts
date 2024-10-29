package main

import (
	"curso-go/mathematic"
	"fmt"

	"github.com/google/uuid"
)

func main() {
	s := mathematic.Sum(10, 20)
	fmt.Println("Resultado: ", s)

	fmt.Println(uuid.New())
}
