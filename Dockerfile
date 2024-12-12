# Stage 1: Build the application
FROM golang:1.23-alpine as builder

# Set environment variables for Go build
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# Install necessary build tools and remove cache to reduce image size
RUN apk --no-cache add git

# Set the working directory inside the container
WORKDIR /app

# Copy Go modules manifests and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application source code
COPY . .

# Build the application binary
RUN go build -o app .

# Stage 2: Final runtime image
FROM gcr.io/distroless/static:nonroot

# Set working directory in runtime container
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/app .

# Copy only the necessary environment and migration files
COPY .env /app/.env
COPY db/migrations /app/db/migrations

# Expose the application's port
EXPOSE 3000

# Command to run the application
CMD ["/app/app"]



