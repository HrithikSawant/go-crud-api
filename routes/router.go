package routes

import (
	"github.com/HrithikSawant/go-crud-api/handlers"
	"github.com/HrithikSawant/go-crud-api/routes/api"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, appHandler *handlers.ApplicationHandler) {
	api.RegisterRoutes(router, appHandler)
}
