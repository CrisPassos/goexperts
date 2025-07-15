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

type Blog struct {
	Title string
}

// fiz um attache a minha struct Blog para implementar o ServeHTTP
func (b *Blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.Title + "\n"))
}

func main() {
	mux := http.NewServeMux()

	//criando meu mux personalizado
	mux.HandleFunc("/", HomeHandler)
	// como usar: http://localhost:8080/cep?cep=01001000
	mux.HandleFunc("/cep", BuscaCepHandler)

	// ao adicionar uma sruct que implementa o ServeHTTP, posso usar o mux para definir uma rota
	// Vantagem, customização do handler
	mux.Handle("/blog", &Blog{Title: "Blog do Go Experts"})

	http.ListenAndServe(":8080", mux)

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Bem-vindo ao meu servidor personalizado!"))
	})

	http.ListenAndServe(":8081", mux2)

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bem-vindo à API de Busca de CEP!"))
}

func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/cep" {
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

	cep, err := BuscaCep(cepParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao buscar CEP"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cep)
}

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
	error = json.Unmarshal(body, &c)

	if error != nil {
		return nil, error
	}

	return &c, nil

}
