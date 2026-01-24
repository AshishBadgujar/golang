# Load Balancer (reverse proxy, round-robin)

A simple **HTTP reverse-proxy load balancer** example using Go’s `net/http` and `httputil.ReverseProxy`.

It forwards incoming requests to a list of upstream servers using a basic **round-robin** strategy and skips servers that fail a health check.

## Run

From `loadbalancer/`:

```bash
go run .
```

It listens on:

- `http://localhost:8000`

## How it works

- Upstreams are hardcoded in `main.go` (example: Facebook, Bing, DuckDuckGo)
- Health check: performs a `GET` to the upstream with a 5s timeout and treats **2xx** as alive
- Proxying: forwards your request to the selected upstream

## Try it

```bash
curl -I http://localhost:8000
```


