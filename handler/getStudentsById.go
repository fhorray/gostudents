package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
