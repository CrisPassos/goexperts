package main

import (
	"1_Client_Server_API/shared"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

func main() {
	// Open a connection with /server
	c := http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8080/cotacao", nil)
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
