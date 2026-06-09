# Feature Specification: AI Card Review Web App

**Feature Branch**: `001-card-review-app`

**Created**: 2026-06-09

**Status**: Draft

**Input**: User description: "创建一个web app 用于制作卡片进行复习，目前想到的功能点包括：资料管理，支持文件/网页/文字输入，添加多个总标签，输入提示词告诉 AI 如何分类并提供默认提示词，AI 分析资料并拆分成多个知识点录入系统；卡片生成，输入提示词，选择标签，由 AI 生成卡组和卡片；卡片/卡组管理，多个卡组可以合并为一个卡组，也可以拆分恢复成多个卡组，卡组中可以创建/编辑卡片，可切换卡组视图和卡片视图；复习，直接选择卡组进行复习，或制定多个复习计划，可以手动/输入提示词由 AI 规划，指定学习计划、学习天数等，最终使用遗忘曲线进行归纳整理，生成统计视图，由 AI 或脚本再次整理优化复习计划，一个复习计划可以查看规划修改历史。后续补充：进一步查看与分析，使用最合理的功能完善现有规格需求，使用合理默认值，不提出澄清问题。"

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Import materials and extract knowledge points (Priority: P1)

As a learner, I want to add study materials from files, web pages, or pasted text, tag them,
and guide AI classification with a prompt so that the app can turn raw material into organized
knowledge points ready for card creation.

**Why this priority**: Material intake and knowledge point extraction are the foundation for all
later card generation and review workflows.

**Independent Test**: A learner can submit one text material with tags and the default prompt,
review the extracted knowledge points, edit or reject them, and save approved points to the
library without using card generation or review planning.

**Acceptance Scenarios**:

1. **Given** a learner has text material and two tags, **When** they submit the material with
   the default classification prompt, **Then** the system creates a material record, shows
   processing status, and presents extracted knowledge points for review.
2. **Given** extracted knowledge points are shown, **When** the learner edits one point,
   rejects another, and approves the rest, **Then** only the approved and edited points are
   saved with the selected tags and source reference.
3. **Given** a web page or file cannot be read, **When** the learner submits it, **Then** the
   learner receives a clear failure reason and can retry, replace the source, or enter text
   manually.
4. **Given** a learner submits material similar to an existing source, **When** the analysis is
   prepared, **Then** the system warns about possible duplication and lets the learner continue,
   cancel, or associate the new material with the existing source.

---

### User Story 2 - Curate knowledge before card generation (Priority: P2)

As a learner, I want a knowledge point library where I can search, filter, merge, split, and
mark knowledge points so that AI-generated content stays accurate before it becomes cards.

**Why this priority**: Knowledge points are the bridge between raw materials and cards; poor
curation produces poor review quality.

**Independent Test**: Starting from extracted knowledge points, a learner can find duplicate or
related points, merge two points, split one overloaded point, mark one as needing review, and
then use only approved points for card generation.

**Acceptance Scenarios**:

1. **Given** multiple extracted knowledge points share similar wording, **When** the learner
   opens the library, **Then** the system highlights possible duplicates or related points for
   review.
2. **Given** a knowledge point contains multiple ideas, **When** the learner splits it, **Then**
   each resulting point keeps the original source reference and selected tags.
3. **Given** a knowledge point is marked as needing review, **When** the learner generates
   cards, **Then** that point is excluded unless the learner explicitly includes unapproved
   content.

---

### User Story 3 - Generate decks and cards from knowledge points (Priority: P3)

As a learner, I want to choose tags and provide a generation prompt so that AI can create card
decks and cards from existing knowledge points, while allowing me to review and adjust the
result before using it.

**Why this priority**: The core product value is turning organized knowledge into reusable
review cards quickly.

**Independent Test**: Starting from existing knowledge points, a learner can select one or more
tags, run card generation with a prompt, review a proposed deck and cards, edit card content,
and save the result as a usable deck.

**Acceptance Scenarios**:

