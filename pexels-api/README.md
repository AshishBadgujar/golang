# Pexels API (Go client demo)

A small Go program that calls the **Pexels REST API** and prints results.

Current demo behavior: fetches a **random curated photo** and prints it to stdout.

## Run

From `pexels-api/`:

```bash
go run .
```

## Token

The Pexels API requires an API key.

This project currently sets `PexelsToken` inside `main.go` (via `os.Setenv`). A better practice is to set it in your shell and remove the hardcoded token:

```bash
export PexelsToken="YOUR_PEXELS_API_KEY"
go run .
```

## Notes

- Base URL: `https://api.pexels.com/v1`
- Authorization header: `Authorization: <token>`


