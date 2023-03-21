package database

import (
	"github.com/lanpaiva/api-user/internal/models"
	"gorm.io/gorm"
)

type Task struct {
	DB *gorm.DB
}

func NewTask(db *gorm.DB) *Task {
	return &Task{DB: db}
}

func (t *Task) Create(task *models.Task) error {
	return t.DB.Create(task).Error
}

func (t *Task) FindById(id string) (*models.Task, error) {
	var task models.Task
	err := t.DB.First(&task, "id = ?", id).Error
	return &task, err
}

func (t *Task) Update(task *models.Task) error {
	_, err := t.FindById(task.ID.String())
	if err != nil {
		panic(err)
	}
	return t.DB.Save(task).Error
}

func (t *Task) Delete(id string) error {
	task, err := t.FindById(id)
	if err != nil {
		panic(err)
	}
	return t.DB.Delete(task).Error
}
