package v1

import (
	"github.com/HrithikSawant/go-crud-api/handlers"
	health_route "github.com/HrithikSawant/go-crud-api/routes/api/v1/health"
	student_route "github.com/HrithikSawant/go-crud-api/routes/api/v1/student"
	"github.com/gin-gonic/gin"
)

// RegisterV1Routes registers all version 1 routes for the API
func RegisterV1Routes(apiGroup *gin.RouterGroup, appHandler *handlers.ApplicationHandler) {
	v1Group := apiGroup.Group("/v1") // Grouping routes under /api/v1
	student_route.RegisterStudentRoutes(v1Group, appHandler)
	// Register health check routes
	health_route.RegisterHealthRoutes(v1Group, appHandler)
}
