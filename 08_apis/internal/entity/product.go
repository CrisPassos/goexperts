package entity

import (
	"errors"
	"time"

	"github.com/CrisPassos/goexperts/08_apis/pkg/entity"
)

// CRUD COMPLETO
// vamos começar pela validacao de erros
var (
	ErrIDisRequired    = errors.New("ID is required")
	ErrInvalidID       = errors.New("invalid ID")
	ErrNameIsRequired  = errors.New("name is required")
	ErrPriceIsRequired = errors.New("price is required")
	ErrInvalidPrice    = errors.New("invalid price")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

// criando o produto
func NewProduct(name string, price float64) (*Product, error) {
	product := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}

	err := product.Validate()

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrIDisRequired
	}

	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrInvalidID
	}

	if p.Name == "" {
		return ErrNameIsRequired
	}

	if p.Price == 0 {
		return ErrPriceIsRequired
	}

	if p.Price < 0 {
		return ErrInvalidPrice
	}

	return nil
}
