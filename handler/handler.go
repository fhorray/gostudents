package handler

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

func GetStudents(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Students)
}

func GetStudentsById(ctx *gin.Context) {
	id := ctx.Param("id")
	for _, student := range Students {
		if student.ID == id {
			ctx.JSON(http.StatusOK, student)
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"message": "Student ID not found!"})
}

func CreateStudent(ctx *gin.Context) {
	var newStudent Student
	if err := ctx.BindJSON(&newStudent); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	newStudent.ID = uuid.New().String()

	Students = append(Students, newStudent)
	ctx.JSON(http.StatusCreated, newStudent)
}

func UpdatedStudent(ctx *gin.Context) {
	id := ctx.Param("id")
	var updatedStudent Student
	if err := ctx.BindJSON(&updatedStudent); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	for i, student := range Students {
		if student.ID == id {
			Students[i] = updatedStudent
			ctx.JSON(http.StatusOK, updatedStudent)
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"message": "Student not found!"})
}

func DeleteStudent(ctx *gin.Context) {
	id := ctx.Param("id")
	for i, student := range Students {
		if student.ID == id {
			Students = append(Students[:i], Students[i+1:]... )
			ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"error": "Student not found!"})
}
