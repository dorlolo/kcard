package service

import (
	"context"
	"testing"

	"kcardDesgin/backend/internal/domain"
)

func TestKnowledgeSplitPreservesSource(t *testing.T) {
	svc := &KnowledgeService{Points: map[domain.ID]domain.KnowledgePoint{"kp1": {ID: "kp1", LearnerWorkspaceID: "w1", SourceMaterialID: "m1", Content: "A and B"}}}
	points, err := svc.Split(context.Background(), "kp1", []string{"A", "B"})
	if err != nil {
		t.Fatal(err)
	}
	if len(points) != 2 {
		t.Fatalf("points=%d", len(points))
	}
	if points[0].SourceMaterialID != "m1" {
		t.Fatal("source not preserved")
	}
}

func TestKnowledgeMergeRequiresTwoPoints(t *testing.T) {
	svc := &KnowledgeService{Points: map[domain.ID]domain.KnowledgePoint{"kp1": {ID: "kp1", LearnerWorkspaceID: "w1", Content: "A"}}}
	_, err := svc.Merge(context.Background(), []domain.ID{"kp1"}, "A")
	if err == nil {
		t.Fatal("expected error")
	}
}
