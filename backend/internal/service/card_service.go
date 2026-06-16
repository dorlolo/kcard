// Package service 提供领域服务层的实现，包括资料、卡片、知识图谱等核心业务逻辑。
package service

import (
	"errors"
	"kcardDesgin/backend/internal/domain"
)

// CardService 处理卡片的保存和归档操作。
type CardService struct{ Cards map[domain.ID]domain.Card }

// Save 保存卡片到内存存储，若已存在则覆盖。
func (s *CardService) Save(card domain.Card) domain.Card {
	if s.Cards == nil {
		s.Cards = map[domain.ID]domain.Card{}
	}
	s.Cards[card.ID] = card
	return card
}
// Archive 根据卡片 ID 将卡片标记为已归档状态。
func (s *CardService) Archive(id domain.ID) (domain.Card, error) {
	card, ok := s.Cards[id]
	if !ok {
		return domain.Card{}, errors.New("card not found")
	}
	card.Status = domain.CardArchived
	s.Cards[id] = card
	return card, nil
}
