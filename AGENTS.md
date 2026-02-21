# AGENTS.md - SWJTU CTF OJ Project Context & Rules

> **Note to AI Agents:** This file defines the technical stack, coding standards, and architectural guidelines for the SWJTU CTF OJ project. Please read this before generating or modifying code.

## 1. Project Overview
**Name:** SWJTU CTF OJ (Online Judge)
**Type:** Cybersecurity Training Platform
**Architecture:** Monorepo (`frontend/` + `backend/`)
**Core Features:**
- **User Side:** Training (Docker container management), Contests, Forum.
- **Admin Side:** Challenge management, Container monitoring, User management.
- **Infrastructure:** Dynamic Docker container scheduling for CTF challenges.

### Challenge Types
- **Attachment:** Static files (Misc, OSINT) — download only, no container.
- **Static Container:** Pre-built Docker image, fixed flag per challenge.
- **Dynamic Container:** Per-user container with unique generated flag (future).

## 2. Tech Stack

### Frontend (`frontend/`)
- **Framework:** Vue 3 (Composition API) + TypeScript + Vite.
- **State Management:** Pinia.
- **UI Library:** Arco Design Vue (for complex components like Tables, Forms, Modals).
- **Styling:** Tailwind CSS (for Layout, Spacing, Typography, Colors).
- **Package Manager:** pnpm.
- **Deployment:** Cloudflare Pages / Nginx / Docker.

### Backend (`backend/`)
- **Language:** Go (Golang) >= 1.20.
- **Web Framework:** Gin.
- **Database ORM:** GORM (PostgreSQL).
- **Cache/KV:** Redis (for session, rate limiting, port pool).
- **Containerization:** Docker SDK for Go (moby/moby).

### Infrastructure
- **Deployment:** Docker Compose (single server).
- **Database:** PostgreSQL 15.
- **Cache:** Redis 7.
- **Container Runtime:** Docker (host socket mount).

## 3. Directory Structure Rules

### Root (Monorepo)
```text
/
├── frontend/               # Vue 3 Frontend Application
├── backend/                # Go Backend Application
├── docker-compose.yml      # Full-stack orchestration
├── .env.example            # Environment variable template
├── AGENTS.md               # This file
└── .gitignore
```

### Frontend (`frontend/src/`)
Please adhere to the following modular structure. Do not create flat structures in `views/` or `components/`.

```text
frontend/
├── index.html
├── package.json
├── vite.config.ts
├── tsconfig.json
└── src/
    ├── api                 # API definitions (Axios wrappers), NO mock data in production logic
    ├── assets              # Static assets
    ├── components          # Global shared components ONLY
    │   ├── common          # Basic UI wrappers
    │   └── layout          # Header, Sidebar, Footer
    ├── composables         # Global hooks (useTheme, etc.)
    ├── constants           # Global constants (Enums)
    ├── layouts             # Layout containers
    │   ├── UserLayout.vue  # For User Interface (TopNav)
    │   └── AdminLayout.vue # For Admin Interface (Sidebar + Header)
    ├── router              # Vue Router config
    ├── stores              # Pinia stores
    ├── types               # TypeScript interfaces/types
    ├── utils               # Helper functions (request.ts, format.ts)
    └── views               # Page Views
        ├── admin           # Admin Dashboard Pages
        │   ├── dashboard
        │   ├── challenges  # Challenge Management
        │   └── containers  # Container Monitoring
        ├── user            # User Interface Pages
        │   ├── home
        │   ├── training    # Training List & Detail
        │   └── contest
        └── auth            # Login/Register
```

### Backend (`backend/`)
Follow the Standard Go Project Layout.

```text
backend/
├── cmd/
│   └── server/             # Main application entry point
│       └── main.go
├── config/
│   └── config.yaml         # Application config (DB, Redis, Docker, JWT)
├── internal/               # Private application code
│   ├── api/                # Gin Handlers (Controllers)
│   ├── middleware/         # JWT auth, CORS, rate limiting
│   ├── service/            # Business Logic (Docker scheduling here)
│   ├── repository/         # Data Access Layer (GORM/Redis)
│   └── model/              # Struct definitions (DTOs & DB Models)
├── pkg/                    # Public library code (Utils)
│   └── utils/
├── go.mod
└── go.sum
```

