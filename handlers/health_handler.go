package handlers

import (
	"net/http"

	service "github.com/HrithikSawant/go-crud-api/services"
	"github.com/gin-gonic/gin"
)

// HealthHandler handles health check requests.
type HealthHandler struct {
	healthService *service.HealthService
}

// NewHealthHandler creates a new instance of HealthHandler.
func NewHealthHandler(healthService *service.HealthService) *HealthHandler {
	return &HealthHandler{healthService: healthService}
}

// RegisterRoutes registers all routes for the health resource.
func (h *HealthHandler) RegisterRoutes(r *gin.Engine) {
	healthRoutes := r.Group("/health") // Grouping the health check routes
	{
		// Health check route to check the application status
		healthRoutes.GET("/", h.HealthCheck)
	}
}

// HealthCheck handles health check requests and returns the health status.
func (h *HealthHandler) HealthCheck(c *gin.Context) {
	// Get the health status and error from the HealthService
	healthStatus, err := h.healthService.CheckHealth()
	if err != nil {
		// If there's an error (e.g., database connection issue), return a 500 status
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  healthStatus.Status,
			"message": healthStatus.Message,
		})
		return
	}

	// If there's no error, return the healthy status
	c.JSON(http.StatusOK, healthStatus)
}
