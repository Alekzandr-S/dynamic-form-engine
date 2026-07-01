# Dynamic Form Engine

A full-stack, configuration-driven Dynamic Form Engine built with **Go**, **PostgreSQL**, **React**, and **TypeScript**.

Instead of hardcoding forms into application code or database tables, this application stores form configurations as JSON, dynamically renders them on the frontend, validates submissions using JSON Schema, and persists responses as JSONB while preserving the template version that produced each submission.

---

# Live Demo

## Frontend

https://dashboard-six-blond-71.vercel.app/

## Backend API

https://dynamic-form-engine-api.onrender.com

Health Check

https://dynamic-form-engine-api.onrender.com/health

---

# Features

## Backend

- Dynamic form definitions
- Versioned form templates
- Draft / Publish workflow
- JSON Schema validation
- Dynamic submission storage
- PostgreSQL JSONB persistence
- RESTful API
- Layered architecture
- Docker support
- Environment-based configuration

---

## Frontend

- React + TypeScript
- Dynamic form rendering
- Field Registry pattern
- API-driven UI
- Loading, success and error states
- Component-based architecture
- Responsive layout

---

# Architecture

```
                    React Frontend
                           │
                           ▼
                    HTTP Handlers
                           │
                           ▼
                     Service Layer
                           │
                           ▼
                  Repository Layer
                           │
                           ▼
                     PostgreSQL
```

The application follows a layered architecture to separate HTTP concerns, business logic and persistence.

---

# Data Model

The system separates a logical form from its published versions.

```
Form Definition

        │

        ▼

Form Template Version

        │

        ▼

Form Submission
```

## form_definitions

Stores the logical definition of a form.

| Column | Description |
|---------|-------------|
| id | Primary Key |
| name | Form name |
| description | Form description |
| created_at | Created timestamp |
| updated_at | Updated timestamp |

---

## form_template_versions

Stores versioned templates.

| Column | Description |
|---------|-------------|
| id | Primary Key |
| definition_id | FK → Form Definition |
| version | Version number |
| status | Draft / Published / Archived |
| ui_schema | JSON UI definition |
| validation_schema | JSON Schema |
| created_at | Created timestamp |

---

## form_submissions

Stores user submissions.

| Column | Description |
|---------|-------------|
| id | Primary Key |
| template_version_id | FK → Published Template |
| data | JSONB submission |
| created_at | Created timestamp |

---

# Dynamic Workflow

## Administrator

```
Create Definition

↓

Create Template Version

↓

Publish Version
```

---

## User

```
Request Form

↓

Backend returns UI Schema

↓

React renders dynamically

↓

User submits form

↓

Backend validates JSON Schema

↓

Submission stored
```

---

# Technology Stack

## Backend

- Go
- Chi Router
- PostgreSQL
- pgx
- JSON Schema Validator

## Frontend

- React
- TypeScript
- Vite
- Axios
- Tailwind CSS
- shadcn/ui

## Deployment

- Render
- Vercel
- Neon PostgreSQL
- Docker

---

# Running Locally

## Prerequisites

- Go
- Node.js
- PostgreSQL
- Docker (optional)

---

## Backend

Create

```
.env
```

```env
PORT=8080

DATABASE_URL=<postgres connection string>
```

Install dependencies

```bash
go mod tidy
```

Run

```bash
go run ./cmd/api
```

Backend

```
http://localhost:8080
```

---

## Frontend

Create

```
frontend/.env
```

```env
VITE_API_URL=http://localhost:8080
```

Install

```bash
npm install
```

Run

```bash
npm run dev
```

Frontend

```
http://localhost:5173
```

---

# Docker

## Backend

```bash
docker build -t dynamic-form-engine-backend .

docker run \
-p 8080:8080 \
--env-file .env \
dynamic-form-engine-backend
```

---

## Frontend

```bash
cd frontend

docker build -t dynamic-form-engine-frontend .

docker run -p 3000:80 dynamic-form-engine-frontend
```

---

# API

## Form Definitions

| Method | Endpoint |
|---------|----------|
| GET | /definitions |
| POST | /definitions |

---

## Template Versions

| Method | Endpoint |
|---------|----------|
| POST | /definitions/{id}/versions |
| POST | /definitions/{id}/publish |

---

## Forms

| Method | Endpoint |
|---------|----------|
| GET | /forms/{definitionId} |

---

## Submissions

| Method | Endpoint |
|---------|----------|
| POST | /forms/{definitionId}/submissions |

---

# Example Flow

1. Create a Form Definition.

2. Create Version 1.

3. Publish Version 1.

4. React fetches the published UI Schema.

5. The form is rendered dynamically.

6. User submits data.

7. Backend validates against the stored JSON Schema.

8. Submission is stored in PostgreSQL.

---

# Design Decisions

## Why JSONB?

JSONB allows arbitrary form structures to be stored without requiring schema migrations whenever new fields are introduced.

---

## Why Separate Definitions and Versions?

Separating definitions from template versions preserves historical integrity. Existing submissions remain associated with the exact template version that generated them, even after newer versions are published.

---

## Why JSON Schema?

Validation rules are stored alongside each template version, enabling runtime validation without hardcoding field-specific logic into the application.

---

## Why Layered Architecture?

The application separates:

- HTTP Handlers
- Services
- Repositories

This improves maintainability, testability and keeps business logic independent from transport and persistence concerns.

---

# Trade-offs

Given the assessment timeframe, the implementation focuses on delivering a complete end-to-end dynamic form engine.

The following were intentionally deferred:

- Authentication & authorization
- Role-based administration
- Drag-and-drop form designer
- Automated unit and integration tests
- Form analytics
- Conditional field visibility
- Audit logging
- Multi-step forms

The architecture was designed so these features can be introduced with minimal refactoring.

---

# AI Usage

AI assistance was used throughout the project in accordance with the assessment guidelines.

## Tools Used

- ChatGPT (GPT-5.5)

## How AI Was Used

- Architectural discussions
- Project planning
- Code review
- Debugging
- Repository pattern guidance
- Docker configuration
- README generation
- Frontend scaffolding
- Backend scaffolding
- API design

## What Was Verified

All generated code was manually reviewed, integrated, compiled, tested, and modified where necessary. The complete application—including the architecture, business logic, API behavior, and deployment—was implemented, understood, and verified before submission.

---

# Future Improvements

Given additional development time, the project could be extended with:

- Authentication & RBAC
- Admin Portal
- Visual Form Builder
- Conditional fields
- Multi-step forms
- File uploads
- Audit trail
- Automated testing
- CI/CD pipeline
- Docker Compose deployment
- Kubernetes deployment

---

# Author

**Alekzandr S**

Dynamic Form Engine Assessment