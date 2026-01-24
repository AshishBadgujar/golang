# Go Concepts (learning exercises)

A collection of small, focused Go programs demonstrating core language concepts (input, conversions, slices, maps, structs, concurrency, etc.).

Each numbered folder is a standalone mini-program you can run with `go run .` from inside that folder.

## Run an exercise

Example:

```bash
cd 01user-input
go run .
```

## What’s inside

- `01user-input` → reading from stdin
- `02conversion` → type conversions
- `03time` → time utilities
- `04pointers` → pointers
- `05array`, `06slices`, `07map` → collections
- `08struct`, `12methods` → structs & methods
- `22goroutines`, `23mutexAndAwaiGroups`, `24channels` → concurrency
- `21mongoapi`, `20buildapi`, `19modules` → bigger modules-based examples (have their own `go.mod`)

## lcoserver (Node/Express)

`lcoserver/` is a small Node.js Express server (not Go).

Run it:

```bash
cd lcoserver
npm install
npm run start
```


