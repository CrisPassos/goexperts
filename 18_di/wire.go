//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/cristianapassos/18-DI/product"
	"github.com/google/wire"
)

// resolve as dependencias no build
func NewUseCase(db *sql.DB) *product.ProductUseCase {
	wire.Build(product.NewProductRepository, product.NewProdutUseCase)
	return &product.ProductUseCase{}
}
