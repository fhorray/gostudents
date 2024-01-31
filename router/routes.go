package router

import (
	"studentsapi/handler"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(route *gin.Engine) {
	v1 := route.Group("/api/v1/")
	{
			// Get All Student
			v1.GET("/students/", handler.GetStudents)

			// Get Student by ID
			v1.GET("/students/:id", handler.GetStudentsById)

			// Create Student
			v1.POST("/students/", handler.CreateStudent)

			// Update Student
			v1.PUT("/students/:id", handler.UpdatedStudent)

			// Delete Student
			v1.DELETE("/students/:id", handler.DeleteStudent)
	}
}
