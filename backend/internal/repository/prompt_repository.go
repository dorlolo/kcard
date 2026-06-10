package repository

import "gorm.io/gorm"

type PromptRepository struct{ db *gorm.DB }

func NewPromptRepository(db *gorm.DB) PromptRepository { return PromptRepository{db: db} }
