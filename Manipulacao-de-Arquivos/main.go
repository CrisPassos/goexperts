package main

import (
	"bufio"
	"fmt"
	"os"
)

// criando e manipulando um arquivo
func main() {
	f, err := os.Create("arquivo.txt")

	if err != nil {
		panic(err)
	}

	//escreve uma string
	// f.WriteString("Hello World")

	//escreve um bye
	tamanho, err := f.Write([]byte("Escrevendo dados no arquivo"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes\n", tamanho)

	f.Close()

	//Leitura
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf(string(arquivo))

	//Leitura em blocks, quando o arquivo é maior que a memória o ideal
	//é abrir o arquivos em partes
	blocos, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(blocos)
	buffer := make([]byte, 3)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}

		fmt.Println(string(buffer[:n]))
	}

	//removendo arquivos
	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}

}
