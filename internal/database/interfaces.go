package database

import "github.com/lanpaiva/api-user/internal/models"

type UserInterface interface {
	Create(user *models.User) error
	FindByEmail(email string) (*models.User, error)
}

type TaskInterface interface {
	Create(task *models.Task) error
	FindById(id string) (*models.Task, error)
	Update(task *models.Task) error
	Delete(id string) error
}
