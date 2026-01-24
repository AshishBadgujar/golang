# Mailcheck (DNS record checker)

A small CLI that reads domain names from **stdin** and prints whether they have:

- **MX** records
- **SPF** TXT record (`v=spf1...`)
- **DMARC** TXT record (`_dmarc.<domain>`, `v=DMARC1...`)

Output is CSV-like:

`domain,hasMX,hasSPF,spfRecord,hasDMARC,dmarcRecord`

## Run

From `mailcheck/`:

```bash
go run .
```

Then type domains (one per line), for example:

```bash
echo -e "gmail.com\nexample.com" | go run .
```


