# Tasks: AI Card Review Web App

**Input**: Design documents from `/specs/001-card-review-app/`

**Prerequisites**: plan.md (required), spec.md (required for user stories), research.md, data-model.md, contracts/

**Tests**: Tests are REQUIRED by the constitution for each feature, bug fix, and refactor. Include automated tests for success, failure, and relevant edge cases. If automation is not feasible, include an explicit manual validation task with rationale and owner.

**Organization**: Tasks are grouped by user story to enable independent implementation and testing of each story.

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (e.g., US1, US2, US3)
- Include exact file paths in descriptions

## Path Conventions

- **Backend**: `backend/` with Go packages under `backend/internal/`
- **Frontend**: `frontend/` with Vue 3 feature modules under `frontend/src/features/`
- **Contracts**: `specs/001-card-review-app/contracts/openapi.yaml`
- **Tests**: backend tests beside Go packages or under `backend/test/`; frontend tests under `frontend/src/**/__tests__/` and `frontend/e2e/`

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Initialize repository structure, tooling, local services, and shared validation commands.

- [X] T001 Create backend and frontend directory structure from plan in `backend/`, `frontend/`, `scripts/`, and `contracts/`
- [X] T002 Initialize Go module and backend command entrypoints in `backend/go.mod`, `backend/cmd/api/main.go`, and `backend/cmd/worker/main.go`
- [X] T003 Initialize Vue 3 TypeScript app scaffolding in `frontend/package.json`, `frontend/vite.config.ts`, and `frontend/src/app/main.ts`
- [X] T004 [P] Add backend dependencies for Gin, GORM, PostgreSQL, Redis, CloudWeGo Eino, Anthropic Go SDK, validation, and logging in `backend/go.mod`
- [X] T005 [P] Add frontend dependencies for Vue Router, Pinia, server-state queries, charts, graph visualization, testing, linting, formatting, and Playwright in `frontend/package.json`
- [X] T006 [P] Configure backend lint, format, static analysis, and test scripts in `backend/Makefile`
- [X] T007 [P] Configure frontend lint, format, unit, component, and e2e scripts in `frontend/package.json`
- [X] T008 Create local development service orchestration for PostgreSQL, Redis, backend, worker, and frontend in `scripts/dev.sh`
- [X] T009 Create full validation wrapper for backend tests, frontend tests, contract checks, and e2e checks in `scripts/test.sh`
- [X] T010 [P] Copy or link OpenAPI contract into implementation contract location in `contracts/openapi.yaml`
- [X] T011 [P] Add environment example with database, Redis, storage, Anthropic, CORS, and frontend API settings in `.env.example`
- [X] T012 [P] Add project README setup notes for running the separated app in `README.md`

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Build shared backend/frontend foundations that MUST complete before any user story implementation.

**⚠️ CRITICAL**: No user story work can begin until this phase is complete.

- [X] T013 Create backend configuration loader and validation in `backend/internal/config/config.go`
- [X] T014 Create backend application dependency container in `backend/internal/app/app.go`
- [X] T015 Create PostgreSQL connection and migration runner in `backend/internal/repository/postgres.go`
- [X] T016 Create Redis client and health check wiring in `backend/internal/jobs/redis.go`
- [X] T017 [P] Create storage abstraction for uploads and exports in `backend/internal/storage/storage.go`
- [X] T018 Create shared domain base types, IDs, timestamps, archive helpers, and workspace scoping in `backend/internal/domain/common.go`
- [X] T019 Create learner workspace and preference domain models in `backend/internal/domain/workspace.go`
- [X] T020 Create visual theme palette domain model and constants for #fff8e7, #f8e7ff, and #e7fff8 in `backend/internal/domain/theme.go`
- [X] T021 Create initial database migration for workspaces, preferences, theme palettes, tags, materials, knowledge relationships, prompts, jobs, drafts, decks, cards, review, plans, statistics, and exports in `backend/migrations/001_initial_schema.sql`
- [X] T022 Create GORM model mappings for shared entities in `backend/internal/repository/models.go`
- [X] T023 Create workspace repository with workspace scoping helpers in `backend/internal/repository/workspace_repository.go`
- [X] T024 Create authentication placeholder middleware for private learner workspace access in `backend/internal/transport/http/middleware/auth.go`
- [X] T025 Create Gin router, health endpoint, request ID, error handling, and JSON response helpers in `backend/internal/transport/http/router.go`
- [X] T026 [P] Create backend structured logging and request/job correlation helpers in `backend/internal/observability/logging.go`
- [X] T027 Create OpenAPI request/response DTO package skeleton in `backend/internal/transport/http/dto/dto.go`
- [X] T028 Create AI client adapter interface and Anthropic Go SDK adapter skeleton in `backend/internal/ai/client.go`
- [X] T029 Create Eino workflow package skeleton for classification, card generation, plan generation, and plan optimization in `backend/internal/ai/workflows.go`
- [X] T030 Create AI structured output schema definitions in `backend/internal/ai/schemas/schemas.go`
- [X] T031 Create Redis-backed job enqueue, status, progress, idempotency, and retry primitives in `backend/internal/jobs/queue.go`
- [X] T032 [P] Create frontend router, app layout shell, and protected workspace route placeholder in `frontend/src/app/router.ts`
- [X] T033 [P] Create frontend API client base with request IDs, error normalization, and auth token hook in `frontend/src/services/apiClient.ts`
- [X] T034 [P] Create frontend shared UX state components for loading, empty, error, partial success, disabled, undo, and success states in `frontend/src/components/state/StateViews.vue`
- [X] T035 [P] Create frontend design tokens and required palette variables in `frontend/src/styles/tokens.css`

