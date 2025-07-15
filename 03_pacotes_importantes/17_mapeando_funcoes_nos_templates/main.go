package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

// apenas um exemplo
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}
	t := template.New("content.html")
	// t.Func, eu consigo colocar um mapa de função que pode ser mapeado dentro do HTML
	// mapa chave e valor
	t.Funcs(template.FuncMap{"ToUpper": ToUpper})
	t = template.Must(t.ParseFiles(templates...))

	err := t.Execute(os.Stdout, Cursos{
		{"Go Avançado", 40},
		{"JAVA", 80},
		{"Python", 15},
	})

	if err != nil {
		fmt.Println(err)
	}

}
