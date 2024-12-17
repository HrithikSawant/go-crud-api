# GO CRUD REST API Deployment

![Go CRUD REST API Logo](https://github.com/HrithikSawant/go-crud-api/blob/level-4/.assets/GoRest.png)


![Go Modules](https://img.shields.io/github/go-mod/go-version/HrithikSawant/go-crud-api)
[![GoDoc](https://pkg.go.dev/badge/github.com/HrithikSawant/go-crud-api)](https://pkg.go.dev/github.com/HrithikSawant/go-crud-api)
![Docker](https://img.shields.io/badge/Docker-%230db7ed?style=flat&logo=docker&logoColor=white)
[![Made with Gin](https://img.shields.io/badge/made_with-Gin-00b59c.svg)](https://gin-gonic.com/)
[![PostgreSQL Supported](https://img.shields.io/badge/PostgreSQL-Supported-4169e1.svg)](https://www.postgresql.org/)
[![Go Migration Supported](https://img.shields.io/badge/Go_Migrations-Supported-63d368.svg)](https://github.com/golang-migrate/migrate)
[![GORM Supported](https://img.shields.io/badge/GORM-Supported-9c1e5e.svg)](https://gorm.io/)
![CI](https://img.shields.io/badge/CI-Passing-brightgreen)
![CD](https://img.shields.io/badge/CD-Active-brightgreen)
![Vagrant](https://img.shields.io/badge/Vagrant-Enabled-brightgreen)
![NGINX](https://img.shields.io/badge/NGINX-Running-brightgreen)
[![Make Supported](https://img.shields.io/badge/Make-Supported-2d3138.svg)](https://www.gnu.org/software/make/)
![License](https://img.shields.io/badge/license-MIT-green.svg)

---

## Overview

This repository demonstrates the **bare-metal deployment** of the **Go CRUD REST API** project using Vagrant. It includes the following services:

- **Nginx** (Port `8080`) - Handles incoming requests and performs load balancing.
- **2 API Instances** (Ports `8081` & `8082`) - Go REST API instances.
- **PostgreSQL Database** (Port `5432`) - For persistent data storage.

---

## Architecture

Below is the architecture diagram for reference:

![Deployment Diagram](https://github.com/HrithikSawant/go-crud-api/blob/level-5/.assets/bare-metal.png)

- Nginx acts as the **Proxy & load balancer** for the two API instances.
- Incoming HTTP requests hit port `8080` on Nginx.
- Nginx forwards the requests to API services (`8081`, `8082`).
- API services interact with a PostgreSQL database running on port `5432`.

---

## Prerequisites

Ensure the following tools are installed:

1. **Vagrant**: [Install Vagrant](https://www.vagrantup.com/docs/installation)
2. **VirtualBox**: For Vagrant to provision virtual machines. 

---

## Setup Instructions

### 1. Clone the Repository

```bash
git clone https://github.com/HrithikSawant/go-crud-api.git
cd go-crud-api
```

---

### 2. Vagrant Setup

Spin up the Vagrant box that acts as your production environment:

```bash
vagrant up
```

This command performs the following:
- Creates a Vagrant-managed virtual machine.
- Executes a **bash script** to install all required dependencies (Go, Nginx, PostgreSQL, etc.).

---

### 3. Nginx Configuration

The Nginx configuration for load balancing is located in `nginx/nginx.conf`. It ensures requests are forwarded to the API instances:

```nginx
events {}

http {
    upstream go-crud-app {
            # Point to both app instances (Load Balancing)
        server go_crud_api_1:3000; # Modify as per your need
        server go_crud_api_2:3001; # Modify as per your need
    }

    server {
        listen 80;

        location / {
            proxy_pass http://go-crud-app;  # Use upstream to forward to Go API
            proxy_set_header Host $host;
            proxy_set_header X-Real-IP $remote_addr;
            proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
            proxy_set_header X-Forwarded-Proto $scheme;
        }

        location /health {
            proxy_pass http://go-crud-app/health;  # Health check route
        }
    }
}

```

---

## Testing the Setup

1. **Access the API** via port `8080` on the Vagrant box:

   ```bash
   curl http://127.0.0.1:8080/health/
   ```

   You should receive a `200 OK` response.

2. Use **Postman** or a browser to test API endpoints exposed via Nginx load balancing.

---

## Accessing the Application

Once deployed, the application will be accessible as follows:

- **API Endpoint**: `http://127.0.0.1:8080/students/`
- **Health Check**: `http://127.0.0.1:8080/health/`

Nginx will load balance requests between the two API services running on ports `8081` and `8082`.

---

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.

---

With this setup:
- A Vagrant box acts as the production environment.
- Nginx handles load balancing for high availability.
- Dependencies are installed using a bash script.
