# Feature Specification: AI Card Review Web App

**Feature Branch**: `001-card-review-app`

**Created**: 2026-06-09

**Status**: Draft

**Input**: User description: "创建一个web app 用于制作卡片进行复习，目前想到的功能点包括：资料管理，支持文件/网页/文字输入，添加多个总标签，输入提示词告诉 AI 如何分类并提供默认提示词，AI 分析资料并拆分成多个知识点录入系统；卡片生成，输入提示词，选择标签，由 AI 生成卡组和卡片；卡片/卡组管理，多个卡组可以合并为一个卡组，也可以拆分恢复成多个卡组，卡组中可以创建/编辑卡片，可切换卡组视图和卡片视图；复习，直接选择卡组进行复习，或制定多个复习计划，可以手动/输入提示词由 AI 规划，指定学习计划、学习天数等，最终使用遗忘曲线进行归纳整理，生成统计视图，由 AI 或脚本再次整理优化复习计划，一个复习计划可以查看规划修改历史。"

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

---

### User Story 2 - Generate decks and cards from knowledge points (Priority: P2)

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

---

### User Story 3 - Manage cards and decks flexibly (Priority: P3)

As a learner, I want to browse, create, edit, merge, split, and restore decks and cards so that
my review materials stay accurate and organized as my study goals change.

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

---

### User Story 4 - Review with direct sessions and adaptive plans (Priority: P4)

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
   learner can answer cards, record difficulty or result, and finish with a session summary.
2. **Given** a learner wants a 14-day plan, **When** they select decks, enter study goals, and
   request AI-assisted planning, **Then** the system proposes a daily review schedule that the
   learner can edit before activation.
3. **Given** review sessions have been completed, **When** the learner opens the statistics
   view, **Then** the system shows progress, recall performance, overdue work, and suggested
   plan adjustments based on forgetting-curve timing.
4. **Given** a plan has been optimized multiple times, **When** the learner views plan history,
   **Then** each revision shows what changed, when it changed, and whether it was manual or
   AI-assisted.

---

### Edge Cases

- A submitted file is unreadable, too large for practical processing, duplicated, or contains
  mostly non-text content.
- A web page is unreachable, requires access the learner cannot provide, or contains content
  unrelated to the learner's prompt.
- AI classification or card generation returns low-quality, duplicate, contradictory, or empty
  results.
- A learner applies many tags to materials and later changes or removes a tag used by existing
  knowledge points, decks, cards, or plans.
- A merged deck contains cards that were edited, moved, or deleted after merging before the
  learner attempts to restore original decks.
- A review plan conflicts with another active plan using the same cards or creates more daily
  reviews than the learner's stated capacity.
- A learner misses several scheduled review days and needs the plan to rebalance overdue cards.
- Loading, empty, processing, failure, partial success, disabled, and completion states must be
  visible for material analysis, card generation, deck operations, and plan optimization.
- Views must remain responsive when filtering or switching among up to 500 decks, 5,000 cards,
  1,000 knowledge points, and 200 imported materials in one learner workspace.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The system MUST allow learners to add study material from file upload, web page
  address, or direct text input.
- **FR-002**: The system MUST allow learners to attach one or more tags to each material at
  intake and update those tags later.
- **FR-003**: The system MUST provide a default classification prompt and allow learners to
  customize the prompt before material analysis.
- **FR-004**: The system MUST analyze submitted materials into separate knowledge points that
  include source reference, tags, extracted content, and review status.
- **FR-005**: The system MUST require learner review before extracted knowledge points become
  approved library items.
- **FR-006**: The system MUST allow learners to edit, approve, reject, and search knowledge
  points.
- **FR-007**: The system MUST allow learners to generate decks and cards from approved
  knowledge points using a learner-entered prompt and optional tag filters.
- **FR-008**: The system MUST show generated decks and cards as drafts until the learner
  approves, edits, or discards them.
- **FR-009**: The system MUST preserve traceability from cards to their source knowledge
  points and source materials when generated from existing content.
- **FR-010**: The system MUST allow learners to create, edit, delete, and search cards within
  decks.
- **FR-011**: The system MUST provide both deck view and card view, with consistent filters for
  tag, source, deck, creation status, and review status where applicable.
- **FR-012**: The system MUST allow multiple decks to be merged into a combined deck while
  recording the source decks for later restoration.
- **FR-013**: The system MUST allow a merged deck to be split or restored into its recorded
  source deck grouping while preserving later card edits whenever the card still exists.
- **FR-014**: The system MUST support direct review sessions started from one selected deck.
- **FR-015**: The system MUST record each reviewed card's result, difficulty or confidence,
  review time, and next suggested review timing.
