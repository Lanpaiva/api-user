package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTask(t *testing.T) {
	task, err := NewTask("task 1", "description 1")
	assert.Nil(t, err)
	assert.NotNil(t, task)
	assert.NotEmpty(t, "task 1", task.Name)
	assert.NotEmpty(t, "description 1", task.Description)
	assert.NotEmpty(t, task.ID)
}

func TestValidate(t *testing.T) {
	task, err := NewTask("task 1", "description 1")
	assert.Nil(t, err)
	assert.NotEqual(t, task.ValidateTask(), task)
}
