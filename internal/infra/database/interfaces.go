package database

import (
	"github.com/lanpaiva/api-user/internal/entity"
)

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

type TaskInterface interface {
	Create(task *entity.Task) error
	FindAll(page, limit int, sort string) ([]entity.Task, error)
	FindById(id string) (*entity.Task, error)
	Update(task *entity.Task) error
	Delete(id string) error
}
