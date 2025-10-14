# Task Completion Checklist
- Ensure modified Go files are gofmt formatted (`gofmt -w`).
- Run `go test ./...` to confirm unit tests (once they exist) pass.
- If server behavior affected, consider `go run ./cmd/api` smoke test against local MySQL.
- Update documentation/comments when changing exposed behavior or contracts.
- Summarize changes and note any manual verification steps for the user.