package main

//trabalhar com workspace
//para criar basta usar os comandos
//go mod tidy -e : ignora os pacotes não publicados que é o caso do /3/math
import (
	"fmt"

	"github.com/CrisPassos/goexpert/5_Packaging/3/math"
	"github.com/google/uuid"
)

func main() {
	m := math.Math{A: 2, B: 3}
	fmt.Println(m.Add())
	fmt.Println(uuid.New().String())
}
