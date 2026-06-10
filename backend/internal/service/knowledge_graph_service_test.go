package service

import (
	"context"
	"testing"

	"kcardDesgin/backend/internal/domain"
)

func TestKnowledgeGraphFocusIncludesNeighborhood(t *testing.T) {
	svc := KnowledgeGraphService{Points: map[domain.ID]domain.KnowledgePoint{"a": {ID: "a", LearnerWorkspaceID: "w1", Content: "A", ApprovalStatus: domain.KnowledgeApproved}, "b": {ID: "b", LearnerWorkspaceID: "w1", Content: "B", ApprovalStatus: domain.KnowledgeApproved}}, Relationships: []KnowledgeRelationship{{ID: "r1", WorkspaceID: "w1", SourceID: "a", TargetID: "b", Type: RelationshipRelated}}}
	graph, err := svc.Graph(context.Background(), "w1", "a", 1, false)
	if err != nil {
		t.Fatal(err)
	}
	if len(graph.Nodes) != 2 || len(graph.Edges) != 1 {
		t.Fatalf("nodes=%d edges=%d", len(graph.Nodes), len(graph.Edges))
	}
}

func TestKnowledgeGraphRejectsSelfRelationship(t *testing.T) {
	svc := &KnowledgeGraphService{}
	_, err := svc.AddRelationship(context.Background(), KnowledgeRelationship{WorkspaceID: "w1", SourceID: "a", TargetID: "a", Type: RelationshipRelated})
	if err == nil {
		t.Fatal("expected error")
	}
}