1. **Given** knowledge points exist under selected tags, **When** the learner enters a card
   generation prompt and starts generation, **Then** the system proposes one or more decks
   containing cards linked back to the source knowledge points.
2. **Given** generated cards are shown, **When** the learner edits card wording, deletes weak
   cards, and approves the deck, **Then** the approved deck appears in deck management and its
   cards appear in card view.
3. **Given** no knowledge points match the selected tags, **When** the learner starts
   generation, **Then** the system explains that no source content is available and suggests
   selecting other tags or importing materials.
4. **Given** generated cards include duplicates or cards with unclear answers, **When** the
   learner reviews the draft, **Then** the system flags the issue and lets the learner edit,
   merge, regenerate, or discard the affected cards.

---

### User Story 4 - Manage cards and decks flexibly (Priority: P4)

As a learner, I want to browse, create, edit, merge, split, restore, archive, and organize decks
and cards so that my review materials stay accurate as my study goals change.

**Why this priority**: Review quality depends on learners being able to maintain and reorganize
card content after generation.

**Independent Test**: A learner can create a manual card in a deck, switch between deck and
card views, merge two decks, then restore the merged deck back into its original deck grouping
while preserving card edits.

**Acceptance Scenarios**:

1. **Given** a learner is in deck view, **When** they select multiple decks and merge them,
   **Then** a new combined deck is created and records which original decks contributed cards.
2. **Given** a deck was created by merging other decks, **When** the learner chooses to split
   or restore it, **Then** the original deck grouping can be recreated without losing card edits
   made after the merge.
3. **Given** a learner is browsing cards, **When** they switch from deck view to card view,
   **Then** the same cards remain filterable by deck, tag, source material, and review status.
4. **Given** a deck is no longer needed for active review, **When** the learner archives it,
   **Then** the deck is hidden from active review choices but remains searchable and restorable.

---

### User Story 5 - Review with direct sessions and adaptive plans (Priority: P5)

As a learner, I want to either review a selected deck immediately or maintain multiple review
plans that can be manually configured or AI-assisted, so that my study schedule follows my
goals and adapts to forgetting-curve performance.

**Why this priority**: Reviewing cards and improving future review timing completes the study
loop and creates measurable learning value.

**Independent Test**: A learner can start a direct review session from one deck, create a
multi-day review plan for selected decks, complete reviews, view statistics, optimize the plan,
and inspect the plan's modification history.

**Acceptance Scenarios**:

1. **Given** a deck contains cards, **When** the learner starts direct review, **Then** the
   learner can answer cards, record difficulty or result, pause or resume the session, and
   finish with a session summary.
2. **Given** a learner wants a 14-day plan, **When** they select decks, enter study goals, and
   request AI-assisted planning, **Then** the system proposes a daily review schedule that the
   learner can edit before activation.
3. **Given** review sessions have been completed, **When** the learner opens the statistics
   view, **Then** the system shows progress, recall performance, overdue work, and suggested
   plan adjustments based on forgetting-curve timing.
4. **Given** a plan has been optimized multiple times, **When** the learner views plan history,
   **Then** each revision shows what changed, when it changed, and whether it was manual,
   AI-assisted, or rules-based.
5. **Given** a learner misses scheduled review days, **When** the learner returns, **Then** the
   system shows overdue work, proposes a realistic recovery schedule, and avoids overwhelming
   the learner beyond the stated daily capacity unless the learner overrides it.

---

### User Story 6 - Manage prompts, preferences, and AI draft safety (Priority: P6)

As a learner, I want default prompts, saved prompt presets, and clear AI-draft boundaries so
that I can control how materials, cards, and plans are generated without trusting AI output
blindly.

**Why this priority**: Prompt control and draft review make AI-assisted workflows reliable,
repeatable, and safe for learning.

**Independent Test**: A learner can use a default prompt, save a customized prompt preset,
apply it to card generation, compare the draft with source references, and approve only the
usable result.

**Acceptance Scenarios**:

1. **Given** the learner customizes a classification prompt, **When** they save it as a preset,
   **Then** the preset is available for later material analysis and can be renamed or deleted.
