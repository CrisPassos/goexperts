package main

import (
	"fmt"

	"github.com/CrisPassos/goexperts/tree/main/07_packaging/01_introduzindo_go_mod/math"
)

func main() {
	fmt.Println("Hello, World")
	m := math.NewMath(10, 20)
	fmt.Println(m.Add())
}
