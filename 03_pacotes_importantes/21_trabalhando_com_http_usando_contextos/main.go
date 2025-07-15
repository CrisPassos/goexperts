package main

import (
	"context"
	"io"
	"net/http"
	"time"
)

func main() {
	// vai executar os processos até 1 segudo, apos isso ele vai parar
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Microsecond)

	//se rodar esse função cancela o contexto
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
