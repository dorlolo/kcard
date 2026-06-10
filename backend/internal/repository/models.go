package repository

import "time"

type BaseModel struct {
	ID         string `gorm:"primaryKey;type:uuid"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	ArchivedAt *time.Time
}

type LearnerWorkspaceModel struct {
	BaseModel
	DisplayName   string
	OwnerIdentity string
	PrivacyState  string
}
type LearnerPreferenceModel struct {
	BaseModel
	LearnerWorkspaceID   string
	DailyCapacityDefault int
	ReviewGradingStyle   string
	Timezone             string
	VisualThemePaletteID string
}
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
type TagModel struct {
	BaseModel
	LearnerWorkspaceID string
	Name               string
	Color              string
	Description        string
}
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
	CurrentVersionID   string
}
type MaterialVersionModel struct {
	BaseModel
	SourceMaterialID string
	VersionNumber    int
	ContentDigest    string
	ContentLocation  string
	Summary          string
	CreatedByAction  string
}
type KnowledgePointModel struct {
	BaseModel
	LearnerWorkspaceID string
	SourceMaterialID   string
	MaterialVersionID  string
	Content            string
	Summary            string
	Notes              string
	ApprovalStatus     string
	CreationSource     string
	DuplicateGroupID   string
	GraphLabel         string
	GraphPositionHint  string
	AIJobID            string
	PromptSnapshotID   string
	ApprovedAt         *time.Time
	RejectedAt         *time.Time
}
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
type PromptPresetModel struct {
	BaseModel
	LearnerWorkspaceID string
	Name               string
	Purpose            string
	PromptText         string
	IsDefault          bool
	VersionNumber      int
}
type PromptSnapshotModel struct {
	BaseModel
	PromptPresetID  string
	Purpose         string
	PromptText      string
	ModelID         string
	SchemaVersion   string
	CreatedForJobID string
}
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
