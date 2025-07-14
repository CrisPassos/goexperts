package main

import "net/http"

func main() {
	http.HandleFunc("/", BuscaCepHandler)
	http.ListenAndServe(":8080", nil)

}

func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {
	// se for diferente de "/", retorna 404
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Página não encontrada"))
		return
	}
	//manipulação de strings
	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("CEP não informado"))
		return
	}
	// faz com que retorne um status em json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World!"))
}
