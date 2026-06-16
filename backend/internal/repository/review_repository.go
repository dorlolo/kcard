// Package repository 提供数据模型定义和数据库访问接口的实现。
package repository

import "gorm.io/gorm"

// ReviewRepository 处理复习数据的持久化操作。
type ReviewRepository struct{ db *gorm.DB }

// NewReviewRepository 创建 ReviewRepository 实例。
func NewReviewRepository(db *gorm.DB) ReviewRepository { return ReviewRepository{db: db} }
