package repository

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type KnowledgeSearchFilter struct {
	WorkspaceID     string
	Query           string
	ApprovalStatus  string
	TagIDs          []string
	IncludeArchived bool
}

type ExtendedKnowledgeRepository struct{ db *gorm.DB }

func NewExtendedKnowledgeRepository(db *gorm.DB) ExtendedKnowledgeRepository {
	return ExtendedKnowledgeRepository{db: db}
}

func (r ExtendedKnowledgeRepository) Search(ctx context.Context, filter KnowledgeSearchFilter) ([]KnowledgePointModel, error) {
	var points []KnowledgePointModel
	q := r.db.WithContext(ctx).Where("learner_workspace_id = ?", filter.WorkspaceID)
	if filter.Query != "" {
		q = q.Where("content ILIKE ? OR summary ILIKE ?", "%"+filter.Query+"%", "%"+filter.Query+"%")
	}
	if filter.ApprovalStatus != "" {
		q = q.Where("approval_status = ?", filter.ApprovalStatus)
	}
	if !filter.IncludeArchived {
		q = q.Where("archived_at IS NULL")
	}
	return points, q.Order("updated_at DESC").Find(&points).Error
}

func (r ExtendedKnowledgeRepository) CreateRelationship(ctx context.Context, rel *KnowledgeRelationshipModel) error {
	if rel.ID == "" {
		rel.ID = uuid.NewString()
	}
	return r.db.WithContext(ctx).Create(rel).Error
}

func (r ExtendedKnowledgeRepository) Merge(ctx context.Context, merged *KnowledgePointModel, sourceIDs []string) error {
	if merged.ID == "" {
		merged.ID = uuid.NewString()
	}
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(merged).Error; err != nil {
			return err
		}
		return tx.Model(&KnowledgePointModel{}).Where("id IN ?", sourceIDs).Update("archived_at", gorm.Expr("now()")).Error
	})
}
