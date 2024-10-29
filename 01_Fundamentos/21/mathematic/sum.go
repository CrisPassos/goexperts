package mathematic

// Tudo que começa com a letra maiuscula é EXPORT
func Sum[T int | float32](a T, b T) T {
	return a + b
}
