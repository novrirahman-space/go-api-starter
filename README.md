# 🚀 Go API Starter

[![CI](https://github.com/novrirahman-space/go-api-starter/actions/workflows/ci.yml/badge.svg)](https://github.com/novrirahman-space/go-api-starter/actions/workflows/ci.yml)
[![Release](https://github.com/novrirahman-space/go-api-starter/actions/workflows/release.yml/badge.svg)](https://github.com/novrirahman-space/go-api-starter/actions/workflows/release.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/novrirahman-space/go-api-starter)](https://goreportcard.com/report/github.com/novrirahman-space/go-api-starter)
[![Docker](https://img.shields.io/badge/docker-ghcr.io%2Fnovrirahman--space%2Fgo--api--starter-blue)](https://ghcr.io/novrirahman-space/go-api-starter)
[![License](https://img.shields.io/github/license/novrirahman-space/go-api-starter)](LICENSE)

> 🧩 **Production-ready Go REST API starter kit** — built for real-world services with Docker, GitHub Actions CI/CD, multi-arch releases, Prometheus metrics, structured logging, and OpenAPI documentation.

---

## 📚 Features

✅ **Built-in Middlewares**
- CORS, Rate Limiter, RequestID, Request Logger  
- Prometheus metrics and tracing hooks  

✅ **Observability**
- `/healthz` — Liveness probe  
- `/readyz` — Readiness probe  
- `/metrics` — Prometheus metrics

✅ **Developer Experience**
- Local dev with `air` or `make run`
- Modular internal packages (`config`, `logger`, `middleware`, `handlers`, `server`)
- Ready for `go test ./... -race -cover`

✅ **CI/CD Ready**
- ✅ GitHub Actions for build/test/lint/release  
- ✅ GoReleaser v2 pipeline → binary + multi-arch Docker image  
- ✅ Auto-publish to GHCR (`ghcr.io/novrirahman-space/go-api-starter`)  

✅ **Security**
- Distroless Docker image (no shell, minimal attack surface)  
- Rate limiting & graceful shutdown  

✅ **Docs & API Spec**
- `/docs` (Redoc auto-renderer)
- `api/openapi.yaml` (OpenAPI 3.0.3 spec)

---

## ⚙️ Project Structure

```bash
go-api-starter/
├── cmd/
│   └── server/               # Main server entrypoint
├── internal/
│   ├── config/               # Config loader (env)
│   ├── handlers/             # HTTP handlers (healthz, users, version, etc.)
│   ├── logger/               # Zerolog setup
│   ├── middleware/           # CORS, rate limit, metrics, etc.
│   ├── server/               # Router + HTTP server setup
├── api/
│   └── openapi.yaml          # API spec (served by /docs)
├── Dockerfile                # Local development Dockerfile
├── Dockerfile.goreleaser     # Runtime-only distroless image
├── .goreleaser.yaml          # Release automation config
├── .github/workflows/        # CI/CD pipelines
├── go.mod
├── go.sum
└── README.md
```

---

## 🧪 Local Development

### 1️⃣ Run directly with Go
```bash
go run ./cmd/server
```

### 2️⃣ Using Docker
```bash
docker build -t go-api-starter .
docker run -p 8080:8080 go-api-starter
```

### 3️⃣ Using Docker Compose (recommended for local stack)
```bash
docker compose up --build
```

```bash
This will spin up the API container and expose port 8080.
You can then test the endpoints using curl or Postman.
```

### 4️⃣ Run tests and coverage
```bash
go test ./... -race -coverprofile=coverage.out
go tool cover -html=coverage.out
```

---

## 🌐 API Endpoints

| Endpoint        | Method          | Description                         |
| --------------- | --------------- | ----------------------------------- |
| `/healthz`      | GET             | Liveness probe                      |
| `/readyz`       | GET             | Readiness probe                     |
| `/metrics`      | GET             | Prometheus metrics                  |
| `/docs`         | GET             | OpenAPI Redoc docs                  |
| `/openapi.yaml` | GET             | Raw OpenAPI YAML                    |
| `/v1/users`     | GET/POST/DELETE | Example CRUD (in-memory)            |

---

## 🚀 Deployment

### 🐳 Run from GitHub Container Registry
```bash
docker pull ghcr.io/novrirahman-space/go-api-starter:latest
docker run -p 8080:8080 ghcr.io/novrirahman-space/go-api-starter:latest
```

### ☁️ Multi-architecture Support
The container supports:
- linux/amd64
- linux/arm64
- darwin/amd64
- windows/amd64
The manifest auto-selects the correct architecture for your platform.

### 🔁 CI/CD Overview

| Stage                     | Description                                                              |
| ------------------------- | ------------------------------------------------------------------------ |
| **CI (ci.yml)**           | Runs lint, tests, and uploads coverage                                   |
| **Release (release.yml)** | Tag-based release automation                                             |
| **GoReleaser**            | Builds multi-arch Docker images, pushes to GHCR, creates GitHub Releases |

## 🧩 Observability

### Prometheus Metrics
Available at /metrics

### Health Endpoints

| Endpoint   | Purpose                                     |
| ---------- | ------------------------------------------- |
| `/healthz` | Check if API is alive                       |
| `/readyz`  | Check if dependencies (DB, cache) are ready |


## 🧭 OpenAPI Documentation
OpenAPI spec file:
api/openapi.yaml

Preview via:
```bash
http://localhost:8080/docs
```

## 🧑‍💻 Contribution Guide

### 💬 Conventional Commits
All commits follow Conventional Commits:
```bash
feat(auth): add JWT middleware
fix(router): resolve panic on shutdown
chore(ci): update GoReleaser config
```

### ✅ Lint & Test before push
```bash
go vet ./...
go fmt ./...
go test ./...
```

### 🚧 Branch Naming Convention
Use clear prefixes:
- feat/feature-name
- fix/bug-description
- chore/ci-pipeline
- docs/readme-update

## 🧱 Tech Stack
| Layer         | Tool                        | Purpose                      |
| ------------- | --------------------------- | ---------------------------- |
| **Language**  | Go 1.22+                    | Core API development         |
| **Router**    | Chi v5                      | Lightweight HTTP router      |
| **Logging**   | Zerolog                     | Structured JSON logging      |
| **Metrics**   | Prometheus                  | Observability and alerts     |
| **Docs**      | Redoc / OpenAPI             | API documentation            |
| **CI/CD**     | GitHub Actions + GoReleaser | Automated testing & releases |
| **Container** | Distroless                  | Minimal, secure base image   |
| **Registry**  | GHCR                        | Private/public image hosting |

## 🌟 Support & Feedback
✅ If you find this project useful, give it a ⭐ on GitHub!

✅ If you find bugs or have ideas, please open an issue.