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
	curso := Curso{"Go Avançado", 40}
	t := template.Must(template.New("CursoTemplate").Parse("Curso: {{.Nome}}, Carga Horária: {{.CargaHoraria}} horas"))
	err := t.Execute(os.Stdout, curso)

	if err != nil {
		fmt.Println(err)
	}
}
