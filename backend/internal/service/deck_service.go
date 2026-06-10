package service

import (
	"errors"
	"kcardDesgin/backend/internal/domain"
)

type DeckService struct{ Decks map[domain.ID]domain.Deck }

func (s *DeckService) Merge(id domain.ID, decks []domain.Deck, name string) (domain.Deck, error) {
	if len(decks) < 2 {
		return domain.Deck{}, errors.New("merge requires two decks")
	}
	deck := domain.Deck{ID: id, LearnerWorkspaceID: decks[0].LearnerWorkspaceID, Name: name, CreationSource: domain.CreationManual, Status: domain.CardActive}
	return deck, nil
}
func (s *DeckService) Restore(deck domain.Deck) []domain.Deck { return []domain.Deck{deck} }
