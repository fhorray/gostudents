package router

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

var Students []Student



func Initialize() {
	router := gin.Default()

	router.GET("/students", getStudents)
	router.GET("/students/:id", getStudentsById)
	router.POST("/students/", createStudent)
	router.PUT("/students/:id", updateStudent)
	router.DELETE("/students/:id",deleteStudent)

	router.Run(":8080")
}


func getStudents(c *gin.Context) {
	c.JSON(http.StatusOK, Students)
}

func getStudentsById(c *gin.Context) {
	id := c.Param("id")
	var student Student
	if err := c.BindJSON(&student); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	for i, student := range Students {
		if student.ID == id {
			Students[i] = student
			c.JSON(http.StatusOK, student)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Student ID not found!"})
}

func createStudent(c *gin.Context) {
	var newStudent Student
	if err := c.BindJSON(&newStudent); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	newStudent.ID = uuid.New().String()

	Students = append(Students, newStudent)
	c.JSON(http.StatusCreated, newStudent)
}

func updateStudent(c *gin.Context) {
	id := c.Param("id")
	var updatedStudent Student
	if err := c.BindJSON(&updatedStudent); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	for i, student := range Students {
		if student.ID == id {
			Students[i] = updatedStudent
			c.JSON(http.StatusOK, updatedStudent)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Student not found!"})
}

func deleteStudent(c *gin.Context) {
	id := c.Param("id")
	var student Student
	if err := c.BindJSON(&student); err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		return
	}
	for i, student := range Students {
		if student.ID == id {
			Students = append(Students[:i], Students[i+1:]... )
			c.JSON(http.StatusOK, student)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Student not found!"})
}


