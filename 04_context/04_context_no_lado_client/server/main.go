package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

	//vamos ter 5 segundos para processar a requisição e depois vamos retornar isso ao user
	//se passar os 5 vamos cancelar a req e o user vai receber uma mensagem com o status timeout

}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request iniciada")
	defer log.Println("Request finalizada")

	select {
	// caso seja mais de 5 segundos processada com sucesso, será imprimida na tela do servidor
	case <-time.After(5 * time.Second):
		//imprimi no comand line
		log.Println("Request processada com sucesso")
		//imprimi no browser
		w.Write([]byte("Request processada com sucesso"))
	// usada para casos complexos, como cálculos que exigem um tempo
	case <-ctx.Done():
		log.Println("Request cancelada pelo cliente")
		http.Error(w, "Request cancelada pelo cliente", http.StatusRequestTimeout)
	}
}
