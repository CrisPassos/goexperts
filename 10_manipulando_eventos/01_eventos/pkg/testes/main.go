package main

import "fmt"

func main() {
	evento := []string{"evento1", "evento2", "evento3", "evento4"}
	//dois primeiros registros
	// evento = evento[:2]
	// fmt.Println(evento)

	// //dois ultimos registros
	// evento = evento[2:]
	// fmt.Println(evento)

	//remove o primeiro registro e depois gruda o resto
	//vamos usar o append para remover um evento
	evento = append(evento[:0], evento[1:]...)
	fmt.Println(evento)
}
