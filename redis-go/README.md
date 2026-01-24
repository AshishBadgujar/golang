# Redis Go (videos API)

A small Go HTTP server that stores and reads “video” records in **Redis**.

## Prerequisites

- Redis running locally on `localhost:6379`

Example (Docker):

```bash
docker run --rm -p 6379:6379 redis:7
```

## Run

From `redis-go/`:

```bash
go run .
```

Server listens on `http://localhost:8080`.

## Endpoints

- `GET /` — list all videos (reads all Redis keys)
- `GET /?id=<id>` — get a single video by id
- `POST /update?id=<id>` — upsert a single video (JSON body)
- `POST /update` — upsert multiple videos (JSON array body)

### Example: add/update one video

```bash
curl -s -X POST "http://localhost:8080/update?id=QThadS3Soig" \
  -H 'Content-Type: application/json' \
  -d '{"id":"QThadS3Soig","title":"Kubernetes on Amazon","description":"","imageurl":"","url":"https://youtu.be/QThadS3Soig"}'
```

## Notes

- Uses Redis DB `0` with no password (see `main.go`).
- `videos.json` is sample data; you can POST it to `/update` to populate Redis.