- **FR-016**: The system MUST allow learners to create multiple review plans with selected
  decks, study goals, learning duration, daily capacity, and active or paused status.
- **FR-017**: The system MUST support both manual review plan editing and AI-assisted plan
  creation or optimization from learner prompts.
- **FR-018**: The system MUST organize review timing using forgetting-curve concepts so that
  cards due for reinforcement are prioritized over cards not yet due.
- **FR-019**: The system MUST provide statistics for review progress, recall performance,
  overdue cards, plan completion, and deck-level performance.
- **FR-020**: The system MUST allow review plans to be optimized after reviewing statistics,
  either by learner action, AI-assisted suggestion, or rules-based recalculation.
- **FR-021**: The system MUST keep a modification history for each review plan, including what
  changed, who or what initiated the change, when it changed, and the previous plan state.
- **FR-022**: The system MUST provide consistent loading, empty, error, partial-success,
  disabled, and success states for material intake, AI processing, deck/card management,
  review sessions, statistics, and plan optimization.
- **FR-023**: The system MUST keep learner-facing interactions responsive by showing visible
  feedback within 2 seconds for long-running analysis, generation, merge, restore, and plan
  optimization actions.
- **FR-024**: The system MUST prevent destructive actions such as deleting materials, cards,
  decks, or plan revisions without confirmation and a clear explanation of downstream impact.

### Key Entities *(include if feature involves data)*

- **Learner Workspace**: The learner's private collection of materials, tags, knowledge points,
  decks, cards, review plans, and review history.
- **Source Material**: A file, web page, or text entry submitted by the learner; includes source
  type, title, content summary, tags, processing status, and extraction result summary.
- **Tag**: A learner-defined label used to group materials, knowledge points, decks, cards, and
  review plans.
- **Knowledge Point**: A discrete concept extracted from source material; includes source
  reference, tags, content, approval status, and links to generated cards.
- **Prompt Preset**: A default or learner-edited instruction used for classification, card
  generation, or plan creation.
- **Deck**: A named group of cards used for management or review; may be original, generated,
  manually created, or merged from other decks.
- **Card**: A review item with prompt side, answer side, explanation or notes, tags, source
  links, edit history, and review status.
- **Deck Composition Record**: The relationship data that records how decks were merged or
  split so merged decks can be restored into prior groupings.
- **Review Session**: A direct or plan-based study activity that records reviewed cards,
  learner results, timing, and summary outcomes.
- **Review Plan**: A configurable schedule for one or more decks with goals, duration, daily
  capacity, active state, due-card organization, and optimization settings.
- **Plan Revision**: A historical snapshot of a review plan after manual, AI-assisted, or
  rules-based changes.
- **Review Statistics**: Aggregated progress, recall, overdue, completion, and deck performance
  measures used by the learner and by plan optimization.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: 90% of learners can import a text material, approve at least one extracted
  knowledge point, and find it in the library within 5 minutes on first use.
- **SC-002**: For typical study materials under 5,000 words, learners see analysis started or
  progress feedback within 2 seconds and receive a reviewable result or clear failure state
  within 2 minutes.
- **SC-003**: 85% of generated card drafts from approved knowledge points can be accepted with
  minor edits or no edits during validation sessions.
- **SC-004**: Learners can switch between deck view and card view, apply a tag filter, and open
  an item with visible feedback within 1 second for workspaces containing up to 500 decks and
  5,000 cards.
- **SC-005**: Learners can merge two decks and restore the merged result back to source deck
  grouping without losing surviving card edits in 95% of tested merge/restore scenarios.
- **SC-006**: Learners can create a 14-day review plan for selected decks in under 3 minutes
  using either manual configuration or AI-assisted planning.
- **SC-007**: After completing review sessions, learners can view progress, recall performance,
  overdue cards, and plan adjustment suggestions from a single statistics view.
- **SC-008**: 90% of learners understand whether a review plan change was manual, AI-assisted,
  or rules-based by inspecting the plan revision history.

## Assumptions

- The first release targets individual learners managing their own private study workspace;
  collaboration, public sharing, and classroom administration are outside the initial scope.
- Imported materials are used for personal study, and learners are responsible for having the
  right to process the content they submit.
- AI-assisted outputs are treated as drafts that require learner review before becoming trusted
  knowledge points, cards, or active plan changes.
- The default classification and generation prompts are sufficient for first use but can be
  edited by learners for different subjects or learning styles.
- Review scheduling follows common spaced-repetition and forgetting-curve expectations without
  requiring learners to understand the underlying calculation details.
- Deleted source materials do not automatically delete cards already approved from them, but
  the app must preserve a clear indication when a source reference is no longer available.
- Mobile layout support is desirable for review sessions, but the initial specification is
  centered on a web app experience for desktop and tablet browsers.
