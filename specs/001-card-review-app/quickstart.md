# Quickstart: AI Card Review Web App Validation

## Purpose

This guide describes how to validate the planned feature end-to-end once implementation tasks
are generated and completed. It intentionally avoids service/controller/model implementation
code; use it as a run and validation checklist.

## Prerequisites

- Go 1.24+
- Node.js 22+
- PostgreSQL 16+
- Redis 7+
- An Anthropic API credential available to the backend as an environment variable
- Local object/file storage directory for uploaded materials and exports
- Browser for frontend validation

## Expected Project Commands

The implementation should provide equivalent commands even if the exact wrapper scripts differ:

```bash
# Backend
cd backend
go test ./...
go run ./cmd/api
go run ./cmd/worker

# Frontend
cd frontend
npm install
npm run lint
npm run test
npm run dev
npm run e2e

# Full validation wrapper, if provided
./scripts/test.sh
```

## Environment Setup

Create local configuration for:

- PostgreSQL connection string
- Redis connection string
- upload/export storage location
- Anthropic API credential
- allowed frontend origin
- backend API base URL for frontend

The backend must fail fast with a clear configuration error when required settings are missing.
Do not hardcode secrets in repository files.

## Contract Validation

1. Open `contracts/openapi.yaml`.
2. Validate that backend routes match the OpenAPI operation IDs and schemas.
3. Validate that frontend API client calls use the same request and response shapes.
4. Run contract tests that cover:
   - material creation and job acceptance;
   - knowledge point update, split, and merge;
   - AI deck generation job acceptance;
   - deck merge and restore;
   - card creation and update;
   - review session start, answer recording, pause/resume/complete;
   - review plan create/generate/optimize/revision restore;
   - dashboard and statistics retrieval;
   - export/import job acceptance.

## Scenario 1: Import Material and Curate Knowledge

**Goal**: Prove P1 and P2 flows work independently.

1. Start backend API, worker, PostgreSQL, Redis, and frontend.
2. Open the app as a learner with an empty private workspace.
3. Add a text source under two tags using the default classification prompt.
4. Confirm the UI immediately shows queued/processing feedback.
5. Wait for the analysis job to complete or use a mocked AI fixture in test mode.
6. Review extracted knowledge points.
7. Edit one knowledge point, reject one, and approve the rest.
8. Split one overloaded knowledge point and merge two duplicates.

**Expected outcome**:

- Material exists with source reference and tags.
- Approved knowledge points are searchable in the knowledge library.
- Rejected or needs-review points are excluded from default card generation.
- Duplicate/split/merge operations preserve source references and tags.
- Loading, partial success, empty, and failure states are visible where applicable.

## Scenario 2: Generate and Approve Cards

**Goal**: Prove AI card generation produces reviewable drafts and preserves traceability.

1. Select one or more tags with approved knowledge points.
2. Enter or select a card generation prompt preset.
3. Start deck/card generation.
4. Confirm a job ID and progress state are visible within 2 seconds.
5. Review generated deck/card drafts.
6. Edit one card, discard one weak card, approve the deck.
7. Open card view and filter by tag, source, deck, and review status.

**Expected outcome**:

- Generated cards remain drafts until approved.
- Approved cards appear in deck and card views.
- Each AI-generated card links back to source knowledge/material and prompt snapshot.
- Duplicate or unclear cards are flagged for learner review.

## Scenario 3: Manage Decks and Cards

**Goal**: Prove card/deck management supports organization and recovery.

1. Create a manual card in an existing deck.
2. Create a second deck with at least one card.
3. Merge the two decks into a combined deck.
4. Edit one card inside the merged deck.
5. Restore the merged deck into source deck grouping.
6. Archive one deck and confirm it is hidden from active review selection but searchable.

**Expected outcome**:

- Deck merge records source composition.
- Restore recreates source grouping and preserves surviving card edits.
- Archived decks do not delete cards or review history.
- Destructive or high-impact actions require confirmation.

## Scenario 4: Direct Review Session

**Goal**: Prove review sessions can be started, interrupted, resumed, and completed.

1. Start direct review from one active deck.
2. Answer at least three cards with different results or confidence levels.
3. Pause the session.
4. Refresh or navigate away and return.
5. Resume the session and complete it.

**Expected outcome**:

- Answered-card results are preserved after pause/refresh.
- Each answer records result, confidence/difficulty, response time, and next due timing.
- Completion shows a session summary.
- Review statistics update after completion.

## Scenario 5: Review Plan and Optimization

**Goal**: Prove manual and AI-assisted planning with revision history.

