# AGENTS.md

# Doko

Doko is a backend service for a novel reading platform.

Tech Stack:

* Go
* Fiber
* SQLite
* SQLC
* Viper
* JWT
* Swagger

Goal:

Build a maintainable, feature-based backend for a novel reading website.

---

# Development Rules

## Architecture

Use:

```text
Feature-based Modular Monolith
```

Organize by feature:

```text
internal/

    auth/
    user/
    novel/
    chapter/
    reading/
    bookshelf/
    comment/
    ranking/
    admin/
```

Forbidden:

```text
controller/
service/
repository/
```

as global layers.

Each module should maintain:

```text
handler/
service/
repository/
dto/
routes/
```

Example:

```text
internal/

    auth/
        handler/
        service/
        repository/
        dto/
        routes/

    user/
        handler/
        service/
        repository/
        dto/
        routes/
```

---

## Dependency Rules

Dependency flow:

```text
handler
    ↓
service
    ↓
repository
    ↓
sqlc
```

Allowed:

```text
feature -> pkg
feature -> db/sqlc
feature -> shared utilities (config, response, errors)
```

Cross-module interaction should be done through:

* service layer calls (preferred when needed)
* shared domain models or DTO reuse when appropriate

Avoid direct dependency on internal implementation details of other features.

---

## SQL Rules

The project follows:

```text
SQL First
```

Schema files:

```text
internal/db/schema/
```

Query files:

```text
internal/db/query/
```

Generate code using:

```bash
sqlc generate
```

Rules:

* All data access must go through sqlc-generated code
* No ORM usage
* No GORM
* No scattered SQL outside query directory

---

## Config Rules

Use:

```text
Viper
```

Configuration file:

```text
configs/app.yaml
```

All configuration access must go through:

```text
internal/config
```

Avoid using:

```go
os.Getenv()
```

inside business logic.

---

## API Rules

Unified response format.

Success:

```json
{
  "code": 0,
  "message": "ok",
  "data": {}
}
```

Failure:

```json
{
  "code": 40001,
  "message": "error"
}
```

Rules:

* Never return raw errors directly to clients
* All handlers must use unified response wrapper

---

## JWT Rules

Authentication uses:

```text
JWT Access Token
```

Rules:

* Token generation and validation centralized in pkg/jwt
* User identity injected via middleware into request context
* Handlers must not parse tokens directly

---

## Swagger Rules

Use:

```text
swaggo/swag
fiber-swagger
```

Rules:

* All handlers must include Swagger annotations
* Swagger must stay in sync with code changes

Endpoint:

```text
/swagger/index.html
```

Generation command:

```bash
swag init -g cmd/api/main.go
```

---

## Project Structure

```text
cmd/
    api/
        main.go

configs/
    app.yaml

internal/

    config/

    db/
        schema/
        query/
        sqlc/

    middleware/
    response/
    errors/
    router/
    server/
    health/

    auth/
    user/
    novel/
    chapter/
    reading/
    bookshelf/
    comment/
    ranking/
    admin/

pkg/
    jwt/

docs/

sqlc.yaml
```

---

# Milestones

## P0 — Foundation

Goal:

Build a runnable backend skeleton.

Scope:

* Fiber initialization
* Project structure
* Viper configuration
* SQLite connection
* SQLC setup
* JWT utilities
* Middleware base
* Error handling
* Unified response
* Swagger setup
* Health check

---

### P0.1 — Fiber Bootstrap

New Dependencies:

* Fiber

Scope:

* Initialize Fiber app
* Basic router setup
* Start HTTP server

Acceptance Criteria:

```http
GET /health
```

returns:

```json
{
  "code": 0,
  "message": "ok"
}
```

---

### P0.2 — Config System

New Dependencies:

* Viper

Scope:

* Load configs/app.yaml
* Central config module

Acceptance Criteria:

* Config is loaded successfully at startup
* No hardcoded server config

---

### P0.3 — Database Base

New Dependencies:

* SQLite driver

Scope:

* Initialize SQLite connection
* Graceful shutdown support

Acceptance Criteria:

* DB connection established successfully

---

### P0.4 — SQLC

New Dependencies:

* sqlc

Scope:

* Define schema and queries
* Generate Go code using sqlc

Acceptance Criteria:

```bash
sqlc generate
```

runs successfully

---

### P0.5 — JWT

New Dependencies:

* JWT library

Scope:

* Token generation
* Token validation utilities

Acceptance Criteria:

* JWT functions are usable and testable

---

### P0.6 — Middleware

Scope:

* Recovery middleware
* Logging middleware
* Auth middleware (basic structure)

Acceptance Criteria:

* Middleware chain works in Fiber

---

### P0.7 — Response & Error Handling

Scope:

* Unified response format
* Error wrapper system
* Global error handler

Acceptance Criteria:

* All APIs return standardized JSON format

---

### P0.8 — Swagger

New Dependencies:

* swaggo/swag
* fiber-swagger

Scope:

* Swagger setup
* Handler annotations

Acceptance Criteria:

* Swagger UI accessible at:

```text
/swagger/index.html
```

---

### P0.9 — Health Check

Scope:

* Final integration check

Acceptance Criteria:

```http
GET /health
```

works in production-like run

Swagger works

---

## P1 — Authentication & User

Modules:

```text
auth
user
```

Scope:

* register
* login
* logout
* refresh token
* current user
* update profile

Endpoints:

```http
POST /auth/register
POST /auth/login
POST /auth/logout
POST /auth/refresh

GET /me
PUT /me
```

---

## P2 — Novel Reading Core

Modules:

```text
novel
chapter
reading
```

Scope:

* novel list
* novel detail
* categories
* tags
* chapter list
* chapter reading
* reading progress
* recent reading

Endpoints:

```http
GET /novels
GET /novels/:id

GET /novels/:id/chapters
GET /chapters/:id

GET /me/reading
```

---

## P3 — Bookshelf & Community

Modules:

```text
bookshelf
comment
```

Scope:

* favorite novels
* bookshelf
* comments
* likes

Endpoints:

```http
POST /bookshelf/:id
DELETE /bookshelf/:id

GET /bookshelf

POST /comments
GET /comments
```

---

## P4 — Ranking & Admin

Modules:

```text
ranking
admin
```

Scope:

* hot ranking
* latest ranking
* recommended ranking
* novel management
* chapter management
* user management

---

# Workflow

Before development, confirm:

```text
Current Milestone:
Goal:
Files To Create:
New Dependencies:
Acceptance Criteria:
```

Rules:

1. Only implement current milestone
2. Do not implement future milestones
3. Introduce dependencies incrementally
4. Keep code runnable after each step
5. Keep build green after each commit
6. Swagger must stay in sync with handlers

---

# Done Definition

A task is complete only if:

* build success
* tests pass
* runs locally
* sqlc generate success (if applicable)
* swagger generate success (if applicable)
* system remains runnable
