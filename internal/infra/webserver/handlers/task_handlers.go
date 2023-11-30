package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/lanpaiva/api-user/internal/entity"
	"github.com/lanpaiva/api-user/internal/infra/database"
	"github.com/lanpaiva/api-user/internal/usecase"
)

type TaskHandler struct {
	TaskDB database.TaskInterface
}

func NewTaskHandler(db database.TaskInterface) *TaskHandler {
	return &TaskHandler{
		TaskDB: db,
	}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var taskDto usecase.CreateTaskInput
	err := json.NewDecoder(r.Body).Decode(&taskDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	t, err := entity.NewTask(taskDto.Name, taskDto.Description)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	err = h.TaskDB.Create(t)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
