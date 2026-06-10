package service

import (
	"kcardDesgin/backend/internal/domain"
	"testing"
)

func TestPlanCapacityRequired(t *testing.T) {
	_, err := ReviewPlanService{}.Create(domain.ReviewPlan{Name: "Plan"})
	if err == nil {
		t.Fatal("expected error")
	}
}
