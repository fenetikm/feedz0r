# FeedzOr

RSS feed wrangler CLI.

## Stack
- Go language
- `sqlite`
- `goose` for migrations
- `sqlc` for generating DB models etc. from SQL

## Development
Run:
```sh
go run ./cmd/fz
```

Build:
```sh
go build -o fz ./cmd/fz
```

Clean up:
```sh
go mod tidy
```

### Database
After changing SQL in `internal/db/queries/*.sql`, run:
```sh
sqlc generate
```

Migrations live in `internal/db/schema/`. After adding a migration file, run:
```sh
goose -dir internal/db/schema sqlite3 feedz0r.db up
```
Each migration file name has an incremented prefix.

### Adding a new command
1. Create `internal/commands/<name>/<name>.go` with a `Handle(s *state.State)` function
2. Register it in `cmd/fz/main.go`
