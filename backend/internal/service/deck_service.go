// Package service 提供领域服务层的实现，包括资料、卡片、知识图谱等核心业务逻辑。
package service

import (
	"errors"
	"kcardDesgin/backend/internal/domain"
)

// DeckService 处理卡组的管理操作，包括合并与还原。
type DeckService struct{ Decks map[domain.ID]domain.Deck }

// Merge 执行将多个卡组合并为一个新卡组的操作。
func (s *DeckService) Merge(id domain.ID, decks []domain.Deck, name string) (domain.Deck, error) {
	if len(decks) < 2 {
		return domain.Deck{}, errors.New("merge requires two decks")
	}
	deck := domain.Deck{ID: id, LearnerWorkspaceID: decks[0].LearnerWorkspaceID, Name: name, CreationSource: domain.CreationManual, Status: domain.CardActive}
	return deck, nil
}
// Restore 执行将卡组还原为原始卡组列表的操作。
func (s *DeckService) Restore(deck domain.Deck) []domain.Deck { return []domain.Deck{deck} }
