# Social (Go API skeleton)

A small Go API project scaffold using **chi** with:

- structured routing under `/v1`
- standard HTTP middleware (request id, logging, recovery, timeouts)
- `.env` support via `godotenv`

## Run

From `social/`:

```bash
go run ./cmd/api
```

By default it listens on `:8080`.

You can override with:

```bash
export ADDR=":8080"
go run ./cmd/api
```

## Endpoints

- `GET /v1/health` — returns `OK`

## Notes

- There are stubs for a Postgres-backed storage layer in `internal/store`, but the current `cmd/api` only wires routing/health.


