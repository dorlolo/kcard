package service

import (
	"context"
	"testing"
	"time"

	"kcardDesgin/backend/internal/domain"
)

func TestKnowledgeSearchFiltersByQueryStatusAndTag(t *testing.T) {
	svc := NewKnowledgeService([]domain.KnowledgePoint{
		{ID: "kp1", LearnerWorkspaceID: "w1", Content: "细胞膜控制物质进出", ApprovalStatus: domain.KnowledgeApproved, Tags: []domain.Tag{{ID: "bio", Name: "生物"}}},
		{ID: "kp2", LearnerWorkspaceID: "w1", Content: "历史事件时间线", ApprovalStatus: domain.KnowledgeDraft, Tags: []domain.Tag{{ID: "history", Name: "历史"}}},
	})

	points := svc.Search(context.Background(), domain.KnowledgeFilter{WorkspaceID: "w1", Query: "细胞", ApprovalStatus: domain.KnowledgeApproved, Tag: "生物"})
	if len(points) != 1 || points[0].ID != "kp1" {
		t.Fatalf("unexpected search results: %#v", points)
	}
}

func TestKnowledgeSplitPreservesSourceAndMarksOriginalNeedsReview(t *testing.T) {
	svc := NewKnowledgeService([]domain.KnowledgePoint{{ID: "kp1", LearnerWorkspaceID: "w1", SourceMaterialID: "m1", Content: "A and B", ApprovalStatus: domain.KnowledgeApproved}})
	points, err := svc.Split(context.Background(), "w1", "kp1", []string{"A", "B"})
	if err != nil {
		t.Fatal(err)
	}
	if len(points) != 2 {
		t.Fatalf("points=%d", len(points))
	}
	if points[0].SourceMaterialID != "m1" {
		t.Fatal("source not preserved")
	}
	original, err := svc.Store.Get(context.Background(), "w1", "kp1")
	if err != nil {
		t.Fatal(err)
	}
	if original.ApprovalStatus != domain.KnowledgeNeedsReview {
		t.Fatalf("original status=%s", original.ApprovalStatus)
	}
}

func TestKnowledgeMergePreservesWorkspaceAndMarksSourcesNeedsReview(t *testing.T) {
	svc := NewKnowledgeService([]domain.KnowledgePoint{
		{ID: "kp1", LearnerWorkspaceID: "w1", Content: "A", ApprovalStatus: domain.KnowledgeApproved},
		{ID: "kp2", LearnerWorkspaceID: "w1", Content: "B", ApprovalStatus: domain.KnowledgeApproved},
	})
	merged, err := svc.Merge(context.Background(), "w1", []domain.ID{"kp1", "kp2"}, "A and B")
	if err != nil {
		t.Fatal(err)
	}
	if merged.LearnerWorkspaceID != "w1" || merged.ApprovalStatus != domain.KnowledgeDraft {
		t.Fatalf("merged=%#v", merged)
	}
	first, _ := svc.Store.Get(context.Background(), "w1", "kp1")
	second, _ := svc.Store.Get(context.Background(), "w1", "kp2")
	if first.ApprovalStatus != domain.KnowledgeNeedsReview || second.ApprovalStatus != domain.KnowledgeNeedsReview {
		t.Fatal("source points should need review after merge")
	}
}

func TestKnowledgeUpdateStatusSetsApprovalTimestamps(t *testing.T) {
	now := time.Date(2026, 6, 10, 12, 0, 0, 0, time.UTC)
	svc := NewKnowledgeService([]domain.KnowledgePoint{{ID: "kp1", LearnerWorkspaceID: "w1", Content: "A", ApprovalStatus: domain.KnowledgeDraft}})
	point, err := svc.UpdateStatus(context.Background(), "w1", "kp1", domain.KnowledgeApproved, "ok", now)
	if err != nil {
		t.Fatal(err)
	}
	if point.ApprovedAt == nil || !point.ApprovedAt.Equal(now) {
		t.Fatalf("approvedAt=%v", point.ApprovedAt)
	}
}

func TestKnowledgeMergeRequiresTwoPoints(t *testing.T) {
	svc := NewKnowledgeService([]domain.KnowledgePoint{{ID: "kp1", LearnerWorkspaceID: "w1", Content: "A"}})
	_, err := svc.Merge(context.Background(), "w1", []domain.ID{"kp1"}, "A")
	if err == nil {
		t.Fatal("expected error")
	}
}
