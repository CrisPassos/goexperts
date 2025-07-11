package main

import "fmt"

func main() {
	var minhaVar interface{} = "Cristiana Passos"
	//converter para string
	println(minhaVar.(string))
	//converter para int
	// temos um resultado + um statos (ok ou panic)
	res, ok := minhaVar.(int)
	fmt.Printf("O valor de res é %v e o resultado de ok é %v\n", res, ok)

	//aqui geramos um panic
	res2 := minhaVar.(int)
	fmt.Printf("O valor de res é %v", res2)
}
