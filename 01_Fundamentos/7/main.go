package main

import "fmt"

// Criação de HASHMAP(Chave e Valor)(Key, Value)
func main() {
	salarios := map[string]int{"Cristiana": 1000, "Mariana": 2000, "Juliana": 3000}
	fmt.Println(salarios["Cristiana"])
	//delete
	delete(salarios, "Mariana")
	//criação
	salarios["Fernanda"] = 4000
	fmt.Println(salarios["Fernanda"])

	//outras formas de declarar um HASHMAP
	sal := make(map[string]int)
	sal1 := map[string]int{}

	sal["Aparecida"] = 1000
	sal1["Maria"] = 2000

	//percorrer um array
	for name, salario := range salarios {
		fmt.Printf("O salário de %s é de %d\n", name, salario)
	}

	//se eu quiser mostrar somente o salário
	for _, salario := range salarios {
		fmt.Printf("O salário é de %d\n", salario)
	}

}
