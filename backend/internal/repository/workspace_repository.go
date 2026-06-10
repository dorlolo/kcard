package repository

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WorkspaceRepository struct{ db *gorm.DB }

func NewWorkspaceRepository(db *gorm.DB) WorkspaceRepository { return WorkspaceRepository{db: db} }

func (r WorkspaceRepository) EnsureDefault(ctx context.Context, ownerIdentity, displayName string) (LearnerWorkspaceModel, error) {
	var workspace LearnerWorkspaceModel
	if err := r.db.WithContext(ctx).Where("owner_identity = ?", ownerIdentity).First(&workspace).Error; err == nil {
		return workspace, nil
	}
	workspace = LearnerWorkspaceModel{BaseModel: BaseModel{ID: uuid.NewString()}, OwnerIdentity: ownerIdentity, DisplayName: displayName, PrivacyState: "private"}
	return workspace, r.db.WithContext(ctx).Create(&workspace).Error
}

func ScopeWorkspace(db *gorm.DB, workspaceID string) *gorm.DB {
	return db.Where("learner_workspace_id = ?", workspaceID)
}
