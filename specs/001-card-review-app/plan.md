# Implementation Plan: AI Card Review Web App

**Branch**: `001-card-review-app` | **Date**: 2026-06-09 | **Spec**: [spec.md](./spec.md)

**Input**: Feature specification from `/specs/001-card-review-app/spec.md`

**Note**: This template is filled in by the `/speckit-plan` command. See `.specify/templates/plan-template.md` for the execution workflow.

## Summary

Build a frontend/backend separated web app for individual learners to import study materials,
curate AI-extracted knowledge points, generate decks/cards, manage decks/cards, and review
cards through direct sessions or adaptive review plans. The backend will be a Go service using
Gin for HTTP routing, CloudWeGo Eino for AI workflow orchestration, GORM for persistence,
PostgreSQL for durable data, Redis for queues/cache/session-like ephemeral state, and the
Anthropic Go SDK for Claude-powered classification, card generation, and plan optimization.
The frontend will be a Vue 3 single-page app with a consistent learner dashboard, material
review, card/deck management, and review-plan experience using #fff8e7 as the primary soft
study background/base tone and #f8e7ff plus #e7fff8 as supporting accent backgrounds.

## Technical Context

**Language/Version**: Go 1.24+ for backend; TypeScript 5.x and Vue 3.5+ for frontend

**Primary Dependencies**: Backend: Gin, CloudWeGo Eino, GORM, PostgreSQL driver, Redis client,
Anthropic Go SDK, structured logging, validation, OpenAPI generation/validation. Frontend:
Vue 3, Vite, Vue Router, Pinia, TanStack Query or equivalent server-state layer, UI component
library chosen during implementation, charting library for statistics.

**Storage**: PostgreSQL for learner workspaces, materials, knowledge points, prompt presets,
decks, cards, review sessions, review plans, plan revisions, and statistics snapshots; Redis
for AI job queues, idempotency locks, short-lived progress events, cache entries, and rate-limit
coordination; object/local storage abstraction for uploaded source files and exports.

**Testing**: Backend unit tests with `go test`, handler tests with `httptest`, repository and
contract tests with disposable PostgreSQL/Redis containers, AI workflow tests with mocked model
responses and schema fixtures. Frontend unit/component tests with Vitest and Vue Test Utils,
end-to-end flows with Playwright, contract validation against OpenAPI examples.

**Target Platform**: Browser-based web app for desktop and tablet browsers; backend deployed as
a Linux containerized web service with worker capability for long-running AI jobs.

**Project Type**: Web application with separated backend API, background worker, and Vue 3 SPA.

**Performance Goals**: Show feedback within 2 seconds for long-running AI/import/export/plan
operations; deck/card filtering and view switching visible within 1 second for 500 decks and
5,000 cards; typical material under 5,000 words produces reviewable AI output or clear failure
within 2 minutes; review session answer recording completes within 500 ms from learner action.

**Constraints**: AI outputs are drafts until approved; no collaboration/sharing in first
release; learner workspace content private by default; destructive actions require confirmation
and recoverability where possible; long AI tasks must be asynchronous with resumable progress;
contracts must not expose implementation-specific internals; learner-facing screens must use
#fff8e7 as the primary background/base tone with #f8e7ff and #e7fff8 as supporting accent
backgrounds while preserving readable text, controls, focus indicators, charts, and state
feedback.

**Scale/Scope**: First release supports one private learner workspace per account; target per
workspace up to 200 materials, 1,000 knowledge points, 500 decks, 5,000 cards, 100 prompt
presets, and 20 active or paused review plans.

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

- **Code Quality**: PASS. Backend and frontend will use typed code, clear module boundaries,
  lint/format/static analysis gates, OpenAPI contracts, and AI workflow adapters isolated from
  domain services. New abstractions are justified by the separated web app, asynchronous AI
  jobs, and AI orchestration requirements.
- **Testing Standards**: PASS. Each user story has automated unit, contract, integration, and
  end-to-end validation paths. AI behavior will be tested through deterministic schema fixtures,
  mocked model responses, and worker/job state tests; manual validation is limited to final UX
  review scenarios documented in quickstart.md.
