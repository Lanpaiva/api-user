package entity

import (
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
	return &User{
		ID:       models.NewID(),
		Name:     name,
		Email:    email,
		Password: string(hash),
	}, nil

}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	password = "Password123"
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

	return err == nil
}
