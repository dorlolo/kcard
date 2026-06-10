package service

import (
	"context"
	"testing"

	"kcardDesgin/backend/internal/domain"
)

func graphService(points []domain.KnowledgePoint, relationships []domain.KnowledgeRelationship) *KnowledgeGraphService {
	pointStore := NewMemoryKnowledgeStore(points)
	relationshipStore := NewMemoryKnowledgeRelationshipStore(relationships)
	return NewKnowledgeGraphService(pointStore, relationshipStore)
}

func TestKnowledgeGraphFocusIncludesNeighborhood(t *testing.T) {
	svc := graphService(
		[]domain.KnowledgePoint{{ID: "a", LearnerWorkspaceID: "w1", Content: "A", ApprovalStatus: domain.KnowledgeApproved}, {ID: "b", LearnerWorkspaceID: "w1", Content: "B", ApprovalStatus: domain.KnowledgeApproved}, {ID: "c", LearnerWorkspaceID: "w1", Content: "C", ApprovalStatus: domain.KnowledgeApproved}},
		[]domain.KnowledgeRelationship{{ID: "r1", WorkspaceID: "w1", SourceID: "a", TargetID: "b", Type: RelationshipRelated}, {ID: "r2", WorkspaceID: "w1", SourceID: "b", TargetID: "c", Type: RelationshipPrerequisite}},
	)
	graph, err := svc.Graph(context.Background(), GraphQuery{WorkspaceID: "w1", FocusID: "a", Depth: 1})
	if err != nil {
		t.Fatal(err)
	}
	if len(graph.Nodes) != 2 || len(graph.Edges) != 1 {
		t.Fatalf("nodes=%d edges=%d", len(graph.Nodes), len(graph.Edges))
	}
}

func TestKnowledgeGraphFiltersRelationshipType(t *testing.T) {
	svc := graphService(
		[]domain.KnowledgePoint{{ID: "a", LearnerWorkspaceID: "w1", Content: "A", ApprovalStatus: domain.KnowledgeApproved}, {ID: "b", LearnerWorkspaceID: "w1", Content: "B", ApprovalStatus: domain.KnowledgeApproved}, {ID: "c", LearnerWorkspaceID: "w1", Content: "C", ApprovalStatus: domain.KnowledgeApproved}},
		[]domain.KnowledgeRelationship{{ID: "r1", WorkspaceID: "w1", SourceID: "a", TargetID: "b", Type: RelationshipRelated}, {ID: "r2", WorkspaceID: "w1", SourceID: "a", TargetID: "c", Type: RelationshipPrerequisite}},
	)
	graph, err := svc.Graph(context.Background(), GraphQuery{WorkspaceID: "w1", RelationshipTypes: []RelationshipType{RelationshipPrerequisite}})
	if err != nil {
		t.Fatal(err)
	}
	if len(graph.Edges) != 1 || graph.Edges[0].RelationshipType != RelationshipPrerequisite {
		t.Fatalf("edges=%#v", graph.Edges)
	}
}

func TestKnowledgeGraphHidesRejectedByDefault(t *testing.T) {
	svc := graphService([]domain.KnowledgePoint{{ID: "a", LearnerWorkspaceID: "w1", Content: "A", ApprovalStatus: domain.KnowledgeRejected}}, nil)
	graph, err := svc.Graph(context.Background(), GraphQuery{WorkspaceID: "w1"})
	if err != nil {
		t.Fatal(err)
	}
	if len(graph.Nodes) != 0 {
		t.Fatalf("rejected nodes should be hidden by default")
	}
}

func TestKnowledgeGraphRejectsSelfRelationship(t *testing.T) {
	svc := graphService([]domain.KnowledgePoint{{ID: "a", LearnerWorkspaceID: "w1"}}, nil)
	_, err := svc.AddRelationship(context.Background(), domain.KnowledgeRelationship{WorkspaceID: "w1", SourceID: "a", TargetID: "a", Type: RelationshipRelated})
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestKnowledgeGraphArchivesRelationship(t *testing.T) {
	svc := graphService(nil, []domain.KnowledgeRelationship{{ID: "r1", WorkspaceID: "w1", SourceID: "a", TargetID: "b", Type: RelationshipRelated}})
	rel, err := svc.ArchiveRelationship(context.Background(), "r1")
	if err != nil {
		t.Fatal(err)
	}
	if !rel.Archived {
		t.Fatal("relationship should be archived")
	}
}
