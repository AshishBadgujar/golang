# Todo App (Go + Fiber + Postgres + sqlc + React)

A full-stack Todo app with:

- **Backend**: Go + Fiber + PostgreSQL + sqlc (pgx)
- **Frontend**: React + Vite + TypeScript

## Features

- **CRUD Todos** (via API + UI)
- **PATCH toggle** for completion status
- **Typed DB layer** via `sqlc` generated code
- **CORS enabled** for local frontend dev

## Project structure

- `backend/`: Go/Fiber API
  - `cmd/server/`: app entrypoint
  - `internal/routes/`: HTTP routes
  - `internal/handlers/`: Fiber handlers
  - `internal/services/`: business logic
  - `internal/database/`: Postgres connection
  - `internal/db/schema.sql`: DB schema
  - `internal/db/queries/`: sqlc query files
  - `internal/db/sqlc/`: generated code (do not hand-edit)
- `frontend/`: React UI (Vite)

## Prerequisites

- **Go** (see `backend/go.mod`)
- **Node.js** (for the React frontend)
- **Docker** (for running Postgres locally)
- Optional: **sqlc** CLI (only needed if you change SQL and want to re-generate queries)

## Quickstart (local dev)

### 1) Start Postgres (Docker)

From `fiber-postgres-sqlc/backend/`:

```bash
docker compose up -d
```

This starts Postgres on **localhost:5432** with:

- user: `postgres`
- password: `postgres`
- db: `todo_db`

### 2) Create tables (apply schema)

```bash
docker exec -i todo_postgres psql -U postgres -d todo_db < backend/internal/db/schema.sql
```

### 3) Configure backend environment variables

Create a `.env` file in `fiber-postgres-sqlc/backend/` (the API loads it via `godotenv`):

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

### 4) Run the backend API

```bash
cd backend
go run ./cmd/server
```

Backend listens on `http://localhost:8080` (or your `APP_PORT`).

### 5) Run the frontend (React)

From `fiber-postgres-sqlc/frontend/`:

```bash
cd frontend
npm install
npm run dev
```

Vite will print the local URL (commonly `http://localhost:5173`).

Note: the frontend currently calls the API at `http://localhost:8080/api` (see `frontend/src/api/todos.ts`).

## Backend API

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

- schema: `backend/internal/db/schema.sql`
- queries: `backend/internal/db/queries/`
- config: `backend/sqlc.yaml`

If you edit schema or queries, re-generate:

```bash
cd backend
sqlc generate
```

Generated files go to `backend/internal/db/sqlc/`.

## Notes / Troubleshooting

- **Schema not applied**: if endpoints error with missing table, run the schema command again:

```bash
docker exec -i todo_postgres psql -U postgres -d todo_db < backend/internal/db/schema.sql
```

- **DB connection errors**: confirm env vars and that Postgres is running:

```bash
docker ps | grep todo_postgres
```
