package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Alan Doe", "test@test.com", "Password123")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Email)
}

func TestPassword(t *testing.T) {
	user, err := NewUser("Alan Doe", "test@test.com", "Password123")
	assert.Nil(t, err)
	assert.NotNil(t, user.Password)
	assert.NotEmpty(t, user.Password)
	assert.True(t, user.ValidatePassword("Password123"))
	assert.False(t, user.ValidatePassword("password123"))
	assert.NotEqual(t, user.ValidatePassword("Password123"), user.Password)
}
