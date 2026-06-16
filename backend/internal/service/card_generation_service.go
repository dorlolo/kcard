// Package service 提供领域服务层的实现，包括资料、卡片、知识图谱等核心业务逻辑。
package service

import (
	"errors"
	"kcardDesgin/backend/internal/domain"
)

// CardGenerationService 处理从知识点生成卡片草稿和审批操作。
type CardGenerationService struct{}

// GenerateDrafts 根据已审批的知识点列表生成卡片草稿。
func (CardGenerationService) GenerateDrafts(points []domain.KnowledgePoint) ([]domain.Card, error) {
	approved := ApprovedKnowledgeOnly(points, false)
	if len(approved) == 0 {
		return nil, errors.New("no approved knowledge points")
	}
	cards := []domain.Card{}
	for _, p := range approved {
		cards = append(cards, domain.Card{ID: p.ID + ":card", LearnerWorkspaceID: p.LearnerWorkspaceID, FrontPrompt: p.Summary, BackAnswer: p.Content, Status: domain.CardDraft, ReviewStatus: "new"})
	}
	return cards, nil
}
// Approve 将卡片草稿状态更新为已激活。
func (CardGenerationService) Approve(card domain.Card) domain.Card {
	card.Status = domain.CardActive
	return card
}
