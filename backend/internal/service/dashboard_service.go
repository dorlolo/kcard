// Package service 提供领域服务层的实现，包括资料、卡片、知识图谱等核心业务逻辑。
package service

// DashboardSummary 表示仪表盘概览摘要数据。
type DashboardSummary struct {
	DraftsNeedingApproval int          `json:"draftsNeedingApproval"`
	DueReviews            int          `json:"dueReviews"`
	OverdueReviews        int          `json:"overdueReviews"`
	ActivePlans           int          `json:"activePlans"`
	WeakAreas             []string     `json:"weakAreas"`
	NextActions           []NextAction `json:"nextActions"`
}
// NextAction 表示建议用户执行的下一步操作。
type NextAction struct {
	Type     string `json:"type"`
	Label    string `json:"label"`
	TargetID string `json:"targetId,omitempty"`
}
// DashboardService 处理仪表盘数据的汇总与展示。
type DashboardService struct{}

// Summary 返回仪表盘概览摘要，包含待办事项统计。
func (DashboardService) Summary() DashboardSummary {
	return DashboardSummary{NextActions: []NextAction{{Type: "import", Label: "Import material"}}}
}
