package tax

import "github.com/stretchr/testify/mock"

type TaxRepositoryMock struct {
	mock.Mock
}

// implementando a função SaveTax
func (m *TaxRepositoryMock) SaveTax(tax float64) error {
	args := m.Called(tax)
	return args.Error(0)
}
