package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	c := http.Client{}
	jsonVar := bytes.NewBuffer([]byte(`{"name": "cristiana}`))
	resp, err := c.Post("http://google.com", "application/json", jsonVar)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	// eu escolho onde vou jogar os daddos (os.Stdout) e vou pegar esses dados de resp.Body
	io.CopyBuffer(os.Stdout, resp.Body, nil)
}