**Checkpoint**: Foundation ready - user story implementation can now begin in parallel.

---

## Phase 3: User Story 1 - Import materials and extract knowledge points (Priority: P1) 🎯 MVP

**Goal**: Learners can import file/web/text material, tag it, run AI classification, review extracted knowledge points, and save approved points.

**Independent Test**: Submit one text material with tags and default prompt, receive processing feedback, review extracted points, edit/reject/approve points, and find approved points in the library without card generation or review planning.

### Tests for User Story 1 (REQUIRED) ⚠️

> **NOTE: Write these tests FIRST, ensure they FAIL before implementation**

- [X] T036 [P] [US1] Add contract tests for material create, material detail, reanalysis job, and knowledge point update endpoints in `backend/test/contract/materials_contract_test.go`
- [X] T037 [P] [US1] Add backend unit tests for material status transitions and duplicate policy handling in `backend/internal/service/material_service_test.go`
- [X] T038 [P] [US1] Add AI workflow fixture tests for material classification structured output validation in `backend/internal/ai/classification_workflow_test.go`
- [X] T039 [P] [US1] Add frontend component tests for material import form and processing states in `frontend/src/features/materials/__tests__/MaterialImport.spec.ts`
- [X] T040 [P] [US1] Add Playwright e2e test for text material import and knowledge approval in `frontend/e2e/material-import.spec.ts`

### Implementation for User Story 1

- [X] T041 [P] [US1] Create source material, material version, tag, knowledge point, and tag assignment repositories in `backend/internal/repository/material_repository.go`
- [X] T042 [P] [US1] Create material and knowledge domain logic in `backend/internal/domain/material.go` and `backend/internal/domain/knowledge_point.go`
- [X] T043 [US1] Implement material intake, duplicate detection, analysis enqueue, and reanalysis use cases in `backend/internal/service/material_service.go`
- [X] T044 [US1] Implement AI material classification workflow with default prompt and structured knowledge point drafts in `backend/internal/ai/classification_workflow.go`
- [X] T045 [US1] Implement material analysis worker handler in `backend/internal/jobs/material_analysis_worker.go`
- [X] T046 [US1] Implement material and knowledge HTTP handlers in `backend/internal/transport/http/material_handler.go`
- [X] T047 [US1] Register material, reanalysis, job, and knowledge update routes in `backend/internal/transport/http/router.go`
- [X] T048 [P] [US1] Implement material API service functions in `frontend/src/services/materials.ts`
- [X] T049 [P] [US1] Implement material import page and form in `frontend/src/features/materials/MaterialImportPage.vue`
- [X] T050 [P] [US1] Implement material processing status and retry UI in `frontend/src/features/materials/MaterialProcessingPanel.vue`
- [X] T051 [P] [US1] Implement knowledge point review, edit, approve, and reject UI in `frontend/src/features/knowledge/KnowledgeReviewPanel.vue`
- [X] T052 [US1] Wire material import and knowledge review routes in `frontend/src/app/router.ts`
- [X] T053 [US1] Add learner-readable error and partial-success handling for material analysis in `frontend/src/features/materials/materialStates.ts`
- [X] T054 [US1] Validate US1 performance feedback within 2 seconds in `frontend/e2e/material-import.spec.ts`
- [X] T055 [P] [US1] Add frontend component tests for polished material import layout, source-type selector, prompt card, tag chips, responsive states, and required palette usage in `frontend/src/features/materials/__tests__/MaterialImportLayout.spec.ts`
- [X] T056 [US1] Redesign material import page with structured hero, source-type cards for text/file/web, prompt preset panel, tag input chips, progress preview, and responsive two-column layout in `frontend/src/features/materials/MaterialImportPage.vue`
- [X] T057 [US1] Add material import page-specific responsive spacing, form controls, card surfaces, and #fff8e7/#f8e7ff/#e7fff8 palette refinements in `frontend/src/styles/tokens.css`
- [X] T058 [US1] Validate material import page visual polish, keyboard navigation, empty/loading/error states, and mobile/tablet/desktop layout in `frontend/e2e/material-import.spec.ts`

