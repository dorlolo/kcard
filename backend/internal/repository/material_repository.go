// Package repository 提供数据模型定义和数据库访问接口的实现。
package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"kcardDesgin/backend/internal/domain"
)

// MaterialRepository 处理源资料数据的持久化操作。
type MaterialRepository struct{ db *gorm.DB }

// NewMaterialRepository 创建 MaterialRepository 实例。
func NewMaterialRepository(db *gorm.DB) MaterialRepository { return MaterialRepository{db: db} }

// Create 创建新的源资料记录。
func (r MaterialRepository) Create(ctx context.Context, material *SourceMaterialModel) error {
	if material.ID == "" {
		material.ID = uuid.NewString()
	}
	return r.db.WithContext(ctx).Create(material).Error
}

// FindByDigest 根据内容摘要查询指定工作区中的源资料。
func (r MaterialRepository) FindByDigest(ctx context.Context, workspaceID, digest string) ([]SourceMaterialModel, error) {
	var items []SourceMaterialModel
	err := r.db.WithContext(ctx).Where("learner_workspace_id = ? AND content_digest = ? AND archived_at IS NULL", workspaceID, digest).Find(&items).Error
	return items, err
}

// Get 根据 ID 获取指定工作区中的源资料。
func (r MaterialRepository) Get(ctx context.Context, workspaceID, id string) (SourceMaterialModel, error) {
	var material SourceMaterialModel
	err := r.db.WithContext(ctx).Where("learner_workspace_id = ? AND id = ?", workspaceID, id).First(&material).Error
	return material, err
}

// SaveMaterial 保存源资料领域对象到数据库。
func (r MaterialRepository) SaveMaterial(ctx context.Context, material domain.SourceMaterial) (domain.SourceMaterial, error) {
	model := sourceMaterialToModel(material)
	if model.ID == "" {
		model.ID = uuid.NewString()
		material.ID = domain.ID(model.ID)
	}
	if err := r.db.WithContext(ctx).Save(&model).Error; err != nil {
		return domain.SourceMaterial{}, err
	}
	return sourceMaterialFromModel(model), nil
}

// FindMaterialDigests 根据摘要查找工作区中的源资料领域对象列表。
func (r MaterialRepository) FindMaterialDigests(ctx context.Context, workspaceID domain.ID, digest string) ([]domain.SourceMaterial, error) {
	models, err := r.FindByDigest(ctx, string(workspaceID), digest)
	if err != nil {
		return nil, err
	}
	materials := make([]domain.SourceMaterial, 0, len(models))
	for _, model := range models {
		materials = append(materials, sourceMaterialFromModel(model))
	}
	return materials, nil
}

// GetMaterial 根据领域 ID 获取源资料领域对象。
func (r MaterialRepository) GetMaterial(ctx context.Context, workspaceID domain.ID, id domain.ID) (domain.SourceMaterial, error) {
	model, err := r.Get(ctx, string(workspaceID), string(id))
	if err != nil {
		return domain.SourceMaterial{}, err
	}
	return sourceMaterialFromModel(model), nil
}

// CreateMaterialVersion 创建新的资料版本并更新源资料的当前版本 ID。
func (r MaterialRepository) CreateMaterialVersion(ctx context.Context, version domain.MaterialVersion) (domain.MaterialVersion, error) {
	if version.ID == "" {
		version.ID = domain.ID(uuid.NewString())
	}
	model := materialVersionToModel(version)
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&model).Error; err != nil {
			return err
		}
		return tx.Model(&SourceMaterialModel{}).Where("id = ?", string(version.SourceMaterialID)).Update("current_version_id", model.ID).Error
	})
	if err != nil {
		return domain.MaterialVersion{}, err
	}
	return materialVersionFromModel(model), nil
}

// GetCurrentMaterialVersion 获取指定源资料的当前版本。
func (r MaterialRepository) GetCurrentMaterialVersion(ctx context.Context, workspaceID domain.ID, materialID domain.ID) (domain.MaterialVersion, error) {
	material, err := r.Get(ctx, string(workspaceID), string(materialID))
	if err != nil {
		return domain.MaterialVersion{}, err
	}
	if material.CurrentVersionID == nil || *material.CurrentVersionID == "" {
		return domain.MaterialVersion{}, errors.New("material has no current version")
	}
	var version MaterialVersionModel
	if err := r.db.WithContext(ctx).Where("id = ? AND source_material_id = ?", *material.CurrentVersionID, string(materialID)).First(&version).Error; err != nil {
		return domain.MaterialVersion{}, err
	}
	return materialVersionFromModel(version), nil
}

// UpdateMaterialStatus 更新源资料的处理状态和失败原因。
func (r MaterialRepository) UpdateMaterialStatus(ctx context.Context, workspaceID domain.ID, id domain.ID, status domain.ProcessingStatus, failureReason string) error {
	return r.db.WithContext(ctx).Model(&SourceMaterialModel{}).Where("learner_workspace_id = ? AND id = ?", string(workspaceID), string(id)).Updates(map[string]any{"processing_status": string(status), "failure_reason": failureReason}).Error
}

func sourceMaterialFromModel(model SourceMaterialModel) domain.SourceMaterial {
	return domain.SourceMaterial{ID: domain.ID(model.ID), LearnerWorkspaceID: domain.ID(model.LearnerWorkspaceID), SourceType: domain.SourceType(model.SourceType), Title: model.Title, OriginalLocation: model.OriginalLocation, ContentDigest: model.ContentDigest, ContentSummary: model.ContentSummary, ProcessingStatus: domain.ProcessingStatus(model.ProcessingStatus), FailureReason: model.FailureReason, DuplicateStatus: domain.DuplicateStatus(model.DuplicateStatus), Timestamps: domain.Timestamps{CreatedAt: model.CreatedAt, UpdatedAt: model.UpdatedAt}, ArchiveState: domain.ArchiveState{ArchivedAt: model.ArchivedAt}}
}

func sourceMaterialToModel(material domain.SourceMaterial) SourceMaterialModel {
	return SourceMaterialModel{BaseModel: BaseModel{ID: string(material.ID), CreatedAt: material.CreatedAt, UpdatedAt: material.UpdatedAt, ArchivedAt: material.ArchivedAt}, LearnerWorkspaceID: string(material.LearnerWorkspaceID), SourceType: string(material.SourceType), Title: material.Title, OriginalLocation: material.OriginalLocation, ContentDigest: material.ContentDigest, ContentSummary: material.ContentSummary, ContentStatus: "available", ProcessingStatus: string(material.ProcessingStatus), FailureReason: material.FailureReason, DuplicateStatus: string(material.DuplicateStatus)}
}

func materialVersionFromModel(model MaterialVersionModel) domain.MaterialVersion {
	return domain.MaterialVersion{ID: domain.ID(model.ID), SourceMaterialID: domain.ID(model.SourceMaterialID), VersionNumber: model.VersionNumber, ContentDigest: model.ContentDigest, ContentLocation: model.ContentLocation, Summary: model.Summary, CreatedByAction: model.CreatedByAction}
}

func materialVersionToModel(version domain.MaterialVersion) MaterialVersionModel {
	return MaterialVersionModel{BaseModel: BaseModel{ID: string(version.ID)}, SourceMaterialID: string(version.SourceMaterialID), VersionNumber: version.VersionNumber, ContentDigest: version.ContentDigest, ContentLocation: version.ContentLocation, Summary: version.Summary, CreatedByAction: version.CreatedByAction}
}
