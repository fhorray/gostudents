package main

import (
	// "studentsapi/handler"
	"studentsapi/config"
	"studentsapi/router"
	// "github.com/google/uuid"
)

var (logger *config.Logger)


func main() {
	// handler.Students = append(handler.Students, handler.Student{ID: uuid.New().String(), Name: "Francy", Age: 25, Grade: 8})
	// handler.Students = append(handler.Students, handler.Student{ID: uuid.New().String(), Name: "marcos", Age: 27, Grade: 9})

	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errf("config initialization error: %v", err)
		return
	}


	// Initialize Router
	router.Initialize()

}


