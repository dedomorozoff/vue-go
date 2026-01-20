# Vue + Go Starter Kit

A full-stack project template with a Go (Golang) backend and a Vue.js (Vite) frontend, designed to provide a Laravel-like developer experience.

## 🚀 Features

- **Monolith Structure**: Frontend and backend in a single repository.
- **Modern Tech Stack**: Go 1.25+, Vue 3 (Composition API), Vite.
- **Embedded Frontend**: In production, the Vue application is embedded into the Go binary using `go:embed`.
- **Developer Friendly**: 
    - Automated proxy in dev mode (Vite -> Go).
    - Shared management via `Makefile`.
    - SPA routing support out of the box.

## 🛠 Project Structure

```text
.
├── cmd/server/       # Go entry point (main.go)
├── frontend/         # Vue.js application (Vite)
├── internal/         # Backend business logic, models, services
├── Makefile          # Main command center
└── go.mod            # Go module definition
```

## 🚥 Getting Started

### Prerequisites

- Go (1.25 or later)
- Node.js (v20+) & npm

### Installation

```bash
make install
```

### Development

Run both backend and frontend simultaneously with Hot Module Replacement (HMR):

```bash
make dev
# or
npm run dev
```

- **Frontend**: http://localhost:5173
- **Backend API**: http://localhost:8080 (Proxied via frontend)

### Production Build

Create a single executable binary containing the embedded frontend:

```bash
make build
```
The resulting binary will be in `bin/server`.

## 📡 API Endpoints

- `GET /api/health`: Health check endpoint.

## 📝 License

MIT