- **User Experience Consistency**: PASS. Plan defines shared UX states for loading, empty,
  processing, partial success, disabled, undo, failure, and completion across materials,
  knowledge curation, card/deck management, review sessions, dashboard, and plan optimization.
  It also carries the required soft study palette: #fff8e7 primary background/base tone,
  #f8e7ff and #e7fff8 supporting accent backgrounds, and readable controls/text/charts/focus
  states on all three colors.
- **Performance Requirements**: PASS. Measurable budgets are defined for visible feedback,
  filtering, AI job completion, and review-session interactions. Long-running AI work is
  queued, progress-tracked, and streamed/polled rather than blocking foreground requests.

**Post-design re-check**: PASS. Research, data model, contracts, and quickstart preserve the
same gates. No unresolved clarifications remain.

## Project Structure

### Documentation (this feature)

```text
specs/001-card-review-app/
├── plan.md              # This file (/speckit-plan command output)
├── research.md          # Phase 0 output (/speckit-plan command)
├── data-model.md        # Phase 1 output (/speckit-plan command)
├── quickstart.md        # Phase 1 output (/speckit-plan command)
├── contracts/           # Phase 1 output (/speckit-plan command)
│   └── openapi.yaml     # Backend HTTP API contract for frontend and tests
└── tasks.md             # Phase 2 output (/speckit-tasks command - NOT created by /speckit-plan)
```

### Source Code (repository root)

```text
backend/
├── cmd/
│   ├── api/                 # Gin HTTP server entrypoint
│   └── worker/              # AI/import/export/review-plan background worker entrypoint
├── internal/
│   ├── app/                 # application bootstrap and dependency wiring
│   ├── config/              # environment/config loading and validation
│   ├── domain/              # entities, value objects, domain rules
│   ├── service/             # use cases for materials, cards, review, plans, dashboard
│   ├── repository/          # GORM persistence adapters
│   ├── transport/http/      # Gin handlers, middleware, request/response DTOs
│   ├── ai/                  # Eino workflows, Anthropic client adapter, schemas, prompts
│   ├── jobs/                # Redis-backed jobs, progress events, idempotency
│   ├── storage/             # uploaded material/export storage abstraction
│   └── observability/       # logging, metrics, tracing helpers
├── migrations/              # PostgreSQL migrations
├── test/                    # integration fixtures and contract tests
├── go.mod
└── go.sum

frontend/
├── src/
│   ├── app/                 # app bootstrap, router, providers
│   ├── assets/
│   ├── components/          # shared UI states and reusable components
│   ├── features/
│   │   ├── dashboard/
│   │   ├── materials/
│   │   ├── knowledge/
│   │   ├── cards/
│   │   ├── review/
│   │   ├── plans/
│   │   └── prompts/
│   ├── services/            # API clients generated or validated from OpenAPI
│   ├── stores/              # Pinia stores for local UI state
│   ├── styles/              # design tokens, required palette, and global styles
│   └── tests/               # component and unit test helpers
├── e2e/                     # Playwright end-to-end scenarios
├── package.json
└── vite.config.ts

contracts/
└── openapi.yaml             # optional synced copy or generated artifact from spec contract

scripts/
├── dev.sh                   # local backend/frontend dependencies helper
└── test.sh                  # full validation command wrapper
```

**Structure Decision**: Use a monorepo-style separated web application with `backend/` and
`frontend/` top-level directories. The backend exposes HTTP contracts and owns data/AI jobs;
the frontend owns learner workflows and consumes the contract. A separate `backend/cmd/worker`
entrypoint is required because material analysis, card generation, exports, imports, and plan
optimization exceed interactive request lifetimes.

## Complexity Tracking

> **Fill ONLY if Constitution Check has violations that must be justified**

| Violation | Why Needed | Simpler Alternative Rejected Because |
|-----------|------------|-------------------------------------|
| None | N/A | N/A |
