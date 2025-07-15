package main

import (
	"io"
	"net/http"
)

func main() {
	c := http.Client{}
	//criamos requestes personalizadas ou customizadas
	req, err := http.NewRequest("GET", "http://google.com", nil)

	if err != nil {
		panic(err)
	}

	req.Header.Set("Accept", "application/json")
	// fazemos a junção da request com o header aqui através do metódo DO
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