**Checkpoint**: Material intake and knowledge point extraction work independently.

---

## Phase 4: User Story 2 - Curate knowledge before card generation (Priority: P2)

**Goal**: Learners can search, filter, browse in list view, inspect an Obsidian-like relationship graph, merge, split, and mark knowledge points before they become cards.

**Independent Test**: Starting from extracted knowledge points, use list view to find duplicates, switch to graph view to inspect source/tag/card/semantic relationships, merge two points, split one overloaded point, mark one as needing review, and confirm only approved points are used for card generation.

### Tests for User Story 2 (REQUIRED) ⚠️

- [X] T059 [P] [US2] Add contract tests for knowledge point list, split, merge, update, graph, and relationship endpoints in `backend/test/contract/knowledge_contract_test.go`
- [X] T060 [P] [US2] Add backend unit tests for knowledge split/merge/source preservation in `backend/internal/service/knowledge_service_test.go`
- [X] T061 [P] [US2] Add frontend component tests for knowledge filters, duplicate warnings, list/graph view toggle, split, and merge dialogs in `frontend/src/features/knowledge/__tests__/KnowledgeLibrary.spec.ts`
- [X] T062 [P] [US2] Add Playwright e2e test for knowledge curation, list view, graph exploration, and node opening before card generation in `frontend/e2e/knowledge-curation.spec.ts`
- [X] T063 [P] [US2] Add backend unit tests for knowledge graph filtering, local-neighborhood depth, lineage, and archived/rejected visibility in `backend/internal/service/knowledge_graph_service_test.go`
- [X] T064 [P] [US2] Add frontend component tests for graph rendering, edge labels, node selection, relationship filters, and dense graph fallback in `frontend/src/features/knowledge/__tests__/KnowledgeGraph.spec.ts`

### Implementation for User Story 2

- [X] T065 [P] [US2] Extend knowledge repository with search, filtering, duplicate groups, split, merge, and relationship persistence in `backend/internal/repository/knowledge_repository.go`
- [X] T066 [P] [US2] Create dedicated knowledge graph query repository for typed edges, derived shared-source/shared-tag links, and focused-neighborhood queries in `backend/internal/repository/knowledge_graph_repository.go`
- [X] T067 [US2] Implement knowledge curation use cases for search, filters, split, merge, restore, graph relationship maintenance, and approval status in `backend/internal/service/knowledge_service.go`
- [X] T068 [US2] Implement Obsidian-like graph use cases for searchable nodes, typed edges, filters, clustering warnings, and local-neighborhood depth in `backend/internal/service/knowledge_graph_service.go`
- [X] T069 [US2] Implement knowledge curation, graph, and relationship HTTP handlers in `backend/internal/transport/http/knowledge_handler.go`
- [X] T070 [P] [US2] Implement knowledge list, curation, graph, and relationship API service functions in `frontend/src/services/knowledge.ts`
- [X] T071 [P] [US2] Implement knowledge library page with list view, graph view toggle, filters, and duplicate indicators in `frontend/src/features/knowledge/KnowledgeLibraryPage.vue`
- [X] T072 [P] [US2] Implement split and merge dialogs in `frontend/src/features/knowledge/KnowledgeCurationDialogs.vue`
- [X] T073 [P] [US2] Implement Obsidian-like graph component with node selection, edge labels, filters, and accessible fallback list in `frontend/src/features/knowledge/KnowledgeGraphView.vue`
- [X] T074 [P] [US2] Implement graph relationship editor for learner-created related/prerequisite/supports links in `frontend/src/features/knowledge/KnowledgeRelationshipEditor.vue`
- [X] T075 [US2] Exclude unapproved knowledge points by default in card generation selection logic in `backend/internal/service/knowledge_selection.go`
- [X] T076 [US2] Validate knowledge graph switching and filtering performance within 1 second at 1,000 knowledge points and 10,000 edges in `frontend/e2e/knowledge-curation.spec.ts`
- [X] T077 [US2] Validate knowledge library list/graph empty, loading, error, undo, dense graph, keyboard navigation, and palette states in `frontend/src/features/knowledge/__tests__/KnowledgeLibrary.spec.ts`

**Checkpoint**: Knowledge curation list and graph views work independently and produce approved input for cards.

---

## Phase 5: User Story 3 - Generate decks and cards from knowledge points (Priority: P3)

**Goal**: Learners can generate deck/card drafts from approved knowledge points using prompts and filters, then review, edit, discard, or approve the drafts.

**Independent Test**: Select tags, run card generation, review a proposed deck with linked cards, edit/delete weak cards, approve the deck, and see cards in card view.

### Tests for User Story 3 (REQUIRED) ⚠️

