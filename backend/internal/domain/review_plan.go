package domain

type ReviewPlan struct {
	ID                 ID
	LearnerWorkspaceID ID
	Name               string
	Goal               string
	DailyCapacity      int
	Status             string
	CurrentRevisionID  ID
}
type PlanRevision struct {
	ID             ID
	ReviewPlanID   ID
	RevisionNumber int
	ChangeSource   string
	ChangeSummary  string
	ChangeReason   string
	PlanSnapshot   map[string]any
}
type ReviewStatisticsSnapshot struct {
	ScopeType      string
	ScopeID        ID
	CardsReviewed  int
	RecallRate     float64
	OverdueCount   int
	CompletionRate float64
	WeakAreaScore  float64
}
