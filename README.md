# Vue-Go Starter Kit

A full-stack project template with a Go (Golang) backend and a Vue.js (Vite) frontend, designed to provide a Laravel-like developer experience.

## Features

- **Monolith Structure**: Frontend and backend in a single repository.
- **Modern Tech Stack**: Go 1.25+, Vue 3 (Composition API), Vite.
- **Embedded Frontend**: In production, the Vue application is embedded into the Go binary using `go:embed`.
- **Developer Friendly**: 
    - Automated proxy in dev mode (Vite -> Go).
    - Shared management via `Makefile`.
    - SPA routing support out of the box.
- **Ready-to-go Auth**: JWT authentication with Pinia and Axios already configured.

## Project Structure

```text
.
├── main.go           # Go entry point
├── frontend/         # Vue.js application (Vite)
├── internal/         # Backend business logic, models, handlers
├── Makefile          # Main command center
└── go.mod            # Go module definition
```

## Quick Start (One-liner)

You can scaffold a new project using `degit`. This will download the template without the git history:

```bash
npx degit dedomorozoff/vue-go my-project && cd my-project && make install && make migrate && make admin
```

> **Note:** You can also click the **"Use this template"** button on GitHub to create a new repository based on this one.

## Getting Started

### Prerequisites

- Go (1.25 or later)
- Node.js (v20+) & npm

### Installation from scratch

1. **Install dependencies**:
   ```bash
   make install
   ```

2. **Initialize Database**:
   ```bash
   make migrate
   ```

3. **Set up Admin User**:
   ```bash
   make admin
   ```
   *Follow the interactive prompt to set your password.*

4. **(Optional) Seed test data**:
   ```bash
   make seed
   ```

### Development

Run both backend and frontend simultaneously with Hot Module Replacement (HMR):

```bash
make dev
```

- **Frontend**: http://localhost:5173
- **Backend API**: http://localhost:8080 (Proxied via frontend)
- **Admin**: http://localhost:5173/admin

### Production Build

Create a single executable binary containing the embedded frontend:

```bash
make build
```
The resulting binary will be in `bin/server`.

##  API Endpoints

- `GET /api/health`: Health check.
- `POST /api/auth/login`: Authentication.
- `GET /api/projects`: Protected list of projects.

##  License

MIT
