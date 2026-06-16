// Package repository 提供数据模型定义和数据库访问接口的实现。
package repository

import "gorm.io/gorm"

// CardRepository 处理卡片数据的持久化操作。
type CardRepository struct{ db *gorm.DB }

// NewCardRepository 创建 CardRepository 实例。
func NewCardRepository(db *gorm.DB) CardRepository { return CardRepository{db: db} }
