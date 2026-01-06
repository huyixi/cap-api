# Repository Guidelines

## Project Structure & Module Organization
- `cmd/api/main.go` is the entry point; it wires the HTTP server, database, and routes.
- `internal/handler/` contains Gin HTTP handlers and request/response wiring.
- `internal/service/` holds business logic and data access via `sqlx`.
- `internal/db/schema.sql` defines the SQLite schema loaded at startup.
- `memos.db` is the default local SQLite file used when `DB_PATH` is not set.

## Build, Test, and Development Commands
- `go run ./cmd/api` starts the API on `:8080` using `memos.db` in the repo.
- `DB_PATH=/path/to/memos.db go run ./cmd/api` points the server at another SQLite file.
- `go test ./...` runs all tests (none are present yet).
- `gofmt -w ./...` formats Go code using the standard Go formatter.
- `go mod tidy` cleans up module dependencies after adding or removing packages.

## Coding Style & Naming Conventions
- Use `gofmt` defaults (tabs for indentation, standard Go formatting).
- Keep package names short and lowercase (`handler`, `service`).
- Prefer clear, HTTP-oriented method names in handlers (e.g., `create`, `list`).
- Use struct tags for JSON/DB mapping, matching existing patterns in `internal/service/memos.go`.

## Testing Guidelines
- Use Go’s standard `testing` package.
- Favor table-driven tests for service logic.
- Suggested locations: `internal/service/` for business rules, `internal/handler/` for HTTP behavior.
- Name test files `*_test.go` and test functions `TestXxx`.

## Commit & Pull Request Guidelines
- Only two initial commits exist; there is no established commit message convention.
- Use concise, imperative subjects (e.g., "Add memo search filtering").
- PRs should explain behavior changes, list new endpoints or parameters, and include API examples (curl is fine).
- If a change impacts schema, include the SQL change and note migration steps.

## Configuration & Security Notes
- The API reads `DB_PATH` at startup; default is `./memos.db`.
- Avoid committing local data changes unless they are intentional fixtures.
