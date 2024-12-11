package student

import (
	"github.com/HrithikSawant/go-crud-api/handlers"
	"github.com/gin-gonic/gin"
)

// RegisterStudentRoutes registers the CRUD routes for the student resource
func RegisterStudentRoutes(v1Group *gin.RouterGroup, appHandler *handlers.ApplicationHandler) {
	studentHandler := appHandler.StudentHandler // Assuming StudentHandler is already initialized
	studentGroup := v1Group.Group("/students")  // /api/v1/students

	// Define CRUD operations for the student resource
	studentGroup.POST("/", studentHandler.CreateStudent)      // Create student
	studentGroup.GET("/:id", studentHandler.GetStudentByID)   // Get student by ID
	studentGroup.PUT("/:id", studentHandler.UpdateStudent)    // Update student by ID
	studentGroup.DELETE("/:id", studentHandler.DeleteStudent) // Delete student by ID
}
