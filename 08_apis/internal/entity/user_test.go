package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("John Doe", "j@j.com", "123456")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "j@j.com", user.Email)
}

func TestCheckPassword(t *testing.T) {

	user, err := NewUser("John Doe", "j@j.com", "123456")

	assert.Nil(t, err)
	assert.True(t, user.CheckPassword("123456"))
	assert.False(t, user.CheckPassword("wrongpassword"))
	assert.NotEqual(t, "123456", user.Password)
}
