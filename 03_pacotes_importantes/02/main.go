package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// acessando ao Google
	req, err := http.Get("https://www.google.com")

	if err != nil {
		panic(err)
	}

	//retorno da requisição feita
	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(res))

	//fechando para economizar os recursos, exemplo quando fechamos os arquivos
	req.Body.Close()
}
