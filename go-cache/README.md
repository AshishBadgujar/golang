# Go Cache (LRU example)

A small **LRU (Least Recently Used) cache** example implemented from scratch using:

- a doubly-linked list (queue) to track recency
- a hash map for O(1) lookups

Cache capacity is fixed at `SIZE = 5`.

## Run

From `go-cache/`:

```bash
go run .
```

It runs a small demo sequence and prints cache operations (`add`, `remove`) and the cache contents after each step.


