package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStudents(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Students)
}
