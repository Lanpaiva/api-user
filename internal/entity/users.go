package entity

import (
	"errors"
	"unicode"

	"github.com/lanpaiva/api-user/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       models.ID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
}

func NewUser(name, email, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	if password == "" {
		return nil, errors.New("a senha não pode estar vazia")
	}

	upperCase := false
	lowerCase := false

	for _, char := range password {
		if unicode.IsUpper(char) {
			upperCase = true
		} else if unicode.IsLower(char) {
			lowerCase = true
		}
		if upperCase && lowerCase {
			break
		}
	}

	if !upperCase || !lowerCase {
		return nil, errors.New("a senha deve conter pelo menos um caractere maiúsculo e um minúsculo")
	}
	return &User{
		ID:       models.NewID(),
		Name:     name,
		Email:    email,
		Password: string(hash),
	}, nil

}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	return err == nil
}
