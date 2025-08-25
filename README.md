# 🚀 Modern Go REST API Template

[![Go Version](https://img.shields.io/badge/Go-1.24-00ADD8?style=for-the-badge&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/License-Apache-blue?style=for-the-badge)](LICENSE)
[![OpenAPI](https://img.shields.io/badge/OpenAPI-3.0-6BA539?style=for-the-badge&logo=openapiinitiative)](https://swagger.io/specification/)
[![SQLC](https://img.shields.io/badge/SQLC-Enabled-4285F4?style=for-the-badge)](https://sqlc.dev/)
[![Moon](https://img.shields.io/badge/Moon-Repo-FF6B6B?style=for-the-badge)](https://moonrepo.dev/)

This repository provides a production-ready starter template for building REST APIs in Go following a schema-first approach. It leverages code generation to create a robust and maintainable application structure with modern tooling and development workflows.

## ✨ Features

- 🐹 **Go** for the core application logic
- 📋 **OpenAPI-first development** with `oapi-codegen` to generate server stubs and models from an `openapi.yaml` spec
- 🛡️ **Type-safe database interaction** with `sqlc` to generate Go code from raw SQL queries
- 🔄 **Database migrations** with golang-migrate
- 🔍 **Comprehensive linting** with golangci-lint and sqruff for SQL
- 🎨 **Code formatting** with gofmt and SQL formatting tools
- 🌙 **Modern task management** with Moonrepo for consistent development workflows
- ⚙️ **Tool versioning management** with Moonrepo proto
- 🔧 **Environment-based configuration** using `envconfig` to read configuration from environment variables
- 📁 **Clear project structure** separating API definitions, business logic, database code, and configuration

## 📂 Project Structure

```
.
├── api/                        # API specifications
│   └── openapi.yml             # OpenAPI 3.0 specification
├── cmd/                        # Application entry points
│   └── server/                 # HTTP server application
│       └── main.go             # Main server entry point
├── configs/                    # Configuration files for tools
│   ├── golangci.yml            # Go linting configuration
│   ├── oapi-codegen-client.yml # OpenAPI client generation config
│   ├── oapi-codegen-server.yml # OpenAPI server generation config
│   ├── sqlc.yml                # SQLC configuration
│   └── sqruff.toml             # SQL linting and formatting config
├── db/                         # Database related files
│   ├── migrations/             # Database migration files
│   └── queries/                # SQL queries for SQLC
│       └── users.sql           # User-related queries
├── internal/                   # Private application code
│   ├── config/                 # Configuration management
│   │   └── config.go           # Environment configuration
│   ├── gen/                    # Generated code (auto-generated)
│   │   ├── db/                 # SQLC generated database code
│   │   └── oapi/               # OpenAPI generated server code
│   ├── repository/             # Data access layer
│   │   ├── errors.go           # Repository error definitions
│   │   └── user_repository.go  # User repository implementation
│   └── server/                 # HTTP server implementation
│       ├── middleware.go       # HTTP middleware
│       └── server.go           # Server setup and routing
├── pkg/                        # Public API packages
│   └── gen/                    # Generated client code
├── scripts/                    # Development scripts
│   ├── format.sh               # Code formatting script
│   ├── generate.sh             # Code generation script
│   ├── install.sh              # Tool installation script
│   └── lint.sh                 # Linting script
├── moon.yml                    # Moonrepo task configuration
├── Makefile                    # Make compatibility layer
└── go.mod                      # Go module definition
```

## ⚙️ Configuration

The application can be configured using the following environment variables:

| Variable        | Description                    | Default           | Required |
|-----------------|--------------------------------|-------------------|----------|
| LISTEN_ADDRESS  | The address to bind the server | localhost:8080    | No       |
| DATABASE_URL    | Database connection string     | -                 | Yes      |

## 🔧 Setup

This project uses [Moonrepo Proto](https://moonrepo.dev/proto) to manage tool versions consistently across environments. 

### Quick Start

1. **Install Proto** (if you haven't already):
   ```shell
   curl -fsSL https://moonrepo.dev/install/proto.sh | bash
   ```

2. **Install all required tools**:
   ```shell
   proto use
   ```
   
   Alternatively, use our installation script:
   ```shell
   bash scripts/install.sh
   ```

3. **Run the project** using Moonrepo tasks:
   ```shell
   moon :setup    # Install Go dependencies
   moon :generate # Generate code from OpenAPI and SQL
   moon :serve    # Start development server
   ```

### What Gets Installed

Proto will automatically install and manage the following tools at their specified versions:
- **Go** (1.24) - Core language runtime
- **Node.js** - For additional tooling dependencies
- **golangci-lint** - Go code linting
- **sqlc** - SQL to Go code generation
- **oapi-codegen** - OpenAPI to Go code generation
- **golang-migrate** - Database migration tool

All tool versions are locked in `.prototools` to ensure consistency across development environments.

### Make Compatibility

The `Makefile` is provided purely for compatibility with existing workflows and tooling. All Make targets simply delegate to the corresponding Moon tasks:

```bash
# These are equivalent:
make build    # Runs: moon :build
moon :build    # Direct Moon command
```

We recommend using Moon directly for the best development experience, but Make commands are available for teams transitioning from Make-based workflows.

## 🚀 Running the Application

### Prerequisites
1. Set up your environment variables (see Configuration section)
2. Ensure your database is running and accessible via `DATABASE_URL`

### Development
To run the application in development mode:

```bash
# Using Moon (recommended)
moon :serve

# Using Make (compatibility)
make serve

# Direct Go command
go run cmd/api/main.go
```

### Production Build
```bash
# Build for production
moon :build

# The binary will be available at: bin/api
```

### Database Setup
```bash
# Run database migrations
moon :migrate
```

## 🌐 API Endpoints

The API provides the following endpoints:

- `GET /users` - List all users
- `POST /users` - Create a new user
- `GET /users/{id}` - Get a user by ID
- `PUT /users/{id}` - Update a user
- `DELETE /users/{id}` - Delete a user

For complete API documentation, see the OpenAPI specification in `api/openapi.yml`.