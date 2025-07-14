package main

import (
	"encoding/json"
	"io"
	"net/http"
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

func main() {
	http.HandleFunc("/", BuscaCepHandler)
	http.ListenAndServe(":8080", nil)

}

func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Página não encontrada"))
		return
	}

	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("CEP não informado"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World!"))
}

// retorno é um ponteiro para ViaCep e um erro, isso para ter uma mundaça local
// agora sim estamos buscando o CEP com o HTTP GET
func BuscaCep(cep string) (*ViaCep, error) {
	resp, error := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if error != nil {
		return nil, error
	}
	defer resp.Body.Close()
	body, error := io.ReadAll(resp.Body)

	if error != nil {
		return nil, error
	}

	var c ViaCep
	// decodifica o JSON recebido e armazena na variável c
	// salva o valor do json na variável c
	error = json.Unmarshal(body, &c)

	if error != nil {
		return nil, error
	}

	// retornamos o ponteiro de memória e sem erros
	return &c, nil

}
