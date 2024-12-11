package main

import (
	"log"
	"os"

	db "github.com/HrithikSawant/go-crud-api/db"
	"github.com/HrithikSawant/go-crud-api/handlers"
	"github.com/HrithikSawant/go-crud-api/repository"
	"github.com/HrithikSawant/go-crud-api/server"
	services "github.com/HrithikSawant/go-crud-api/services"
	"github.com/joho/godotenv"
)

func main() {
	// Load enviroment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("DB_MIGRATION_FILE: %s", os.Getenv("DB_MIGRATION_FILE"))

	// Initiate and Migrate DB
	db, err := db.InitiateDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Initialize application repository
	applicationRepository, err := repository.NewApplicationRepository(db)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Initialize application service
	appService := services.NewApplicationService(applicationRepository)

	// Register handlers
	appHandler := handlers.NewApplicationHandler(appService)

	// Initiate server
	serverConfig := server.ServerConfig{
		Port:    os.Getenv("SERVER_PORT"),
		Address: os.Getenv("SERVER_HOST"),
	}
	server.InitiateServer(serverConfig, appHandler)

}
