package tax

import "testing"

// sempre que eu vou trabalhar com Testes eu sempre preciso trabalhar com a palavra
// TEST na frente
func TestCalculate(t *testing.T) {
	amount := 500.0
	expected := 5.0

	result := CalculateTax(amount)

	//opção nativa é com o if
	if result != expected {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}
