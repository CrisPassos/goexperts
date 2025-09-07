package product

import "database/sql"

type ProductRepositoryInterface interface {
	GetProduct(id int) (Product, error)
}

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

// simulando uma conexão com o banco de dados
func (r *ProductRepository) GetProduct(id int) (Product, error) {
	return Product{ID: id, Name: "Sample Product"}, nil
}
