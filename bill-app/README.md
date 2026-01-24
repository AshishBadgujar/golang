# Bill App (CLI)

A small interactive **CLI bill calculator** that lets you:

- create a bill (name)
- add line items (name + price)
- add a tip
- save the bill to a text file in `bills/`

## Run

From `bill-app/`:

```bash
go run .
```

## Usage

When prompted:

- `a` = add item
- `t` = add tip
- `s` = save bill

On save, the bill is written to:

- `bills/<bill-name>.txt`

Example output format includes each item, the tip, and the total.

## Notes

- Prices/tip must be valid numbers (floats).
- Files are saved with mode `0644`.


