package tax

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// quero emular um teste que salva as coisas na base de dados,
// sen necessáriamente ter um banco

func TestCalculateTaxAndSave(t *testing.T) {
	//estou mockando a chamada do repository
	repository := &TaxRepositoryMock{}
	repository.On("SaveTax", 10.0).Return(nil)
	repository.On("SaveTax", 0.0).Return(errors.New("opa temos um error"))

	err := CalculateTaxAndSave(1000.0, repository)
	assert.Nil(t, err)

	err = CalculateTaxAndSave(0.0, repository)
	assert.Error(t, err, "opa temos um error")

	repository.AssertExpectations(t)
}
