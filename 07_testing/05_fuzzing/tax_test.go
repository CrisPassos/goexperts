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

func TestCalculateTaxtBatch(t *testing.T) {
	type calcTax struct {
		amount, expect float64
	}

	//crindo uma matrix com os valores que estou esperando
	//quando eu receber 500 eu vou esperar 5
	table := []calcTax{
		{500.0, 5.0},
		{1000.0, 10.0},
		{1500.0, 10.0},
	}

	//vamos fazer um for para verificar esse valores
	for _, item := range table {
		result := CalculateTax(item.amount)

		if result != item.expect {
			t.Errorf("Expected %f but got %f", item.amount, item.expect)
		}
	}
}

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.0)
	}
}

func BenchmarkCalculateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(500.0)
	}
}

// organizado em dois momento seeds (eu dou algumas informações para ele saber que numeros mandar)
func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2, -2.5, 500.0, 1000.0, 1501.0}
	for _, amount := range seed {
		//o fuzz vai testar com base nesses valores
		f.Add(amount)
	}

	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)

		if amount <= 0 && result != 0 {
			t.Errorf("Received %f but expected 0", result)
		}

		if amount > 20000 && result != 20 {
			t.Errorf("Received %f but expected 20", result)
		}
	})
}
