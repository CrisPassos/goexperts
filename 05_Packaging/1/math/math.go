package math

type Math struct {
	A int
	B int
}

// Letra maisucula pode ser um metódo publico
func (m Math) Add() int {
	return m.A + m.B
}

// Letra minuscula somente visivel na classe
func (m Math) sub() int {
	return m.A - m.B
}
