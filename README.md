# Dynamic Form Engine

A configuration-driven dynamic form engine built with **Go**, **PostgreSQL**, **React**, and **TypeScript**.

The application allows administrators to define forms through configuration instead of code. Published forms are rendered dynamically on the frontend, validated against JSON Schema on the backend, and persisted as versioned submissions.

---

## Features

### Backend

- Dynamic form definitions
- Versioned form templates
- Draft / Publish workflow
- JSON Schema validation
- Dynamic submission storage (JSONB)
- PostgreSQL persistence
- RESTful API
- Layered architecture (Handler → Service → Repository)

### Frontend

- Dynamic form rendering
- TypeScript throughout
- Field Registry pattern
- API-driven UI
- Dynamic submissions
- Component-based architecture

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

---

## Project Structure

```
dynamic-form-engine/

├── cmd/
│   └── api/
│       └── main.go
│
├── internal/
│   ├── config/
│   ├── database/
│   ├── domain/
│   ├── handler/
│   ├── repository/
│   │   ├── interfaces/
│   │   └── postgres/
│   ├── service/
│   └── validator/
│
├── frontend/
│   ├── src/
│   │   ├── api/
│   │   ├── components/
│   │   ├── pages/
│   │   ├── types/
│   │   └── App.tsx
│
├── migrations/
│
└── README.md
```

---

# Database Design

The system separates the logical form from its published versions.

```
Form Definition
        │
        │
        ▼
Form Template Version
        │
        │
        ▼
Form Submission
```

### form_definitions

Stores the logical form.

- id
- name
- description
- created_at
- updated_at

---

### form_template_versions

Stores published versions.

- id
- definition_id
- version
- status
- ui_schema
- validation_schema
- created_at

---

### form_submissions

Stores submitted data.

- id
- template_version_id
- data
- created_at

---

# Request Flow

Creating a form

```
Administrator

↓

Create Definition

↓

Create Draft Version

↓

Publish
```

Submitting a form

```
User

↓

GET Published Form

↓

React renders UI Schema

↓

User submits data

↓

Backend validates JSON Schema

↓

Submission stored
```

---

# Technologies

Backend

- Go
- Chi Router
- PostgreSQL
- pgx
- JSON Schema

Frontend

- React
- TypeScript
- Vite
- Axios
- TailwindCSS
- shadcn/ui

---

# Running the Project

## Backend

Create:

```
.env
```

```
PORT=8080

DATABASE_URL=postgres://username:password@localhost:5432/dynamic_form_engine?sslmode=disable
```

Install dependencies

```bash
go mod tidy
```

Run

```bash
go run ./cmd/api
```

Backend runs on

```
http://localhost:8080
```

---

## Frontend

Create

```
frontend/.env
```

```
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

# API Endpoints

## Form Definitions

| Method | Endpoint |
|---------|----------|
| POST | /definitions |

---

## Template Versions

| Method | Endpoint |
|---------|----------|
| POST | /definitions/{id}/versions |
| POST | /definitions/{id}/publish |

---

## Dynamic Forms

| Method | Endpoint |
|---------|----------|
| GET | /forms/{definitionId} |

---

## Submissions

| Method | Endpoint |
|---------|----------|
| POST | /forms/{definitionId}/submissions |

---

# Example Workflow

1. Create a form definition.

2. Create Version 1 of the template.

3. Publish Version 1.

4. React retrieves the published UI Schema.

5. User completes the form.

6. Backend validates the submission.

7. Submission is stored in PostgreSQL.

---

# Design Decisions

## Why JSONB?

The engine stores form schemas and submissions as JSONB to support arbitrary form structures without requiring database schema changes whenever a new form is introduced.

---

## Why Separate Definitions and Versions?

Separating form definitions from template versions allows:

- historical integrity
- versioning
- draft workflows
- publishing
- future rollback support

without affecting existing submissions.

---

## Why Layered Architecture?

The project separates responsibilities into:

- Handlers
- Services
- Repositories

This keeps HTTP concerns, business logic, and persistence independent, making the application easier to maintain and test.

---

## Future Improvements

Given additional time, the following features could be added:

- Authentication and authorization
- Admin UI for form management
- Drag-and-drop form builder
- Conditional field visibility
- Multi-step forms
- File uploads
- Form analytics and reporting
- Soft deletes and audit logging
- Docker Compose deployment
- Automated tests
- CI/CD pipeline

---

# Author

**Alekzandr S**

Dynamic Form Engine Assessment