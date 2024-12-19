package tax

import (
	"github.com/stretchr/testify/mock"
)

type TaxRepositoryMock struct {
	//simula os métodos da interface
	mock.Mock
}

func (m *TaxRepositoryMock) SaveTax(tax float64) error {
	args := m.Called(tax)
	return args.Error(0)
}
