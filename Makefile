# Load environment variables from .env file
include .env
export $(shell sed 's/=.*//' .env)

# Install Go dependencies
install-go-gets:
	go get -u github.com/gin-gonic/gin
	go get -u gorm.io/gorm
	go get -u gorm.io/driver/postgres
	go get -u github.com/golang-migrate/migrate/v4/database/postgres@v4.18.1
	go get -u github.com/joho/godotenv

# Install migration tool
install-go-migration:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Run migrations up
migrate-up:
	migrate -path $(DB_MIGRATION_FILE) -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" up

# Run migrations down
migrate-down:
	migrate -path $(DB_MIGRATION_FILE) -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" down

migrate-drop:
	migrate -path $(DB_MIGRATION_FILE) -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" drop

# Run migrations down
migrate-version:
	migrate -path $(DB_MIGRATION_FILE) -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" version

migrate-force:
	migrate -path $(DB_MIGRATION_FILE) -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" force 1

migrate-create:
	migrate create -ext sql -dir $(DB_MIGRATION_FILE) -seq student_schema
	$(MAKE) migrate-up

# Test the HTTP CRUD operations
test-http:
	./test_http.sh

# Additional commands
run:
	go run .

build:
	go build -o app .

test:
	go test ./...

# Clean compiled binaries and temporary files
clean:
	rm -f app

# Start the server
start-server:
	go run main.go

# Format the code
format:
	go fmt ./...

# Lint the code
lint:
	golangci-lint run
