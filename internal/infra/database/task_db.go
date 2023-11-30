package database

import (
	"github.com/lanpaiva/api-user/internal/entity"
	"gorm.io/gorm"
)

type TaskDB struct {
	DB *gorm.DB
}

func NewTaskDB(db *gorm.DB) *TaskDB {
	return &TaskDB{DB: db}
}

func (t *TaskDB) Create(task *entity.Task) error {
	return t.DB.Create(task).Error
}

func (t *TaskDB) FindAll(page, limit int, sort string) ([]entity.Task, error) {
	var tasks []entity.Task
	var err error

	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	if page != 0 && limit != 0 {
		err = t.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Find(&tasks).Error
	} else {
		err = t.DB.Order("created_at " + sort).Find(&tasks).Error
	}
	return tasks, err
}

func (t *TaskDB) FindById(id string) (*entity.Task, error) {
	var task entity.Task
	err := t.DB.First(&task, "id = ?", id).Error
	return &task, err
}

func (t *TaskDB) Update(task *entity.Task) error {
	_, err := t.FindById(task.ID.String())
	if err != nil {
		panic(err)
	}
	return t.DB.Save(task).Error
}

func (t *TaskDB) Delete(id string) error {
	task, err := t.FindById(id)
	if err != nil {
		return err
	}
	return t.DB.Delete(task).Error
}
