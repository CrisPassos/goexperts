package main

import (
	"log"
	"net/http"
)

func main() {
	// rodando em arquivos estáticos
	fileServer := http.FileServer(http.Dir("./public"))
	mux := http.NewServeMux()
	mux.Handle("/", fileServer)
	mux.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the blog!"))
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
