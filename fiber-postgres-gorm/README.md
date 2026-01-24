# Todo API (Go + Fiber + Postgres + GORM)

A small **REST Todo API** built with:

- **Go** + **Fiber**
- **PostgreSQL**
- **GORM** (ORM) + Postgres driver

## Features

- **CRUD Todos**: create, list, get-by-id, update, delete
- **Auto-migrations** on startup using `db.AutoMigrate(&models.Todo{})`

## Project structure

- `cmd/server/`: app entrypoint
- `internal/routes/`: HTTP routes
- `internal/handlers/`: Fiber handlers (request/response)
- `internal/services/`: business logic (GORM calls)
- `internal/database/`: Postgres connection (GORM)
- `internal/models/`: GORM models
- `internal/config/`: env config loader

## Prerequisites

- **Go** (see `go.mod`)
- **Docker** (for running Postgres locally)

## Quickstart (local API + dockerized Postgres)

### 1) Start Postgres

From `fiber-postgres-gorm/`:

```bash
docker compose up -d
```

This starts Postgres on **localhost:5432** with:

- user: `postgres`
- password: `postgres`
- db: `todo_db`

### 2) Configure environment variables

Create a `.env` file in `fiber-postgres-gorm/` (the app loads it via `godotenv`):

```bash
cat > .env <<'EOF'
APP_PORT=8080

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=todo_db
EOF
```

### 3) Run the API

On startup, the app will run **GORM auto-migration** to create/update the `todos` table.

```bash
go run ./cmd/server
```

Server listens on `http://localhost:8080` (or your `APP_PORT`).

## API

Base path: `/api/todos`

### Create a todo

`POST /api/todos/`

Request body (example):

```json
{"title":"buy milk","completed":false}
```

```bash
curl -s -X POST http://localhost:8080/api/todos/ \
  -H 'Content-Type: application/json' \
  -d '{"title":"buy milk"}'
```

### List todos

`GET /api/todos/`

```bash
curl -s http://localhost:8080/api/todos/
```

### Get by id

`GET /api/todos/:id`

```bash
curl -s http://localhost:8080/api/todos/1
```

### Update a todo

`PUT /api/todos/:id`

```bash
curl -s -X PUT http://localhost:8080/api/todos/1 \
  -H 'Content-Type: application/json' \
  -d '{"title":"buy milk and eggs","completed":true}'
```

### Delete a todo

`DELETE /api/todos/:id`

```bash
curl -i -X DELETE http://localhost:8080/api/todos/1
```

## Notes / Troubleshooting

- **DB connection errors**: confirm env vars and that Postgres is running:

```bash
docker ps | grep todo_postgres
```

- **Table not created**: the app runs `AutoMigrate` on startup. If Postgres wasn’t reachable when the app started, restart the API after Postgres is up.


