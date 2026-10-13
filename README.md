# Go Service Starter

A REST API built with Go's standard `net/http` package. The project demonstrates HTTP handlers, middleware, repository interfaces, in-memory storage, CRUD operations, and unit testing.

## Features

- Health check endpoint
- Create, list, retrieve, update, and delete tasks
- JSON request and response handling
- Request logging middleware
- Input validation and HTTP error handling
- Thread-safe in-memory repository
- Automated HTTP handler tests

## Requirements

- Go 1.25 or compatible version
- Docker (for containerized execution)

## Run locally

```bash
go run .
```

The server starts on port `8080`.

## API Endpoints

| Method | Endpoint | Purpose |
|---|---|---|
| GET | `/health` | Health check |
| POST | `/tasks` | Create a task |
| GET | `/tasks` | List tasks |
| GET | `/tasks/{id}` | Retrieve a task |
| PUT | `/tasks/{id}` | Update a task |
| DELETE | `/tasks/{id}` | Delete a task |

## Run tests

```bash
go test -v ./...
go test -race ./...
go vet ./...
```

## Example request

```bash
curl.exe -X POST "http://localhost:8080/tasks" -H "Content-Type: application/json" -d "{\"title\":\"Learn Go REST APIs\"}"
```

## Current limitations

Task data is stored in memory and is lost when the application restarts. Database persistence, authentication, and production deployment are outside the scope of this starter project.