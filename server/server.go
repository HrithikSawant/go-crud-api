package server

import (
	"log"

	"github.com/HrithikSawant/go-crud-api/handlers"
	"github.com/gin-gonic/gin"
)

// ServerConfig holds configuration for the server
type ServerConfig struct {
	Port    string
	Address string
}

// InitiateServer starts the Gin server with the given config and handler
func InitiateServer(config ServerConfig, appHandler *handlers.ApplicationHandler) {
	r := gin.Default()

	// Register routes through the handler
	appHandler.RegisterRoutes(r)

	// Start the server
	if err := r.Run(config.Address + ":" + config.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	} else {
		log.Printf("Server started successfully on %s:%s", config.Address, config.Port)
	}

}
