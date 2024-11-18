package main

import (
	"1_Client_Server_API/shared"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 300*time.Millisecond)
	defer cancel()

	// Open a connection with /server
	c := http.Client{}
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	HandleError(err)

	req.Header.Set("Accept", "Application/json")

	resp, err := c.Do(req)
	HandleError(err)

	defer resp.Body.Close()

	// Reading response
	body, err := io.ReadAll(resp.Body)
	HandleError(err)
	println(string(body))

	var data shared.Cotacao
	err = json.Unmarshal(body, &data)
	HandleError(err)

	RunTemplate(&data)
	SaveAsTxtFile(&data)
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

func RunTemplate(data *shared.Cotacao) {
	template := template.Must(template.New("page.html").ParseFiles("page.html"))
	err := template.Execute(os.Stdout, data)
	HandleError(err)
}

func SaveAsTxtFile(data *shared.Cotacao) {
	file, err := os.Create("cotacao.txt")
	HandleError(err)

	defer file.Close()
	_, err = file.WriteString(fmt.Sprintf("Dólar: %s", data.USDBRL.Bid))
	HandleError(err)
}
