package tax

import "testing"

// go test -v (verbose) or go test
// não esquecer de ter um go mod init inicializado
// Test covarage

// go test -coverprofile=coverage.out
// mostra o que não foi coberto
// go tool cover -html=coverage.out
func TestCalculateTax(t *testing.T) {
	amount := 500.00
	expected := 5.0

	result := CalculateTax(amount)

	if result != expected {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount, expected float64
	}

	table := []calcTax{
		{500.00, 5.0},
		{1000.00, 10.0},
		{1500.00, 10.0},
	}

	for _, item := range table {
		result := CalculateTax(item.amount)

		if result != item.expected {
			t.Errorf("Expected %f but got %f", item.expected, result)
		}

	}
}

// go test -bench=.
func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.0)
	}
}

// passa varios dados até quebrar a aplicação
// go test -fuzz=.
func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2, -2.5, 500.0, 1000.0, 1501.0}
	for _, amount := range seed {
		f.Add(amount)
	}

	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)
		if amount <= 0 && result != 0 {
			t.Errorf("Reveived %f but expected 0", result)
		}
		if amount > 20000 && result != 20 {
			t.Errorf("Reveived %f but expected 20", result)
		}
	})

}
