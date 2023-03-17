package models

import (
	"errors"
	"time"

	"github.com/lanpaiva/api-user/pkg/entity"
)

var (
	ErrIdisRequired   = errors.New("id is required")
	ErrIdInvalid      = errors.New("id is invalid")
	ErrNameisRequired = errors.New("name is required")
)

type Task struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func NewTask(name string) (*Task, error) {
	task := &Task{
		ID:        entity.NewID(),
		Name:      name,
		CreatedAt: time.Now(),
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
	if _, err := entity.ParseID(t.ID.String()); err != nil {
		return ErrIdInvalid
	}
	if t.Name == "" {
		return ErrNameisRequired
	}
	return nil
}
