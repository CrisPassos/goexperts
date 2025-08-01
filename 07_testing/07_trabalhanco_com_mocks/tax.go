package tax

import "errors"

// simulando uma conexão com a base de dados, ainda inexistente
type Repository interface {
	SaveTax(amount float64) error
}

// meu objetivo é testar a unidade sem saber a conexão
// precisamos emular esse resultado nos testes
func CalculateTaxAndSave(amount float64, repository Repository) error {
	tax := CalculateTax2(amount)
	return repository.SaveTax(tax)
}

func CalculateTax(amount float64) (float64, error) {
	if amount <= 0 {
		return 0.0, errors.New("opa temos um error")
	}

	if amount >= 1000 && amount < 20000 {
		return 10.0, nil
	}

	if amount >= 20000 {
		return 20.0, nil
	}

	return 5.0, nil
}

func CalculateTax2(amount float64) float64 {
	if amount <= 0 {
		return 0.0
	}

	if amount >= 1000 && amount < 20000 {
		return 10.0
	}

	if amount >= 20000 {
		return 20.0
	}

	return 5.0
}
