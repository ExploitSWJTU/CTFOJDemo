# GEMINI.md - SWJTU CTF OJ Frontend Demo

This document serves as the primary context for Gemini CLI when working on this project.

## Project Overview

**SWJTU CTF OJ Frontend** is a modern Capture The Flag (CTF) Online Judge frontend demo built with **Vue 3** and **TypeScript**. It features a high-performance developer experience using **Vite**, sophisticated styling with **Tailwind CSS v4**, and a rich UI component library provided by **Arco Design Vue**.

### Core Tech Stack

- **Framework:** Vue 3 (Composition API with `<script setup>`)
- **Language:** TypeScript
- **Build System:** Vite
- **UI Components:** [Arco Design Vue](https://arco.design/vue/)
- **Styling:** Tailwind CSS v4
- **Icons:** Lucide Vue Next
- **Routing:** Vue Router

### Key Features

- **Dark Mode:** Persisted theme support with system preference detection.
- **Auto-Imports:** Streamlined development with `unplugin-auto-import` and `unplugin-vue-components`.
- **Responsive Layout:** Mobile-friendly design using Tailwind utility classes.
- **Mock Features:** Dynamic challenge lists, container status monitoring, and navigation for Training, Contests, and Forums.

---

## Getting Started

### Prerequisites

- Node.js (Latest LTS recommended)
- `pnpm` (Preferred package manager)

### Development Commands

| Command         | Description                                                      |
| :-------------- | :--------------------------------------------------------------- |
| `pnpm dev`      | Starts the development server with Hot Module Replacement (HMR). |
| `pnpm build`    | Compiles the project using `vue-tsc` and Vite for production.    |
| `pnpm preview`  | Locally previews the production build.                           |
| `pnpm lint`     | Runs ESLint for code quality checks.                             |
| `pnpm lint:fix` | Automatically fixes linting issues.                              |
| `pnpm format`   | Formats the entire codebase using Prettier.                      |

---

## Project Structure

- `src/main.ts`: Application entry point.
- `src/App.vue`: Main layout wrapper with global theme handling.
- `src/router/index.ts`: Route definitions (Home, Training, Contest, Forum).
- `src/components/`: Reusable Vue components.
  - `Header.vue`: Main navigation, search, and system utilities.
  - `Challenges.vue`: Display logic for CTF tasks.
  - `Sidebar.vue`: Category filters and stats.
- `src/views/`: Page components.
- `vite.config.ts`: Configuration for Vite, including Arco Design resolvers.

---

## Development Conventions

- **Component Standard:** Use SFCs with `<script setup lang="ts">`.
- **Styling:**
  - Prefer Tailwind utility classes.
  - Custom CSS variables for theme-aware Arco Design overrides are located in `src/App.vue`.
- **Linting:** Strict ESLint configuration (`eslint.config.js`) enforces Tailwind class ordering and Vue SFC best practices.
- **Typing:** Avoid `any` where possible; use interface/type definitions for component props and reactive state.
- **Commiting:** Run `pnpm format` and `pnpm lint` before submitting changes.
