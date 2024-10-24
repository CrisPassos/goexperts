package main

type MyNumber int

// constrant
// ~ reconhece que o tipo MyNumber é um int
type Number interface {
	~int | ~float64
}

func Sum[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}

	return soma
}

func Compare[T comparable](a T, b T) bool {
	if a == b {
		return true
	}

	return false
}

// Generics
func main() {
	m := map[string]int{"Cristiana": 1000, "Mariana": 2000, "Carla": 4000}
	m2 := map[string]float64{"Cristiana": 100.50, "Mariana": 200.40, "Carla": 400.60}
	m3 := map[string]MyNumber{"Cristiana": 100, "Mariana": 200, "Carla": 400}

	println(Sum(m))
	println(Sum(m2))
	println(Sum(m3))
	println(Compare(10, 5.5))

}
