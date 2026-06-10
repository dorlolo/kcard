# Research: AI Card Review Web App

## Decision: Separated Vue 3 frontend and Go backend

**Rationale**: The feature combines a rich browser workflow with long-running backend AI jobs,
review scheduling, data persistence, imports/exports, and statistics. A separated frontend and
backend lets the Vue app optimize interactive UX while the Go backend owns authorization,
persistence, AI workflow orchestration, background jobs, and API contracts.

**Alternatives considered**:
- Single full-stack app: simpler deployment, but mixes AI/background job concerns with UI
  delivery and weakens contract testing.
- Backend-rendered pages only: reduces frontend tooling, but is less suitable for interactive
  draft review, card editing, review sessions, dashboards, and progress updates.

## Decision: Go + Gin for backend HTTP API

**Rationale**: Gin provides a lightweight, explicit routing model for JSON APIs, middleware,
request validation, and streaming/progress endpoints. Go gives predictable concurrency for
background work, strong typing for domain/service layers, and efficient runtime characteristics
for the target web service and worker.

**Alternatives considered**:
- Go standard library router only: viable but requires more repeated middleware and validation
  plumbing for a product with many routes.
- Heavier backend framework: more batteries included, but adds complexity beyond the current
  release needs.

## Decision: CloudWeGo Eino for AI workflow orchestration

**Rationale**: AI classification, card generation, and plan optimization are multi-step
workflows involving prompt presets, source context preparation, structured model output,
validation, retry/fallback behavior, and post-processing. Eino is a Go-native AI workflow
orchestration framework that can model these flows as reusable chains/graphs while keeping AI
logic separate from HTTP handlers and domain persistence.

**Alternatives considered**:
- Direct model calls in services: simplest initially, but becomes hard to test and evolve when
  adding prompt presets, structured output validation, partial result handling, and retries.
- External workflow service: powerful but unnecessary for the first release and increases
  operational footprint.

## Decision: GORM + PostgreSQL for durable persistence

**Rationale**: The domain contains relational entities with histories, state transitions, tags,
composition records, and review statistics. PostgreSQL provides reliable transactions,
constraints, JSON support for controlled metadata, and query capability for dashboards. GORM
fits the requested stack and provides migrations/model mapping while still allowing explicit SQL
for performance-sensitive statistics queries.

**Alternatives considered**:
- Document database: flexible for AI drafts, but weaker for relational integrity across source
  materials, knowledge points, cards, decks, plans, and revisions.
- Raw SQL only: maximizes control, but increases repetitive repository code for the first
  release.

## Decision: Redis for queues, progress, idempotency, and short-lived cache

**Rationale**: AI analysis, card generation, import/export, and plan optimization can exceed
foreground request duration. Redis supports lightweight job queues, progress/status events,
idempotency locks, retry coordination, short-lived generated result caches, and rate-limit
coordination for AI calls.

**Alternatives considered**:
- PostgreSQL-only jobs: acceptable for low volume but less convenient for progress events,
  short-lived cache, and bursty AI job coordination.
- External queue service: scalable but unnecessary for the local-first planning scope.

## Decision: Claude API via official Anthropic Go SDK

**Rationale**: The backend is Go, and the official Anthropic Go SDK supports Messages API,
streaming, tool use/manual loops, structured outputs, prompt caching, Files API, token counting,
and beta tool runner patterns. Use the SDK instead of raw HTTP to get typed request/response
handling, retries, and SDK-level API compatibility.

**Model default**: Use `claude-opus-4-8` for AI classification, card generation, and plan
optimization unless product settings explicitly choose another model. It is the latest default
recommended Claude model and supports 1M context, structured outputs, adaptive thinking, and
high effort modes. Cost-sensitive future routes may be configurable, but the plan defaults to
quality for learner-facing generated study content.

**Request guidance**:
- Use `thinking: {type: "adaptive"}` for complex generation and plan optimization.
- Use `output_config.effort` per workflow: `medium` for routine classification cleanup,
  `high` for card generation and plan optimization, and higher only after measurement.
- Do not use `budget_tokens`, `temperature`, `top_p`, or `top_k` with Opus 4.8.
- Use streaming for long outputs or large `max_tokens`; accumulate final message in the Go SDK
  by streaming and `Accumulate`-ing events.
- Use exact model ID `claude-opus-4-8`; do not append a date suffix.

**Alternatives considered**:
- Raw HTTP: usable but creates avoidable error-handling and streaming complexity in Go.
- Provider-neutral wrapper only: useful at boundaries, but initial implementation still needs a
  concrete Anthropic adapter for structured outputs, caching, token counting, and streaming.

## Decision: Structured outputs for AI-produced domain drafts

**Rationale**: Knowledge points, card drafts, deck drafts, review-plan proposals, duplicate
flags, and plan optimization suggestions need deterministic validation before becoming drafts.
The Anthropic API supports structured JSON output through `output_config.format`; schemas should
be versioned in `backend/internal/ai/schemas` and validated before persistence.

**Alternatives considered**:
- Free-form text parsing: fast to prototype but fragile, hard to test, and unsafe for workflows
  that create persistent drafts.
- Tool calls for every output: useful for actions, but structured outputs are simpler for pure
  extraction/generation results.

## Decision: Asynchronous AI jobs with progress events

