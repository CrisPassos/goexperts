package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx := context.Background()
	//vamos ter um context que vai se encerrar depois de 5 segundos
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	//fazendo chamada ao servidor
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080", nil)

	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	io.Copy(os.Stdout, res.Body)
}
