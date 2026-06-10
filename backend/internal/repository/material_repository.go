package repository

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MaterialRepository struct{ db *gorm.DB }

func NewMaterialRepository(db *gorm.DB) MaterialRepository { return MaterialRepository{db: db} }

func (r MaterialRepository) Create(ctx context.Context, material *SourceMaterialModel) error {
	if material.ID == "" {
		material.ID = uuid.NewString()
	}
	return r.db.WithContext(ctx).Create(material).Error
}

func (r MaterialRepository) FindByDigest(ctx context.Context, workspaceID, digest string) ([]SourceMaterialModel, error) {
	var items []SourceMaterialModel
	err := r.db.WithContext(ctx).Where("learner_workspace_id = ? AND content_digest = ? AND archived_at IS NULL", workspaceID, digest).Find(&items).Error
	return items, err
}

func (r MaterialRepository) Get(ctx context.Context, workspaceID, id string) (SourceMaterialModel, error) {
	var material SourceMaterialModel
	err := r.db.WithContext(ctx).Where("learner_workspace_id = ? AND id = ?", workspaceID, id).First(&material).Error
	return material, err
}
