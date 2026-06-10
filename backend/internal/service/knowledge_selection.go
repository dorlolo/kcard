package service

import "kcardDesgin/backend/internal/domain"

func ApprovedKnowledgeOnly(points []domain.KnowledgePoint, includeUnapproved bool) []domain.KnowledgePoint {
	if includeUnapproved {
		return points
	}
	out := []domain.KnowledgePoint{}
	for _, point := range points {
		if point.ApprovalStatus == domain.KnowledgeApproved {
			out = append(out, point)
		}
	}
	return out
}
