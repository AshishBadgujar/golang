# Fiber HRMS (MongoDB)

A small HRMS-style REST API built with **Go Fiber** and **MongoDB**. It manages `Employee` records in the `fiber-hrms` database.

## Start MongoDB (Docker)

From `go-fiber-mongo-hrms/`:

```bash
docker compose up -d
```

- MongoDB: `mongodb://username:password@localhost:27017`
- Optional Mongo Express UI: `http://localhost:8080`

## Run the API

```bash
go run .
```

Server listens on `http://localhost:3000`.

## API

- `GET /employee` — list all employees
- `POST /employee` — create employee
- `PUT /employee/:id` — update employee (Mongo ObjectID hex)
- `DELETE /employee/:id` — delete employee

### Create employee (example)

```bash
curl -s -X POST http://localhost:3000/employee \
  -H 'Content-Type: application/json' \
  -d '{"name":"Alice","age":28,"salary":120000}'
```

## Notes

- Mongo connection string is hardcoded in `main.go` (`mongoURI`).
- Collection: `employees`


