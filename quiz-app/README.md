# Quiz App (CLI)

A small timed quiz CLI that reads questions/answers from a CSV file and quizzes you in the terminal.

## Run

From `quiz-app/`:

```bash
go run .
```

## Options

- `-f` — path to CSV file (default: `quiz.csv`)
- `-t` — quiz timer in seconds (default: `30`)

Example:

```bash
go run . -f quiz.csv -t 60
```

## CSV format

Each line is:

`question,answer`