2. **Given** AI produces knowledge points, cards, or a plan suggestion, **When** the learner
   reviews the output, **Then** the system clearly marks it as a draft and shows source or input
   context where available.
3. **Given** AI generation fails or produces an unusable draft, **When** the learner reviews the
   result, **Then** the system provides retry, edit prompt, use partial result, or discard
   actions.

---

### User Story 7 - Understand progress and focus next actions (Priority: P7)

As a learner, I want a dashboard that summarizes current materials, card readiness, due
reviews, weak areas, and next recommended actions so that I know what to do when I open the app.

**Why this priority**: A learning app must reduce planning friction and keep learners oriented
across materials, cards, and review plans.

**Independent Test**: A learner with imported materials, generated decks, and active review
plans can open the dashboard and see pending knowledge reviews, draft cards, due reviews,
weak tags, and the next recommended action.

**Acceptance Scenarios**:

1. **Given** the learner has pending drafts and due reviews, **When** they open the dashboard,
   **Then** the system shows separate action groups for reviewing AI drafts, continuing review,
   and fixing overdue plan items.
2. **Given** a tag has low recall performance, **When** the learner views statistics, **Then**
   the system identifies the tag or deck as a weak area and suggests focused review or card
   improvement.
3. **Given** there is no content yet, **When** the learner opens the app, **Then** the system
   shows an empty state that guides them to import material or create cards manually.
4. **Given** the learner moves across dashboard, material review, card editing, review, and
   statistics screens, **When** each screen loads, **Then** the interface uses a consistent soft
   study palette with #fff8e7 as the primary background tone and #f8e7ff / #e7fff8 as
   supporting background accents without reducing readability.

---

### Edge Cases

- A submitted file is unreadable, too large for practical processing, duplicated, password
  protected, corrupted, or contains mostly non-text content.
- A web page is unreachable, requires access the learner cannot provide, changes after import,
  or contains content unrelated to the learner's prompt.
- AI classification, card generation, or plan generation returns low-quality, duplicate,
  contradictory, unsupported, unsafe, or empty results.
- A learner edits a source material after knowledge points or cards have already been created
  from it.
- A learner applies many tags to materials and later renames, merges, or removes a tag used by
  existing knowledge points, decks, cards, or plans.
- A learner bulk-approves AI output accidentally and needs to undo or return items to draft
  status.
- A generated card links to a knowledge point that is later merged, split, rejected, or deleted.
- A merged deck contains cards that were edited, moved, archived, or deleted after merging
  before the learner attempts to restore original decks.
- A review plan conflicts with another active plan using the same cards or creates more daily
  reviews than the learner's stated capacity.
- A learner pauses, exits, or loses connection during a review session and later resumes.
- A learner misses several scheduled review days and needs the plan to rebalance overdue cards.
- A learner changes the target completion date or daily capacity of an active plan.
- Loading, empty, processing, failure, partial success, disabled, undo, and completion states
  must be visible for material analysis, card generation, deck operations, review sessions,
  statistics, and plan optimization.
