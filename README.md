# KCard Design

AI-assisted card review web app for individual learners. The project is organized as a separated Go backend and Vue 3 frontend.

## Structure

- `backend/` — Gin API, worker entrypoints, domain, repositories, AI workflow adapters, Redis jobs, storage, migrations.
- `frontend/` — Vue 3 + TypeScript SPA with dashboard, materials, knowledge list/graph, cards, review, plans, and prompts modules.
- `contracts/openapi.yaml` — synced OpenAPI contract used by backend tests and frontend clients.
- `scripts/dev.sh` — local development command notes.
- `scripts/test.sh` — full validation wrapper.

## Local setup

1. Copy `.env.example` to `.env` and fill local PostgreSQL, Redis, storage, and Anthropic settings.
2. Start PostgreSQL 16 and Redis 7.
3. Run backend API with `cd backend && go run ./cmd/api`.
4. Run worker with `cd backend && go run ./cmd/worker`.
5. Run frontend with `cd frontend && npm install && npm run dev`.

## Validation

Use `./scripts/test.sh` for backend and frontend validation once dependencies are installed. See `specs/001-card-review-app/quickstart.md` for end-to-end scenarios.
