package service

import "kcardDesgin/backend/internal/domain"

type StatisticsService struct{}

func (StatisticsService) Summarize(results []domain.ReviewResult) domain.ReviewStatisticsSnapshot {
	return domain.ReviewStatisticsSnapshot{ScopeType: "workspace", CardsReviewed: len(results), RecallRate: 1}
}