- [ ] T078 [P] [US3] Add contract tests for deck generation job, AI draft update, deck list, and card list endpoints in `backend/test/contract/deck_generation_contract_test.go`
- [ ] T079 [P] [US3] Add AI workflow fixture tests for card/deck generation schemas and duplicate card warnings in `backend/internal/ai/card_generation_workflow_test.go`
- [ ] T080 [P] [US3] Add backend unit tests for AI draft approval and card source traceability in `backend/internal/service/card_generation_service_test.go`
- [ ] T081 [P] [US3] Add frontend component tests for deck generation form and draft review in `frontend/src/features/cards/__tests__/DeckGeneration.spec.ts`
- [ ] T082 [P] [US3] Add Playwright e2e test for generating and approving a deck in `frontend/e2e/deck-generation.spec.ts`

### Implementation for User Story 3

- [ ] T083 [P] [US3] Create deck, card, card source link, and AI draft repositories in `backend/internal/repository/card_repository.go`
- [ ] T084 [P] [US3] Create deck, card, card source link, and AI draft domain logic in `backend/internal/domain/card.go`
- [ ] T085 [US3] Implement card generation use cases and draft approval/discard/edit behavior in `backend/internal/service/card_generation_service.go`
- [ ] T086 [US3] Implement AI card generation workflow with structured deck/card drafts and source links in `backend/internal/ai/card_generation_workflow.go`
- [ ] T087 [US3] Implement card generation worker handler in `backend/internal/jobs/card_generation_worker.go`
- [ ] T088 [US3] Implement deck generation, AI draft, deck list, and card list HTTP handlers in `backend/internal/transport/http/card_handler.go`
- [ ] T089 [P] [US3] Implement cards and decks API service functions in `frontend/src/services/cards.ts`
- [ ] T090 [P] [US3] Implement deck generation page with prompt and source filters in `frontend/src/features/cards/DeckGenerationPage.vue`
- [ ] T091 [P] [US3] Implement generated deck/card draft review UI in `frontend/src/features/cards/CardDraftReviewPanel.vue`
- [ ] T092 [P] [US3] Implement deck view and card view shell with shared filters in `frontend/src/features/cards/CardDeckViews.vue`
- [ ] T093 [US3] Validate generated draft traceability and performance feedback in `frontend/e2e/deck-generation.spec.ts`

**Checkpoint**: Approved generated decks and cards are usable independently.

---

## Phase 6: User Story 4 - Manage cards and decks flexibly (Priority: P4)

**Goal**: Learners can browse, create, edit, merge, restore, archive, and organize cards/decks while preserving history and edits.

**Independent Test**: Create a manual card, switch views, merge two decks, edit a card in the merged deck, restore source deck grouping, and archive a deck.

### Tests for User Story 4 (REQUIRED) ⚠️

- [ ] T094 [P] [US4] Add contract tests for deck create/update/merge/restore and card create/update endpoints in `backend/test/contract/card_deck_management_contract_test.go`
- [ ] T095 [P] [US4] Add backend unit tests for deck composition restore and surviving card edits in `backend/internal/service/deck_service_test.go`
- [ ] T096 [P] [US4] Add frontend component tests for deck/card views, merge dialog, restore flow, and archive state in `frontend/src/features/cards/__tests__/CardDeckManagement.spec.ts`
- [ ] T097 [P] [US4] Add Playwright e2e test for manual card, deck merge, restore, and archive in `frontend/e2e/card-deck-management.spec.ts`

### Implementation for User Story 4

- [ ] T098 [P] [US4] Create deck composition repository and persistence helpers in `backend/internal/repository/deck_composition_repository.go`
- [ ] T099 [US4] Implement card create, edit, duplicate, archive, restore, delete, and search use cases in `backend/internal/service/card_service.go`
- [ ] T100 [US4] Implement deck create, update, merge, restore, archive, and composition use cases in `backend/internal/service/deck_service.go`
- [ ] T101 [US4] Extend card/deck HTTP handlers with manual management, merge, restore, and archive actions in `backend/internal/transport/http/card_handler.go`
- [ ] T102 [P] [US4] Implement card editor component in `frontend/src/features/cards/CardEditor.vue`
- [ ] T103 [P] [US4] Implement deck management list, archive, and restore UI in `frontend/src/features/cards/DeckManagementPage.vue`
- [ ] T104 [P] [US4] Implement deck merge and restore dialogs in `frontend/src/features/cards/DeckMergeRestoreDialogs.vue`
- [ ] T105 [US4] Add destructive action confirmation and undo handling for cards/decks in `frontend/src/features/cards/cardDeckActions.ts`
- [ ] T106 [US4] Validate deck/card filtering performance at target scale in `frontend/e2e/card-deck-management.spec.ts`

**Checkpoint**: Card/deck management works independently without requiring review plans.

---

## Phase 7: User Story 5 - Review with direct sessions and adaptive plans (Priority: P5)

**Goal**: Learners can run direct review sessions, create manual or AI-assisted review plans, track statistics, optimize plans, and inspect revision history.

