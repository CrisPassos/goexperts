package main

import (
	"fmt"
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}
	t := template.Must(template.New("content.html").ParseFiles(templates...))

	err := t.Execute(os.Stdout, Cursos{
		{"Go Avançado", 40},
		{"JAVA", 80},
		{"Python", 15},
	})

	if err != nil {
		fmt.Println(err)
	}

}
