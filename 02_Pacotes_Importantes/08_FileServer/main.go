package main

import (
	"log"
	"net/http"
)

func main() {
	// aqui abrimos um ficheiro que está em public na porta :8080/
	fileserver := http.FileServer(http.Dir("./public"))
	mux := http.NewServeMux()
	mux.Handle("/", fileserver)

	// aqui escrevemos um write na porta :8080/blog
	mux.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Meu BLOG"))
	})

	log.Fatal(http.ListenAndServe(":8080", mux))

}
