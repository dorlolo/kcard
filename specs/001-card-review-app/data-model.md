# Data Model: AI Card Review Web App

## Conventions

- All durable records have `id`, `created_at`, `updated_at`, and soft-delete or archive fields
  where recovery is required by the specification.
- Records belong to a `learner_workspace_id` unless explicitly global/default.
- AI-generated records remain drafts until approved by the learner.
- Prompt/model/source metadata is retained for learner auditability.
- PostgreSQL enforces identity, relationships, uniqueness, and state constraints; Redis holds
  short-lived job/progress/cache state only.

## Entity: Learner Workspace

**Purpose**: Private boundary for a learner's materials, cards, plans, statistics, prompts, and
preferences.

**Fields**:
- `id`
- `display_name`
- `owner_identity` or account reference
- `default_review_capacity_per_day`
- `default_review_grading_style` (`binary`, `confidence`, `difficulty`)
- `dashboard_focus` preferences
- `privacy_state` (`private` by default)

**Relationships**:
- Has many source materials, tags, knowledge points, prompt presets, decks, cards, review
  sessions, review plans, exports, and statistics snapshots.

**Validation Rules**:
- Workspace content is private by default.
- First release assumes one private workspace per learner account.

## Entity: Learner Preference

**Purpose**: Stores learner defaults used across AI and review workflows.

**Fields**:
- `id`
- `learner_workspace_id`
- `default_classification_prompt_preset_id`
- `default_card_generation_prompt_preset_id`
- `default_plan_prompt_preset_id`
- `daily_capacity_default`
- `review_grading_style`
- `timezone`
- `dashboard_sections`
- `visual_theme_palette_id`

**Relationships**:
- Belongs to learner workspace.
- References optional prompt presets.
- References the required visual theme palette.

**Validation Rules**:
- Daily capacity must be positive.
- Timezone must be valid.

## Entity: Visual Theme Palette

**Purpose**: Defines the required learner-facing soft study palette and readability constraints
used across frontend screens.

**Fields**:
- `id`
- `name`
- `primary_background` (`#fff8e7`)
- `accent_background_1` (`#f8e7ff`)
- `accent_background_2` (`#e7fff8`)
- `semantic_warning_color`
- `semantic_error_color`
- `semantic_success_color`
- `readability_notes`

**Relationships**:
- Referenced by learner preference and frontend design tokens.
- Applies to dashboard, materials, knowledge, cards, review, plans, prompts, statistics, and
  shared UI states.

**Validation Rules**:
- Primary background must remain `#fff8e7` unless a future specification amendment changes the
  visual identity.
- Accent backgrounds must remain `#f8e7ff` and `#e7fff8` for supporting surfaces.
- Text, controls, charts, focus indicators, validation messages, and review feedback must be
  distinguishable on each palette background.
- Palette color alone cannot be the only indicator of status or required action.

## Entity: Source Material

**Purpose**: Represents imported file, web page, or pasted text.

**Fields**:
- `id`
- `learner_workspace_id`
- `source_type` (`file`, `web_page`, `text`)
- `title`
- `original_location` (filename, URL, or text label)
- `content_digest`
- `content_summary`
- `content_status` (`available`, `unavailable`, `deleted`, `changed`)
- `processing_status` (`draft`, `queued`, `processing`, `needs_review`, `processed`, `failed`)
- `failure_reason`
- `duplicate_status` (`unchecked`, `possible_duplicate`, `confirmed_duplicate`, `unique`)
- `current_version_id`
- `archived_at`

**Relationships**:
- Belongs to learner workspace.
- Has many material versions.
- Has many tags through tag assignments.
- Has many knowledge points.
- May have many AI jobs.

**Validation Rules**:
- Material must have one source type.
- Web page sources must retain the submitted URL.
- Text sources must retain a content digest for duplicate detection.
- Failed processing must include a learner-readable failure reason.

**State Transitions**:
- `draft` → `queued` → `processing` → `needs_review` → `processed`
- `queued`/`processing` → `failed`
- `processed` → `queued` when re-analysis is requested
- Any active state → archived when no longer used actively

## Entity: Material Version

