# Go Beautiful Blogs (static site + contacts API)

A small Go web server that:

- serves a **static “beautiful blog” landing page** (HTML/CSS/JS + images)
- exposes a tiny **contacts API** backed by a local `contacts.json` file

## Run

From `go-beutiful-blogs/`:

```bash
go run .
```

Open:

- `http://localhost:8080/`

## API

- `GET /api/` — returns all contacts as JSON
- `POST /api/submit` — appends a contact to `contacts.json`

Example:

```bash
curl -s -X POST http://localhost:8080/api/submit \
  -H 'Content-Type: application/json' \
  -d '{"id":"id_test","name":"Alice","email":"alice@example.com","message":"Hello!"}'
```

## Data storage

Contacts are stored in `contacts.json` in the project root (simple file-based storage).


