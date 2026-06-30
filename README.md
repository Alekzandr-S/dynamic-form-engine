# Dynamic Form Engine

A configuration-driven Dynamic Form Builder Engine built with **Go**, **PostgreSQL**, **React**, and **TypeScript**.

This application demonstrates how dynamic forms can be defined, versioned, validated, rendered, and submitted without requiring application code changes for every new form. Form layouts and validation rules are stored as JSON documents, allowing the frontend to render forms dynamically while the backend validates submissions using JSON Schema.

---

# Features

## Backend

- Dynamic form definitions
- Versioned form templates
- Draft → Publish workflow
- Runtime JSON Schema validation
- PostgreSQL JSONB storage
- RESTful API
- Layered architecture
- Docker support

## Frontend

- React + TypeScript
- Dynamic form rendering
- Field Registry pattern
- API-driven UI
- Dynamic submissions
- Loading, success and error states

---

# Technology Stack

| Layer | Technology |
|--------|------------|
| Backend | Go |
| Router | Chi |
| Database | PostgreSQL |
| Driver | pgx |
| Validation | JSON Schema |
| Frontend | React |
| Language | TypeScript |
| Build Tool | Vite |
| HTTP Client | Axios |
| Styling | TailwindCSS + shadcn/ui |
| Containers | Docker + Docker Compose |

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

Each layer has a single responsibility:

- **Handlers** receive HTTP requests.
- **Services** implement business rules.
- **Repositories** handle persistence.
- **Domain** models represent the business entities.

---

# Project Structure

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
│   │
│   ├── Dockerfile
│   └── nginx.conf
│
├── migrations/
│   ├── 001_create_tables.sql
│   └── 002_seed.sql
│
├── Dockerfile
├── docker-compose.yml
└── README.md
```

---

# Data Model

The engine separates logical forms, published templates, and submitted data.

```
Form Definition
        │
        ▼
Template Version
        │
        ▼
Submission
```

## form_definitions

Represents the logical form.

| Column |
|---------|
| id |
| name |
| description |
| created_at |
| updated_at |

---

## form_template_versions

Stores versioned form configurations.

| Column |
|---------|
| id |
| definition_id |
| version |
| status |
| ui_schema |
| validation_schema |
| created_at |

---

## form_submissions

Stores submitted user data.

| Column |
|---------|
| id |
| template_version_id |
| data |
| created_at |

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

## Form Submissions

| Method | Endpoint |
|---------|----------|
| POST | /forms/{definitionId}/submissions |

---

# Running Locally

## 1. Clone the Repository

```bash
git clone <repository-url>

cd dynamic-form-engine
```

---

## 2. Configure the Backend

Create a `.env` file:

```env
PORT=8080

DATABASE_URL=postgres://postgres:postgres@localhost:5432/dynamic_form_engine?sslmode=disable
```

---

## 3. Configure the Frontend

Create:

```
frontend/.env
```

```env
VITE_API_URL=http://localhost:8080

VITE_DEMO_FORM_ID=11111111-1111-1111-1111-111111111111
```

---

## 4. Create the Database

```sql
CREATE DATABASE dynamic_form_engine;
```

Run the migration:

```bash
psql \
-U postgres \
-d dynamic_form_engine \
-f migrations/001_create_tables.sql
```

Load the demo data:

```bash
psql \
-U postgres \
-d dynamic_form_engine \
-f migrations/002_seed.sql
```

---

## 5. Start the Backend

```bash
go mod tidy

go run ./cmd/api
```

Backend:

```
http://localhost:8080
```

---

## 6. Start the Frontend

```bash
cd frontend

npm install

npm run dev
```

Frontend:

```
http://localhost:5173
```

---

# Running with Docker

Build and start all services:

```bash
docker compose up --build
```

Services:

| Service | Port |
|----------|------|
| PostgreSQL | 5432 |
| Backend | 8080 |
| Frontend | 5173 |

After the containers are running, execute the migration and seed scripts:

```bash
psql \
-U postgres \
-d dynamic_form_engine \
-f migrations/001_create_tables.sql

psql \
-U postgres \
-d dynamic_form_engine \
-f migrations/002_seed.sql
```

---

# Demo Workflow

1. Create a form definition.
2. Create a draft template.
3. Publish the template.
4. Open the React application.
5. The frontend requests the published form.
6. The form is rendered dynamically.
7. Submit the completed form.
8. The backend validates the payload against the stored JSON Schema.
9. The submission is stored in PostgreSQL.

---

# Design Decisions

## JSONB Storage

Form templates and submissions are stored using PostgreSQL's JSONB type. This allows new forms to be introduced without modifying the database schema.

---

## Versioned Templates

Separating definitions from template versions allows:

- Draft editing
- Publishing
- Historical integrity
- Future rollback support

while ensuring submissions always remain linked to the version that produced them.

---

## Layered Architecture

The application separates:

- HTTP
- Business Logic
- Persistence

This makes the project easier to test, maintain and extend.

---

## Dynamic Rendering

The frontend renders forms dynamically using a **Field Registry** pattern. Each field type maps to a dedicated React component, allowing additional field types to be added with minimal code changes.

---

# Trade-offs

Given the assessment timeframe, the implementation prioritizes delivering a complete, working dynamic form engine over additional features.

The following were intentionally left as future enhancements:

- Authentication and authorization
- Drag-and-drop form designer
- Conditional field visibility
- Multi-step forms
- File uploads
- Automated tests
- Audit logging
- Cloud deployment
- CI/CD pipeline

The architecture was designed to accommodate these enhancements without significant refactoring.

---

# AI Usage

This project was developed with assistance from **ChatGPT (OpenAI)**.

AI assistance was used for:

- Architectural discussions
- Reviewing implementation approaches
- Explaining unfamiliar concepts
- Debugging compile/runtime issues
- Reviewing code quality
- Drafting documentation

All generated code was reviewed, adapted where necessary, compiled, tested, and understood before being included in the final submission.

---

# Future Improvements

- Role-based authentication
- Drag-and-drop form builder
- Conditional rendering rules
- Multi-page forms
- File upload fields
- Form analytics
- Full test suite
- CI/CD pipeline
- Production cloud deployment

---

# Author

**Alekzandr S**

Dynamic Form Engine — Full Stack Technical Assessment