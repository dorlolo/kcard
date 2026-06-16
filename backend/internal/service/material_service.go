// Package service 提供领域服务层的实现，包括资料、卡片、知识图谱等核心业务逻辑。
package service

import (
	"context"
	"errors"
	"strings"

	"github.com/google/uuid"

	"kcardDesgin/backend/internal/domain"
)

// MaterialStore 定义资料存储的持久化接口。
type MaterialStore interface {
	SaveMaterial(ctx context.Context, material domain.SourceMaterial) (domain.SourceMaterial, error)
	FindMaterialDigests(ctx context.Context, workspaceID domain.ID, digest string) ([]domain.SourceMaterial, error)
	GetMaterial(ctx context.Context, workspaceID domain.ID, id domain.ID) (domain.SourceMaterial, error)
	CreateMaterialVersion(ctx context.Context, version domain.MaterialVersion) (domain.MaterialVersion, error)
	UpdateMaterialStatus(ctx context.Context, workspaceID domain.ID, id domain.ID, status domain.ProcessingStatus, failureReason string) error
}

// MaterialAnalysisRequest 表示资料分析任务的请求参数。
type MaterialAnalysisRequest struct {
	WorkspaceID       domain.ID
	MaterialID        domain.ID
	MaterialVersionID domain.ID
	Text              string
	ContentLocation   string
	Prompt            string
}

// JobEnqueuer 定义将任务加入队列的接口。
type JobEnqueuer interface {
	EnqueueMaterialAnalysis(ctx context.Context, input MaterialAnalysisRequest) (string, error)
}

// MaterialService 处理资料导入和管理的核心业务逻辑。
type MaterialService struct {
	Store MaterialStore
	Jobs  JobEnqueuer
}

// CreateMaterialInput 表示创建资料请求的输入参数。
type CreateMaterialInput struct {
	WorkspaceID     domain.ID
	SourceType      domain.SourceType
	Title           string
	Text            string
	URL             string
	Tags            []domain.Tag
	PromptText      string
	DuplicatePolicy string
}
// CreateMaterialResult 表示创建资料请求的返回结果。
type CreateMaterialResult struct {
	Material         domain.SourceMaterial
	JobID            string
	DuplicateWarning bool
}

// Create 执行资料创建流程，包括去重检查和异步分析任务入队。
func (s MaterialService) Create(ctx context.Context, input CreateMaterialInput) (CreateMaterialResult, error) {
	if input.WorkspaceID == "" {
		return CreateMaterialResult{}, errors.New("workspace id is required")
	}
	if input.SourceType != domain.SourceText {
		return CreateMaterialResult{}, errors.New("only text material intake is implemented in the MVP skeleton")
	}
	material, err := domain.NewTextMaterial(input.WorkspaceID, input.Title, input.Text, input.Tags)
	if err != nil {
		return CreateMaterialResult{}, err
	}
	material.ID = domain.ID(uuid.NewString())
	matches, err := s.Store.FindMaterialDigests(ctx, input.WorkspaceID, material.ContentDigest)
	if err != nil {
		return CreateMaterialResult{}, err
	}
	duplicate := len(matches) > 0
	if duplicate && strings.EqualFold(input.DuplicatePolicy, "warn") {
		material = material.Queue(domain.DuplicatePossible)
	} else {
		material = material.Queue(domain.DuplicateUnique)
	}
	saved, err := s.Store.SaveMaterial(ctx, material)
	if err != nil {
		return CreateMaterialResult{}, err
	}
	version := domain.MaterialVersion{
		ID:               domain.ID(uuid.NewString()),
		SourceMaterialID: saved.ID,
		VersionNumber:    1,
		ContentDigest:    saved.ContentDigest,
		ContentLocation:  "inline:" + string(saved.ID),
		Summary:          saved.ContentSummary,
		CreatedByAction:  "initial_import",
	}
	version, err = s.Store.CreateMaterialVersion(ctx, version)
	if err != nil {
		return CreateMaterialResult{}, err
	}
	jobID := ""
	if s.Jobs != nil {
		jobID, err = s.Jobs.EnqueueMaterialAnalysis(ctx, MaterialAnalysisRequest{WorkspaceID: input.WorkspaceID, MaterialID: saved.ID, MaterialVersionID: version.ID, Text: input.Text, ContentLocation: version.ContentLocation, Prompt: input.PromptText})
		if err != nil {
			return CreateMaterialResult{}, err
		}
	}
	return CreateMaterialResult{Material: saved, JobID: jobID, DuplicateWarning: duplicate}, nil
}

