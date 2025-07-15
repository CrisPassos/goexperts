package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/// criaçao de arquivos ///
	f, err := os.Create("arquivo.txt")

	if err != nil {
		panic(err)
	}

	/// escrever no arquivo ///
	tamanho, err := f.WriteString("Hello, world!")

	if err != nil {
		panic(err)
	}

	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes\n", tamanho)

	//posso escrever com bytes também
	tamanhoBytes, err := f.Write([]byte("Escrevendo dados num arquivo\n"))

	if err != nil {
		panic(err)
	}

	fmt.Printf("Arquivo criado com sucesso! tamanha: %d bytes\n", tamanhoBytes)

	////leitura de um arquivo////
	arquivo, err := os.ReadFile("arquivo.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(arquivo))

	//Lendo um arquivo maior do que o esperado, fazemos isso em batches ou blocos
	bloco, err := os.Open("arquivo.txt")

	if err != nil {
		panic(err)
	}

	//depois de que o arquivo estiver aberto eu preciso ler arquivo em arquivo e para isso usamos
	//a biblioteca chamada buffer.io
	reader := bufio.NewReader(bloco)
	buffer := make([]byte, 10)

	// faco um arquivo que le o arquivo e depois imprimi
	// um dos erros é quando o arquivo acabar, ele vai parar
	for {
		n, err := reader.Read(buffer)

		if err != nil {
			break
		}
		// a cada 10 bytes ele quebra o conteudo do arquivo
		fmt.Println(string(buffer[:n]))
	}

	f.Close()

	//remover o arquivo
	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}
}
