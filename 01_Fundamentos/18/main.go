package main

import "fmt"

// interface vazia: aceita qualquer coisa(type) e implementa todo mundo, uso moderado
// lembra um pouco any do TS
func main() {
	var x interface{} = 10
	var y interface{} = "Hello World"

	showType(x)
	showType(y)

	printValue("Cristiana") //imprime uma string
	printValue(1000)        //imprime um número
	printValue(false)       //imprime um bool
	printValue([]int{1, 2}) //imprime um slice

}

func showType(t interface{}) {
	fmt.Printf("O tipo da varíavel é %T e o valor é %v\n", t, t)
}

func printValue(value interface{}) {
	fmt.Println(value)
}
