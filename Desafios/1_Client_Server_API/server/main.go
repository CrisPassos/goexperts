package main

import (
	"1_Client_Server_API/shared"
	"encoding/json"
	"io"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/cotacao", CotacaoHandler)

	http.ListenAndServe(":8080", mux)
}

func CotacaoHandler(w http.ResponseWriter, r *http.Request) {
	cotacao, error := GetCotacao()

	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(cotacao)
}

func GetCotacao() (*shared.Cotacao, error) {
	resp, err := http.Get("https://economia.awesomeapi.com.br/json/last/USD-BRL")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, error := io.ReadAll(resp.Body)

	if error != nil {
		return nil, error
	}

	var c shared.Cotacao
	error = json.Unmarshal(body, &c)

	if error != nil {
		return nil, error
	}

	return &c, nil
}
