package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Student representa a estrutura de dados de um estudante
type Student struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Grade int    `json:"grade"`
}

var students []Student

func main() {
	students = append(students, Student{ID: "1", Name: "Francy", Age: 25, Grade: 8})
	students = append(students, Student{ID: "2", Name: "marcos", Age: 27, Grade: 9})

	router := gin.Default()

	router.GET("/students", getStudents)
	router.POST("/students/", createStudent)

	router.Run(":8080")
}

func getStudents(c *gin.Context) {
	c.JSON(http.StatusOK, students)
}

func createStudent(c *gin.Context) {
	var newStudent Student
	if err := c.BindJSON(&newStudent); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	newStudent.ID = uuid.New().String()

	students = append(students, newStudent)
	c.JSON(http.StatusCreated, newStudent)
}

