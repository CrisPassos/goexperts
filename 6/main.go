package main

import "fmt"

func main() {
	//declaração de um slice (por de trás dos panos é executado um array)
	s := []int{10, 20, 30, 40, 50, 60, 100}

	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	//o : faz um corte ou um slide, para direita ou para esquerda
	//nesse caso ele fez o corte para a esquerda, por isso o resultado é vazio
	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])
	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4])

	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])

	//aumentar a capacidade do slice, ele sempre duplica o valor original, por isso
	//a capacidade será de 14 e não de 8 no final
	s = append(s, 100)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

/**
 * Claro! Em Go, slices são uma estrutura de dados que oferecem uma maneira flexível e poderosa de trabalhar com coleções de elementos. Eles são construídos sobre arrays, mas diferentemente dos arrays, slices têm um tamanho dinâmico. Aqui estão alguns pontos chave sobre slices:

Definição: Um slice é uma referência a uma sequência contínua de elementos de um array. Ele contém um ponteiro para o array, a quantidade de elementos (length) e a capacidade (capacity), que é o tamanho máximo que o slice pode crescer sem precisar criar um novo array subjacente.

Criação: Slices podem ser criados de várias maneiras, incluindo a partir de arrays, usando a função make, ou simplesmente declarando e inicializando com valores.

// A partir de um array
arr := [5]int{1, 2, 3, 4, 5}
slice := arr[1:4] // Contém [2, 3, 4]

// Usando make
slice2 := make([]int, 3, 5) // Cria um slice de length 3 e capacity 5
Manipulação: Slices podem ser manipulados usando funções embutidas como append, que adiciona elementos ao final do slice, e copy, que copia elementos de um slice para outro.

Capacidade e Comprimento: O comprimento de um slice é o número de elementos que ele contém, enquanto a capacidade é o número de elementos no array subjacente a partir do primeiro elemento do slice.

Referência: Como slices são referências a arrays, modificações feitas em um slice afetam o array subjacente e, portanto, outros slices que referenciam o mesmo array.

Slices são uma parte fundamental da linguagem Go, oferecendo uma maneira eficiente de manipular coleções de dados. Se precisar de mais detalhes ou exemplos específicos, sinta-se à vontade para perguntar!
*/