**Independent Test**: Start direct review, answer cards, pause/resume, create a 14-day plan, complete sessions, inspect statistics, optimize a plan, compare revisions, and restore a compatible prior revision.

### Tests for User Story 5 (REQUIRED) ⚠️

- [ ] T107 [P] [US5] Add contract tests for review session start/answer/update endpoints in `backend/test/contract/review_session_contract_test.go`
- [ ] T108 [P] [US5] Add contract tests for plan create/generate/update/optimize/revisions/statistics endpoints in `backend/test/contract/review_plan_contract_test.go`
- [ ] T109 [P] [US5] Add backend unit tests for review result scheduling and forgetting-curve due logic in `backend/internal/service/review_scheduler_test.go`
- [ ] T110 [P] [US5] Add backend unit tests for plan conflict prevention, missed-day recovery, revision compare, and restore in `backend/internal/service/review_plan_service_test.go`
- [ ] T111 [P] [US5] Add AI workflow fixture tests for plan generation and optimization structured outputs in `backend/internal/ai/plan_workflow_test.go`
- [ ] T112 [P] [US5] Add frontend component tests for review session, plan editor, statistics, and revision history in `frontend/src/features/review/__tests__/ReviewPlan.spec.ts`
- [ ] T113 [P] [US5] Add Playwright e2e test for direct review pause/resume and plan optimization in `frontend/e2e/review-plan.spec.ts`

### Implementation for User Story 5

- [ ] T114 [P] [US5] Create review session, review result, review plan, plan revision, and statistics repositories in `backend/internal/repository/review_repository.go`
- [ ] T115 [P] [US5] Create review and plan domain logic in `backend/internal/domain/review.go` and `backend/internal/domain/review_plan.go`
- [ ] T116 [US5] Implement review scheduler with due-card prioritization and next due calculation in `backend/internal/service/review_scheduler.go`
- [ ] T117 [US5] Implement direct review session start, answer record, pause, resume, complete, and abandon use cases in `backend/internal/service/review_session_service.go`
- [ ] T118 [US5] Implement review plan create, update, conflict detection, missed-day recovery, revision history, compare, and restore use cases in `backend/internal/service/review_plan_service.go`
- [ ] T119 [US5] Implement review statistics aggregation for workspace, tag, deck, plan, and card scopes in `backend/internal/service/statistics_service.go`
- [ ] T120 [US5] Implement AI plan generation and optimization workflows in `backend/internal/ai/plan_workflow.go`
- [ ] T121 [US5] Implement plan generation and optimization worker handlers in `backend/internal/jobs/plan_worker.go`
- [ ] T122 [US5] Implement review, plan, revision, and statistics HTTP handlers in `backend/internal/transport/http/review_handler.go`
- [ ] T123 [P] [US5] Implement review API service functions in `frontend/src/services/review.ts`
- [ ] T124 [P] [US5] Implement direct review session UI with pause/resume protection in `frontend/src/features/review/ReviewSessionPage.vue`
- [ ] T125 [P] [US5] Implement review plan editor and AI-assisted plan draft review UI in `frontend/src/features/plans/ReviewPlanEditor.vue`
- [ ] T126 [P] [US5] Implement statistics view with weak areas, overdue work, and plan suggestions in `frontend/src/features/plans/ReviewStatisticsPage.vue`
- [ ] T127 [P] [US5] Implement plan revision history, comparison, and restore UI in `frontend/src/features/plans/PlanRevisionHistory.vue`
- [ ] T128 [US5] Validate review interruption recovery and answer preservation in `frontend/e2e/review-plan.spec.ts`

**Checkpoint**: Review sessions and plans work independently with statistics and revisions.

---

## Phase 8: User Story 6 - Manage prompts, preferences, and AI draft safety (Priority: P6)

**Goal**: Learners can manage prompt presets and understand AI draft provenance before trusting AI-produced content.

**Independent Test**: Use a default prompt, save a customized preset, apply it to generation, review source/prompt context for a draft, and approve only usable output.

### Tests for User Story 6 (REQUIRED) ⚠️

- [ ] T129 [P] [US6] Add contract tests for prompt preset list/create and AI draft update endpoints in `backend/test/contract/prompts_and_drafts_contract_test.go`
- [ ] T130 [P] [US6] Add backend unit tests for prompt preset defaults, prompt snapshots, and AI draft audit metadata in `backend/internal/service/prompt_service_test.go`
- [ ] T131 [P] [US6] Add frontend component tests for prompt preset management and AI draft provenance display in `frontend/src/features/prompts/__tests__/PromptPreset.spec.ts`
- [ ] T132 [P] [US6] Add Playwright e2e test for saving a prompt preset and approving an AI draft with source context in `frontend/e2e/prompts-and-ai-drafts.spec.ts`

### Implementation for User Story 6

