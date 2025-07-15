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

func main() {
	curso := Curso{
		Nome:         "Go Avançado",
		CargaHoraria: 40,
	}
	// Criando um template, e parseando o template
	// para trabalhar nos templates usamos sempre .
	temp := template.New("CursoTemplate")
	temp, _ = temp.Parse("Curso: {{.Nome}}, Carga Horária: {{.CargaHoraria}} horas")

	//executar o processo
	err := temp.Execute(os.Stdout, curso)

	if err != nil {
		fmt.Println(err)
	}
}
