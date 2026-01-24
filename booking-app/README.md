# Booking App (CLI)

A simple **conference ticket booking** CLI written in Go. It collects user input, books tickets in-memory, and simulates sending a ticket using a goroutine.

## Run

From `booking-app/`:

```bash
go run .
```

## How it works

- Prompts for **first name**, **last name**, and **number of tickets**
- Validates requested tickets are `> 0` and `<= remainingTickets`
- Books tickets (updates remaining count + stores booking in memory)
- Starts a goroutine to **simulate sending tickets** (waits ~10 seconds)

## Notes

- Bookings are stored **in memory only** (no DB/file persistence).
- Uses `sync.WaitGroup` to wait for the “send ticket” goroutine to finish.


