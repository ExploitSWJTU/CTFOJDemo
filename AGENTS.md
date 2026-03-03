# AGENTS.md - SWJTU CTF OJ Project

> **Note:** This file defines cross-cutting architecture. See `frontend/AGENTS.md` and `backend/AGENTS.md` for domain-specific conventions.

**Generated:** 2026-03-03 | **Commit:** 26cbfb8 | **Branch:** main

## OVERVIEW
CTF Training Platform with Docker container scheduling. Vue 3 frontend + Go backend, PostgreSQL + Redis infrastructure.

## STRUCTURE
```
/
├── frontend/       # Vue 3 app → frontend/AGENTS.md
├── backend/        # Go app → backend/AGENTS.md
├── scripts/        # docker-start.sh, docker-stop.sh, pre-commit.mjs
├── docs/           # Documentation
├── docker-compose.yml
├── .env.example
└── AGENTS.md
```

## WHERE TO LOOK
| Task | Location |
|------|----------|
| Add frontend feature | `frontend/src/views/`, `frontend/src/stores/` |
| Add backend API | `backend/internal/api/`, `backend/internal/service/` |
| Add database model | `backend/internal/model/` |
| Modify routes | `frontend/src/router/index.ts`, `backend/cmd/server/main.go` |
| Change styling | `frontend/` — Tailwind classes, Arco components |

## TECH STACK
| Layer | Technology |
|-------|------------|
| Frontend | Vue 3 + TypeScript + Vite + Pinia + Arco Design + Tailwind |
| Backend | Go 1.25 + Gin + GORM + PostgreSQL + Redis |
| Infra | Docker Compose, Docker SDK for Go |

## DATA MODELS

### API Response
```json
{ "code": 200, "msg": "success", "data": {...} }
```

### GORM Models
```go
User{ID, Username, Role}           // Role: "admin" | "user"
Challenge{ID, Title, Category, Score, ContainerConfig, Metadata, Type}
Container{ID, UserID, ChallengeID, ContainerID, PortMapping, Flag, ExpiresAt}
```

### Core Endpoints
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/auth/login` | JWT login |
| GET | `/api/challenges` | List challenges |
| POST | `/api/challenges/:id/container` | Start container |
| POST | `/api/challenges/:id/flag` | Submit flag |

## COMMANDS
```bash
pnpm dev:all           # Full stack (DB + Redis + Backend + Frontend)
pnpm dev:frontend      # Frontend only (localhost:5173)
pnpm dev:backend       # Backend with air hot reload
pnpm build             # Production build
pnpm lint              # ESLint (frontend)
pnpm type-check        # TypeScript + Go vet
```

## KNOWN ISSUES
| Issue | Location | Impact |
|-------|----------|--------|
| Missing `Dockerfile` | `backend/` | `docker compose up` fails |
| `backend/package.json` exists | `backend/` | Unusual for Go (npm scripts only) |
| `frontend/src/mock/` present | `frontend/src/` | Violates "no mock in production" |
| No tests | Both | `*_test.go`, `*.spec.ts` absent |
| 8 TODOs | Frontend stores/views | API integration incomplete |

## ENVIRONMENT
See `.env.example` for: `DB_*`, `REDIS_PORT`, `BACKEND_PORT`, `JWT_SECRET`, `DOCKER_PORT_RANGE_*`