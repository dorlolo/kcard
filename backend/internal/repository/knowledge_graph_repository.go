// Package repository 提供数据模型定义和数据库访问接口的实现。
package repository

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"kcardDesgin/backend/internal/domain"
)

// GraphFilter 表示知识图谱的筛选条件。
type GraphFilter struct {
	WorkspaceID           domain.ID
	FocusKnowledgePointID domain.ID
	Depth                 int
	RelationshipTypes     []domain.RelationshipType
	ApprovalStatus        domain.ApprovalStatus
	Query                 string
	IncludeArchived       bool
	IncludeRejected       bool
	MaxNodes              int
	MaxEdges              int
}

// KnowledgeGraphRepository 处理知识图谱数据的持久化操作。
type KnowledgeGraphRepository struct{ db *gorm.DB }

// NewKnowledgeGraphRepository 创建 KnowledgeGraphRepository 实例。
func NewKnowledgeGraphRepository(db *gorm.DB) KnowledgeGraphRepository {
	return KnowledgeGraphRepository{db: db}
}

// ListPoints 根据图谱筛选条件返回知识点列表。
func (r KnowledgeGraphRepository) ListPoints(ctx context.Context, filter GraphFilter) ([]domain.KnowledgePoint, error) {
	return NewKnowledgeRepository(r.db).Search(ctx, domain.KnowledgeFilter{WorkspaceID: filter.WorkspaceID, Query: filter.Query, ApprovalStatus: filter.ApprovalStatus, IncludeArchived: filter.IncludeArchived, IncludeRejected: filter.IncludeRejected})
}

// ListRelationships 返回指定工作区中的知识关系列表。
func (r KnowledgeGraphRepository) ListRelationships(ctx context.Context, workspaceID domain.ID, relationshipTypes []domain.RelationshipType, includeArchived bool, max int) ([]domain.KnowledgeRelationship, error) {
	var models []KnowledgeRelationshipModel
	q := r.db.WithContext(ctx).Where("learner_workspace_id = ?", string(workspaceID))
	if len(relationshipTypes) > 0 {
		types := make([]string, 0, len(relationshipTypes))
		for _, relationshipType := range relationshipTypes {
			types = append(types, string(relationshipType))
		}
		q = q.Where("relationship_type IN ?", types)
	}
	if !includeArchived {
		q = q.Where("archived_at IS NULL")
	}
	limit := max
	if limit <= 0 {
		limit = 1000
	}
	if err := q.Limit(limit).Find(&models).Error; err != nil {
		return nil, err
	}
	relationships := make([]domain.KnowledgeRelationship, 0, len(models))
	for _, model := range models {
		relationships = append(relationships, relationshipFromModel(model))
	}
	return relationships, nil
}

// CreateRelationship 创建新的知识关系记录。
func (r KnowledgeGraphRepository) CreateRelationship(ctx context.Context, relationship domain.KnowledgeRelationship) (domain.KnowledgeRelationship, error) {
	if relationship.ID == "" {
		relationship.ID = domain.ID(uuid.NewString())
	}
	model := relationshipToModel(relationship)
	if err := r.db.WithContext(ctx).Create(&model).Error; err != nil {
		return domain.KnowledgeRelationship{}, err
	}
	return relationshipFromModel(model), nil
}

// ArchiveRelationship 根据 ID 归档知识关系。
func (r KnowledgeGraphRepository) ArchiveRelationship(ctx context.Context, id domain.ID) (domain.KnowledgeRelationship, error) {
	var model KnowledgeRelationshipModel
	if err := r.db.WithContext(ctx).Where("id = ?", string(id)).First(&model).Error; err != nil {
		return domain.KnowledgeRelationship{}, err
	}
	if err := r.db.WithContext(ctx).Model(&model).Update("archived_at", gorm.Expr("now()")).Error; err != nil {
		return domain.KnowledgeRelationship{}, err
	}
	model.ArchivedAt = &model.UpdatedAt
	return relationshipFromModel(model), nil
}

func relationshipFromModel(model KnowledgeRelationshipModel) domain.KnowledgeRelationship {
	return domain.KnowledgeRelationship{ID: domain.ID(model.ID), WorkspaceID: domain.ID(model.LearnerWorkspaceID), SourceID: domain.ID(model.SourceKnowledgePointID), TargetID: domain.ID(model.TargetKnowledgePointID), Type: domain.RelationshipType(model.RelationshipType), Label: model.Label, Weight: model.Weight, SourceType: model.SourceType, Archived: model.ArchivedAt != nil}
}

func relationshipToModel(rel domain.KnowledgeRelationship) KnowledgeRelationshipModel {
	return KnowledgeRelationshipModel{BaseModel: BaseModel{ID: string(rel.ID)}, LearnerWorkspaceID: string(rel.WorkspaceID), SourceKnowledgePointID: string(rel.SourceID), TargetKnowledgePointID: string(rel.TargetID), RelationshipType: string(rel.Type), Label: rel.Label, Weight: rel.Weight, SourceType: rel.SourceType}
}
