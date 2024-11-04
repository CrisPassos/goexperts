package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {
	curso := Curso{"Go", 40}
	// o comando Must trata os erros
	t := template.Must(
		template.New("CursoTemplate").Parse("Curso: {{.Nome}} -- Carga Horária {{.CargaHoraria}}\n"))
	err := t.Execute(os.Stdout, curso)

	if err != nil {
		panic(err)
	}
}