// Get 根据工作区和资料 ID 获取单个资料。
func (s MaterialService) Get(ctx context.Context, workspaceID domain.ID, materialID domain.ID) (domain.SourceMaterial, error) {
	return s.Store.GetMaterial(ctx, workspaceID, materialID)
}

// MemoryMaterialStore 实现 MaterialStore 接口的内存存储，用于测试和开发。
type MemoryMaterialStore struct {
	Materials []domain.SourceMaterial
	Versions  []domain.MaterialVersion
}

// SaveMaterial 保存资料到内存存储，若已存在则更新。
func (m *MemoryMaterialStore) SaveMaterial(ctx context.Context, material domain.SourceMaterial) (domain.SourceMaterial, error) {
	for index, existing := range m.Materials {
		if existing.ID == material.ID {
			m.Materials[index] = material
			return material, nil
		}
	}
	m.Materials = append(m.Materials, material)
	return material, nil
}
// FindMaterialDigests 根据工作区和内容摘要查找匹配的资料列表。
func (m *MemoryMaterialStore) FindMaterialDigests(ctx context.Context, workspaceID domain.ID, digest string) ([]domain.SourceMaterial, error) {
	var out []domain.SourceMaterial
	for _, item := range m.Materials {
		if item.LearnerWorkspaceID == workspaceID && item.ContentDigest == digest {
			out = append(out, item)
		}
	}
	return out, nil
}
// GetMaterial 从内存存储中根据工作区和 ID 获取单个资料。
func (m *MemoryMaterialStore) GetMaterial(ctx context.Context, workspaceID domain.ID, id domain.ID) (domain.SourceMaterial, error) {
	for _, item := range m.Materials {
		if item.LearnerWorkspaceID == workspaceID && item.ID == id {
			return item, nil
		}
	}
	return domain.SourceMaterial{}, errors.New("material not found")
}
// CreateMaterialVersion 创建资料版本记录到内存存储。
func (m *MemoryMaterialStore) CreateMaterialVersion(ctx context.Context, version domain.MaterialVersion) (domain.MaterialVersion, error) {
	m.Versions = append(m.Versions, version)
	return version, nil
}
// UpdateMaterialStatus 更新资料的处理状态和失败原因。
func (m *MemoryMaterialStore) UpdateMaterialStatus(ctx context.Context, workspaceID domain.ID, id domain.ID, status domain.ProcessingStatus, failureReason string) error {
	material, err := m.GetMaterial(ctx, workspaceID, id)
	if err != nil {
		return err
	}
	material.ProcessingStatus = status
	material.FailureReason = failureReason
	_, err = m.SaveMaterial(ctx, material)
	return err
}

// MemoryJobEnqueuer 实现 JobEnqueuer 接口的内存队列，用于测试和开发。
type MemoryJobEnqueuer struct{ Jobs []MaterialAnalysisRequest }

// EnqueueMaterialAnalysis 将资料分析任务加入内存队列并返回任务 ID。
func (m *MemoryJobEnqueuer) EnqueueMaterialAnalysis(ctx context.Context, input MaterialAnalysisRequest) (string, error) {
	m.Jobs = append(m.Jobs, input)
	return string(input.MaterialID) + ":analysis", nil
}
