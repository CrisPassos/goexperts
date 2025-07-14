package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCep struct {
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

// digita o número do Cep e retorna rua, bairro, cidade e estado
func main() {
	//pega dos os argumentos do inicio ao fim e printa exemplo:
	//go run main.go cristiana ou o cep no caso
	for _, cep := range os.Args[1:] {
		req, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
		}

		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v\n", err)
		}

		// os dados voltam como como json e eu preciso passar para struct
		var data ViaCep
		err = json.Unmarshal(res, &data)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer parse da resposta: %v\n", err)
		}

		//Exemplo go run main.go https://viacep.com.br/ws/01001000/json/
		fmt.Println(data)
		fmt.Println(data.Localidade)

		//armazenar num arquivo
		file, err := os.Create("cidade.txt")

		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar arquivo: %v\n", err)
		}

		defer file.Close()
		_, err = file.WriteString(fmt.Sprintf("CEP: %s\nRua: %s\nBairro: %s\nCidade: %s\nEstado: %s\n",
			data.Cep, data.Logradouro, data.Bairro, data.Localidade, data.Uf))

		fmt.Println("Arquivo criado com sucesso!")
		fmt.Println("Cidade:", data.Localidade)
	}
}
