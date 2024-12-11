package health

import (
	"github.com/HrithikSawant/go-crud-api/handlers"
	"github.com/gin-gonic/gin"
)

// RegisterHealthRoutes registers the health check routes
func RegisterHealthRoutes(v1Group *gin.RouterGroup, appHandler *handlers.ApplicationHandler) {
	healthHandler := appHandler.HealthHandler // Assuming HealthHandler is already initialized
	healthGroup := v1Group.Group("/health")   // /api/v1/health

	// Define the health check operation
	healthGroup.GET("/", healthHandler.HealthCheck) // Health check endpoint
}
