package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	//os.Args pega tudo que foi digitado
	for _, cep := range os.Args[1:] {
		req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")

		if err != nil {
			//O F na frente joga a resposta em algum lugar
			fmt.Fprintf(os.Stderr, "Erro ao fazer a requisição: %v\n", err)
		}

		defer req.Body.Close()
		//lê a request
		res, err := io.ReadAll(req.Body)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer a requisição: %v\n", err)
		}

		//deserializa em JSON
		//usar o Json convert to GO Struct
		var dataCep ViaCEP
		err = json.Unmarshal(res, &dataCep)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer parse da resposta: %v\n", err)
		}

		//armazena em um ficheiro
		arquivo, err := os.Create("rua.txt")

		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro em criar um arquivo: %v\n", err)
		}
		defer arquivo.Close()

		//no ficheiro escrever da seguinte forma
		_, err = arquivo.WriteString(fmt.Sprintf("CEP: %s, Localidade: %s, UF: %s", dataCep.Cep, dataCep.Localidade, dataCep.Uf))
		fmt.Println("Arquivo criado com sucesso!")
		fmt.Println("Cidade: ", dataCep.Localidade)

	}
}
