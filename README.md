# Task Management API

This is a simple REST API for managing tasks, built with Go, SQLite, and Swagger for API documentation. The API allows you to create, read, update, and delete tasks. Each task has a title, description, and a completed status.

## Features

- Create a new task
- Get a list of all tasks
- Get a specific task by ID
- Update a task by ID
- Delete a task by ID
- SQLite database for persistent storage
- Swagger UI for API documentation

## Prerequisites

- **Go** (version 1.16 or higher recommended)
- **SQLite** (SQLite3 required for persistent storage)
- **Swagger** (for API documentation)

## Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/fuseraft/go-tasks.git
   cd go-tasks
   ```

2. **Install dependencies:**
   Make sure you have the required Go packages installed:

   ```bash
   go get github.com/gorilla/mux
   go get github.com/mattn/go-sqlite3
   go get -u github.com/swaggo/http-swagger
   go get -u github.com/swaggo/swag/cmd/swag
   ```

3. **Initialize Swagger documentation:**
   After the Go dependencies are installed, generate Swagger documentation:

   ```bash
   swag init
   ```

   This will create the `docs` folder in your project.

## Configuration

1. **Configure routes:**
   
   The `config.yaml` file contains configurable routes for the API. You can update this file to add or modify the existing routes.

   Example `config.yaml`:

   ```yaml
   routes:
     - path: /tasks
       method: POST
       handler: createTask
     - path: /tasks
       method: GET
       handler: getTasks
     - path: /tasks/{id}
       method: GET
       handler: getTask
     - path: /tasks/{id}
       method: PUT
       handler: updateTask
     - path: /tasks/{id}
       method: DELETE
       handler: deleteTask
   ```

2. **Database:**
   
   The API uses SQLite as the database. When you run the API for the first time, it will automatically create an `tasks.db` file in the root directory to store your tasks.

## Running the API

1. **Start the server:**
   
   Run the following command to start the API server:

   ```bash
   go run main.go
   ```

   The server will start at `http://localhost:8080`.

2. **Access Swagger UI:**
   
   After the server is running, visit `http://localhost:8080/swagger/index.html` to view the Swagger UI and test the API endpoints.

## API Endpoints

Here are the main API endpoints available:

- **POST /tasks** - Create a new task
- **GET /tasks** - Get all tasks
- **GET /tasks/{id}** - Get a task by ID
- **PUT /tasks/{id}** - Update a task by ID
- **DELETE /tasks/{id}** - Delete a task by ID

### Example JSON for Creating/Updating a Task:

```json
{
  "title": "Buy groceries",
  "description": "Milk, eggs, bread",
  "completed": false
}
```

## Example Usage

### 1. Create a new task:
```bash
curl -X POST http://localhost:8080/tasks -d '{"title":"Learn Go","description":"Complete Go course","completed":false}' -H "Content-Type: application/json"
```

### 2. Get all tasks:
```bash
curl http://localhost:8080/tasks
```

### 3. Get a specific task:
```bash
curl http://localhost:8080/tasks/1
```

### 4. Update a task:
```bash
curl -X PUT http://localhost:8080/tasks/1 -d '{"title":"Learn Go","description":"Finish Go tutorial","completed":true}' -H "Content-Type: application/json"
```

### 5. Delete a task:
```bash
curl -X DELETE http://localhost:8080/tasks/1
```

