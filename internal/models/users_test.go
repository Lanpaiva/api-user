package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Alan Doe", "alan@test.com", "Senha123")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Email)
}

func TestPassword(t *testing.T) {
	user, err := NewUser("Alan Doe", "alan@test.com", "Senha123")
	assert.Nil(t, err)
	assert.NotNil(t, user.Password)
	assert.NotEmpty(t, user.Password)
	assert.True(t, user.ValidatePassword("Senha123"))
	assert.False(t, user.ValidatePassword("senha123"))
	assert.NotEqual(t, user.ValidatePassword("Senha123"), user)
}
