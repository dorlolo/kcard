package service

import (
	"context"
	"errors"
	"sort"
	"strings"

	"kcardDesgin/backend/internal/domain"
)

type RelationshipType = domain.RelationshipType

const (
	RelationshipRelated      = domain.RelationshipRelated
	RelationshipPrerequisite = domain.RelationshipPrerequisite
	RelationshipDuplicate    = domain.RelationshipDuplicate
	RelationshipSimilar      = domain.RelationshipSimilar
	RelationshipSplitFrom    = domain.RelationshipSplitFrom
	RelationshipMergedFrom   = domain.RelationshipMergedFrom
	RelationshipSupports     = domain.RelationshipSupports
	RelationshipContradicts  = domain.RelationshipContradicts
)

type KnowledgeRelationship = domain.KnowledgeRelationship

type KnowledgeGraphNode struct {
	ID             domain.ID             `json:"id"`
	Label          string                `json:"label"`
	NodeType       string                `json:"nodeType"`
	ApprovalStatus domain.ApprovalStatus `json:"status"`
	Weight         float64               `json:"weight"`
}

type KnowledgeGraphEdge struct {
	ID               domain.ID        `json:"id"`
	SourceID         domain.ID        `json:"sourceId"`
	TargetID         domain.ID        `json:"targetId"`
	RelationshipType RelationshipType `json:"relationshipType"`
	Label            string           `json:"label"`
	Weight           float64          `json:"weight"`
	SourceType       string           `json:"sourceType"`
}

type KnowledgeGraphResult struct {
	Nodes    []KnowledgeGraphNode `json:"nodes"`
	Edges    []KnowledgeGraphEdge `json:"edges"`
	Warnings []string             `json:"warnings"`
}

type GraphQuery struct {
	WorkspaceID       domain.ID
	FocusID           domain.ID
	Depth             int
	Query             string
	ApprovalStatus    domain.ApprovalStatus
	RelationshipTypes []RelationshipType
	IncludeArchived   bool
	IncludeRejected   bool
	MaxNodes          int
	MaxEdges          int
}

type KnowledgeRelationshipStore interface {
	ListRelationships(ctx context.Context, workspaceID domain.ID, relationshipTypes []domain.RelationshipType, includeArchived bool, max int) ([]domain.KnowledgeRelationship, error)
	CreateRelationship(ctx context.Context, relationship domain.KnowledgeRelationship) (domain.KnowledgeRelationship, error)
	ArchiveRelationship(ctx context.Context, id domain.ID) (domain.KnowledgeRelationship, error)
}

type KnowledgeGraphService struct {
	PointStore        KnowledgeStore
	RelationshipStore KnowledgeRelationshipStore
}

func NewKnowledgeGraphService(pointStore KnowledgeStore, relationshipStore KnowledgeRelationshipStore) *KnowledgeGraphService {
	return &KnowledgeGraphService{PointStore: pointStore, RelationshipStore: relationshipStore}
}

func (s KnowledgeGraphService) Graph(ctx context.Context, query GraphQuery) (KnowledgeGraphResult, error) {
	if query.WorkspaceID == "" {
		return KnowledgeGraphResult{}, errors.New("workspace is required")
	}
	if query.Depth < 1 {
		query.Depth = 1
	}
	if query.Depth > 3 {
		query.Depth = 3
	}
	if query.MaxNodes <= 0 {
		query.MaxNodes = 250
	}
	if query.MaxEdges <= 0 {
		query.MaxEdges = 1000
	}
	if s.PointStore == nil || s.RelationshipStore == nil {
		return KnowledgeGraphResult{}, errors.New("knowledge graph stores are not configured")
	}

	relationships, err := s.RelationshipStore.ListRelationships(ctx, query.WorkspaceID, query.RelationshipTypes, query.IncludeArchived, query.MaxEdges)
	if err != nil {
		return KnowledgeGraphResult{}, err
	}
	included := includedIDs(query, relationships)
	points, err := s.PointStore.Search(ctx, domain.KnowledgeFilter{WorkspaceID: query.WorkspaceID, Query: query.Query, ApprovalStatus: query.ApprovalStatus, IncludeRejected: query.IncludeRejected, IncludeArchived: query.IncludeArchived})
	if err != nil {
		return KnowledgeGraphResult{}, err
	}

	result := KnowledgeGraphResult{}
	lowerQuery := strings.ToLower(strings.TrimSpace(query.Query))
	for _, point := range points {
		if query.FocusID != "" && !included[point.ID] {
			continue
		}
		if lowerQuery != "" && !strings.Contains(strings.ToLower(point.Content+" "+point.Summary+" "+point.Notes), lowerQuery) {
			continue
		}
		result.Nodes = append(result.Nodes, KnowledgeGraphNode{ID: point.ID, Label: graphLabel(point), NodeType: "knowledge_point", ApprovalStatus: point.ApprovalStatus, Weight: 1})
	}
	sort.Slice(result.Nodes, func(i, j int) bool { return result.Nodes[i].ID < result.Nodes[j].ID })
	if len(result.Nodes) > query.MaxNodes {
		result.Nodes = result.Nodes[:query.MaxNodes]
		result.Warnings = append(result.Warnings, "节点过多，已截断显示，请使用搜索或聚焦节点。")
	}

	nodeSet := map[domain.ID]bool{}
	for _, node := range result.Nodes {
		nodeSet[node.ID] = true
	}
	for _, rel := range relationships {
		if !nodeSet[rel.SourceID] || !nodeSet[rel.TargetID] {
			continue
		}
		result.Edges = append(result.Edges, KnowledgeGraphEdge{ID: rel.ID, SourceID: rel.SourceID, TargetID: rel.TargetID, RelationshipType: rel.Type, Label: rel.Label, Weight: relationshipWeight(rel), SourceType: sourceType(rel)})
	}
	if len(result.Edges) > query.MaxEdges {
		result.Edges = result.Edges[:query.MaxEdges]
		result.Warnings = append(result.Warnings, "关系过多，已截断显示，请缩小关系类型或聚焦节点。")
	}
	return result, nil
}

