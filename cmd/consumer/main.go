package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/lanpaiva/api-user/config"
	"github.com/lanpaiva/api-user/internal/entity"
	"github.com/lanpaiva/api-user/internal/infra/database"
	"github.com/lanpaiva/api-user/internal/infra/webserver/handlers"
	"github.com/lanpaiva/api-user/internal/usecase"
	"github.com/lanpaiva/api-user/pkg/models/kafka"
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

	user := database.NewUserDB(db)
	usecase := usecase.NewUser{UserDb: user}

	msgChanKafka := make(chan *ckafka.Message)
	topics := []string{"users"}
	servers := "host.docker.internal:9094"
	fmt.Println("Kafka consumer has started")
	go kafka.Consume(topics, servers, msgChanKafka)
	go kafkaWorker(msgChanKafka, usecase)

	http.ListenAndServe(":8000", r)
}

func kafkaWorker(msgChan chan *ckafka.Message, uc usecase.NewUser) {
	fmt.Println("Kafka worker has started")
	for msg := range msgChan {
		var CreateUserInput usecase.CreateUserOutput
		err := json.Unmarshal(msg.Value, &CreateUserInput)
		if err != nil {
			panic(err)
		}
		outputDto, err := uc.Create(usecase.CreateUserInput{
			Name:     "Alan Doe",
			Email:    "allanpaiva17@gmail.com",
			Password: "70567254Password",
		})
		if err != nil {
			panic(err)
		}
		fmt.Printf("Kafka has processed order %s\n", outputDto.Name)
	}
}
