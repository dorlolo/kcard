// Package repository 提供数据模型定义和数据库访问接口的实现。
package repository

import "gorm.io/gorm"

// PromptRepository 处理提示词数据的持久化操作。
type PromptRepository struct{ db *gorm.DB }

// NewPromptRepository 创建 PromptRepository 实例。
func NewPromptRepository(db *gorm.DB) PromptRepository { return PromptRepository{db: db} }
