# Go CRUD App - Task Manager

A simple CRUD (Create, Read, Update, Delete) task manager built with Go. No external dependencies — just the Go standard library.

## Project Structure

```
intro/
├── main.go         # Backend: REST API with CRUD handlers
├── frontend.go     # Frontend: HTML/CSS/JS embedded in Go
├── go.mod          # Go module definition
└── README.md
```

## Running

```bash
go build -o crud-app .
./crud-app
```

Then open http://localhost:8080 in your browser.

## API Endpoints

| Method   | Endpoint          | Description       |
|----------|-------------------|-------------------|
| `GET`    | `/api/tasks`      | List all tasks    |
| `POST`   | `/api/tasks`      | Create a task     |
| `GET`    | `/api/tasks/:id`  | Get a single task |
| `PUT`    | `/api/tasks/:id`  | Update a task     |
| `DELETE` | `/api/tasks/:id`  | Delete a task     |

## API Examples

```bash
# Create a task
curl -X POST http://localhost:8080/api/tasks \
  -H 'Content-Type: application/json' \
  -d '{"title":"Learn Go"}'

# List all tasks
curl http://localhost:8080/api/tasks

# Mark a task as done
curl -X PUT http://localhost:8080/api/tasks/1 \
  -H 'Content-Type: application/json' \
  -d '{"done":true}'

# Delete a task
curl -X DELETE http://localhost:8080/api/tasks/1
```

## Key Go Concepts Covered

- **Structs & Methods** — `Task` struct with handler methods on `Store`
- **HTTP Server** — `net/http` for routing and serving
- **JSON** — `encoding/json` for request/response handling
- **Concurrency** — `sync.Mutex` for thread-safe in-memory storage
- **Pointers** — Used for partial updates (only update fields that are sent)
