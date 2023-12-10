package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/lanpaiva/api-user/internal/entity"
	"github.com/lanpaiva/api-user/internal/infra/database"
	"github.com/lanpaiva/api-user/internal/usecase"
	IdPkg "github.com/lanpaiva/api-user/pkg/models"
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

func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	task, err := h.TaskDB.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 0
	}

	sort := r.URL.Query().Get("sort")
	tasks, err := h.TaskDB.FindAll(pageInt, limitInt, sort)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) TaskUpdate(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var task entity.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	task.ID, _ = IdPkg.ParseID(id)

	_, err = h.TaskDB.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	err = h.TaskDB.Update(&task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *TaskHandler) TaskDelete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err := h.TaskDB.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}
