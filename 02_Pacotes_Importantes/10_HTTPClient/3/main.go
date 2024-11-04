package main

import (
	"io"
	"net/http"
)

// Funcão usada para fazer mais sets nos header
func main() {
	c := http.Client{}
	req, err := http.NewRequest("GET", "http://google.com", nil)

	if err != nil {
		panic(err)
	}

	req.Header.Set("Accept", "Application/json")
	//meu cliente execute essa request
	resp, err := c.Do(req)

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
