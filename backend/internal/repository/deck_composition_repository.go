package repository

import "gorm.io/gorm"

type DeckCompositionRepository struct{ db *gorm.DB }

func NewDeckCompositionRepository(db *gorm.DB) DeckCompositionRepository {
	return DeckCompositionRepository{db: db}
}
