package entity

import (
	"errors"
	"time"

	"github.com/lanpaiva/api-user/pkg/models"
)

var (
	ErrIdisRequired          = errors.New("id is required")
	ErrIdInvalid             = errors.New("id is invalid")
	ErrNameisRequired        = errors.New("name is required")
	ErrDescriptionIsRequired = errors.New("Description is required")
)

type Task struct {
	ID          models.ID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewTask(name, description string) (*Task, error) {
	task := &Task{
		ID:          models.NewID(),
		Name:        name,
		Description: description,
		CreatedAt:   time.Now(),
	}
	err := task.ValidateTask()
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (t *Task) ValidateTask() error {
	if t.ID.String() == "" {
		return ErrIdisRequired
	}
	if _, err := models.ParseID(t.ID.String()); err != nil {
		return ErrIdInvalid
	}
	if t.Name == "" {
		return ErrNameisRequired
	}
	if t.Description == "" {
		return ErrDescriptionIsRequired
	}
	return nil
}
