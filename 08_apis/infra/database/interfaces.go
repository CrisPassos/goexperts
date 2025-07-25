package database

import "github.com/CrisPassos/goexperts/08_apis/internal/entity"

// desacoplando as coisas para fazer tudo via interfaces
type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(emaild string) (*entity.User, error)
}

// interface para CRUD do produ
type ProductInterface interface {
	Create(product *entity.Product) error
	FindAll(page, limit int, sort string) (entity.Product, error)
	FindById(id string) (*entity.Product, error)
	Update(product *entity.Product) error
	Delete(id string) error
}
