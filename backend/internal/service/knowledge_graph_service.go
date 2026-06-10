package service

import (
	"context"
	"errors"

	"kcardDesgin/backend/internal/domain"
)

type RelationshipType string

const (
	RelationshipRelated      RelationshipType = "related"
	RelationshipPrerequisite RelationshipType = "prerequisite"
	RelationshipDuplicate    RelationshipType = "duplicate"
	RelationshipSplitFrom    RelationshipType = "split_from"
	RelationshipMergedFrom   RelationshipType = "merged_from"
)

type KnowledgeGraphNode struct {
	ID             domain.ID             `json:"id"`
	Label          string                `json:"label"`
	NodeType       string                `json:"nodeType"`
	ApprovalStatus domain.ApprovalStatus `json:"status"`
}
type KnowledgeGraphEdge struct {
	ID               domain.ID        `json:"id"`
	SourceID         domain.ID        `json:"sourceId"`
	TargetID         domain.ID        `json:"targetId"`
	RelationshipType RelationshipType `json:"relationshipType"`
	Label            string           `json:"label"`
	Weight           float64          `json:"weight"`
}
type KnowledgeGraphResult struct {
	Nodes    []KnowledgeGraphNode `json:"nodes"`
	Edges    []KnowledgeGraphEdge `json:"edges"`
	Warnings []string             `json:"warnings"`
}
type KnowledgeRelationship struct {
	ID          domain.ID
	WorkspaceID domain.ID
	SourceID    domain.ID
	TargetID    domain.ID
	Type        RelationshipType
	Label       string
	Weight      float64
	Archived    bool
}

type KnowledgeGraphService struct {
	Points        map[domain.ID]domain.KnowledgePoint
	Relationships []KnowledgeRelationship
}

func (s KnowledgeGraphService) Graph(ctx context.Context, workspaceID domain.ID, focus domain.ID, depth int, includeArchived bool) (KnowledgeGraphResult, error) {
	if depth < 1 {
		depth = 1
	}
	if depth > 3 {
		depth = 3
	}
	included := map[domain.ID]bool{}
	if focus != "" {
		included[focus] = true
	}
	for i := 0; i < depth; i++ {
		for _, rel := range s.Relationships {
			if rel.WorkspaceID != workspaceID || (!includeArchived && rel.Archived) {
				continue
			}
			if focus == "" || included[rel.SourceID] || included[rel.TargetID] {
				included[rel.SourceID] = true
				included[rel.TargetID] = true
			}
		}
	}
	result := KnowledgeGraphResult{}
	for id, point := range s.Points {
		if point.LearnerWorkspaceID != workspaceID {
			continue
		}
		if focus != "" && !included[id] {
			continue
		}
		if !includeArchived && point.ApprovalStatus == domain.KnowledgeRejected {
			continue
		}
		label := point.GraphLabel
		if label == "" {
			label = point.Summary
		}
		if label == "" {
			label = point.Content
		}
		result.Nodes = append(result.Nodes, KnowledgeGraphNode{ID: id, Label: label, NodeType: "knowledge_point", ApprovalStatus: point.ApprovalStatus})
	}
	for _, rel := range s.Relationships {
		if rel.WorkspaceID != workspaceID || (!includeArchived && rel.Archived) {
			continue
		}
		if focus != "" && !(included[rel.SourceID] && included[rel.TargetID]) {
			continue
		}
		result.Edges = append(result.Edges, KnowledgeGraphEdge{ID: rel.ID, SourceID: rel.SourceID, TargetID: rel.TargetID, RelationshipType: rel.Type, Label: rel.Label, Weight: rel.Weight})
	}
	if len(result.Edges) > 10000 {
		result.Warnings = append(result.Warnings, "Dense graph truncated; use filters or focus a node.")
	}
	return result, nil
}

func (s *KnowledgeGraphService) AddRelationship(ctx context.Context, rel KnowledgeRelationship) (KnowledgeRelationship, error) {
	if rel.WorkspaceID == "" || rel.SourceID == "" || rel.TargetID == "" {
		return KnowledgeRelationship{}, errors.New("workspace, source, and target are required")
	}
	if rel.SourceID == rel.TargetID {
		return KnowledgeRelationship{}, errors.New("relationship cannot target the same point")
	}
	if rel.Weight == 0 {
		rel.Weight = 1
	}
	s.Relationships = append(s.Relationships, rel)
	return rel, nil
}