## 4. Coding Standards & Best Practices

### Frontend (Vue + Tailwind)
1.  **Tailwind First:** Use Tailwind utility classes for layout (`flex`, `grid`), spacing (`p-4`, `m-2`), and colors (`text-blue-600`).
    -   *Bad:* `<div style="margin: 20px; color: blue;">`
    -   *Good:* `<div class="m-5 text-blue-600">`
2.  **Arco Design:** Use Arco for interactive components (`a-table`, `a-modal`, `a-form`).
    -   Customize Arco styles using Tailwind classes where possible, or scoped CSS if necessary.
3.  **TypeScript:**
    -   Avoid `any`. Define interfaces in `src/types/`.
    -   Use `<script setup lang="ts">`.
4.  **API Handling:**
    -   Do not call `axios` directly in components. Define functions in `src/api/` and import them.

### Backend (Go)
1.  **Error Handling:** Always check `if err != nil`. Return standardized JSON errors to the frontend.
2.  **Database:** Use GORM models. Use `datatypes.JSON` for flexible configuration fields (e.g., ContainerConfig).
3.  **Docker SDK:** Ensure Docker clients are closed or reused properly. Handle context timeouts for API calls.

## 5. Data Models (Context)

### Standard API Response
```json
{
  "code": 200,          // 200 = Success, others = Error
  "msg": "success",     // Error message if code != 200
  "data": { ... }       // Payload
}
```

### Core API Endpoints
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/auth/login` | JWT login |
| POST | `/api/auth/register` | User registration |
| GET | `/api/challenges` | List challenges (with filters) |
| GET | `/api/challenges/:id` | Challenge detail |
| POST | `/api/challenges/:id/container` | Start container instance |
| DELETE | `/api/challenges/:id/container` | Destroy container |
| POST | `/api/challenges/:id/container/extend` | Extend container lifetime |
| POST | `/api/challenges/:id/flag` | Submit flag |

### Core Database Schema (GORM)
*Reference these models when writing backend logic.*

```go
// User
type User struct {
    ID       uint   `gorm:"primarykey"`
    Username string `gorm:"uniqueIndex"`
    Role     string `gorm:"default:'user'"` // 'admin' or 'user'
}

// Challenge
type Challenge struct {
    ID              uint           `gorm:"primarykey"`
    Title           string
    Category        string         // Web, Pwn, Crypto...
    Score           int
    ContainerConfig datatypes.JSON `gorm:"type:jsonb"` // {"image": "nginx", "port": 80}
    Metadata        datatypes.JSON `gorm:"type:jsonb"` // {"author": "admin", "tags": ["easy"]}
    Type            string         // 'attachment', 'static' or 'dynamic'
}

// Container (Runtime)
type Container struct {
    ID          uint      `gorm:"primarykey"`
    UserID      uint
    ChallengeID uint
    ContainerID string    // Docker ID
    PortMapping datatypes.JSON `gorm:"type:jsonb"` // {"80/tcp": 32001}
    Flag        string    // The actual flag
    ExpiresAt   time.Time
}
```

## 6. Design System Guidelines (UI/UX)
*   **Theme:** Modern SaaS style.
*   **Colors:**
    *   Primary: Blue (`blue-600` / `#2563eb`).
    *   Warning: Amber (`amber-500`).
    *   Success: Emerald (`emerald-500`).
    *   Background: Slate (`slate-50` for light mode).
*   **Components:** Large border-radius (`rounded-2xl` or `rounded-3xl`), soft shadows (`shadow-lg`).

## 7. Development & Deployment

### Local Development
```bash
# Frontend
cd frontend && pnpm install && pnpm dev

# Backend
cd backend && go run cmd/server/main.go

# Infrastructure (DB + Redis)
docker compose up db redis
```

### Production Deployment
```bash
cp .env.example .env   # Edit with real secrets
docker compose up -d
```
