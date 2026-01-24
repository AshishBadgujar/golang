# Go Todo (static UI + JSON API)

A simple Todo web app:

- serves a static HTML/CSS/JS UI
- provides a small JSON API backed by a local `todo.json` file

## Run

From `go-todo/`:

```bash
go run .
```

Open:

- `http://localhost:8080/`

## API

- `GET /api/` — list todos
- `POST /api/submit` — create todo (JSON body)
- `DELETE /api/delete?id=<id>` — delete todo by id

Example:

```bash
curl -s -X POST http://localhost:8080/api/submit \
  -H 'Content-Type: application/json' \
  -d '{"id":"id_test","todo":"buy milk"}'
```

## Notes

- Data is stored in `todo.json` in the project root.
- `Dockerfile` currently appears to be copied from another project (it builds `/contacts`), so local `go run .` is the recommended way to run this one.


