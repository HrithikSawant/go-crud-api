package handlers

import (
	services "github.com/HrithikSawant/go-crud-api/services"
	"github.com/gin-gonic/gin"
)

// ApplicationHandler struct holds all handlers, like student handler
type ApplicationHandler struct {
	StudentHandler *StudentHandler
	HealthHandler  *HealthHandler
}

// NewApplicationHandler initializes the ApplicationHandler with services
func NewApplicationHandler(appService *services.ApplicationService) *ApplicationHandler {
	return &ApplicationHandler{
		StudentHandler: NewStudentHandler(appService.StudentService),
		HealthHandler:  NewHealthHandler(appService.HealthService),
	}
}

// RegisterRoutes sets up routes for all the entities handled by this handler
func (h *ApplicationHandler) RegisterRoutes(r *gin.Engine) {
	// Register routes for student
	h.StudentHandler.RegisterRoutes(r)
	h.HealthHandler.RegisterRoutes(r)
}
