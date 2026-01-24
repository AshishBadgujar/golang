# Weather API (open-meteo)

A small Go HTTP server that returns the current temperature for a given city using the **Open-Meteo** APIs:

- Geocoding: `geocoding-api.open-meteo.com`
- Forecast: `api.open-meteo.com`

## Run

From `weather-api/`:

```bash
go run .
```

Server listens on `http://localhost:8080`.

## Endpoints

- `GET /hello` — returns a hello message
- `GET /weather/{city}` — returns JSON including current temperature

Example:

```bash
curl -s http://localhost:8080/weather/london
```


