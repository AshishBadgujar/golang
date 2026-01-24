# Todo API (Go + Fiber + Postgres + sqlc)

A small **REST Todo API** built with:

- **Go** + **Fiber**
- **PostgreSQL**
- **sqlc** (typed queries) + **pgx/v5** (connection pool)

## Features

- **CRUD Todos**: create, list, get-by-id, update, delete
- **Typed DB layer** via `sqlc` generated code (`internal/db/sqlc`)
- **Graceful shutdown** on SIGINT/SIGTERM

## Project structure

- `cmd/server/`: app entrypoint
- `internal/routes/`: HTTP routes
- `internal/handlers/`: Fiber handlers (request/response)
- `internal/services/`: business logic
- `internal/database/`: Postgres connection (pgx pool)
- `internal/db/schema.sql`: DB schema
- `internal/db/queries/`: sqlc query files
- `internal/db/sqlc/`: generated code (do not hand-edit)

## Prerequisites

- **Go** (see `go.mod`)
- **Docker** (for running Postgres locally)
- Optional: **sqlc** CLI (only needed if you change SQL and want to re-generate)

## Quickstart (local API + dockerized Postgres)

### 1) Start Postgres

From `fiber-postgres-sqlc/`:

```bash
docker compose up -d
```

This starts Postgres on **localhost:5432** with:

- user: `postgres`
- password: `postgres`
- db: `todo_db`

### 2) Create tables (apply schema)

```bash
docker exec -i todo_postgres psql -U postgres -d todo_db < internal/db/schema.sql
```

### 3) Configure environment variables

Create a `.env` file in `fiber-postgres-sqlc/` (the app loads it via `godotenv`):

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

### 4) Run the API

```bash
go run ./cmd/server
```

Server listens on `http://localhost:8080` (or your `APP_PORT`).

## API

Base path: `/api/todos`

### Create a todo

`POST /api/todos/`

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

## sqlc (generate typed queries)

SQL sources:

- schema: `internal/db/schema.sql`
- queries: `internal/db/queries/`
- config: `sqlc.yaml`

If you edit schema or queries, re-generate:

```bash
sqlc generate
```

Generated files go to `internal/db/sqlc/`.

## Notes / Troubleshooting

- **Schema not applied**: if endpoints error with missing table, run the schema command again:

```bash
docker exec -i todo_postgres psql -U postgres -d todo_db < internal/db/schema.sql
```

- **DB connection errors**: confirm env vars and that Postgres is running:

```bash
docker ps | grep todo_postgres
```
