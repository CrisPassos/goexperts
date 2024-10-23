package main

import "fmt"

// Type assertion (forçar um tipo na interface vazia)
func main() {
	var minhavar interface{} = "Cristiana Passos"
	println(minhavar)
	println(minhavar.(string))

	res, ok := minhavar.(int)
	fmt.Printf("O valor de res %v e o resultado de ok é %v\n", res, ok)

	res2 := minhavar.(int)
	fmt.Println(res2)

}
