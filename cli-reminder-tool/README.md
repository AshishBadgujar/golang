# CLI Reminder Tool

A small command-line reminder app that schedules a desktop notification for a future time.

It parses natural-ish time input (like `14:30`, `tomorrow 9am`) and then shows a notification using `beeep`.

## Run

From `cli-reminder-tool/`:

```bash
go run . "<time>" "<message...>"
```

Example:

```bash
go run . "14:30" "Stand up and stretch"
```

## Notes

- Uses a child process (re-exec) to run in the background and then sleeps until the target time.
- Notification icon path is `assets/information.png`. If you don’t have that file, notifications may fail or show without an icon (depending on OS).


