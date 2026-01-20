# AI Agents Context

This file documents the AI agents (like Antigravity) that assist in the development of this project. It provides context, rules, and history to ensure consistent collaboration.

## 🤖 Current Agent: Antigravity

**Role**: Lead AI Developer / Pair Programmer
**Stack Expertise**: 
- Backend: Go (chi, gorm, embed)
- Frontend: Vue 3 (Vite, Pinia, Vue Router)
- Infrastructure: Make, Docker, CI/CD

### 📜 Project Guidelines for Agents

1.  **Uniformity**: Keep the "Laravel-like" feel. Commands should be centralized in the `Makefile` or root `package.json`.
2.  **Embedded First**: Always remember that the Go backend is responsible for serving the frontend in production via `go:embed`.
3.  **Clean Code**: Use `internal/` for Go logic to keep `cmd/` thin. Use Vue Composition API (`<script setup>`) for the frontend.
4.  **No Placeholders**: When adding features, provide working implementations or UI mocks using actual designs, not just "TODO" comments.

### 🛠 Active Workflows

- **Dev Workflow**: `make dev` runs both Vite and Go. Vite proxies `/api` to `:8080`.
- **Build Workflow**: `make build` compiles frontend, then builds Go with the static assets embedded.

### 📅 Session History Summary

- **2026-01-21**: Project initialization. 
    - Set up Go server with `chi` and CORS.
    - Scaffolded Vue frontend with Vite.
    - Configured Vite Proxy and Go SPA routing with `go:embed`.
    - Created automation via `Makefile`.
