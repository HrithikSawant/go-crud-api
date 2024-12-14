// Package routes defines the application's HTTP routes and request handling.
package routes

import (
	"github.com/HrithikSawant/go-crud-api/handlers"
	"github.com/HrithikSawant/go-crud-api/routes/api"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up all the routes for the application's API, binding them to the appropriate handlers.
func RegisterRoutes(router *gin.Engine, appHandler *handlers.ApplicationHandler) {
	api.RegisterRoutes(router, appHandler)
}
