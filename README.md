# Modern Go REST API
This repository provides a production-ready starter template for building REST APIs in Go following a schema-first approach. It leverages code generation to create a robust and maintainable application structure.

## Features
- **Go** for the core application logic.
- **OpenAPI-first development** with `oapi-codegen` to generate server stubs and models from an `openapi.yaml` spec.
- **Type-safe database interaction** with `sqlc` to generate Go code from raw SQL queries.
- **Clear project structure** separating API definitions, business logic, database code, and configuration.
- **Environment-based configuration** using `envconfig` to read configuration from environment variables.

## Configuration
The application can be configured using the following environment variables:

| Variable | Description | Default |
|----------|-------------|---------|
| PORT     | The port to run the HTTP server on | 8080 |
| HOST     | The host to bind the HTTP server to | localhost |

## Running the Application
To run the application with default settings:

```bash
go run cmd/main.go
```

To run the application with custom settings:

```bash
PORT=9000 HOST=0.0.0.0 go run cmd/main.go
```

## API Endpoints
The API provides the following endpoints:

- `GET /users` - List all users
- `POST /users` - Create a new user
- `GET /users/{id}` - Get a user by ID
- `PUT /users/{id}` - Update a user
- `DELETE /users/{id}` - Delete a user

For more details, see the OpenAPI specification in `api/openapi.yml`.
