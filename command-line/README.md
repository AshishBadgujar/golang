# Videos CLI (command-line)

A simple **CLI tool** to manage a list of videos stored in a local `videos.json` file.

## Commands

### Get videos

- Get all:

```bash
go run . get --all
```

- Get by id:

```bash
go run . get --id QThadS3Soig
```

### Add a video

```bash
go run . add \
  --id "abc123" \
  --title "My Video" \
  --url "https://youtu.be/abc123" \
  --imageurl "https://i.ytimg.com/vi/abc123/mqdefault.jpg" \
  --desc "Short description"
```

This appends a new entry to `videos.json`.

## Docker

Build:

```bash
docker build -t videos-cli .
```

Run:

```bash
docker run --rm -v "$PWD:/videos" -w /videos videos-cli get --all
```

## Notes

- Data lives in `videos.json` in the project root.
- The binary name is `videos` (see `Dockerfile` / `run.sh`).


