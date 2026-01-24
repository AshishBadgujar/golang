# Go Basic Server

A minimal **net/http** server that serves static files and a couple of simple endpoints.

## Run

From `go-basic-server/`:

```bash
go run .
```

Server starts on `http://localhost:8080`.

## Endpoints

- `GET /`: serves static files from `./static` (opens `static/index.html`)
- `GET /hello`: returns `Hello!`
- `POST /form`: handles a form submission and echoes `name` and `address`

## Try it

- Open the static site:

```bash
open http://localhost:8080
```

- Test `/hello`:

```bash
curl -s http://localhost:8080/hello
```

- Test form handler:

```bash
curl -s -X POST http://localhost:8080/form \
  -d "name=Alice" \
  -d "address=NYC"
```


