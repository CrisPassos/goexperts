package main

import (
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	// primeiro eu faço um setup antes de chamar a api
	// tipo de erro: context deadline exceeded (Client.Timeout exceeded while awaiting headers)
	c := http.Client{
		Timeout: time.Duration(1) * time.Microsecond,
	}

	resp, err := c.Get("http://google.com")

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	println(string(body))
}
