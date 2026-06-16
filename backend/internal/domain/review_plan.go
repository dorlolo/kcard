// Package domain 提供复习计划相关类型。
package domain

// ReviewPlan 表示一个复习计划。
type ReviewPlan struct {
	ID                 ID
	LearnerWorkspaceID ID
	Name               string
	Goal               string
	DailyCapacity      int
	Status             string
	CurrentRevisionID  ID
}
// PlanRevision 表示复习计划的修订版本。
type PlanRevision struct {
	ID             ID
	ReviewPlanID   ID
	RevisionNumber int
	ChangeSource   string
	ChangeSummary  string
	ChangeReason   string
	PlanSnapshot   map[string]any
}
// ReviewStatisticsSnapshot 表示复习统计的快照。
type ReviewStatisticsSnapshot struct {
	ScopeType      string
	ScopeID        ID
	CardsReviewed  int
	RecallRate     float64
	OverdueCount   int
	CompletionRate float64
	WeakAreaScore  float64
}