- [ ] T133 [P] [US6] Create prompt preset and prompt snapshot repositories in `backend/internal/repository/prompt_repository.go`
- [ ] T134 [P] [US6] Create prompt preset and prompt snapshot domain logic in `backend/internal/domain/prompt.go`
- [ ] T135 [US6] Implement prompt preset create, list, rename, archive, default selection, and snapshot use cases in `backend/internal/service/prompt_service.go`
- [ ] T136 [US6] Implement AI draft provenance and draft status use cases in `backend/internal/service/ai_draft_service.go`
- [ ] T137 [US6] Implement prompt preset and AI draft HTTP handlers in `backend/internal/transport/http/prompt_handler.go`
- [ ] T138 [P] [US6] Implement prompt and draft API service functions in `frontend/src/services/prompts.ts`
- [ ] T139 [P] [US6] Implement prompt preset management page in `frontend/src/features/prompts/PromptPresetPage.vue`
- [ ] T140 [P] [US6] Implement AI draft provenance panel showing source, prompt snapshot, model ID, and approval state in `frontend/src/features/prompts/AIDraftProvenancePanel.vue`
- [ ] T141 [US6] Add retry, edit prompt, use partial result, and discard handling for AI drafts in `frontend/src/features/prompts/aiDraftActions.ts`

**Checkpoint**: Prompt presets and AI draft safety are independently usable.

---

## Phase 9: User Story 7 - Understand progress and focus next actions (Priority: P7)

**Goal**: Learners can open a visually consistent dashboard to see drafts, due reviews, overdue reviews, active plans, weak areas, and next actions using the required palette.

**Independent Test**: With imported materials, generated decks, active plans, pending drafts, due reviews, and weak tags, open the dashboard and identify next recommended actions within 30 seconds.

### Tests for User Story 7 (REQUIRED) ⚠️

- [ ] T142 [P] [US7] Add contract tests for dashboard and statistics endpoints in `backend/test/contract/dashboard_contract_test.go`
- [ ] T143 [P] [US7] Add backend unit tests for dashboard next-action ranking and weak-area summarization in `backend/internal/service/dashboard_service_test.go`
- [ ] T144 [P] [US7] Add frontend component tests for dashboard panels, empty state, next actions, and required palette usage in `frontend/src/features/dashboard/__tests__/Dashboard.spec.ts`
- [ ] T145 [P] [US7] Add frontend accessibility tests for #fff8e7, #f8e7ff, and #e7fff8 readability and focus states in `frontend/src/styles/__tests__/paletteAccessibility.spec.ts`
- [ ] T146 [P] [US7] Add Playwright e2e test for dashboard next action discovery and visual palette consistency in `frontend/e2e/dashboard-palette.spec.ts`

### Implementation for User Story 7

- [ ] T147 [P] [US7] Implement dashboard summary service for drafts, due reviews, overdue reviews, active plans, weak areas, and next actions in `backend/internal/service/dashboard_service.go`
- [ ] T148 [US7] Implement dashboard HTTP handler in `backend/internal/transport/http/dashboard_handler.go`
- [ ] T149 [P] [US7] Implement dashboard API service functions in `frontend/src/services/dashboard.ts`
- [ ] T150 [P] [US7] Implement dashboard page with next action groups in `frontend/src/features/dashboard/DashboardPage.vue`
- [ ] T151 [P] [US7] Implement empty workspace guidance panel in `frontend/src/features/dashboard/EmptyWorkspacePanel.vue`
- [ ] T152 [P] [US7] Implement weak-area and plan-status panels in `frontend/src/features/dashboard/ProgressPanels.vue`
- [ ] T153 [US7] Implement required palette tokens, utility classes, semantic state colors, and focus styles in `frontend/src/styles/tokens.css`
- [ ] T154 [US7] Apply required #fff8e7 base and #f8e7ff/#e7fff8 accent backgrounds in `frontend/src/app/AppShell.vue`
- [ ] T155 [US7] Validate dashboard next-action timing and palette consistency in `frontend/e2e/dashboard-palette.spec.ts`

**Checkpoint**: Dashboard and required visual identity are independently testable.

---

## Phase 10: Polish & Cross-Cutting Concerns

**Purpose**: Improvements that affect multiple user stories and prepare the feature for delivery.

