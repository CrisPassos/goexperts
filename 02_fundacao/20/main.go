package main

// Contraints
type MyNumber int

type Number interface {
	int | float64
}

func SumInt(m map[string]int) int {
	var soma int

	for _, v := range m {
		soma += v
	}

	return soma
}

// O T significa que eu posso usar o int ou o float tudo depende, usamos para não repetir codigo
// podemos usar assim:
// func Sum[T int | float64](m map[string]T) T {
// ou assim:
func Sum[T Number](m map[string]T) T {
	var soma T

	for _, v := range m {
		soma += v
	}

	return soma
}

// aceita os termos que são comparaveis entre si
func Comparar[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"Cristiana": 1000, "Joana": 3000, "Maria": 3000}
	m2 := map[string]float64{"Cristiana": 1000.50, "Joana": 3000.50, "Maria": 3000}
	println(SumInt(m))

	//usando genericos
	println(Sum(m))
	println(Sum(m2))

	// m3 := map[string]MyNumber{"Cristiana": 1000.50, "Joana": 3000.50, "Maria": 3000}
	// não vai funcionar pq MyNumber não é um tipo Number
	// isso devido a caracteristica fortmente tipada
	// println(Sum(m3))

	println(Comparar(10, 10.0))

}
