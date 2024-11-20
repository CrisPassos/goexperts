package main

import (
	"fmt"

	"github.com/CrisPassos/goexpert/5_Packaging/1/math"
)

func main() {
	m := math.Math{A: 2, B: 3}
	fmt.Println(m.Add())
}
