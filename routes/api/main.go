package api

import (
	"github.com/HrithikSawant/go-crud-api/handlers"
	v1 "github.com/HrithikSawant/go-crud-api/routes/api/v1"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers all API routes, including versioned routes.
func RegisterRoutes(r *gin.Engine, appHandler *handlers.ApplicationHandler) {
	// Group all routes under /api
	apiGroup := r.Group("/api")
	v1.RegisterV1Routes(apiGroup, appHandler)
}
