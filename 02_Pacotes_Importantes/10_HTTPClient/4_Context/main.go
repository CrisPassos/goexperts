package main

import (
	"context"
	"io"
	"net/http"
	"time"
)

// Trabalhando com Contextos, lembra um pouco os Pipes do Angular
// Context: Contextos em Go são uma forma de gerenciar e controlar
// o tempo de vida de processos, especialmente em aplicações que lidam
// com operações assíncronas ou que podem ser canceladas, como requisições HTTP.
// O pacote context do Go fornece funcionalidades para passar valores,
// sinais de cancelamento e deadlines entre diferentes partes de um programa.
func main() {
	ctx := context.Background()
	//Funciona como um timeout, testar como Microsecond
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://google.com", nil)

	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	println(string(body))

}