**Purpose**: Snapshot or update marker when a material is edited or re-imported.

**Fields**:
- `id`
- `source_material_id`
- `version_number`
- `content_digest`
- `content_location`
- `summary`
- `created_by_action` (`import`, `manual_edit`, `reimport`)

**Relationships**:
- Belongs to source material.
- Knowledge points reference the version that produced them.

**Validation Rules**:
- Version number is unique within a source material.
- Content digest is required for duplicate/change detection.

## Entity: Tag

**Purpose**: Learner-defined label for organizing materials, knowledge points, cards, decks,
plans, and statistics.

**Fields**:
- `id`
- `learner_workspace_id`
- `name`
- `color`
- `description`
- `archived_at`

**Relationships**:
- Belongs to learner workspace.
- Many-to-many with source materials, knowledge points, cards, decks, and review plans.

**Validation Rules**:
- Tag name is unique case-insensitively within a workspace.
- Archived tags remain visible in historical records but hidden from default pickers.

## Entity: Knowledge Point

**Purpose**: Curated learning concept extracted from material or created manually.

**Fields**:
- `id`
- `learner_workspace_id`
- `source_material_id`
- `material_version_id`
- `content`
- `summary`
- `notes`
- `approval_status` (`draft`, `approved`, `rejected`, `needs_review`)
- `creation_source` (`ai_generated`, `manual`, `imported`)
- `duplicate_group_id`
- `ai_job_id`
- `prompt_snapshot_id`
- `approved_at`
- `rejected_at`
- `archived_at`

**Relationships**:
- Belongs to learner workspace.
- Optionally belongs to source material and material version.
- Has many tags.
- Has many cards through card source links.
- May belong to a duplicate group.

**Validation Rules**:
- Approved knowledge point must have non-empty content.
- Rejected knowledge point is excluded from default card generation.
- Split knowledge points inherit source references and selected tags.

**State Transitions**:
- `draft` → `approved`
- `draft` → `rejected`
- `draft` → `needs_review`
- `needs_review` → `approved` or `rejected`
- `approved` → `needs_review` when source changes or learner marks it uncertain

## Entity: Prompt Preset

**Purpose**: Reusable learner-editable prompt for classification, card generation, plan
creation, plan optimization, or cleanup.

**Fields**:
- `id`
- `learner_workspace_id`
- `name`
- `purpose` (`classification`, `card_generation`, `plan_creation`, `plan_optimization`, `cleanup`)
- `prompt_text`
- `is_default`
- `version_number`
- `archived_at`

**Relationships**:
- Belongs to learner workspace.
- Has many prompt snapshots.

**Validation Rules**:
- Name is unique per workspace and purpose.
- Prompt text cannot be empty.
- Only one active default per purpose per workspace.

## Entity: Prompt Snapshot

**Purpose**: Immutable prompt text used for a specific AI action.

**Fields**:
- `id`
- `prompt_preset_id`
- `purpose`
- `prompt_text`
- `model_id`
- `schema_version`
- `created_for_job_id`

**Relationships**:
- Optionally references prompt preset.
- Referenced by AI drafts, cards, knowledge points, review plans, and jobs.

**Validation Rules**:
- Snapshot is immutable after creation.
- Model ID must be recorded for AI-assisted outputs.

## Entity: AI Job

**Purpose**: Tracks long-running AI or import/export work.

**Fields**:
- `id`
- `learner_workspace_id`
- `job_type` (`material_analysis`, `card_generation`, `plan_generation`, `plan_optimization`, `import`, `export`)
- `status` (`queued`, `running`, `partial_success`, `succeeded`, `failed`, `cancelled`)
- `progress_percent`
- `current_step`
- `input_summary`
- `error_category`
- `error_message`
- `idempotency_key`
- `started_at`
- `finished_at`

**Relationships**:
- Belongs to learner workspace.
- References prompt snapshot when AI-assisted.
- Produces AI drafts or export/import results.

**Validation Rules**:
- Failed jobs must include an error message safe to show to learners.
- Idempotency key prevents duplicate job creation for repeated submissions.

**State Transitions**:
- `queued` → `running` → `succeeded`
- `running` → `partial_success`
- `queued`/`running` → `failed`
- `queued`/`running` → `cancelled`

