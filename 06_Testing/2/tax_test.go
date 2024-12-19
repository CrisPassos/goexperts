package tax

import (
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
