package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type EnderecoViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

type EnderecoBrasilAPI struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
}

func consultarViaCEP(cep string, ch chan<- string) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var endereco EnderecoViaCEP
	body, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &endereco); err == nil {
		ch <- fmt.Sprintf("ViaCEP: %s - %s, %s - %s/%s", endereco.Cep, endereco.Logradouro, endereco.Bairro, endereco.Localidade, endereco.Uf)
	}
}

func consultarBrasilAPI(cep string, ch chan<- string) {
	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var endereco EnderecoBrasilAPI
	body, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &endereco); err == nil {
		ch <- fmt.Sprintf("BrasilAPI: %s - %s, %s - %s/%s", endereco.Cep, endereco.Street, endereco.Neighborhood, endereco.City, endereco.State)
	}
}

func main() {
	cep := "08471010"
	fmt.Println("Consultando CEP:", cep)

	// Canais para receber os resultados das APIs
	viaCepCh := make(chan string)
	brasilAPICh := make(chan string)

	// Goroutines para consultar as APIs
	go consultarViaCEP(cep, viaCepCh)
	go consultarBrasilAPI(cep, brasilAPICh)

	// Vai informar quem é o mais rápido a responder
	select {
	case resultado := <-viaCepCh:
		fmt.Println("Resultado ViaCEP:")
		fmt.Println(resultado)

	case resultado := <-brasilAPICh:
		fmt.Println("Resultado BrasilAPI:")
		fmt.Println(resultado)

	case <-time.After(1 * time.Second):
		fmt.Println("Timeout: Nenhuma API respondeu a tempo.")
	}
}
