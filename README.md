# 🔗 URL Shortener Service

A minimal REST-based URL shortener written in Go.  
It provides endpoints to shorten URLs, redirect to original URLs, and fetch domain usage metrics.

---

## 📦 Features

- 🏷️ Shorten long URLs to unique keys
- 🔁 Reuse the same short URL for identical input
- 🚀 Redirect to original URL using short path
- 📊 View top 3 domains that were shortened most
- 🧠 In-memory storage (for now)
- ✅ Unit tests included
- 🐳 Docker + Makefile support

---

## 🚀 Running the Project

### 📁 Prerequisites

- [`Colima`](https://github.com/abiosoft/colima) (or Docker)
- `docker-compose`
- `make`

---

### 🐳 Using Docker Compose (Recommended)

Start and stop the service:

```bash
make docker-build 
make docker-up
make docker-down

```

Help for Make commands

```bash
make help
```

Shorten 

```bash 
curl -X POST http://localhost:8080/shorten \
  -H "Content-Type: application/json" \
  -d '{"url": "https://www.youtube.com/watch?v=liY3Lii_o1c"}'
```

Redirect 

```bash 
curl -v http://localhost:8080/abc123 
```


Top 3 Domains 

```bash
curl http://localhost:8080/metrics/top-domains
```

