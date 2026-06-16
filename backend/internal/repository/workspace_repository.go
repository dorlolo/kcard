// Package repository 提供数据模型定义和数据库访问接口的实现。
package repository

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// WorkspaceRepository 处理工作区数据的持久化操作。
type WorkspaceRepository struct{ db *gorm.DB }

// NewWorkspaceRepository 创建 WorkspaceRepository 实例。
func NewWorkspaceRepository(db *gorm.DB) WorkspaceRepository { return WorkspaceRepository{db: db} }

// EnsureDefault 确保指定所有者存在默认工作区，不存在则创建。
func (r WorkspaceRepository) EnsureDefault(ctx context.Context, ownerIdentity, displayName string) (LearnerWorkspaceModel, error) {
	var workspace LearnerWorkspaceModel
	if err := r.db.WithContext(ctx).Where("owner_identity = ?", ownerIdentity).First(&workspace).Error; err == nil {
		return workspace, nil
	}
	workspace = LearnerWorkspaceModel{BaseModel: BaseModel{ID: uuid.NewString()}, OwnerIdentity: ownerIdentity, DisplayName: displayName, PrivacyState: "private"}
	return workspace, r.db.WithContext(ctx).Create(&workspace).Error
}

// ScopeWorkspace 返回按工作区 ID 筛选的 GORM 查询范围。
func ScopeWorkspace(db *gorm.DB, workspaceID string) *gorm.DB {
	return db.Where("learner_workspace_id = ?", workspaceID)
}
