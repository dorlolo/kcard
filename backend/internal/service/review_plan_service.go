package service

import (
	"errors"
	"kcardDesgin/backend/internal/domain"
)

type ReviewPlanService struct{}

func (ReviewPlanService) Create(plan domain.ReviewPlan) (domain.ReviewPlan, error) {
	if plan.DailyCapacity <= 0 {
		return domain.ReviewPlan{}, errors.New("daily capacity must be positive")
	}
	plan.Status = "active"
	return plan, nil
}
func (ReviewPlanService) RestoreRevision(plan domain.ReviewPlan, revision domain.PlanRevision) domain.ReviewPlan {
	plan.CurrentRevisionID = revision.ID
	return plan
}
