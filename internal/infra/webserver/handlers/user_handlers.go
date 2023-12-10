package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/lanpaiva/api-user/internal/entity"
	"github.com/lanpaiva/api-user/internal/infra/database"
	"github.com/lanpaiva/api-user/internal/usecase"
)

type UserHandler struct {
	UserDB database.UserInterface
}

func NewUserHandler(userDB database.UserInterface) *UserHandler {
	return &UserHandler{UserDB: userDB}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userDto usecase.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&userDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u, err := entity.NewUser(userDto.Name, userDto.Email, userDto.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.UserDB.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
