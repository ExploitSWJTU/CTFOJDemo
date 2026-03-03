# AGENTS.md - Frontend (Vue 3)

**Generated:** 2026-03-03
**Commit:** 26cbfb8

## OVERVIEW
Vue 3 + TypeScript + Vite frontend for SWJTU CTF OJ. Uses Arco Design for UI, Tailwind for styling, Pinia for state.

## STRUCTURE
```
frontend/src/
├── api/           # Axios wrappers (NO axios in components)
├── components/    # Organized by domain: admin/, user/, shared/
├── constants/     # Enums (e.g., category.ts)
├── layouts/       # UserLayout.vue, AdminLayout.vue
├── mock/          # Development mock data (remove for production)
├── router/        # Vue Router config with lazy-loaded views
├── stores/        # Pinia stores (challenge, user, contest, etc.)
├── types/         # TypeScript interfaces
└── views/         # Page components: admin/*, user/*
```

## WHERE TO LOOK
| Task | Location |
|------|----------|
| Add new API endpoint | `src/api/` |
| Add new page | `src/views/user/` or `src/views/admin/` |
| Add global component | `src/components/shared/` |
| Add state | `src/stores/` |
| Add route | `src/router/index.ts` |
| Add type | `src/types/` |

## CONVENTIONS

### Styling
- **Tailwind First**: `class="m-5 text-blue-600"` NOT `style="margin: 20px"`
- **Arco for Interactive**: Use `a-table`, `a-modal`, `a-form` for complex UI
- **Design Tokens**: `blue-600` primary, `amber-500` warning, `emerald-500` success

### TypeScript
- Strict mode enabled. Avoid `any` — define interfaces in `src/types/`
- Use `<script setup lang="ts">` for all Vue components
- Auto-imports enabled for Vue APIs and Arco components

### API Pattern
```ts
// src/api/example.ts
import request from './request';
export const getItems = () => request.get('/items');
```

### Component Naming
- Multi-word preferred but single-word allowed (ESLint rule off)
- Views: `XxxView.vue`, Admin views: `AdminXxxView.vue`

## ANTI-PATTERNS
- **NEVER** call axios directly in components — use `src/api/`
- **NEVER** use `any` type — define proper interfaces
- **NEVER** inline styles when Tailwind classes exist
- `src/mock/` violates "no mock in production" — remove before deploy

## COMMANDS
```bash
pnpm dev          # Vite dev server (localhost:5173)
pnpm build        # Type-check + build
pnpm lint         # ESLint check
pnpm type-check   # TypeScript validation
```

## NOTES
- Auto-generated: `auto-imports.d.ts`, `components.d.ts` (don't edit)
- Large files: `TeamView.vue` (643L), `TrainingView.vue` (625L) — consider splitting
- 8 TODOs found — all API integration tasks in stores/views