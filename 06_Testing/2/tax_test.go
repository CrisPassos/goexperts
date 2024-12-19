package tax

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	tax, error := CalculateTax(1000.0)
	assert.Nil(t, error)
	assert.Equal(t, 10.0, tax)

	tax, error = CalculateTax(0)
	assert.Error(t, error, "amount must be greater than 0")
	assert.Equal(t, 0.0, tax)
}

func TestCalculateTaxAndSave(t *testing.T) {
	//tudo fake
	repository := &TaxRepositoryMock{}
	// aqui estou esperando que haja um valor e por isso não retorna error
	// cria as regras do que ocorreria numa chamada ao banco de dados
	// ao retornar 10, o valor é nulo
	// ao retornar 0, o valor é um erro
	repository.On("SaveTax", 10.0).Return(nil)
	repository.On("SaveTax", 0.0).Return(errors.New("Error saving tax"))

	err := CalculateTaxAndSave(1000.0, repository)
	assert.Nil(t, err)

	err = CalculateTaxAndSave(0.0, repository)
	assert.Error(t, err, "Error saving tax")

	repository.AssertExpectations(t)
}
