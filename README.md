# 🚀 User with DOB and Calculated Age API - Go Backend

<div align="center">

![Go Version](https://img.shields.io/badge/Go-1.24+-00ADD8?style=for-the-badge&logo=go)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15+-316192?style=for-the-badge&logo=postgresql)
![Fiber](https://img.shields.io/badge/Fiber-v2-00ACD7?style=for-the-badge)
![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)

A production-ready RESTful API built with Go, featuring user management with dynamic age calculation. Built with modern Go best practices, clean architecture, and type-safe database operations.

[Features](#-features) • [Quick Start](#-quick-start) • [API Documentation](#-api-documentation) • [Architecture](#-architecture) • [Contributing](#-contributing)

</div>

---

## 📋 Table of Contents

- [Features](#-features)
- [Tech Stack](#-tech-stack)
- [Prerequisites](#-prerequisites)
- [Project Structure](#-project-structure)
- [Quick Start](#-quick-start)
- [Configuration](#-configuration)
- [API Documentation](#-api-documentation)
- [Development](#-development)
- [Testing](#-testing)
- [Deployment](#-deployment)
- [Troubleshooting](#-troubleshooting)
- [Contributing](#-contributing)
- [License](#-license)

---

## ✨ Features

- ✅ **CRUD Operations** - Complete user management (Create, Read, Update, Delete)
- 🎂 **Dynamic Age Calculation** - Automatically calculates age from date of birth
- 🚄 **High Performance** - Built with GoFiber for lightning-fast HTTP handling
- 🔒 **Type-Safe Database** - Uses SQLC for compile-time SQL query validation
- 📝 **Structured Logging** - Uber Zap logger with request tracking
- ✔️ **Input Validation** - Comprehensive validation using go-playground/validator
- 🐳 **Docker Ready** - Includes Docker and Docker Compose configurations
- 📊 **Pagination Support** - Efficient list endpoints with pagination
- 🎯 **Clean Architecture** - Follows industry-standard layered architecture
- 🔐 **Request ID Tracking** - Built-in middleware for request tracing

---

## 🛠 Tech Stack

| Component | Technology |
|-----------|------------|
| **Language** | Go 1.24+ |
| **Web Framework** | [GoFiber v2](https://gofiber.io/) |
| **Database** | PostgreSQL 15+ |
| **Query Builder** | [SQLC](https://sqlc.dev/) |
| **Validation** | [go-playground/validator](https://github.com/go-playground/validator) |
| **Logging** | [Uber Zap](https://github.com/uber-go/zap) |
| **Configuration** | [godotenv](https://github.com/joho/godotenv) |
| **Containerization** | Docker & Docker Compose |

---

## 📦 Prerequisites

Before you begin, ensure you have the following installed:

- **Go** 1.24 or later ([Download](https://golang.org/dl/))
- **PostgreSQL** 15+ ([Download](https://www.postgresql.org/download/))
- **Git** ([Download](https://git-scm.com/downloads))
- **SQLC** (Optional, for regenerating DB code)
  ```bash
  go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
  ```
- **Docker** (Optional, for containerized setup) ([Download](https://www.docker.com/products/docker-desktop))

---

## 📁 Project Structure

```
user-api-go/
├── cmd/
│   └── server/
│       └── main.go              # Application entry point
├── config/
│   └── config.go                # Configuration management
├── db/
│   ├── migrations/              # SQL migration files
│   │   └── 001_create_users_table.sql
│   ├── queries/                 # SQL queries for SQLC
│   │   └── users.sql
│   └── sqlc/                    # Generated type-safe DB code
├── internal/
│   ├── handler/                 # HTTP request handlers
│   │   └── user_handler.go
│   ├── repository/              # Database access layer
│   │   └── user_repository.go
│   ├── service/                 # Business logic layer
│   │   └── user_service.go
│   ├── routes/                  # Route definitions
│   │   └── routes.go
│   ├── middleware/              # HTTP middleware
│   │   └── logger_middleware.go
│   ├── models/                  # Data models and DTOs
│   │   └── user.go
│   └── logger/                  # Logger configuration
│       └── logger.go
├── .env                         # Environment variables (create this)
├── .gitignore                   # Git ignore rules
├── sqlc.yaml                    # SQLC configuration
├── go.mod                       # Go module dependencies
├── go.sum                       # Dependency checksums
└── README.md                    # This file
```

---

## 🚀 Quick Start

### Option 1: Local Setup (Recommended for Development)

#### 1️⃣ Clone the Repository

```bash
# Clone the repository
git clone https://github.com/mars-alien/user-api-go.git
cd user-api-go
```

#### 2️⃣ Set Up Environment Variables

Create a `.env` file in the root directory:

```env
# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=user_api_db

# Server Configuration
SERVER_PORT=3000
```

#### 3️⃣ Start PostgreSQL Database

**Using Docker (Easiest):**
```bash
docker run --name user-api-postgres \
  -e POSTGRES_PASSWORD=password \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_DB=user_api_db \
  -p 5432:5432 \
  -d postgres:15
```

**Or install PostgreSQL locally** and create the database:
```bash
createdb user_api_db
```

#### 4️⃣ Run Database Migration

```bash
# Using psql
psql -h localhost -p 5432 -U postgres -d user_api_db -f db/migrations/001_create_users_table.sql
```

**Windows (PowerShell):**
```powershell
Get-Content db\migrations\001_create_users_table.sql | psql -h localhost -p 5432 -U postgres -d user_api_db
```

**Or use a GUI tool** like pgAdmin, DBeaver, or TablePlus to execute the migration file.

#### 5️⃣ Install Dependencies

```bash
go mod download
```

#### 6️⃣ (Optional) Generate SQLC Code

Only needed if you modify SQL files in `db/queries/`:

```bash
sqlc generate
```

#### 7️⃣ Run the Application

```bash
# Build and run
go build -o bin/user-api ./cmd/server
./bin/user-api

# Or run directly
go run ./cmd/server
```

🎉 **Server is running!** Open [http://localhost:3000](http://localhost:3000)

---

### Option 2: Docker Compose (Recommended for Production)

#### 1️⃣ Clone and Configure

```bash
git clone https://github.com/mars-alien/user-api-go.git
cd user-api-go
```

Create `.env` file (same as above).

#### 2️⃣ Start Everything with Docker Compose

```bash
docker-compose up --build
```

This will:
- Build the Go application
- Start PostgreSQL database
- Run migrations automatically
- Start the API server

🎉 **Access the API at** [http://localhost:3000](http://localhost:3000)

To stop:
```bash
docker-compose down
```

---

## ⚙️ Configuration

All configuration is managed through environment variables. Create a `.env` file or set them in your environment.

### Environment Variables

| Variable | Description | Default | Required |
|----------|-------------|---------|----------|
| `DB_HOST` | PostgreSQL host | `localhost` | Yes |
| `DB_PORT` | PostgreSQL port | `5432` | Yes |
| `DB_USER` | Database user | `postgres` | Yes |
| `DB_PASSWORD` | Database password | - | Yes |
| `DB_NAME` | Database name | `user_api_db` | Yes |
| `DB_SSLMODE` | SSL mode for connection | `disable` | No |
| `SERVER_PORT` | HTTP server port | `3000` | No |

### Example .env File

```env
# Development
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=mypassword
DB_NAME=user_api_db
DB_SSLMODE=disable
SERVER_PORT=3000

# Production (example)
# DB_HOST=prod-db.example.com
# DB_PORT=5432
# DB_USER=app_user
# DB_PASSWORD=secure_password
# DB_NAME=user_api_prod
# DB_SSLMODE=require
# SERVER_PORT=8080
```

---

## 📚 API Documentation

Base URL: `http://localhost:3000`

### Endpoints

#### 1. Create User

Creates a new user with name and date of birth.

**Request:**
```http
POST /users
Content-Type: application/json

{
  "name": "Abhay",
  "dob": "1990-05-15"
}
```

**Response:** `201 Created`
```json
{
  "id": 1,
 "name": "Abhay",
  "dob": "1990-05-15"
}
```

**cURL Example:**
```bash
curl -X POST http://localhost:3000/users \
  -H "Content-Type: application/json" \
  -d '{"name": "Abhay","dob": "1990-05-15"}'
```

**PowerShell Example:**
```powershell
Invoke-RestMethod -Uri http://localhost:3000/users -Method POST -ContentType "application/json" -Body '{"name":"Alice Smith","dob":"1990-05-15"}'
```

---

#### 2. Get User by ID

Retrieves a single user with calculated age.

**Request:**
```http
GET /users/:id
```

**Response:** `200 OK`
```json
{
  "id": 1,
  "name": "Abhay",
  "dob": "1990-05-15",
  "age": 34
}
```

**cURL Example:**
```bash
curl http://localhost:3000/users/1
```

---

#### 3. List All Users

Retrieves all users with pagination support and calculated ages.

**Request:**
```http
GET /users?page=1&page_size=10
```

**Query Parameters:**
- `page` (optional): Page number, default `1`
- `page_size` (optional): Items per page, default `10`, max `100`

**Response:** `200 OK`
```json
[
  {
    "id": 1,
    "name": "Abhay",
    "dob": "1990-05-15",
    "age": 34
  },
  {
    "id": 2,
    "name": "Royal",
    "dob": "2003-08-20",
    "age": 22
  }
]
```

**cURL Example:**
```bash
curl "http://localhost:3000/users?page=1&page_size=10"
```

---

#### 4. Update User

Updates an existing user's information.

**Request:**
```http
PUT /users/:id
Content-Type: application/json

{
 "name": "Abhay singh",
  "dob": "1990-05-15"
}
```

**Response:** `200 OK`
```json
{
  "id": 1,
   "name": "Abhay singh",
    "dob": "1990-05-15",
}
```

**cURL Example:**
```bash
curl -X PUT http://localhost:3000/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "Abhay", "dob": "1990-05-15"}'
```

---

#### 5. Delete User

Deletes a user from the system.

**Request:**
```http
DELETE /users/:id
```

**Response:** `204 No Content`

**cURL Example:**
```bash
curl -X DELETE http://localhost:3000/users/1
```

---

### Error Responses

The API returns standard HTTP status codes and JSON error messages:

**400 Bad Request** - Invalid input
```json
{
  "error": "Validation failed: name is required"
}
```

**404 Not Found** - Resource not found
```json
{
  "error": "User not found"
}
```

**500 Internal Server Error** - Server error
```json
{
  "error": "Internal server error"
}
```

---

## 🏗 Architecture

This project follows **Clean Architecture** principles with clear separation of concerns:

### Layers

```
┌─────────────────────────────────────┐
│         HTTP Layer (Fiber)          │
│           handler/                  │
└─────────────────┬───────────────────┘
                  │
┌─────────────────▼───────────────────┐
│       Business Logic Layer          │
│           service/                  │
└─────────────────┬───────────────────┘
                  │
┌─────────────────▼───────────────────┐
│      Data Access Layer (SQLC)       │
│          repository/                │
└─────────────────┬───────────────────┘
                  │
┌─────────────────▼───────────────────┐
│          PostgreSQL Database        │
└─────────────────────────────────────┘
```

### Layer Responsibilities

1. **Handler Layer** (`internal/handler/`)
   - Receives HTTP requests
   - Validates input
   - Calls service layer
   - Returns HTTP responses

2. **Service Layer** (`internal/service/`)
   - Contains business logic
   - Orchestrates repository calls
   - Transforms data between layers
   - Handles business rules (e.g., age calculation)

3. **Repository Layer** (`internal/repository/`)
   - Database operations (CRUD)
   - Uses SQLC-generated type-safe queries
   - Returns domain models

4. **Models** (`internal/models/`)
   - Request/Response DTOs
   - Domain entities
   - Validation rules

### Why This Architecture?

✅ **Testability** - Each layer can be tested independently  
✅ **Maintainability** - Clear responsibilities  
✅ **Scalability** - Easy to add new features  
✅ **Flexibility** - Easy to swap implementations  

---

## 🔧 Development

### Modifying Database Schema

1. Create new migration file in `db/migrations/`:
```sql
-- db/migrations/002_add_email_to_users.sql
ALTER TABLE users ADD COLUMN email VARCHAR(255);
```

2. Update queries in `db/queries/users.sql`

3. Regenerate SQLC code:
```bash
sqlc generate
```

4. Update models and handlers accordingly

### Adding New Endpoints

1. **Add SQL queries** in `db/queries/`
2. **Generate SQLC code**: `sqlc generate`
3. **Add repository method** in `internal/repository/`
4. **Add service method** in `internal/service/`
5. **Add handler** in `internal/handler/`
6. **Register route** in `internal/routes/`

### Code Style

This project follows standard Go conventions:
- Use `gofmt` for formatting
- Follow [Effective Go](https://golang.org/doc/effective_go)
- Use meaningful variable names
- Add comments for exported functions

```bash
# Format code
go fmt ./...

# Run linter
go vet ./...

# Check for common issues
golangci-lint run
```

---


## 🚢 Deployment

### Using Docker

#### Build Image

```bash
docker build -t user-api-go:latest .
```

#### Run Container

```bash
docker run -d \
  --name user-api \
  -p 3000:3000 \
  -e DB_HOST=your-db-host \
  -e DB_PASSWORD=your-password \
  user-api-go:latest
```

### Using Docker Compose (Production)

```yaml
# docker-compose.prod.yml
version: '3.8'

services:
  app:
    build: .
    ports:
      - "3000:3000"
    environment:
      DB_HOST: postgres
      DB_PASSWORD: ${DB_PASSWORD}
    depends_on:
      - postgres
    restart: unless-stopped

  postgres:
    image: postgres:15
    environment:
      POSTGRES_PASSWORD: ${DB_PASSWORD}
    volumes:
      - postgres_data:/var/lib/postgresql/data
    restart: unless-stopped

volumes:
  postgres_data:
```

```bash
docker-compose -f docker-compose.prod.yml up -d
```

### Manual Deployment

```bash
# Build binary
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/user-api ./cmd/server

# Copy to server
scp bin/user-api user@server:/opt/user-api/

# Run with systemd or supervisor
```

---

## 🔍 Troubleshooting

### Common Issues

#### 1. Database Connection Failed

**Error:** `Failed to connect to database`

**Solutions:**
- Verify PostgreSQL is running: `docker ps` or `pg_isready`
- Check `.env` file has correct credentials
- Test connection: `psql -h localhost -U postgres -d user_api_db`
- Check firewall settings

#### 2. Port Already in Use

**Error:** `listen tcp :3000: bind: address already in use`

**Solutions:**
```bash
# Find process using port 3000
lsof -i :3000  # macOS/Linux
netstat -ano | findstr :3000  # Windows

# Kill the process or change SERVER_PORT in .env
```

#### 3. SQLC Generation Fails

**Error:** `sqlc: command not found`

**Solution:**
```bash
# Install SQLC
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

# Ensure $GOPATH/bin is in your PATH
export PATH=$PATH:$(go env GOPATH)/bin
```

#### 4. Migration Fails

**Error:** `relation "users" already exists`

**Solution:**
```bash
# Drop and recreate database
dropdb user_api_db
createdb user_api_db
psql -d user_api_db -f db/migrations/001_create_users_table.sql
```

#### 5. Validation Errors

**Error:** `Validation failed: dob must be in format YYYY-MM-DD`

**Solution:**
- Ensure date format is `YYYY-MM-DD` (e.g., `1990-05-15`)
- Name must be 1-100 characters
- All required fields must be present

### Debug Mode

Enable verbose logging:

```go
// In logger/logger.go
Log, err = zap.NewDevelopment() // Instead of NewProduction()
```

### Health Check

```bash
# Check if server is running
curl http://localhost:3000/users

# Check database connection
psql -h localhost -U postgres -d user_api_db -c "SELECT 1"
```

---

## 🤝 Contributing

Contributions are welcome! Please follow these guidelines:

### How to Contribute

1. **Fork the repository**
2. **Create a feature branch**
   ```bash
   git checkout -b feature/amazing-feature
   ```
3. **Make your changes**
4. **Run tests**
   ```bash
   go test ./...
   go fmt ./...
   go vet ./...
   ```
5. **Commit your changes**
   ```bash
   git commit -m "Add amazing feature"
   ```
6. **Push to your fork**
   ```bash
   git push origin feature/amazing-feature
   ```
7. **Open a Pull Request**

### Coding Standards

- Follow [Effective Go](https://golang.org/doc/effective_go)
- Write tests for new features
- Update documentation
- Use meaningful commit messages
- Keep functions small and focused

### Reporting Issues

Please include:
- Go version: `go version`
- Operating system
- Steps to reproduce
- Expected vs actual behavior
- Error messages/logs

---

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

```
MIT License

Copyright (c) 2024 Royal Sachan

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction...
```

---

## 📞 Support

- **Documentation:** [Go Documentation](https://golang.org/doc/)
- **Issues:** [GitHub Issues](https://github.com/mars-alien/user-api-go/issues)
- **Discussions:** [GitHub Discussions](https://github.com/mars-alien/user-api-go/discussions)

---

## 🙏 Acknowledgments

- [GoFiber](https://gofiber.io/) - Express-inspired web framework
- [SQLC](https://sqlc.dev/) - Type-safe SQL code generator
- [Uber Zap](https://github.com/uber-go/zap) - Blazing fast structured logger
- [PostgreSQL](https://www.postgresql.org/) - The world's most advanced open source database

---

<div align="center">

**Made with ❤️ using Go**

[⬆ Back to Top](#-user-management-api---go-backend)

</div>



