package main

import (
	"github.com/CrisPassos/goexperts/tree/main/07_packaging/03/math"
	"github.com/google/uuid"
)

func main() {
	m := math.NewMath(1, 4)

	println(m.Add())

	println(uuid.New().String())
}
