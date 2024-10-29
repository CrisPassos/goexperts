package main

import (
	"encoding/json"
	"os"
)

// Data Annotation ou decorators do Go
type Account struct {
	Number  int
	Balance int `json:"bal"`
}

func main() {
	account := Account{Number: 010, Balance: 1000}

	//Marshal converte em JSON: guarda o valor
	res, err := json.Marshal(account)

	if err != nil {
		panic(err)
	}

	println(res)
	println(string(res))

	//encoder: entrega o valor já pronto
	//o retorno de newEncoder é um erro
	json.NewEncoder(os.Stdout).Encode(account)

	//decoder: recebe um json e transfora em uma struct ou objeto
	//string não deve conter espaços
	jsonPure := []byte(`{"Number":20,"bal":2000}`)
	var accountX Account
	//faz a deseralização e aponta para o item de memória em accountX
	//Unmarshal retorna um erro
	err = json.Unmarshal(jsonPure, &accountX)
	if err != nil {
		println(err)
	}
	println(accountX.Balance)
}
