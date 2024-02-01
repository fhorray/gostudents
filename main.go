package main

import (
	"studentsapi/config"
	"studentsapi/handler"
	"studentsapi/router"

	"github.com/google/uuid"
)



func main() {
	handler.Students = append(handler.Students, handler.Student{ID: uuid.New().String(), Name: "Francy", Age: 25, Grade: 8})
	handler.Students = append(handler.Students, handler.Student{ID: uuid.New().String(), Name: "marcos", Age: 27, Grade: 9})

logger := config.GetLogger("main")



	// Initialize configs
	db, err := config.Init()
	if err != nil {
		logger.Errf("Config initialization Error: %v", err)
		return
	}




	// Initialize Router
	router.Initialize()

}


