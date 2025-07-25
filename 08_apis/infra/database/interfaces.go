package database

import "github.com/CrisPassos/goexperts/08_apis/internal/entity"

// desacoplando as coisas para fazer tudo via interfaces
type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(emaild string) (*entity.User, error)
}
