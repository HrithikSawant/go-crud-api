# GO CRUD REST API

<<<<<<< HEAD
![Go CRUD REST API Logo](https://github.com/HrithikSawant/go-crud-api/blob/main/.assets/GoRest.jpeg)
=======
![Go CRUD REST API Logo](https://github.com/HrithikSawant/go-crud-api/raw/master/.assets/GoRest.jpeg)
>>>>>>> 8f83ee6 (go crud init)


[![GoDoc](https://pkg.go.dev/badge/github.com/HrithikSawant/go-crud-api)](https://pkg.go.dev/github.com/HrithikSawant/go-crud-api)
![Supported Go Versions](https://img.shields.io/badge/Go-1.20%2C%201.21%2C%201.22-lightgrey.svg)
[![Made with Gin](https://img.shields.io/badge/made_with-Gin-00b59c.svg)](https://gin-gonic.com/)
[![PostgreSQL Supported](https://img.shields.io/badge/PostgreSQL-Supported-4169e1.svg)](https://www.postgresql.org/)
[![Go Migration Supported](https://img.shields.io/badge/Go_Migrations-Supported-63d368.svg)](https://github.com/golang-migrate/migrate)
[![GORM Supported](https://img.shields.io/badge/GORM-Supported-9c1e5e.svg)](https://gorm.io/)
[![Make Supported](https://img.shields.io/badge/Make-Supported-2d3138.svg)](https://www.gnu.org/software/make/)
![License](https://img.shields.io/badge/license-MIT-green.svg)




## Overview
This repository provides a simple REST API for managing student records. The API is built using Golang, Gin framework, GORM for ORM, and PostgreSQL for database management.

## Features
The API supports the following operations:
- Add a new student
- Retrieve all students
- Retrieve a student by ID
- Update student information
- Delete a student record

Additionally, the API includes:
- API versioning (`api/v1`)
- A `/healthcheck` endpoint

## Prerequisites
Ensure the following are installed on your machine:
- [Golang](https://golang.org/dl/) (version 1.20 or later)
- [PostgreSQL](https://www.postgresql.org/)
- [Make](https://www.gnu.org/software/make/)
- [Migrate](https://github.com/golang-migrate/migrate) (postgresql version v4) or use Make install-go-migration

## Setup Instructions

### 1. Clone the Repository
```bash
git clone https://github.com/HrithikSawant/go-crud-api.git
cd go-crud-api
```

### 2. Set Up Environment Variables
- Refer `.env.example` file to create your own `.env` file in the root of the project.

| ENV_VARIABLE | Description |
| :-------- | :------------------------- |
| `DB_HOST` | Host for database connection 
| `DB_USER` | Database user
| `DB_PASSWORD` | Database password
| `DB_PORT` | Database port
| `DB_NAME` | Database name
| `SERVER_PORT` | Server port
| `SERVER_HOST` | Server Host
| `DB_MIGRATION_FILE` | DB migration file path

### 3. Run Database Migrations
Execute the database migrations to set up the `students` table:
```bash
make migrate-create:
```

### 5. Start the Server
Start the REST API server:
```bash
make run
```
The server will be accessible at `http://localhost:3000`.

## API Endpoints

### Base URL
`http://localhost:3000/api/v1`

### Endpoints

#### 1. Add a New Student
**POST** `/students`

**Request Body:**
```json
{
    "first_name": "Hrithik",
    "last_name": "Sawant",
    "email": "hrithikrandommail@gmail.com"
  }
```

#### 2. Get All Students
**GET** `/students`

#### 3. Get a Student by ID
**GET** `/students/{id}`

#### 4. Update a Student
**PUT** `/students/{id}`

**Request Body:**
```json
{
    "first_name": "Hrithik",
    "last_name": "Sawant",
    "email": "hrithikgolangdev@gmail.com"
  }
```

#### 5. Delete a Student
**DELETE** `/students/{id}`

#### 6. Health Check
**GET** `/healthcheck`

## Testing
Run the unit tests for the project:
```bash
make test-http
```

## License
This project is licensed under the MIT License. See the `LICENSE` file for details.
