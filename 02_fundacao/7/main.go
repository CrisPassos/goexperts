package main

import "fmt"

func main() {
	//HashMap com chave e valores
	salaries := map[string]float64{
		"John": 50000,
		"Jane": 60000,
		"Bob":  55000,
	}

	fmt.Println("Salaries:", salaries["John"])
	delete(salaries, "John")
	fmt.Println("Salaries:", salaries["John"])
	salaries["John"] = 70000
	fmt.Println("Salaries:", salaries["John"])

	// função make somente cria o map inicial
	// prepara um map vazio e depois eu preencho
	// é o mesmo que fazer map[string]float64{}
	sal := make(map[string]float64)
	sal["Alice"] = 80000

	fmt.Println("Salaries:", sal["Alice"])

	for name, salary := range salaries {
		fmt.Printf("Name: %s, Salary: %.2f\n", name, salary)
	}

	for _, salary := range salaries {
		fmt.Printf("Salary: %.2f\n", salary)
	}

}