## Entity: AI Draft

**Purpose**: Proposed knowledge points, cards, decks, or review plans pending learner approval.

**Fields**:
- `id`
- `learner_workspace_id`
- `draft_type` (`knowledge_point`, `card`, `deck`, `review_plan`)
- `job_id`
- `payload`
- `validation_status` (`valid`, `warning`, `invalid`)
- `validation_messages`
- `status` (`draft`, `approved`, `discarded`, `superseded`)
- `approved_record_id`

**Relationships**:
- Belongs to learner workspace and AI job.
- May produce a knowledge point, card, deck, or review plan.

**Validation Rules**:
- Draft payload must match its schema version.
- Approved draft must point to the durable record it created or updated.

## Entity: Deck

**Purpose**: Named group of cards used for organization or review.

**Fields**:
- `id`
- `learner_workspace_id`
- `name`
- `description`
- `creation_source` (`manual`, `ai_generated`, `merged`, `imported`)
- `status` (`draft`, `active`, `archived`, `deleted`)
- `ai_job_id`
- `prompt_snapshot_id`
- `archived_at`

**Relationships**:
- Belongs to learner workspace.
- Has many cards through deck card memberships.
- Has many tags.
- Has many deck composition records when merged/split.
- May be referenced by review plans.

**Validation Rules**:
- Active deck must have a name.
- Archived deck is excluded from default review selection but remains searchable.

**State Transitions**:
- `draft` → `active`
- `active` → `archived`
- `archived` → `active`
- `active`/`archived` → `deleted` when confirmed

## Entity: Card

**Purpose**: Review item with prompt and answer content.

**Fields**:
- `id`
- `learner_workspace_id`
- `front_prompt`
- `back_answer`
- `explanation`
- `learner_notes`
- `difficulty` (`unrated`, `easy`, `medium`, `hard`)
- `creation_source` (`manual`, `ai_generated`, `imported`)
- `status` (`draft`, `active`, `archived`, `deleted`)
- `review_status` (`new`, `learning`, `due`, `not_due`, `suspended`)
- `ai_job_id`
- `prompt_snapshot_id`
- `archived_at`

**Relationships**:
- Belongs to learner workspace.
- Has many deck memberships.
- Has many tags.
- Has many source links to knowledge points and materials.
- Has many review results.

**Validation Rules**:
- Active card requires front prompt and back answer.
- Draft card must be approved before appearing in review queues.
- Deleted card is excluded from future scheduling but historical results remain.

**State Transitions**:
- `draft` → `active`
- `active` → `archived`
- `archived` → `active`
- `active`/`archived` → `deleted` when confirmed

## Entity: Card Source Link

**Purpose**: Preserves traceability from a card to source knowledge/material and prompt.

**Fields**:
- `id`
- `card_id`
- `knowledge_point_id`
- `source_material_id`
- `material_version_id`
- `source_quote`
- `confidence_note`

**Relationships**:
- Belongs to card.
- Optionally references knowledge point, source material, and material version.

**Validation Rules**:
- AI-generated cards should have at least one source link.
- If source material is deleted, link remains with unavailable status.

## Entity: Deck Composition Record

**Purpose**: Records deck merge/split relationships for restoration.

**Fields**:
- `id`
- `learner_workspace_id`
- `operation_type` (`merge`, `split`, `restore`)
- `result_deck_id`
- `source_deck_ids`
- `card_membership_snapshot`
- `operation_note`

**Relationships**:
- Belongs to learner workspace.
- References source and result decks.

**Validation Rules**:
- Merge requires at least two source decks.
- Restore must preserve surviving card edits and explain missing deleted cards.

## Entity: Review Session

**Purpose**: Direct or plan-based study activity.

**Fields**:
- `id`
- `learner_workspace_id`
- `session_type` (`direct`, `plan_based`)
- `deck_id`
- `review_plan_id`
- `status` (`active`, `paused`, `completed`, `abandoned`)
- `started_at`
- `paused_at`
- `completed_at`
- `summary`

**Relationships**:
- Belongs to learner workspace.
- Has many review results.
- Optionally belongs to deck or review plan.

