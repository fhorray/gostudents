package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func initializeRoutes(route *gin.Engine) {
	v1 := route.Group("/api/v1/")
	{
			// Get All Student
			v1.GET("/students/", func (ctx *gin.Context)  {
				ctx.JSON(http.StatusOK, Students)
			})

			// Get Student by ID
			v1.GET("/students/:id", func(ctx *gin.Context) {
				id := ctx.Param("id")
				var student Student
				if err := ctx.BindJSON(&student); err != nil {
					ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
					return
				}
				for i, student := range Students {
					if student.ID == id {
						Students[i] = student
						ctx.JSON(http.StatusOK, student)
						return
					}
				}
				ctx.JSON(http.StatusNotFound, gin.H{"message": "Student ID not found!"})
			})

			// Create Student
			v1.POST("/students/", func(ctx *gin.Context) {
				var newStudent Student
				if err := ctx.BindJSON(&newStudent); err != nil {
					ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
					return
				}

				newStudent.ID = uuid.New().String()

				Students = append(Students, newStudent)
				ctx.JSON(http.StatusCreated, newStudent)
			})

			// Update Student
			v1.PUT("/students/:id", func(ctx *gin.Context) {
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
			})

			// Delete Student
			v1.DELETE("/students/:id", func(ctx *gin.Context) {
				id := ctx.Param("id")
				var student Student
				if err := ctx.BindJSON(&student); err != nil {
					ctx.JSON(http.StatusBadGateway, err.Error())
					return
				}
				for i, student := range Students {
					if student.ID == id {
						Students = append(Students[:i], Students[i+1:]... )
						ctx.JSON(http.StatusOK, student)
						return
					}
				}
				ctx.JSON(http.StatusNotFound, gin.H{"error": "Student not found!"})
			})
	}
}
