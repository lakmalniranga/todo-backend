# Todo API Backend

Hello, you can use this project as a classic Todo API backend for experiments and learning purposes. I personally use it for GitOps testing. Oh yes, code is AI generated, but it has been reviewed by myself.

## Features

- Create, Read, Update, Delete (CRUD) operations for todos
- MongoDB database for persistent storage
- Dockerized application with docker-compose for easy deployment
- Environment variable configuration
- Kubernetes deployment with liveness and readiness probes

## Prerequisites

- [Go](https://golang.org/dl/) (version 1.20 or later)
- [Docker](https://docs.docker.com/get-docker/) and [Docker Compose](https://docs.docker.com/compose/install/)
- [MongoDB](https://www.mongodb.com/try/download/community) (if running locally without Docker)

## Project Structure

```
todo-backend/
├── configs/       # Database configuration
├── controllers/   # Request handlers
├── models/        # Data models
├── routes/        # API routes
├── Dockerfile     # Docker image definition
├── docker-compose.yml  # Docker Compose definition
├── go.mod         # Go module file
├── go.sum         # Go dependencies checksum
├── main.go        # Application entry point
└── README.md      # This file
```

## Environment Variables

The application can be configured using the following environment variables:

- `PORT`: Server port (default: `8080`)
- `MONGO_URI`: MongoDB connection string (default: `mongodb://localhost:27017`)
- `MONGO_DB_NAME`: MongoDB database name (default: `todoDB`)

You can set these variables in multiple ways:
- Create a `.env` file in the project root
- Set them in your shell before running the application
- Define them in the `docker-compose.yml` file for containerized deployment

## Running Locally

### Without Docker

1. Start MongoDB locally on port 27017
2. Clone the repository
3. Navigate to the project directory
4. Configure environment variables (optional):
   - Create a `.env` file based on `.env.example`
   - Customize the MongoDB connection and other settings
5. Run the application:

```bash
cd todo-backend
go mod download
go run main.go
```

### With Docker Compose

1. Clone the repository
2. Navigate to the project directory
3. Configure environment variables (optional):
   - Edit the `environment` section in `docker-compose.yml`
   - Or create a `.env` file (Docker Compose will use it automatically)
4. Build and start the containers:

```bash
cd todo-backend
docker-compose up -d
```

## API Endpoints

| Method | URL              | Description       | Request Body                                | Status Codes    |
|--------|------------------|-------------------|---------------------------------------------|-----------------|
| POST   | `/api/todos`     | Create a todo     | `{"title":"Task","description":"Details","completed":false}` | 201, 400, 500 |
| GET    | `/api/todos`     | Get all todos     | -                                           | 200, 500        |
| GET    | `/api/todos/{id}`| Get a todo by ID  | -                                           | 200, 404, 500   |
| PUT    | `/api/todos/{id}`| Update a todo     | `{"title":"Updated","description":"New details","completed":true}` | 200, 404, 500 |
| DELETE | `/api/todos/{id}`| Delete a todo     | -                                           | 200, 404, 500   |

## Health Check Endpoints

- `GET /health`: Returns status 200 OK with message "API is running" if the server is up (legacy)
- `GET /health/live`: Liveness probe for Kubernetes - checks if the server is running
- `GET /health/ready`: Readiness probe for Kubernetes - checks if the application is ready to receive traffic (including MongoDB connection)

## Building and Running Tests

### Build

```bash
go build -o todo-api
```

### Run

```bash
./todo-api
```

## Docker Commands

### Build the Docker Image

```bash
docker build -t todo-api .
```

### Run the Docker Container

```bash
docker run -p 8080:8080 -e MONGO_URI=mongodb://host.docker.internal:27017 -e MONGO_DB_NAME=todoDB todo-api
```

## Docker Compose Commands

### Start Services

```bash
docker-compose up -d
```

### Stop Services

```bash
docker-compose down
```

### View Logs

```bash
docker-compose logs -f
```

## License

This project is licensed under the MIT License.
