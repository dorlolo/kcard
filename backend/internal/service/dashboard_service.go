package service

type DashboardSummary struct {
	DraftsNeedingApproval int          `json:"draftsNeedingApproval"`
	DueReviews            int          `json:"dueReviews"`
	OverdueReviews        int          `json:"overdueReviews"`
	ActivePlans           int          `json:"activePlans"`
	WeakAreas             []string     `json:"weakAreas"`
	NextActions           []NextAction `json:"nextActions"`
}
type NextAction struct {
	Type     string `json:"type"`
	Label    string `json:"label"`
	TargetID string `json:"targetId,omitempty"`
}
type DashboardService struct{}

func (DashboardService) Summary() DashboardSummary {
	return DashboardSummary{NextActions: []NextAction{{Type: "import", Label: "Import material"}}}
}
