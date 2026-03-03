# AGENTS.md - Backend (Go)

**Generated:** 2026-03-03
**Commit:** 26cbfb8

## OVERVIEW
Go 1.25 + Gin + GORM backend for SWJTU CTF OJ. PostgreSQL for persistence, Redis for caching, Docker SDK for challenge containers.

## STRUCTURE
```
backend/
├── cmd/server/main.go     # Entry point
├── config/                # (Empty — uses env vars instead)
├── internal/
│   ├── api/               # Gin handlers (controllers)
│   ├── middleware/        # JWT auth, CORS
│   ├── model/             # GORM models + DTOs
│   ├── repository/        # Data access layer
│   └── service/           # Business logic (Docker scheduling)
├── pkg/utils/             # Shared utilities
├── .air.toml              # Hot reload config
└── .golangci.yml          # Linter config
```

## WHERE TO LOOK
| Task | Location |
|------|----------|
| Add new API endpoint | `internal/api/`, register in `main.go` |
| Add new model | `internal/model/` |
| Add business logic | `internal/service/` |
| Add DB query | `internal/repository/` |
| Add middleware | `internal/middleware/` |

## CONVENTIONS

### Error Handling
```go
if err != nil {
    return c.JSON(500, gin.H{"code": 500, "msg": err.Error()})
}
```
Always return standardized JSON: `{"code": 200, "msg": "success", "data": {...}}`

### GORM Models
- Use `datatypes.JSON` for flexible fields (ContainerConfig, Metadata)
- Auto-migrate in `main.go` — add new models there

### Config
- **Uses environment variables**, NOT `config/config.yaml`
- See `.env.example` for required vars
- Load via `repository.LoadConfigFromEnv()`

### Docker SDK
- Close clients after use
- Handle context timeouts for API calls
- Port pool managed in Redis

## ANTI-PATTERNS
- **NEVER** suppress errors — always check `if err != nil`
- **NEVER** use `config/config.yaml` — project uses env vars
- **NEVER** leave Docker clients open — always defer close

## COMMANDS
```bash
go run cmd/server/main.go   # Run directly
air                         # Hot reload (via .air.toml)
go test ./...               # Run tests
golangci-lint run           # Lint check
```

## NOTES
- `backend/package.json` exists for npm scripts only (air, lint) — unusual for Go
- No tests currently (`*_test.go` files absent)
- Missing `Dockerfile` — docker-compose.yml references it but file doesn't exist
- CORS allows `localhost:5173` and `localhost:3000` only