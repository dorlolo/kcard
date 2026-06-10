package service

import (
	"errors"
	"kcardDesgin/backend/internal/domain"
)

type CardService struct{ Cards map[domain.ID]domain.Card }

func (s *CardService) Save(card domain.Card) domain.Card {
	if s.Cards == nil {
		s.Cards = map[domain.ID]domain.Card{}
	}
	s.Cards[card.ID] = card
	return card
}
func (s *CardService) Archive(id domain.ID) (domain.Card, error) {
	card, ok := s.Cards[id]
	if !ok {
		return domain.Card{}, errors.New("card not found")
	}
	card.Status = domain.CardArchived
	s.Cards[id] = card
	return card, nil
}
