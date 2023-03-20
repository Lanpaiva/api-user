package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTask(t *testing.T) {
	task, err := NewTask("service")
	assert.Nil(t, err)
	assert.NotNil(t, task)
	assert.NotEmpty(t, task.Name)
	assert.NotEmpty(t, task.ID)
}

func TestValidate(t *testing.T) {
	task, err := NewTask("service")
	assert.Nil(t, err)
	assert.NotEqual(t, task.ValidateTask(), task)
}
