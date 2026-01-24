# Go Database (file-based JSON storage)

A tiny “database” implementation that stores records as **JSON files on disk**.

It provides basic operations:

- write a record (`Write`)
- read a record (`Read`)
- list all records in a collection (`ReadAll`)
- delete a record/collection (`Delete`)

The demo in `main.go` writes user records into the `users/` collection.

## Run

From `go-database/`:

```bash
go run .
```

After running, you’ll see JSON files created/updated under:

- `./users/*.json`

## Notes

- Uses per-collection mutexes for basic concurrency safety.
- Writes are done via a `*.tmp` file + rename for safer persistence.