func includedIDs(query GraphQuery, relationships []domain.KnowledgeRelationship) map[domain.ID]bool {
	included := map[domain.ID]bool{}
	if query.FocusID == "" {
		return included
	}
	included[query.FocusID] = true
	for i := 0; i < query.Depth; i++ {
		frontier := map[domain.ID]bool{}
		for id := range included {
			frontier[id] = true
		}
		changed := false
		for _, rel := range relationships {
			if rel.WorkspaceID != query.WorkspaceID || (!query.IncludeArchived && rel.Archived) {
				continue
			}
			if frontier[rel.SourceID] && !included[rel.TargetID] {
				included[rel.TargetID] = true
				changed = true
			}
			if frontier[rel.TargetID] && !included[rel.SourceID] {
				included[rel.SourceID] = true
				changed = true
			}
		}
		if !changed {
			break
		}
	}
	return included
}

func (s *KnowledgeGraphService) AddRelationship(ctx context.Context, rel KnowledgeRelationship) (KnowledgeRelationship, error) {
	if rel.WorkspaceID == "" || rel.SourceID == "" || rel.TargetID == "" {
		return KnowledgeRelationship{}, errors.New("workspace, source, and target are required")
	}
	if rel.SourceID == rel.TargetID {
		return KnowledgeRelationship{}, errors.New("relationship cannot target the same point")
	}
	if s.PointStore == nil || s.RelationshipStore == nil {
		return KnowledgeRelationship{}, errors.New("knowledge graph stores are not configured")
	}
	if _, err := s.PointStore.Get(ctx, rel.WorkspaceID, rel.SourceID); err != nil {
		return KnowledgeRelationship{}, errors.New("source knowledge point not found")
	}
	if _, err := s.PointStore.Get(ctx, rel.WorkspaceID, rel.TargetID); err != nil {
		return KnowledgeRelationship{}, errors.New("target knowledge point not found")
	}
	if rel.Type == "" {
		rel.Type = RelationshipRelated
	}
	if rel.Weight == 0 {
		rel.Weight = 1
	}
	if rel.SourceType == "" {
		rel.SourceType = "learner_created"
	}
	return s.RelationshipStore.CreateRelationship(ctx, rel)
}

func (s *KnowledgeGraphService) ArchiveRelationship(ctx context.Context, id domain.ID) (KnowledgeRelationship, error) {
	if s.RelationshipStore == nil {
		return KnowledgeRelationship{}, errors.New("knowledge graph store is not configured")
	}
	return s.RelationshipStore.ArchiveRelationship(ctx, id)
}

func graphLabel(point domain.KnowledgePoint) string {
	if point.GraphLabel != "" {
		return point.GraphLabel
	}
	if point.Summary != "" {
		return point.Summary
	}
	return summarize(point.Content)
}

func relationshipWeight(rel KnowledgeRelationship) float64 {
	if rel.Weight == 0 {
		return 1
	}
	return rel.Weight
}

func sourceType(rel KnowledgeRelationship) string {
	if rel.SourceType == "" {
		return "system_derived"
	}
	return rel.SourceType
}

type MemoryKnowledgeRelationshipStore struct {
	Relationships []domain.KnowledgeRelationship
}

func NewMemoryKnowledgeRelationshipStore(relationships []domain.KnowledgeRelationship) *MemoryKnowledgeRelationshipStore {
	return &MemoryKnowledgeRelationshipStore{Relationships: relationships}
}

func (s *MemoryKnowledgeRelationshipStore) ListRelationships(ctx context.Context, workspaceID domain.ID, relationshipTypes []domain.RelationshipType, includeArchived bool, max int) ([]domain.KnowledgeRelationship, error) {
	typeSet := map[domain.RelationshipType]bool{}
	for _, relationshipType := range relationshipTypes {
		typeSet[relationshipType] = true
	}
	out := []domain.KnowledgeRelationship{}
	for _, rel := range s.Relationships {
		if rel.WorkspaceID != workspaceID || (!includeArchived && rel.Archived) {
			continue
		}
		if len(typeSet) > 0 && !typeSet[rel.Type] {
			continue
		}
		out = append(out, rel)
		if max > 0 && len(out) >= max {
			break
		}
	}
	return out, nil
}

func (s *MemoryKnowledgeRelationshipStore) CreateRelationship(ctx context.Context, relationship domain.KnowledgeRelationship) (domain.KnowledgeRelationship, error) {
	if relationship.ID == "" {
		relationship.ID = domain.ID(string(relationship.SourceID) + ":" + string(relationship.TargetID) + ":" + string(relationship.Type))
	}
	s.Relationships = append(s.Relationships, relationship)
	return relationship, nil
}

func (s *MemoryKnowledgeRelationshipStore) ArchiveRelationship(ctx context.Context, id domain.ID) (domain.KnowledgeRelationship, error) {
	for index, rel := range s.Relationships {
		if rel.ID == id {
			rel.Archived = true
			s.Relationships[index] = rel
			return rel, nil
		}
	}
	return domain.KnowledgeRelationship{}, errors.New("relationship not found")
}
