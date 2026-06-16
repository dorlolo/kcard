// Package repository 提供数据模型定义和数据库访问接口的实现。
package repository

import "time"

// BaseModel 表示基础模型，包含 ID、创建时间和更新时间的公共字段。
type BaseModel struct {
	ID         string `gorm:"primaryKey;type:uuid"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	ArchivedAt *time.Time
}

// LearnerWorkspaceModel 表示学习者工作区模型。
type LearnerWorkspaceModel struct {
	BaseModel
	DisplayName   string
	OwnerIdentity string
	PrivacyState  string
}
// LearnerPreferenceModel 表示学习者偏好模型。
type LearnerPreferenceModel struct {
	BaseModel
	LearnerWorkspaceID   string
	DailyCapacityDefault int
	ReviewGradingStyle   string
	Timezone             string
	VisualThemePaletteID string
}
// VisualThemePaletteModel 表示视觉主题调色板模型。
type VisualThemePaletteModel struct {
	BaseModel
	Name                 string
	PrimaryBackground    string
	AccentBackground1    string
	AccentBackground2    string
	SemanticWarningColor string
	SemanticErrorColor   string
	SemanticSuccessColor string
	ReadabilityNotes     string
}
// TagModel 表示标签模型。
type TagModel struct {
	BaseModel
	LearnerWorkspaceID string
	Name               string
	Color              string
	Description        string
}
// SourceMaterialModel 表示源资料模型。
type SourceMaterialModel struct {
	BaseModel
	LearnerWorkspaceID string
	SourceType         string
	Title              string
	OriginalLocation   string
	ContentDigest      string
	ContentSummary     string
	ContentStatus      string
	ProcessingStatus   string
	FailureReason      string
	DuplicateStatus    string
	CurrentVersionID   *string
}
// MaterialVersionModel 表示资料版本模型。
type MaterialVersionModel struct {
	BaseModel
	SourceMaterialID string
	VersionNumber    int
	ContentDigest    string
	ContentLocation  string
	Summary          string
	CreatedByAction  string
}
// KnowledgePointModel 表示知识点模型。
type KnowledgePointModel struct {
	BaseModel
	LearnerWorkspaceID string
	SourceMaterialID   *string
	MaterialVersionID  *string
	Content            string
	Summary            string
	Notes              string
	ApprovalStatus     string
	CreationSource     string
	DuplicateGroupID   *string
	GraphLabel         string
	GraphPositionHint  string
	AIJobID            *string
	PromptSnapshotID   *string
	ApprovedAt         *time.Time
	RejectedAt         *time.Time
}
// KnowledgeRelationshipModel 表示知识关系模型。
type KnowledgeRelationshipModel struct {
	BaseModel
	LearnerWorkspaceID     string
	SourceKnowledgePointID string
	TargetKnowledgePointID string
	RelationshipType       string
	Label                  string
	Weight                 float64
	SourceType             string
	SourceMaterialID       string
	TagID                  string
	CardID                 string
	Confidence             float64
}

// TableName 返回 learner_workspaces 表名。
func (LearnerWorkspaceModel) TableName() string      { return "learner_workspaces" }
// TableName 返回 learner_preferences 表名。
func (LearnerPreferenceModel) TableName() string     { return "learner_preferences" }
// TableName 返回 visual_theme_palettes 表名。
func (VisualThemePaletteModel) TableName() string    { return "visual_theme_palettes" }
// TableName 返回 tags 表名。
func (TagModel) TableName() string                   { return "tags" }
// TableName 返回 source_materials 表名。
func (SourceMaterialModel) TableName() string        { return "source_materials" }
// TableName 返回 material_versions 表名。
func (MaterialVersionModel) TableName() string       { return "material_versions" }
// TableName 返回 knowledge_points 表名。
func (KnowledgePointModel) TableName() string        { return "knowledge_points" }
// TableName 返回 knowledge_relationships 表名。
func (KnowledgeRelationshipModel) TableName() string { return "knowledge_relationships" }

// PromptPresetModel 表示提示预设模型。
type PromptPresetModel struct {
	BaseModel
	LearnerWorkspaceID string
	Name               string
	Purpose            string
	PromptText         string
	IsDefault          bool
	VersionNumber      int
}

// TableName 返回 prompt_presets 表名。
func (PromptPresetModel) TableName() string { return "prompt_presets" }

// PromptSnapshotModel 表示提示快照模型。
type PromptSnapshotModel struct {
	BaseModel
	PromptPresetID  string
	Purpose         string
	PromptText      string
	ModelID         string
	SchemaVersion   string
	CreatedForJobID string
}

// TableName 返回 prompt_snapshots 表名。
func (PromptSnapshotModel) TableName() string { return "prompt_snapshots" }

// AIJobModel 表示 AI 任务模型。
type AIJobModel struct {
	BaseModel
	LearnerWorkspaceID string
	JobType            string
	Status             string
	ProgressPercent    int
	CurrentStep        string
	InputSummary       string
	ErrorCategory      string
	ErrorMessage       string
	IdempotencyKey     string
	StartedAt          *time.Time
	FinishedAt         *time.Time
}

// TableName 返回 ai_jobs 表名。
func (AIJobModel) TableName() string { return "ai_jobs" }

// AIDraftModel 表示 AI 草稿模型。
type AIDraftModel struct {
	BaseModel
	LearnerWorkspaceID string
	DraftType          string
	JobID              string
	Payload            []byte
	ValidationStatus   string
	ValidationMessages []byte
	Status             string
	ApprovedRecordID   string
}

// TableName 返回 ai_drafts 表名。
func (AIDraftModel) TableName() string { return "ai_drafts" }
