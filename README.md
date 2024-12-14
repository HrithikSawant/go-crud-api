# GO CRUD REST API

![Go CRUD REST API Logo](https://github.com/HrithikSawant/go-crud-api/blob/level-4/.assets/GoRest.jpeg)

![CI Status](https://img.shields.io/github/workflow/status/HrithikSawant/go-crud-api/ci?label=CI)
![Go Modules](https://img.shields.io/github/go-mod/go-version/HrithikSawant/go-crud-api)
[![GoDoc](https://pkg.go.dev/badge/github.com/HrithikSawant/go-crud-api)](https://pkg.go.dev/github.com/HrithikSawant/go-crud-api)
![Docker](https://img.shields.io/badge/Docker-%230db7ed?style=flat&logo=docker&logoColor=white)
[![Made with Gin](https://img.shields.io/badge/made_with-Gin-00b59c.svg)](https://gin-gonic.com/)
[![PostgreSQL Supported](https://img.shields.io/badge/PostgreSQL-Supported-4169e1.svg)](https://www.postgresql.org/)
[![Go Migration Supported](https://img.shields.io/badge/Go_Migrations-Supported-63d368.svg)](https://github.com/golang-migrate/migrate)
[![GORM Supported](https://img.shields.io/badge/GORM-Supported-9c1e5e.svg)](https://gorm.io/)
[![Make Supported](https://img.shields.io/badge/Make-Supported-2d3138.svg)](https://www.gnu.org/software/make/)
![License](https://img.shields.io/badge/license-MIT-green.svg)

## Overview

This branch introduces a Continuous Integration (CI) pipeline using **GitHub Actions** to automate the build, test, lint, and deployment process for the **Go CRUD API** project. It also integrates a **self-hosted GitHub runner** to run the pipeline locally on your machine. The CI pipeline is triggered on changes to the code directory and can be manually triggered as well.

## Features

- **CI Pipeline** using GitHub Actions for automating build, test, lint, Docker login, and Docker push.
- **Self-Hosted GitHub Runner** to run the pipeline on your local machine.
- **Docker Integration** for building and pushing the Go API Docker image.
- **Makefile** to simplify running the commands for building, testing, linting, Docker operations, and database migrations.

## Prerequisites

Ensure the following tools are installed on your system:
- **Docker** (for containerization)
- **Docker Compose** (for managing multi-container applications)
- **Make** (for automating commands)
- **GitHub Actions Runner** (for using a self-hosted runner on your local machine)

You can follow the installation instructions for these tools here:
- [Docker Installation](https://docs.docker.com/get-docker/)
- [Docker Compose Installation](https://docs.docker.com/compose/install/)
- [Make Installation](https://www.gnu.org/software/make/)
- [GitHub Actions Self-hosted Runner](https://docs.github.com/en/actions/hosting-your-own-runners)

## CI Pipeline Overview

The CI pipeline in this repository is defined using **GitHub Actions** and automates the following stages:

1. **Build API** - Uses the `make build` target to build the Go application.
2. **Run Tests** - Runs tests using the `make test` target.
3. **Perform Code Linting** - Lints the code using `make lint`.
4. **Docker Login** - Logs into Docker Hub or GitHub Docker Registry using the credentials stored in GitHub Secrets.
5. **Docker Build and Push** - Builds the Docker image and pushes it to a registry.

### GitHub Actions Workflow (`.github/workflows/ci.yml`)

The pipeline is defined in the `ci.yml` file and includes the following steps:

- **Checkout Code**: Pulls the latest code from the repository.
- **Set up Go Environment**: Configures the Go environment for the project.
- **Build the API**: Builds the Go application.
- **Run Tests**: Executes the tests.
- **Lint Code**: Runs the code linting.
- **Docker Login**: Logs into Docker using GitHub Secrets for the credentials.
- **Docker Build and Push**: Builds the Docker image and pushes it to Docker Hub or GitHub Container Registry.

### Manual Trigger

The CI pipeline can also be triggered manually through the GitHub Actions UI using the `workflow_dispatch` trigger.

## Setting up the CI Pipeline with Self-Hosted Runner

### Step 1: Setup the GitHub Actions Runner Locally

Follow the instructions to set up a **self-hosted GitHub runner** on your local machine. This allows the CI pipeline to run on your local machine instead of GitHub-hosted runners. 

### Step 2: Configure GitHub Secrets

To allow the pipeline to authenticate with Docker Hub or GitHub Container Registry, you need to configure the following secrets in your GitHub repository:

- **DOCKER_USERNAME**: Your Docker Hub or GitHub Container Registry username.
- **DOCKER_PASSWORD**: Your Docker Hub or GitHub Container Registry password or access token.

To set these secrets, go to **Settings > Secrets and variables > Actions** in your GitHub repository, and add the secrets accordingly.

## Building and Running with Docker

### 1. Build the Docker Image

Build a versioned Docker image using the Makefile:

```bash
make docker-build VERSION=<your_version>
```

For example:

```bash
make docker-build VERSION=1.0.0
```

### 2. Run the Docker Container

Run the API in a Docker container:

```bash
make docker-run VERSION=<your_version>
```

For example:

```bash
make docker-run VERSION=1.0.0
```

This command uses the `.env` file to inject environment variables into the container.

### 3. Stop the Docker Container

Stop the running container:

```bash
make docker-stop
```

### 4. Push Docker Image

Push the tagged image to a container registry:

```bash
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
| `docker-run`   | Run the application in a Docker container.                |
| `docker-push`  | Push the Docker image to a container registry.            |

## License

This project is licensed under the MIT License. See the LICENSE file for details.
