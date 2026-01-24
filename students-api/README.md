# Students API

A REST API for managing students (CRUD) built with Go’s `net/http` and a **SQLite** storage backend (via `modernc.org/sqlite`).

## Run

From `students-api/`:

```bash
go run ./cmd/students-api
```

By default (see `config/local.yaml`) it starts on:

- `http://localhost:8080`

## Configuration

Config file: `config/local.yaml`

- `http_server.address`: server address
- `storage_path`: path to SQLite DB file (default: `storage/storage.db`)

## API

Base path: `/api/students`

- `POST /api/students` — create student
- `GET /api/students/{id}` — get by id
- `GET /api/students` — list all
- `PUT /api/students/{id}` — update
- `DELETE /api/students/{id}` — delete

### Create student (example)

```bash
curl -s -X POST http://localhost:8080/api/students \
  -H 'Content-Type: application/json' \
  -d '{"name":"Alice","email":"alice@example.com","age":21}'
```


