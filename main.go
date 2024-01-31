package main

import (
	"studentsapi/router"

	"github.com/google/uuid"
)



func main() {
	router.Students = append(router.Students, router.Student{ID: uuid.New().String(), Name: "Francy", Age: 25, Grade: 8})
	router.Students = append(router.Students, router.Student{ID: uuid.New().String(), Name: "marcos", Age: 27, Grade: 9})

	router.Initialize()

}


