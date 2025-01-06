package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	p, err := NewProduct("Product 1", 100)
	assert.Nil(t, err)                   // sem errors
	assert.NotNil(t, p)                  // produto preenchido
	assert.NotEmpty(t, p.ID)             // id preenchido
	assert.Equal(t, "Product 1", p.Name) // nome do produto
	assert.Equal(t, 100.0, p.Price)      // preço do produto
}

func TestProductWhenNameIsRequired(t *testing.T) {
	_, err := NewProduct("", 100)
	assert.NotNil(t, err) // tem erro, pq o valor é nulo
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	_, err := NewProduct("Product 1", 0)
	assert.NotNil(t, err) // tem erro, pq o valor é nulo
	assert.Equal(t, ErrPriceIsRequired, err)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	_, err := NewProduct("Product 1", -10)
	assert.NotNil(t, err) // tem erro, pq o valor é nulo
	assert.Equal(t, ErrInvalidPrice, err)
}

func TestProductValidate(t *testing.T) {
	p, err := NewProduct("Product 1", 100)
	assert.Nil(t, err)  // sem errors
	assert.NotNil(t, p) // produto preenchido
	assert.Nil(t, p.Validate())
}