1. Create a 14-day review plan for selected decks with daily capacity.
2. Confirm the plan prevents duplicate active scheduling unless overlap is explicitly allowed.
3. Complete review sessions for the plan.
4. Open statistics and inspect progress, recall rate, overdue cards, and weak areas.
5. Run AI-assisted or rules-based optimization.
6. Compare the new plan revision with the previous revision.
7. Restore a compatible prior revision.

**Expected outcome**:

- Active plans show daily workload and due-card organization.
- Optimization respects daily capacity unless the learner overrides it.
- Every plan change records source, timestamp, summary, reason where available, and snapshot.
- Restore does not erase completed review history.

## Scenario 6: Dashboard Next Actions

**Goal**: Prove the learner can quickly understand what to do next.

1. Prepare a workspace with pending AI drafts, due cards, overdue cards, at least one active
   plan, and weak tags/decks from statistics.
2. Open the dashboard.
3. Verify next action groups are visible.
4. Select actions for approving drafts, continuing review, and fixing overdue plan items.

**Expected outcome**:

- Dashboard gives a clear next recommended action within 30 seconds.
- Empty workspace shows import/manual card creation guidance.
- Weak areas link to focused review or card improvement flows.

## Scenario 7: Visual Palette and UX Consistency

**Goal**: Prove the required frontend palette and shared UI states are consistent and readable.

1. Open dashboard, material review, knowledge library, card view, deck view, direct review,
   review plan detail, statistics, prompt preset management, and empty workspace screens.
2. Confirm #fff8e7 is used as the primary background/base tone across learner-facing pages.
3. Confirm #f8e7ff and #e7fff8 appear as supporting accent backgrounds for grouped content,
   secondary surfaces, highlights, status panels, or empty states.
4. Validate primary text, controls, focus indicators, validation messages, error states, and
   chart labels on each of the three palette backgrounds.
5. Confirm warning, error, and success states use labels, icons, borders, or text in addition to
   color and do not rely only on background color.
6. Run the relevant frontend component and accessibility checks for shared layout, cards,
   dashboard panels, forms, review feedback, and chart components.

**Expected outcome**:

- Learner-facing screens share the requested soft study visual identity.
- The palette does not reduce readability in dense lists, charts, forms, or long review sessions.
- Keyboard focus and validation feedback remain visible on #fff8e7, #f8e7ff, and #e7fff8.
- Unrelated dominant background colors are not used in core workflows except for clear semantic
  states such as destructive warning, critical error, or success confirmation.

## Scenario 8: Export and Import

**Goal**: Prove learner portability and backup.

1. Queue an export including materials, knowledge points, prompts, decks, cards, review plans,
   review history, and statistics.
2. Confirm export progress and completion.
3. Import the export into an empty test workspace in preview mode.
4. Review conflicts and confirm import.

**Expected outcome**:

- Export includes schema version and selected study content.
- Import preview identifies conflicts before changes are applied.
- Confirmed import preserves cards, decks, tags, review plans, and review history in validation
  cases.

## AI Workflow Validation

For each AI workflow, validate both a real-provider path and a deterministic mocked path:

- Material classification returns structured knowledge point drafts matching schema.
- Card generation returns structured deck/card drafts matching schema.
- Plan generation returns a review-plan draft matching schema.
- Plan optimization returns a revision proposal with change summary and reason.
- Invalid or empty AI output becomes a failed or warning draft state, not trusted content.
- Learner approval is required before AI drafts become trusted records.

## Performance Validation

Validate against the success criteria and constitution budgets:

- Long-running operations show visible feedback within 2 seconds.
- Typical material under 5,000 words returns reviewable result or clear failure within 2 minutes.
- Deck/card view switching and filtering responds within 1 second at target workspace scale.
- Review answer recording completes within 500 ms from learner action.
- Pause/resume preserves answered-card results in interruption scenarios.

## Accessibility and UX Validation

- All major workflows are navigable by keyboard.
- Loading, empty, processing, failure, partial success, disabled, undo, and completion states are
  consistent across features.
- #fff8e7 is used as the primary background/base tone, with #f8e7ff and #e7fff8 used as
  supporting soft accent backgrounds.
- Text, controls, icons, focus indicators, validation messages, and chart labels remain readable
  and distinguishable on all three required palette backgrounds.
- Error messages are learner-readable and include retry or recovery actions.
- Destructive actions explain downstream impact before confirmation.

## Completion Criteria

The feature is ready for task execution only when:

- `plan.md`, `research.md`, `data-model.md`, `contracts/openapi.yaml`, and `quickstart.md` are
  present and internally consistent.
- Constitution Check has no unresolved violations.
- The OpenAPI contract covers all primary frontend/backend interactions.
- Quickstart scenarios cover the seven user stories in the specification plus the synchronized
  visual palette validation path.
