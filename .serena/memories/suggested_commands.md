# Suggested Commands
- `go run ./cmd/api` – start the HTTP API (uses env vars `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASS`, `DB_NAME`; defaults to localhost MySQL, auto-migrates + seeds plans).
- `go build ./cmd/api` – compile the server binary.
- `go test ./...` – run all Go unit tests (none yet, but command stays standard).
- `gofmt -w <files>` – format Go source before committing.
- `golangci-lint run` – if linting is added; current project does not vendor config.