**Rationale**: Material analysis, card generation, import/export, and plan optimization may take
seconds to minutes. Foreground API requests should enqueue a job and return a job ID. The UI
polls or subscribes to progress and can show retry, partial result, failure, or completion states
within 2 seconds.

**Alternatives considered**:
- Synchronous requests: simpler, but violates UX/performance requirements and risks timeouts.
- Fully managed external workflow engine: not needed for the first release.

## Decision: Prompt presets as first-class user-managed records

**Rationale**: The spec requires default prompts plus learner-customized prompts for material
classification, card generation, and plan creation/optimization. Persisting prompt presets makes
AI behavior repeatable and auditable. AI-generated trusted items should store the prompt preset
or ad hoc prompt snapshot used to create them.

**Alternatives considered**:
- Hard-coded prompts only: simpler, but fails user customization and audit requirements.
- Free text per run only: flexible, but not reusable or traceable.

## Decision: OpenAPI contract for backend/frontend integration

**Rationale**: A separated frontend/backend project needs a stable contract for handlers,
frontend API clients, contract tests, and future task generation. OpenAPI documents request and
response shapes while leaving implementation details in the plan/tasks.

**Alternatives considered**:
- Informal endpoint list: insufficient for contract validation.
- GraphQL: powerful for flexible querying, but adds complexity and is not necessary for the
  current CRUD/job/review workflows.

## Decision: Vue 3 SPA with explicit feature modules and shared UX states

**Rationale**: The frontend has several independent but connected workflows: dashboard,
materials, knowledge curation, cards/decks, prompts, review sessions, and plans/statistics.
Vue 3 with TypeScript, Vite, Vue Router, Pinia, and a server-state query layer supports clear
feature boundaries, testability, and responsive UI states.

**Alternatives considered**:
- Minimal vanilla frontend: inadequate for the amount of stateful interaction.
- Multiple separate apps: unnecessary for a single learner product and would harm UX
  consistency.

## Decision: Obsidian-like knowledge relationship graph

**Rationale**: Recorded knowledge needs both a dense list view for fast review and a graph view
for understanding relationships among concepts, sources, tags, duplicate clusters, split/merge
lineage, prerequisites, and generated cards. Store typed relationship edges in PostgreSQL and
serve a filtered graph payload to the Vue frontend. The graph should default to a readable scope
(current search/filter, selected source/tag, or local neighborhood) rather than rendering every
edge at once, with clustering or edge-type filters for dense workspaces. Vue Flow or Cytoscape.js
are suitable implementation choices; select during implementation based on component ergonomics,
layout controls, and accessibility support.

**Alternatives considered**:
- Deriving relationships only in the browser from list data: simpler initially, but misses
  persisted learner-created relationships, split/merge lineage, and scalable filtering.
- Full graph database: powerful for deep graph traversal, but unnecessary for the first-release
  1,000-node / 10,000-edge workspace target and adds operational complexity.
- Static diagram export only: useful for sharing, but does not meet the interactive Obsidian-like
  exploration requirement.

## Decision: Soft study visual palette

**Rationale**: The learner-facing frontend will use #fff8e7 as the primary background/base tone
because it gives the application a warm, low-glare study environment. #f8e7ff and #e7fff8 will
be used as supporting soft accent backgrounds for grouped content, secondary surfaces,
highlights, calm status panels, and empty states. This palette must be represented as design
tokens so dashboard, material review, card editing, review sessions, statistics, and error states
stay visually consistent.

**Accessibility guidance**:
- Text, controls, focus indicators, validation messages, and chart labels must remain readable
  on #fff8e7, #f8e7ff, and #e7fff8.
- Do not use the accent backgrounds as the only indicator of state; pair color with labels,
  icons, borders, or text.
- Use unrelated strong colors only for clear semantic states such as destructive warning,
  critical error, or success confirmation.

**Alternatives considered**:
- Default UI library palette: faster to adopt, but would weaken the requested product identity.
- Many additional pastel backgrounds: visually flexible, but would reduce consistency and make
  accessibility validation harder.

## Decision: Testing strategy

**Rationale**: The constitution requires tests as done criteria. Use layered testing:
- backend domain/service unit tests for state transitions and scheduling rules;
- repository tests with PostgreSQL and Redis containers;
- handler and contract tests against OpenAPI examples;
- AI workflow tests with mocked Claude responses and schema fixtures;
- frontend component tests for UI states;
- Playwright end-to-end tests for import, draft approval, card generation, direct review,
  plan creation, and dashboard flows.

**Alternatives considered**:
- E2E-only testing: too slow and poor at isolating domain/AI workflow errors.
- Unit-only testing: misses contract and integration issues between frontend, backend, and
  persistence.

## Decision: Observability and auditability for AI actions

**Rationale**: Learners need to understand whether items are manual, AI-generated,
AI-optimized, or rules-generated. Backend logs and domain records should include request IDs,
job IDs, prompt snapshots, model ID, schema version, source references, approval actor/action,
and error categories. Do not store secret API keys or raw hidden reasoning.

**Alternatives considered**:
- Minimal logs only: fails audit and debugging requirements for AI-generated study material.
- Store every raw interaction indefinitely: increases privacy and storage risk; retain only what
  is needed for learner audit, debugging, and reproducibility.
