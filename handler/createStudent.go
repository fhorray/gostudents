package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

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
