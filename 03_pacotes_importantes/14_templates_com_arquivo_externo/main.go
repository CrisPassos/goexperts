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
	t := template.Must(template.New("template.html").ParseFiles("template.html"))

	// o range é feito no time Cursos
	// Exemplo de uso emails
	err := t.Execute(os.Stdout, Cursos{
		{"Go Avançado", 40},
		{"JAVA", 80},
		{"Python", 15},
	})

	if err != nil {
		fmt.Println(err)
	}
}
