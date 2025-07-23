package entity

import (
	"github.com/CrisPassos/goexperts/08_apis/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

//esse user vai ser gravado na base de dados e validações
//a entidade é onde fica as regras de negócios

type User struct {
	ID       entity.ID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
}

// para gerar a senha criptografada, vamos usar o bcrypt
// nunca guardamos a senha e sim criptografamos ela
func NewUser(name, email, password string) (*User, error) {
	//genrando um hash da senha
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	return &User{
		ID:       entity.NewID(),
		Name:     name,
		Email:    email,
		Password: string(hash),
	}, nil
}

// verificando se a senha digitada bate com a senha registrada do usuer
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
