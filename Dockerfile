# Stage 1: Build the application
FROM golang:1.23-alpine as builder

# Set environment variables for Go build
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# Update and install necessary build tools, git, and a shell (bash)
RUN apk update && apk --no-cache add git build-base bash

# Install golang-migrate tool (for running migrations)
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Set the working directory inside the container
WORKDIR /app

# Copy Go modules manifests and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application source code
COPY . .

# Build the application binary
RUN go build -o app .

# Stage 2: Final runtime image with bash and migrate tool
FROM alpine:3.18

# Install bash and migrate tool in the runtime image
RUN apk --no-cache add bash

# Set working directory in runtime container
WORKDIR /app

# Copy the compiled binary and migrate tool from the builder stage
COPY --from=builder /app/app .
COPY --from=builder /go/bin/migrate /usr/local/bin/migrate

# Copy environment and migration files
COPY .env /app/.env
COPY db/migrations /app/db/migrations

# Expose the application's port
EXPOSE 3000

# Command to run the application
CMD ["/app/app"]
