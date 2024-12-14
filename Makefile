# Load environment variables from .env file
include .env
export $(shell sed 's/=.*//' .env)

# Variables
APP_NAME := go-crud-api
VERSION := 1.0.0
DOCKER_IMAGE := $(APP_NAME):$(VERSION)

# Default Target
.PHONY: all
all: build

# Build the Go application locally
.PHONY: build
build:
	@echo "Building the Go application..."
	@if [ ! -d ./bin ]; then mkdir ./bin; fi
	go build -o bin/$(APP_NAME) .

# Run the application locally
.PHONY: run
run: build
	@echo "Running the application..."
	./bin/$(APP_NAME)

# Test the application
.PHONY: test
test:
	@echo "Running tests..."
	./test_http.sh

# Docker Targets
# Build the Docker image
.PHONY: docker-build
docker-build:
	@echo "Building the Docker image..."
	docker build -t $(DOCKER_IMAGE) .

# Run the Docker container
.PHONY: docker-run
docker-run:
	@echo "Running the Docker container..."
	docker run -it --network=host $(DOCKER_IMAGE)

# Push Docker image to a registry
.PHONY: docker-push
docker-push:
	@echo "Pushing the Docker image to the registry..."
	docker push $(DOCKER_IMAGE)

# Docker Clean (optional if needed)
.PHONY: docker-clean
docker-clean:
	@echo "Removing dangling Docker images..."
	docker image prune -f

# Run database migrations
# Install migration tool
migrate-install:
	@echo "Installing PostgreSQL v4 database migrations..."
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

migrate-create:
	migrate create -ext sql -dir $(DB_MIGRATION_FILE) -seq student_schema
	$(MAKE) migrate-up

.PHONY: migrate-up
migrate-up:
	@echo "Running database migrations..."
	migrate -path $(DB_MIGRATION_FILE) -database $(DATABASE_URL) up

# Rollback database migrations
.PHONY: migrate-down
migrate-down:
	@echo "Rolling back database migrations..."
	migrate -path $(DB_MIGRATION_FILE) -database $(DATABASE_URL) down

.PHONY: migrate-drop
migrate-drop:
	@echo "Dropping database migrations..."
	migrate -path $(DB_MIGRATION_FILE) -database $(DATABASE_URL) drop

# Run migrations down
migrate-version:
	@echo "Printing database migrations version..."
	migrate -path $(DB_MIGRATION_FILE) -database $(DATABASE_URL) version

# Run migrations down
migrate-force:
	@echo "forcing database migrations to 1..."
	migrate -path $(DB_MIGRATION_FILE) -database $(DATABASE_URL) force 1

# Run the Docker container using Docker Compose
.PHONY: docker-compose-up
docker-compose-up:
	@echo "Starting the app and db containers with Docker Compose..."
	docker-compose up --build

# Stop the Docker containers using Docker Compose
.PHONY: docker-compose-down
docker-compose-down:
	@echo "Stopping the app and db containers with Docker Compose..."
	docker-compose down

# Lint the code
.PHONY: lint
lint:
	@echo "Linting the Go code..."
	golint ./...

# Help message
.PHONY: help
help:
	@echo "Available commands:"
	@echo "  build               - Build the Go application locally"
	@echo "  run                 - Run the application locally"
	@echo "  test                - Run tests"
	@echo "  clean               - Clean up build artifacts"
	@echo "  docker-build        - Build the Docker image"
	@echo "  docker-run          - Run the Docker container"
	@echo "  docker-push         - Push the Docker image to a registry"
	@echo "  docker-clean        - Remove dangling Docker images"
	@echo "  docker-compose-up   - Start app and db containers using Docker Compose"
	@echo "  docker-compose-down - Stop app and db containers using Docker Compose"
	@echo "  migrate-create      - Generate a new database migration"
	@echo "  migrate-up          - Run database migrations"
	@echo "  migrate-down        - Rollback database migrations"
	@echo "  lint                - Lint the Go code"

# Install Go dependencies
install-go-gets:
	go get -u github.com/gin-gonic/gin
	go get -u gorm.io/gorm
	go get -u gorm.io/driver/postgres
	go get -u github.com/golang-migrate/migrate/v4/database/postgres@v4.18.1
	go get -u github.com/joho/godotenv






