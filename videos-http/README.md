# Videos HTTP (JSON file API)

A small Go HTTP server that:

- serves a list of videos from `videos.json`
- accepts updates and writes them to `videos-updated.json`

## Run

From `videos-http/`:

```bash
go run .
```

Server listens on `http://localhost:8080`.

## Endpoints

- `GET /` — returns videos from `videos.json`
- `POST /update` — accepts a JSON array of videos and saves to `videos-updated.json`

### Update example

```bash
curl -s -X POST http://localhost:8080/update \
  -H 'Content-Type: application/json' \
  -d @videos.json
```

## Notes

- This is a demo project; there is no DB, only JSON files on disk.


