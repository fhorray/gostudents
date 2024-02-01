package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
