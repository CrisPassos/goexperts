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
	//atrasa a execução da request, isso permite que os processamentos sejam feitos
	//quando acabar tudo, as linhas 21 ao 27 ai sim o defer é executado
	//defer sempre é executado no final
	defer req.Body.Close()

	//retorno da requisição feita
	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(res))

	defer fmt.Println("Primeira Linha")
	fmt.Println("Segunda Linha")
	fmt.Println("Terceira Linha")
}
