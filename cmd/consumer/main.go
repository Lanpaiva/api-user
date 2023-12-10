package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/lanpaiva/api-user/config"
	"github.com/lanpaiva/api-user/internal/entity"
	"github.com/lanpaiva/api-user/internal/infra/database"
	"github.com/lanpaiva/api-user/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	_, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Task{}, &entity.User{})

	taskDB := database.NewTaskDB(db)
	userDB := database.NewUserDB(db)

	taskHandler := handlers.NewTaskHandler(taskDB)
	userHandler := handlers.NewUserHandler(userDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/users", userHandler.CreateUser)

	r.Post("/tasks", taskHandler.CreateTask)
	r.Get("/tasks/{id}", taskHandler.GetTask)
	r.Get("/tasks", taskHandler.GetAllTasks)
	r.Put("/tasks/{id}", taskHandler.TaskUpdate)
	r.Delete("/tasks/{id}", taskHandler.TaskDelete)

	http.ListenAndServe(":8000", r)
}
