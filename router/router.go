package router

import (
	"github.com/gin-gonic/gin"
)

// Student representa a estrutura de dados de um estudante
type Student struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Grade int    `json:"grade"`
}

var Students []Student


func Initialize() {
	router := gin.Default()

	initializeRoutes(router)

	// router.GET("/students", getStudents)
	// router.GET("/students/:id", getStudentsById)
	// router.POST("/students/", createStudent)
	// router.PUT("/students/:id", updateStudent)
	// router.DELETE("/students/:id", deleteStudent)

	router.Run(":8080")
}
