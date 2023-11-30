package database

import (
	"fmt"
	"testing"

	"github.com/lanpaiva/api-user/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestNewTask(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Task{})
	task, err := entity.NewTask("task 1", "description 1")
	assert.NoError(t, err)
	taskDB := NewTaskDB(db)

	err = taskDB.Create(task)
	assert.NoError(t, err)
	assert.NotEmpty(t, task.ID)
}

func TestFindAllProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Task{})
	for i := 1; i < 24; i++ {
		task, err := entity.NewTask(fmt.Sprintf("task %d", i), fmt.Sprintf("description %d", i))
		assert.NoError(t, err)
		db.Create(task)
	}
	TaskDB := NewTaskDB(db)
	tasks, err := TaskDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)

	// for i, task := range tasks {
	// 	fmt.Printf("Task %d: %s\n", i+1, task.Name)
	// }
	assert.Len(t, tasks, 10)
	assert.Equal(t, "task 1", tasks[0].Name)
	assert.Equal(t, "task 10", tasks[9].Name)

	tasks, err = TaskDB.FindAll(2, 10, "asc")
	assert.NoError(t, err)

	// for i, task := range tasks {
	// 	fmt.Printf("Task %d: %s\n", i+1, task.Name)
	// }
	assert.Len(t, tasks, 10)
	assert.Equal(t, "task 11", tasks[0].Name)
	assert.Equal(t, "task 20", tasks[9].Name)

	tasks, err = TaskDB.FindAll(3, 10, "asc")
	assert.NoError(t, err)

	// for i, task := range tasks {
	// 	fmt.Printf("Task %d: %s\n", i+1, task.Name)
	// }
	assert.Len(t, tasks, 3)
	assert.Equal(t, "task 21", tasks[0].Name)
	assert.Equal(t, "task 23", tasks[2].Name)
}

func TestFindById(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Task{})
	task, err := entity.NewTask("task 1", "description 1")
	assert.NoError(t, err)
	assert.Equal(t, "task 1", task.Name)
	db.Create(task)
	TaskDB := NewTaskDB(db)
	task, err = TaskDB.FindById(task.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, "task 1", task.Name)
}

func TestUpdateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Task{})
	task, err := entity.NewTask("task 1", "description 1")
	assert.NoError(t, err)
	db.Create(task)
	TaskDB := NewTaskDB(db)
	task.Name = "task 2"
	err = TaskDB.Update(task)
	assert.NoError(t, err)
	task, err = TaskDB.FindById(task.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, "task 2", task.Name)
}

func TestDeleteProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Task{})
	task, err := entity.NewTask("task 1", "description 1")
	assert.NoError(t, err)
	db.Create(task)
	TaskDB := NewTaskDB(db)

	err = TaskDB.Delete(task.ID.String())
	assert.NoError(t, err)

	_, err = TaskDB.FindById(task.ID.String())
	assert.Error(t, err)

}
