package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request iniciada")

	defer log.Println("Request finalizada")

	select {
	//se passar de 5 segundos quer dizer que a chamada foi processada
	case <-time.After(5 * time.Second):
		// Imprime no comand line stdout
		log.Println("Request processada com sucesso")
		// Imprime no browser
		w.Write([]byte("Request processada com sucesso"))

	case <-ctx.Done():
		log.Println("Request cancelada pelo cliente")
	}
}