- [ ] T156 [P] Add portable workspace export worker and service integration in `backend/internal/jobs/export_worker.go`
- [ ] T157 [P] Add workspace import preview and confirm worker integration in `backend/internal/jobs/import_worker.go`
- [ ] T158 Add export/import HTTP handlers and route wiring in `backend/internal/transport/http/portability_handler.go`
- [ ] T159 [P] Add frontend export/import page in `frontend/src/features/settings/PortabilityPage.vue`
- [ ] T160 Add e2e validation for export/import preserving cards, decks, tags, plans, and review history in `frontend/e2e/export-import.spec.ts`
- [ ] T161 Validate backend OpenAPI contract coverage against `specs/001-card-review-app/contracts/openapi.yaml` in `backend/test/contract/openapi_coverage_test.go`
- [ ] T162 Validate frontend generated or handwritten API clients against OpenAPI schemas in `frontend/src/services/__tests__/openapiClient.spec.ts`
- [ ] T163 Add backend performance test fixtures for material jobs, filtering, and review answer recording in `backend/test/performance/performance_test.go`
- [ ] T164 Add frontend performance checks for 500 decks and 5,000 cards in `frontend/e2e/performance.spec.ts`
- [ ] T165 Add cross-feature UX consistency review tests for loading, empty, error, partial-success, disabled, undo, success, and palette states in `frontend/src/components/state/__tests__/StateViews.spec.ts`
- [ ] T166 Run backend lint, static analysis, migrations test, and full `go test ./...` validation via `scripts/test.sh`
- [ ] T167 Run frontend lint, typecheck, unit, component, accessibility, and Playwright validation via `scripts/test.sh`
- [ ] T168 Update implementation notes and quickstart references after final validation in `README.md`

---

## Dependencies & Execution Order

### Phase Dependencies

- **Setup (Phase 1)**: No dependencies - can start immediately.
- **Foundational (Phase 2)**: Depends on Setup completion - BLOCKS all user stories.
- **User Stories (Phase 3+)**: All depend on Foundational phase completion.
  - US1 is the MVP and enables real knowledge data for later stories.
  - US2 depends on US1 knowledge point data.
  - US3 depends on US2 approved knowledge point selection.
  - US4 depends on US3 card/deck entities but can use manual decks/cards for partial work.
  - US5 depends on active cards/decks from US3/US4.
  - US6 depends on shared AI job/draft/prompt foundations and can proceed after US1 draft flow exists.
  - US7 depends on statistics, drafts, review state, and visual palette foundations from earlier stories.
- **Polish (Final Phase)**: Depends on all desired user stories being complete.

### User Story Dependencies

- **User Story 1 (P1)**: Starts after Foundational. No dependency on other stories. MVP scope.
- **User Story 2 (P2)**: Starts after US1 data model/services exist. Independently testable through seeded knowledge points and relationships.
- **User Story 3 (P3)**: Starts after US2 approved knowledge point selection exists. Independently testable through seeded approved points.
- **User Story 4 (P4)**: Starts after card/deck base from US3 exists. Independently testable with manually created cards/decks.
- **User Story 5 (P5)**: Starts after active cards/decks exist. Independently testable with seeded active cards/decks.
- **User Story 6 (P6)**: Starts after AI job/draft/prompt foundations exist. Integrates with US1/US3/US5 AI outputs.
- **User Story 7 (P7)**: Starts after dashboard input sources exist. Visual palette tasks can begin after foundational style tokens.

### Within Each User Story

- Tests MUST be written and fail before implementation.
- Domain/repository tasks before services.
- Services before HTTP handlers.
- HTTP handlers before frontend API integration.
- Frontend API integration before pages/components.
- UX consistency, palette, accessibility, and performance validation before story completion.
- Story complete before moving to the next priority for single-developer MVP delivery.

### Parallel Opportunities

- Setup tasks T004-T007 and T010-T012 can run in parallel after T001-T003.
- Foundational tasks T017, T026, T032-T035 can run in parallel with other foundation work after core structure exists.
- Test tasks within each user story are parallelizable before implementation starts.
- Backend repository/domain tasks and frontend component shell tasks in each story can run in parallel when they touch different files.
- US6 prompt management can proceed in parallel with later review/dashboard work after shared AI draft foundations exist.
- Polish tasks T149-T150 and T152 can run in parallel before integration tasks T151/T153.

---

## Parallel Example: User Story 1

```bash
# Launch all tests for User Story 1 together:
Task: "T036 [US1] Add contract tests for material create, material detail, reanalysis job, and knowledge point update endpoints in backend/test/contract/materials_contract_test.go"
Task: "T037 [US1] Add backend unit tests for material status transitions and duplicate policy handling in backend/internal/service/material_service_test.go"
Task: "T038 [US1] Add AI workflow fixture tests for material classification structured output validation in backend/internal/ai/classification_workflow_test.go"
Task: "T039 [US1] Add frontend component tests for material import form and processing states in frontend/src/features/materials/__tests__/MaterialImport.spec.ts"
Task: "T040 [US1] Add Playwright e2e test for text material import and knowledge approval in frontend/e2e/material-import.spec.ts"

# Launch independent implementation tasks after tests are in place:
Task: "T041 [US1] Create source material, material version, tag, knowledge point, and tag assignment repositories in backend/internal/repository/material_repository.go"
Task: "T042 [US1] Create material and knowledge domain logic in backend/internal/domain/material.go and backend/internal/domain/knowledge_point.go"
Task: "T048 [US1] Implement material API service functions in frontend/src/services/materials.ts"
Task: "T049 [US1] Implement material import page and form in frontend/src/features/materials/MaterialImportPage.vue"
```