**Validation Rules**:
- Session must reference a deck, due-card list, filtered set, or review plan.
- Completed session must have a summary.

**State Transitions**:
- `active` → `paused`
- `paused` → `active`
- `active`/`paused` → `completed`
- `active`/`paused` → `abandoned`

## Entity: Review Result

**Purpose**: Learner's answer outcome for one card in one session.

**Fields**:
- `id`
- `review_session_id`
- `card_id`
- `result` (`again`, `hard`, `good`, `easy`, `correct`, `incorrect`)
- `confidence`
- `elapsed_response_ms`
- `reviewed_at`
- `next_due_at`
- `schedule_reason`

**Relationships**:
- Belongs to review session and card.
- Contributes to review statistics and plan optimization.

**Validation Rules**:
- Result must match workspace review grading style.
- Next due date is required for active cards.

## Entity: Review Plan

**Purpose**: Configurable multi-day schedule for decks/cards.

**Fields**:
- `id`
- `learner_workspace_id`
- `name`
- `goal`
- `start_date`
- `target_completion_date`
- `learning_duration_days`
- `daily_capacity`
- `status` (`draft`, `active`, `paused`, `completed`, `archived`)
- `overlap_policy` (`prevent`, `allow_with_warning`)
- `optimization_mode` (`manual`, `ai_assisted`, `rules_based`)
- `current_revision_id`

**Relationships**:
- Belongs to learner workspace.
- Has many deck assignments.
- Has many plan revisions.
- Has many review sessions.
- Has many tags.

**Validation Rules**:
- Active plan requires selected decks or due-card criteria.
- Daily capacity must be positive.
- Active plans cannot double-schedule the same card unless overlap is explicitly allowed.

**State Transitions**:
- `draft` → `active`
- `active` → `paused`
- `paused` → `active`
- `active` → `completed`
- `active`/`paused`/`completed` → `archived`

## Entity: Plan Revision

**Purpose**: Immutable snapshot of review plan changes.

**Fields**:
- `id`
- `review_plan_id`
- `revision_number`
- `change_source` (`manual`, `ai_assisted`, `rules_based`)
- `change_summary`
- `change_reason`
- `plan_snapshot`
- `prompt_snapshot_id`
- `created_at`

**Relationships**:
- Belongs to review plan.
- Optionally references prompt snapshot.

**Validation Rules**:
- Revision number is unique within plan.
- Snapshot is immutable.
- Restore of revision cannot erase completed review history.

## Entity: Review Statistics Snapshot

**Purpose**: Aggregated progress and weak-area metrics for dashboard and optimization.

**Fields**:
- `id`
- `learner_workspace_id`
- `scope_type` (`workspace`, `tag`, `deck`, `plan`, `card`)
- `scope_id`
- `period_start`
- `period_end`
- `cards_reviewed`
- `recall_rate`
- `overdue_count`
- `completion_rate`
- `weak_area_score`
- `generated_at`

**Relationships**:
- Belongs to learner workspace.
- References optional scope entity.

**Validation Rules**:
- Period end must be after period start.
- Metrics must be reproducible from review results for the same period.

## Entity: Workspace Export

**Purpose**: Backup or migration package for study content.

**Fields**:
- `id`
- `learner_workspace_id`
- `status` (`queued`, `running`, `ready`, `failed`, `expired`)
- `included_content` flags
- `file_location`
- `expires_at`
- `failure_reason`

**Relationships**:
- Belongs to learner workspace.
- May be produced by an AI/import/export job.

**Validation Rules**:
- Ready export must have a file location and expiration.
- Export should include schema version for future imports.

## Cross-Entity Rules

- AI-generated trusted items must retain `creation_source`, `prompt_snapshot_id`, and model ID
  through prompt snapshot.
- Tags can be archived but historical references remain intact.
- Source deletion does not delete approved cards; card source links indicate source
  unavailability.
- Review history is append-only except for explicit, audited deletion flows.
- Drafts are excluded from active review queues until approved.
- Archive is preferred over deletion where recovery is expected.
- Long-running AI/import/export operations write durable `AI Job` records and short-lived Redis
  progress events.
