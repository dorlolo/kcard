package service

import (
	"errors"
	"kcardDesgin/backend/internal/domain"
)

type CardGenerationService struct{}

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
func (CardGenerationService) Approve(card domain.Card) domain.Card {
	card.Status = domain.CardActive
	return card
}
