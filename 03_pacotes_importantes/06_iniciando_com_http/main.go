package main

import "net/http"

func main() {
	// ao acessar a raiz do servidor, a função BuscaCep será chamada
	// é possivel trabalhar com funções anônimas também
	http.HandleFunc("/", BuscaCep)
	// Define a porta onde o servidor irá escutar
	http.ListenAndServe(":8080", nil)

}

// controllers
func BuscaCep(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
