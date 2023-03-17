package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Alan Doe", "alan@test.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Email)
}

func TestPassword(t *testing.T) {
	user, err := NewUser("Alan Doe", "alan@test.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user.Password)
	assert.NotEmpty(t, user.Password)
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("1234567"))
	assert.NotEqual(t, user.ValidatePassword("123456"), user)
}
