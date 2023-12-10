package usecase

import (
	"github.com/lanpaiva/api-user/internal/entity"
	"github.com/lanpaiva/api-user/internal/infra/database"
)

type CreateTaskInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateTaskInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserOutput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type NewUser struct {
	UserDb database.UserInterface
}

func (n *NewUser) Create(input CreateUserInput) (*CreateUserOutput, error) {
	user, err := entity.NewUser(input.Name, input.Email, input.Password)
	if err != nil {
		return nil, err
	}
	err = n.UserDb.Create(user)
	if err != nil {
		return nil, err
	}
	return &CreateUserOutput{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}
