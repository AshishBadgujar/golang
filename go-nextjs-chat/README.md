# Go + Next.js Chat App

A simple chat app with:

- **Backend**: Go + Fiber + WebSocket (`/ws`)
- **Frontend**: Next.js + Tailwind (connects as a WebSocket client)

## Run backend

From `go-nextjs-chat/backend/`:

```bash
go run .
```

Backend listens on `http://localhost:8000` and exposes:

- `GET /` — health message
- `GET /ws?username=<name>` — WebSocket endpoint

## Run frontend

From `go-nextjs-chat/frontend/`:

```bash
npm install
npm run dev
```

Frontend dev server runs at `http://localhost:3000`.

## Notes

- Start the **backend first**, then the frontend.
- The frontend project contains its own `README.md` with Next.js details.


