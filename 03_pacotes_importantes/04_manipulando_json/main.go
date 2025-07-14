package main

import (
	"encoding/json"
	"os"
)

// Tags usadas para personalizar a serialização
// `json:"numero"` significa que o campo "Numero" será serializado como "numero
// o json:"s"` significa que o campo "Saldo" será serializado como "s"
// Tags são usadas para mapear os campos da struct para os nomes desejados no JSON
// se eu usar - json:"-" o campo não será serializado, ele não aparecerá no JSON
// podemos usar validações também, por exemplo: `json:"s,omitempty"`
// tipo annotations ou decorators em outras linguagens
type Conta struct {
	Numero int
	Saldo  int    `json:"s"`
	Vazion string `json:"-"` // este campo não será serializado
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100}
	//marshal faz a serialização
	//quando eu uso o marshal eu guardo o valor para mim
	res, err := json.Marshal(conta)

	if err != nil {
		println(err)
	}
	//Json sempre retorna em bytes
	println(res)
	println(string(res))

	//encoder: não armazena o valor já usa diretamente
	//os.Stdout, vai sair para a saida padrão e mostrar o resultado
	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		println(err)
	}

	//decoder: quando eu faço o processo ao contrário transforma do json para struct
	jsonPuro := []byte(`{"Numero":1,"s":200}`)
	//por padrão a conta está vazia (alocação vazia na memória)
	var contaX Conta
	//quando usamos o Unmarshal estamos falando, vai naquele espaço vazio e atribui um valor
	//salva o valor do jsonPuro dentro da alocação da memória do contaX
	//não esquecer de passar o endereço da variável
	err = json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		println(err)
	}

	println(contaX.Saldo)

}
