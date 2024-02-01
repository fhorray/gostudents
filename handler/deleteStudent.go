package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


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
