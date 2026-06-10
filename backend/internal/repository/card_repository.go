package repository

import "gorm.io/gorm"

type CardRepository struct{ db *gorm.DB }

func NewCardRepository(db *gorm.DB) CardRepository { return CardRepository{db: db} }
