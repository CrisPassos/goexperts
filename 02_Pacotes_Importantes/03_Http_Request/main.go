package main

import (
	"io"
	"net/http"
)

// requests
func main() {
	req, err := http.Get("https://www.google.com")

	if err != nil {
		panic(err)
	}

	//fechando a request
	//defer vai executar na última linha independente de onde estivermos
	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)

	if err != nil {
		panic(err)
	}

	println(string(res))
}
