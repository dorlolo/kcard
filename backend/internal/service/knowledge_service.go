package service

import (
	"context"
	"errors"
	"strings"

	"kcardDesgin/backend/internal/domain"
)

type KnowledgeService struct {
	Points map[domain.ID]domain.KnowledgePoint
}

func (s *KnowledgeService) Search(ctx context.Context, workspaceID domain.ID, query string, status domain.ApprovalStatus) []domain.KnowledgePoint {
	var out []domain.KnowledgePoint
	for _, point := range s.Points {
		if point.LearnerWorkspaceID != workspaceID {
			continue
		}
		if status != "" && point.ApprovalStatus != status {
			continue
		}
		if query != "" && !strings.Contains(strings.ToLower(point.Content+" "+point.Summary), strings.ToLower(query)) {
			continue
		}
		out = append(out, point)
	}
	return out
}

func (s *KnowledgeService) Split(ctx context.Context, id domain.ID, contents []string) ([]domain.KnowledgePoint, error) {
	source, ok := s.Points[id]
	if !ok {
		return nil, errors.New("knowledge point not found")
	}
	if len(contents) < 2 {
		return nil, errors.New("split requires at least two points")
	}
	var out []domain.KnowledgePoint
	for idx, content := range contents {
		point := source
		point.ID = domain.ID(string(id) + ":split:" + string(rune('a'+idx)))
		point.Content = content
		point.ApprovalStatus = domain.KnowledgeDraft
		s.Points[point.ID] = point
		out = append(out, point)
	}
	return out, nil
}

func (s *KnowledgeService) Merge(ctx context.Context, ids []domain.ID, content string) (domain.KnowledgePoint, error) {
	if len(ids) < 2 {
		return domain.KnowledgePoint{}, errors.New("merge requires at least two points")
	}
	first, ok := s.Points[ids[0]]
	if !ok {
		return domain.KnowledgePoint{}, errors.New("knowledge point not found")
	}
	merged := first
	merged.ID = domain.ID(string(ids[0]) + ":merged")
	merged.Content = content
	merged.ApprovalStatus = domain.KnowledgeDraft
	s.Points[merged.ID] = merged
	return merged, nil
}
