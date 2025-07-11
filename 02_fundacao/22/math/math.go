package math

func Sum[T int | float64](a, b T) T {
	return a + b
}

var A int = 10

type Carro struct {
	//acessivel fora da classe
	Marca string
	//não acessivel fora da classe
	cambio string
}