## Parallel Example: User Story 5

```bash
# Launch review-plan test tasks together:
Task: "T100 [US5] Add contract tests for review session start/answer/update endpoints in backend/test/contract/review_session_contract_test.go"
Task: "T101 [US5] Add contract tests for plan create/generate/update/optimize/revisions/statistics endpoints in backend/test/contract/review_plan_contract_test.go"
Task: "T102 [US5] Add backend unit tests for review result scheduling and forgetting-curve due logic in backend/internal/service/review_scheduler_test.go"
Task: "T104 [US5] Add AI workflow fixture tests for plan generation and optimization structured outputs in backend/internal/ai/plan_workflow_test.go"
Task: "T106 [US5] Add Playwright e2e test for direct review pause/resume and plan optimization in frontend/e2e/review-plan.spec.ts"

# Launch independent UI implementation tasks together after API shapes are stable:
Task: "T117 [US5] Implement direct review session UI with pause/resume protection in frontend/src/features/review/ReviewSessionPage.vue"
Task: "T118 [US5] Implement review plan editor and AI-assisted plan draft review UI in frontend/src/features/plans/ReviewPlanEditor.vue"
Task: "T119 [US5] Implement statistics view with weak areas, overdue work, and plan suggestions in frontend/src/features/plans/ReviewStatisticsPage.vue"
Task: "T120 [US5] Implement plan revision history, comparison, and restore UI in frontend/src/features/plans/PlanRevisionHistory.vue"
```

## Parallel Example: User Story 7

```bash
# Launch dashboard and palette validation tests together:
Task: "T135 [US7] Add contract tests for dashboard and statistics endpoints in backend/test/contract/dashboard_contract_test.go"
Task: "T137 [US7] Add frontend component tests for dashboard panels, empty state, next actions, and required palette usage in frontend/src/features/dashboard/__tests__/Dashboard.spec.ts"
Task: "T138 [US7] Add frontend accessibility tests for #fff8e7, #f8e7ff, and #e7fff8 readability and focus states in frontend/src/styles/__tests__/paletteAccessibility.spec.ts"
Task: "T139 [US7] Add Playwright e2e test for dashboard next action discovery and visual palette consistency in frontend/e2e/dashboard-palette.spec.ts"

# Launch independent palette/dashboard implementation tasks together:
Task: "T143 [US7] Implement dashboard page with next action groups in frontend/src/features/dashboard/DashboardPage.vue"
Task: "T144 [US7] Implement empty workspace guidance panel in frontend/src/features/dashboard/EmptyWorkspacePanel.vue"
Task: "T145 [US7] Implement weak-area and plan-status panels in frontend/src/features/dashboard/ProgressPanels.vue"
Task: "T146 [US7] Implement required palette tokens, utility classes, semantic state colors, and focus styles in frontend/src/styles/tokens.css"
```

---

## Implementation Strategy

### MVP First (User Story 1 Only)

1. Complete Phase 1: Setup.
2. Complete Phase 2: Foundational infrastructure.
3. Complete Phase 3: User Story 1 material import and knowledge extraction.
4. **STOP and VALIDATE**: Run US1 backend, AI fixture, frontend component, and Playwright tests.
5. Demo material import, AI processing feedback, knowledge edit/reject/approve, and search.

### Incremental Delivery

1. Setup + Foundational → app shell, backend API, DB, Redis, AI/job infrastructure ready.
2. US1 → material intake and approved knowledge points.
3. US2 → curated knowledge library and duplicate/split/merge quality controls.
4. US3 → AI-generated deck/card drafts and approval.
5. US4 → manual card/deck maintenance, merge, restore, archive.
6. US5 → review sessions, review plans, statistics, and plan optimization.
7. US6 → prompt presets and AI draft provenance hardening.
8. US7 → dashboard and synchronized visual identity.
9. Polish → portability, performance, contract coverage, full validation.

### Parallel Team Strategy

With multiple developers:

1. Team completes Setup + Foundational together.
2. Backend A: material/knowledge/card/review domain and repositories.
3. Backend B: AI workflows, Redis jobs, handlers, contracts.
4. Frontend A: shared app shell, palette tokens, state components, dashboard.
5. Frontend B: materials, knowledge, cards, review, plans pages.
6. QA/Automation: contract tests, fixtures, e2e tests, accessibility/performance checks.

---

## Notes

- [P] tasks = different files, no dependencies on incomplete tasks.
- [Story] label maps task to a specific user story for traceability.
- Each user story is independently completable and testable with seeded data where needed.
- Verify tests fail before implementing each story.
- Keep AI outputs as drafts until learner approval.
- Preserve source and prompt traceability for AI-generated trusted items.
- Validate #fff8e7, #f8e7ff, and #e7fff8 palette readability before completing UI stories, including dense knowledge graph views.
- Avoid vague tasks, same-file conflicts, and cross-story dependencies that break independence.
