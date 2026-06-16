// Package repository 提供数据模型定义和数据库访问接口的实现。
package repository

import (
	"context"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"kcardDesgin/backend/internal/domain"
)

// KnowledgeRepository 处理知识点数据的持久化操作。
type KnowledgeRepository struct{ db *gorm.DB }

// NewKnowledgeRepository 创建 KnowledgeRepository 实例。
func NewKnowledgeRepository(db *gorm.DB) KnowledgeRepository {
	return KnowledgeRepository{db: db}
}

// Search 根据筛选条件搜索知识点列表。
func (r KnowledgeRepository) Search(ctx context.Context, filter domain.KnowledgeFilter) ([]domain.KnowledgePoint, error) {
	var models []KnowledgePointModel
	q := r.db.WithContext(ctx).Where("learner_workspace_id = ?", string(filter.WorkspaceID))
	if filter.Query != "" {
		like := "%" + strings.TrimSpace(filter.Query) + "%"
		q = q.Where("content ILIKE ? OR summary ILIKE ? OR notes ILIKE ?", like, like, like)
	}
	if filter.ApprovalStatus != "" {
		q = q.Where("approval_status = ?", string(filter.ApprovalStatus))
	}
	if !filter.IncludeRejected {
		q = q.Where("approval_status <> ?", string(domain.KnowledgeRejected))
	}
	if !filter.IncludeArchived {
		q = q.Where("archived_at IS NULL")
	}
	if err := q.Order("updated_at DESC").Find(&models).Error; err != nil {
		return nil, err
	}
	points := make([]domain.KnowledgePoint, 0, len(models))
	for _, model := range models {
		points = append(points, knowledgePointFromModel(model))
	}
	return points, nil
}

// Get 根据 ID 获取指定工作区中的知识点。
func (r KnowledgeRepository) Get(ctx context.Context, workspaceID domain.ID, id domain.ID) (domain.KnowledgePoint, error) {
	var model KnowledgePointModel
	err := r.db.WithContext(ctx).Where("learner_workspace_id = ? AND id = ?", string(workspaceID), string(id)).First(&model).Error
	if err != nil {
		return domain.KnowledgePoint{}, err
	}
	return knowledgePointFromModel(model), nil
}

// Create 创建新的知识点记录。
func (r KnowledgeRepository) Create(ctx context.Context, point domain.KnowledgePoint) (domain.KnowledgePoint, error) {
	if point.ID == "" {
		point.ID = domain.ID(uuid.NewString())
	}
	model := knowledgePointToModel(point)
	if err := r.db.WithContext(ctx).Create(&model).Error; err != nil {
		return domain.KnowledgePoint{}, err
	}
	return knowledgePointFromModel(model), nil
}

// Save 保存知识点记录（创建或更新）。
func (r KnowledgeRepository) Save(ctx context.Context, point domain.KnowledgePoint) (domain.KnowledgePoint, error) {
	model := knowledgePointToModel(point)
	if err := r.db.WithContext(ctx).Save(&model).Error; err != nil {
		return domain.KnowledgePoint{}, err
	}
	return knowledgePointFromModel(model), nil
}

// UpdateStatus 更新知识点的审批状态和备注。
func (r KnowledgeRepository) UpdateStatus(ctx context.Context, workspaceID domain.ID, id domain.ID, status domain.ApprovalStatus, notes string, now time.Time) (domain.KnowledgePoint, error) {
	point, err := r.Get(ctx, workspaceID, id)
	if err != nil {
		return domain.KnowledgePoint{}, err
	}
	point.ApprovalStatus = status
	if notes != "" {
		point.Notes = notes
	}
	switch status {
	case domain.KnowledgeApproved:
		point = point.Approve(now)
	case domain.KnowledgeRejected:
		point = point.Reject(now)
	case domain.KnowledgeNeedsReview:
		point = point.MarkNeedsReview()
	}
	return r.Save(ctx, point)
}

func knowledgePointFromModel(model KnowledgePointModel) domain.KnowledgePoint {
	return domain.KnowledgePoint{
		ID:                 domain.ID(model.ID),
		LearnerWorkspaceID: domain.ID(model.LearnerWorkspaceID),
		SourceMaterialID:   domain.ID(stringValue(model.SourceMaterialID)),
		MaterialVersionID:  domain.ID(stringValue(model.MaterialVersionID)),
		Content:            model.Content,
		Summary:            model.Summary,
		Notes:              model.Notes,
		ApprovalStatus:     domain.ApprovalStatus(model.ApprovalStatus),
		CreationSource:     domain.CreationSource(model.CreationSource),
		DuplicateGroupID:   domain.ID(stringValue(model.DuplicateGroupID)),
		GraphLabel:         model.GraphLabel,
		ApprovedAt:         model.ApprovedAt,
		RejectedAt:         model.RejectedAt,
	}
}

func knowledgePointToModel(point domain.KnowledgePoint) KnowledgePointModel {
	return KnowledgePointModel{
		BaseModel:          BaseModel{ID: string(point.ID)},
		LearnerWorkspaceID: string(point.LearnerWorkspaceID),
		SourceMaterialID:   optionalString(point.SourceMaterialID),
		MaterialVersionID:  optionalString(point.MaterialVersionID),
		Content:            point.Content,
		Summary:            point.Summary,
		Notes:              point.Notes,
		ApprovalStatus:     string(point.ApprovalStatus),
		CreationSource:     string(point.CreationSource),
		DuplicateGroupID:   optionalString(point.DuplicateGroupID),
		GraphLabel:         point.GraphLabel,
		ApprovedAt:         point.ApprovedAt,
		RejectedAt:         point.RejectedAt,
	}
}

func optionalString(id domain.ID) *string {
	if id == "" {
		return nil
	}
	value := string(id)
	return &value
}

func stringValue(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}