- The required soft palette (#fff8e7 primary, #f8e7ff and #e7fff8 accents) must remain readable
  in dense views, long review sessions, error states, disabled controls, charts, and empty
  states, including for learners using keyboard navigation or high-contrast display settings.
- Views must remain responsive when filtering or switching among up to 500 decks, 5,000 cards,
  1,000 knowledge points, 200 imported materials, 100 prompt presets, and 20 active or paused
  review plans in one learner workspace.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The system MUST allow learners to add study material from file upload, web page
  address, or direct text input.
- **FR-002**: The system MUST allow learners to attach one or more tags to each material at
  intake and update those tags later.
- **FR-003**: The system MUST provide default prompts for material classification, card
  generation, and review plan generation or optimization.
- **FR-004**: The system MUST allow learners to customize prompts before each AI-assisted
  action and save reusable prompt presets.
- **FR-005**: The system MUST analyze submitted materials into separate knowledge points that
  include source reference, tags, extracted content, draft or approval status, and review notes.
- **FR-006**: The system MUST require learner review before extracted knowledge points become
  approved library items.
- **FR-007**: The system MUST allow learners to edit, approve, reject, restore, search, filter,
  merge, and split knowledge points.
- **FR-008**: The system MUST identify likely duplicate materials, knowledge points, and cards
  and allow the learner to resolve duplicates without losing source references.
- **FR-009**: The system MUST allow learners to generate decks and cards from approved
  knowledge points using a learner-entered prompt and optional filters for tags, source
  materials, knowledge status, and existing decks.
- **FR-010**: The system MUST show generated decks, cards, knowledge points, and review plans
  as drafts until the learner approves, edits, or discards them.
- **FR-011**: The system MUST preserve traceability from cards to their source knowledge points,
  source materials, and generation prompt when generated from existing content.
- **FR-012**: The system MUST allow learners to create, edit, duplicate, archive, restore,
  delete, and search cards within decks.
- **FR-013**: The system MUST support common card content fields: front prompt, back answer,
  explanation, source reference, tags, difficulty, and learner notes.
- **FR-014**: The system MUST provide both deck view and card view, with consistent filters for
  tag, source, deck, draft state, creation source, difficulty, due state, and review status
  where applicable.
- **FR-015**: The system MUST allow multiple decks to be merged into a combined deck while
  recording the source decks and card membership for later restoration.
- **FR-016**: The system MUST allow a merged deck to be split or restored into its recorded
  source deck grouping while preserving later card edits whenever the card still exists.
- **FR-017**: The system MUST allow learners to archive decks without deleting cards or review
  history.
- **FR-018**: The system MUST support direct review sessions started from one selected deck,
  filtered card set, or due-card list.
- **FR-019**: The system MUST allow review sessions to be paused, resumed, completed, or
  abandoned with clear impact on recorded results.
- **FR-020**: The system MUST record each reviewed card's result, difficulty or confidence,
  review time, elapsed response time, session context, and next suggested review timing.
- **FR-021**: The system MUST allow learners to create multiple review plans with selected
  decks, goals, learning duration, daily capacity, start date, target completion date, and
  active, paused, completed, or archived status.
- **FR-022**: The system MUST support both manual review plan editing and AI-assisted plan
  creation or optimization from learner prompts.
- **FR-023**: The system MUST organize review timing using forgetting-curve concepts so that
  cards due for reinforcement are prioritized over cards not yet due.
- **FR-024**: The system MUST avoid double-scheduling the same card across active plans unless
  the learner explicitly allows overlap.
- **FR-025**: The system MUST provide statistics for review progress, recall performance,
  overdue cards, plan completion, weak tags, weak decks, knowledge coverage, and deck-level
  performance.
- **FR-026**: The system MUST allow review plans to be optimized after reviewing statistics,
  either by learner action, AI-assisted suggestion, or rules-based recalculation.
- **FR-027**: The system MUST keep a modification history for each review plan, including what
  changed, who or what initiated the change, why the change was suggested when available, when
  it changed, and the previous plan state.
- **FR-028**: The system MUST allow learners to compare plan revisions and restore a prior
  revision when doing so does not conflict with completed review history.
- **FR-029**: The system MUST provide a dashboard that shows draft items needing approval, due
  reviews, overdue reviews, active plan status, weak areas, and recommended next actions.
- **FR-030**: The system MUST provide consistent loading, empty, error, partial-success,
  disabled, undo, and success states for material intake, AI processing, knowledge curation,
  deck/card management, review sessions, statistics, and plan optimization.
- **FR-031**: The system MUST use a consistent soft study visual palette across all learner-facing
  screens: #fff8e7 as the primary background or base tone, with #f8e7ff and #e7fff8 as
  supporting accent backgrounds for grouped content, highlights, status panels, or secondary
  surfaces.
- **FR-032**: The system MUST ensure text, controls, icons, charts, focus indicators, validation
  messages, and review feedback remain readable and distinguishable when shown on #fff8e7,
  #f8e7ff, or #e7fff8 backgrounds.
- **FR-033**: The system MUST avoid introducing unrelated dominant background colors for core
  learner workflows unless the color communicates a clear state such as destructive warning,
  critical error, or success confirmation.
- **FR-034**: The system MUST keep learner-facing interactions responsive by showing visible
  feedback within 2 seconds for long-running analysis, generation, merge, restore, search,
  filter, import, export, and plan optimization actions.
- **FR-035**: The system MUST prevent destructive actions such as deleting materials, knowledge
  points, cards, decks, prompt presets, sessions, or plan revisions without confirmation and a
  clear explanation of downstream impact.
- **FR-036**: The system MUST provide undo or restore paths for accidental archive, reject,
  draft discard, merge, split, and delete actions where the item can be safely recovered.
- **FR-037**: The system MUST make AI-assisted actions transparent by showing whether content
  was manually created, AI-generated, AI-optimized, or rules-generated.
- **FR-038**: The system MUST keep learner workspace content private by default and must not
  expose imported materials, generated cards, or review statistics to other learners.
- **FR-039**: The system MUST allow learners to export their cards, decks, knowledge points,
  review plans, and review statistics in a portable study format suitable for backup or later
  migration.
- **FR-040**: The system MUST allow learners to import previously exported study content and
  preview conflicts before confirming the import.
- **FR-041**: The system MUST provide accessible keyboard navigation and readable status
  feedback for material review, card editing, deck management, review sessions, and statistics.
- **FR-042**: The system MUST record enough audit information for AI-assisted changes so a
  learner can understand which prompt, source content, and approval action produced a trusted
  item.
- **FR-043**: The system MUST protect review continuity by preserving in-progress work when the
  learner refreshes the page, navigates away, or temporarily loses connectivity.

### Key Entities *(include if feature involves data)*

- **Learner Workspace**: The learner's private collection of materials, tags, knowledge points,
  prompt presets, decks, cards, review plans, review history, statistics, and preferences.
- **Learner Preference**: The learner's saved defaults for prompts, review capacity, display
  preferences, review grading style, and dashboard focus.
- **Visual Theme Palette**: The required soft study palette for learner-facing screens, using
  #fff8e7 as the primary background/base tone and #f8e7ff plus #e7fff8 as supporting accent
  backgrounds while preserving readable text and accessible state feedback.
- **Source Material**: A file, web page, or text entry submitted by the learner; includes source
  type, title, content summary, tags, processing status, duplicate status, extraction summary,
  and source availability.
- **Material Version**: A snapshot or update marker for source material when the learner edits
  text content or re-imports a changed source.
- **Tag**: A learner-defined label used to group materials, knowledge points, decks, cards,
  review plans, and statistics.
- **Knowledge Point**: A discrete concept extracted from source material; includes source
  reference, tags, content, approval status, notes, duplicate relationship, and links to
  generated cards.
- **Prompt Preset**: A default or learner-edited instruction used for classification, card
  generation, plan creation, plan optimization, or result cleanup.
- **AI Draft**: A proposed knowledge point, card, deck, or review plan that remains untrusted
  until the learner approves, edits, or discards it.
- **Deck**: A named group of cards used for management or review; may be original, generated,
  manually created, archived, or merged from other decks.
- **Card**: A review item with front prompt, back answer, explanation, source links, tags,
  difficulty, learner notes, edit history, and review status.
- **Deck Composition Record**: The relationship data that records how decks were merged or
  split so merged decks can be restored into prior groupings.
- **Review Session**: A direct or plan-based study activity that records reviewed cards,
  learner results, timing, pauses, resumes, and summary outcomes.
- **Review Result**: A learner's answer outcome for one card in one session, including
  confidence or difficulty, elapsed response time, and next suggested timing.
- **Review Plan**: A configurable schedule for one or more decks with goals, duration, daily
  capacity, active state, due-card organization, conflict handling, and optimization settings.
- **Plan Revision**: A historical snapshot of a review plan after manual, AI-assisted, or
  rules-based changes.
- **Review Statistics**: Aggregated progress, recall, overdue, completion, weak-area,
  knowledge-coverage, and deck performance measures used by the learner and by plan
  optimization.
- **Workspace Export**: A learner-created backup or transfer package containing selected study
  content and metadata.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: 90% of learners can import a text material, approve at least one extracted
  knowledge point, and find it in the library within 5 minutes on first use.
- **SC-002**: For typical study materials under 5,000 words, learners see analysis started or
  progress feedback within 2 seconds and receive a reviewable result or clear failure state
  within 2 minutes.
- **SC-003**: 85% of generated card drafts from approved knowledge points can be accepted with
  minor edits or no edits during validation sessions.
- **SC-004**: Learners can find and resolve a duplicate knowledge point or card within 2 minutes
  after the app flags the possible duplicate.
- **SC-005**: Learners can switch between deck view and card view, apply a tag filter, and open
  an item with visible feedback within 1 second for workspaces containing up to 500 decks and
  5,000 cards.
- **SC-006**: Learners can merge two decks and restore the merged result back to source deck
  grouping without losing surviving card edits in 95% of tested merge/restore scenarios.
- **SC-007**: Learners can create a 14-day review plan for selected decks in under 3 minutes
  using either manual configuration or AI-assisted planning.
- **SC-008**: Learners can pause or exit a review session and resume it later with no loss of
  answered-card results in 99% of tested interruption scenarios.
- **SC-009**: After completing review sessions, learners can view progress, recall performance,
  overdue cards, weak areas, and plan adjustment suggestions from a single statistics view.
- **SC-010**: 90% of learners understand whether a trusted item or review plan change was
  manual, AI-assisted, or rules-based by inspecting item details or plan revision history.
- **SC-011**: 90% of learners can identify their next recommended action from the dashboard
  within 30 seconds after opening a workspace containing drafts and due reviews.
- **SC-012**: Learners can export a workspace backup and import it into an empty workspace while
  preserving cards, decks, tags, review plans, and review history in 95% of validation cases.
- **SC-013**: 90% of learners describe the interface as visually consistent across dashboard,
  material review, card editing, review, and statistics screens, with the #fff8e7 primary tone
  and #f8e7ff / #e7fff8 accents recognizable in usability review.
- **SC-014**: All primary text, controls, focus indicators, error states, and chart labels remain
  readable and distinguishable on #fff8e7, #f8e7ff, and #e7fff8 backgrounds during accessibility
  validation.

## Assumptions

- The first release targets individual learners managing their own private study workspace;
  collaboration, public sharing, marketplace features, and classroom administration are outside
  the initial scope.
- Imported materials are used for personal study, and learners are responsible for having the
  right to process the content they submit.
- AI-assisted outputs are treated as drafts that require learner review before becoming trusted
  knowledge points, cards, decks, or active plan changes.
- Default prompts are available for first use, while custom prompt presets support different
  subjects, exam types, memory styles, and card formats.
- Review scheduling follows common spaced-repetition and forgetting-curve expectations without
  requiring learners to understand the underlying calculation details.
- The learner can manually override AI-assisted or rules-based recommendations when they accept
  responsibility for the resulting plan workload.
- Deleted source materials do not automatically delete cards already approved from them, but
  the app must preserve a clear indication when a source reference is no longer available.
- Archive is the default safe alternative to deletion for decks, cards, prompt presets, and
  review plans that may be needed later.
- Mobile layout support is desirable for review sessions and quick checks, but the initial
  specification is centered on a web app experience for desktop and tablet browsers.
- The learner-facing visual identity uses #fff8e7 as the primary background/base tone, with
  #f8e7ff and #e7fff8 used as supporting soft accent backgrounds rather than unrelated dominant
  page colors.
- Export and import are intended for learner backup and portability, not public sharing or
  multi-user collaboration in the first release.
