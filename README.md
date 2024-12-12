# GO CRUD REST API

![Go CRUD REST API Logo](https://github.com/HrithikSawant/go-crud-api/blob/level-3/.assets/GoRest.jpeg)


[![GoDoc](https://pkg.go.dev/badge/github.com/HrithikSawant/go-crud-api)](https://pkg.go.dev/github.com/HrithikSawant/go-crud-api)
![Supported Go Versions](https://img.shields.io/badge/Go-1.20%2C%201.21%2C%201.22-lightgrey.svg)
![Docker](https://img.shields.io/badge/Docker-%230db7ed?style=flat&logo=docker&logoColor=white)
[![Made with Gin](https://img.shields.io/badge/made_with-Gin-00b59c.svg)](https://gin-gonic.com/)
[![PostgreSQL Supported](https://img.shields.io/badge/PostgreSQL-Supported-4169e1.svg)](https://www.postgresql.org/)
[![Go Migration Supported](https://img.shields.io/badge/Go_Migrations-Supported-63d368.svg)](https://github.com/golang-migrate/migrate)
[![GORM Supported](https://img.shields.io/badge/GORM-Supported-9c1e5e.svg)](https://gorm.io/)
[![Make Supported](https://img.shields.io/badge/Make-Supported-2d3138.svg)](https://www.gnu.org/software/make/)
![License](https://img.shields.io/badge/license-MIT-green.svg)



## Overview

This branch introduces a one-click local development setup with **Docker Compose** for running the Go API and its dependent services (e.g., PostgreSQL database). It aims to simplify the process of setting up the environment for developers by automating the setup using a `Makefile`, allowing for the start-up of the services with minimal configuration.

## Features

- **Docker Compose** for managing multiple containers (Go API and PostgreSQL).
- **Makefile** targets for setting up and running the containers with simple commands.
- **Automatic database migration** on container start.
- One-click setup to initialize and run the application.

## Prerequisites

Ensure the following tools are installed on your system:
- **Docker** (for containerization)
- **Docker Compose** (for managing multi-container applications)
- **Make** (for automating commands)

You can follow the installation instructions for these tools here:
- [Docker Installation](https://docs.docker.com/get-docker/)
- [Docker Compose Installation](https://docs.docker.com/compose/install/)
- [Make Installation](https://www.gnu.org/software/make/)

## Local Development Setup

### Step 1: Clone the Repository

First, clone the repository to your local machine:

```bash
git clone https://github.com/HrithikSawant/go-crud-api.git
cd go-crud-api
```

## Building and Running with Docker

### 1. Build the Docker Image

Build a versioned Docker image using the Makefile:

```
make docker-build VERSION=<your_version>
```

For example:

```
make docker-build VERSION=1.0.0
```

### 2. Run the Docker Container

Run the API in a container:

```
make docker-run VERSION=<your_version>
```

For example:

```
make docker-run VERSION=1.0.0
```

This command uses the .env file to inject environment variables into the container.

### 3. Stop the Docker Container

Stop the running container:

```
make docker-stop
```

### 4. Push Docker Image

Push the tagged image to a container registry:

```
make docker-push VERSION=<your_version>
```

Ensure you are logged in to your container registry before running this command.

## Multi-Stage Dockerfile

The multi-stage Dockerfile ensures:
- A smaller final image size by separating build and runtime stages.
- Faster builds and easier debugging.
- Security by excluding development tools and files from the runtime image.

## Updated Makefile Commands

| Command        | Description                                               |
|----------------|-----------------------------------------------------------|
| `docker-build` | Build a Docker image.                                     |
| `docker-run`   | Run the application in a Docker container.
| `docker-push`  | Push the Docker image to a container registry.            |

## License

This project is licensed under the MIT License. See the LICENSE file for details.
