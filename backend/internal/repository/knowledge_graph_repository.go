package repository

import (
	"context"

	"gorm.io/gorm"
)

type GraphFilter struct {
	WorkspaceID           string
	FocusKnowledgePointID string
	Depth                 int
	RelationshipTypes     []string
	ApprovalStatus        string
	IncludeArchived       bool
}

type GraphNode struct {
	ID             string
	Label          string
	NodeType       string
	ApprovalStatus string
	Weight         float64
}
type GraphEdge struct {
	ID               string
	SourceID         string
	TargetID         string
	RelationshipType string
	Label            string
	Weight           float64
	SourceType       string
}
type KnowledgeGraph struct {
	Nodes    []GraphNode
	Edges    []GraphEdge
	Warnings []string
}

type KnowledgeGraphRepository struct{ db *gorm.DB }

func NewKnowledgeGraphRepository(db *gorm.DB) KnowledgeGraphRepository {
	return KnowledgeGraphRepository{db: db}
}

func (r KnowledgeGraphRepository) Graph(ctx context.Context, filter GraphFilter) (KnowledgeGraph, error) {
	var points []KnowledgePointModel
	pointQuery := r.db.WithContext(ctx).Where("learner_workspace_id = ?", filter.WorkspaceID)
	if filter.ApprovalStatus != "" {
		pointQuery = pointQuery.Where("approval_status = ?", filter.ApprovalStatus)
	}
	if filter.FocusKnowledgePointID != "" {
		pointQuery = pointQuery.Where("id = ? OR id IN (SELECT target_knowledge_point_id FROM knowledge_relationships WHERE source_knowledge_point_id = ?)", filter.FocusKnowledgePointID, filter.FocusKnowledgePointID)
	}
	if !filter.IncludeArchived {
		pointQuery = pointQuery.Where("archived_at IS NULL")
	}
	if err := pointQuery.Limit(250).Find(&points).Error; err != nil {
		return KnowledgeGraph{}, err
	}
	var edges []KnowledgeRelationshipModel
	edgeQuery := r.db.WithContext(ctx).Where("learner_workspace_id = ?", filter.WorkspaceID)
	if len(filter.RelationshipTypes) > 0 {
		edgeQuery = edgeQuery.Where("relationship_type IN ?", filter.RelationshipTypes)
	}
	if !filter.IncludeArchived {
		edgeQuery = edgeQuery.Where("archived_at IS NULL")
	}
	if err := edgeQuery.Limit(1000).Find(&edges).Error; err != nil {
		return KnowledgeGraph{}, err
	}
	graph := KnowledgeGraph{}
	for _, point := range points {
		label := point.GraphLabel
		if label == "" {
			label = point.Summary
		}
		if label == "" {
			label = point.Content
		}
		graph.Nodes = append(graph.Nodes, GraphNode{ID: point.ID, Label: label, NodeType: "knowledge_point", ApprovalStatus: point.ApprovalStatus, Weight: 1})
	}
	for _, edge := range edges {
		graph.Edges = append(graph.Edges, GraphEdge{ID: edge.ID, SourceID: edge.SourceKnowledgePointID, TargetID: edge.TargetKnowledgePointID, RelationshipType: edge.RelationshipType, Label: edge.Label, Weight: edge.Weight, SourceType: edge.SourceType})
	}
	if len(edges) >= 1000 {
		graph.Warnings = append(graph.Warnings, "Graph was limited to 1000 edges. Narrow filters or focus a node for better readability.")
	}
	return graph, nil
}
