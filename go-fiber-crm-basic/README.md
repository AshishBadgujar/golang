# Fiber CRM (basic)

A basic CRM-style REST API built with **Go Fiber** and **GORM (SQLite)**. It manages `Lead` records and stores them in a local `leads.db` file.

## Run

From `go-fiber-crm-basic/`:

```bash
go run .
```

Server listens on `http://localhost:3000`.

On startup, it runs `AutoMigrate` to create/update the `leads` table in `leads.db`.

## API

- `GET /api/lead` — list all leads
- `GET /api/lead/:id` — get lead by id
- `POST /api/lead` — create a lead
- `DELETE /api/lead/:id` — delete a lead

### Create a lead (example)

```bash
curl -s -X POST http://localhost:3000/api/lead \
  -H 'Content-Type: application/json' \
  -d '{"name":"Alice","company":"Acme","email":"alice@example.com","phone":1234567890}'
```

## Notes

- Uses a local SQLite DB file: `leads.db`
- No authentication/validation (demo project)


