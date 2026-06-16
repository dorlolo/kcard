// Package domain 提供知识点相关领域类型。
package domain

import "time"

// ApprovalStatus 表示知识点的审批状态。
type ApprovalStatus string

const (
	// KnowledgeDraft 常量表示知识点为草稿状态。
	KnowledgeDraft ApprovalStatus = "draft"
	// KnowledgeApproved 常量表示知识点已通过审批。
	KnowledgeApproved ApprovalStatus = "approved"
	// KnowledgeRejected 常量表示知识点已被驳回。
	KnowledgeRejected ApprovalStatus = "rejected"
	// KnowledgeNeedsReview 常量表示知识点需要进一步审核。
	KnowledgeNeedsReview ApprovalStatus = "needs_review"
)

// CreationSource 表示知识点的创建来源。
type CreationSource string

const (
	// CreationAIGenerated 常量表示由 AI 生成。
	CreationAIGenerated CreationSource = "ai_generated"
	// CreationManual 常量表示手动创建。
	CreationManual CreationSource = "manual"
	// CreationImported 常量表示从外部导入。
	CreationImported CreationSource = "imported"
)

// RelationshipType 表示知识点之间的关系类型。
type RelationshipType string

const (
	// RelationshipRelated 常量表示知识点之间相互关联。
	RelationshipRelated RelationshipType = "related"
	// RelationshipPrerequisite 常量表示一个知识点是另一个的前置条件。
	RelationshipPrerequisite RelationshipType = "prerequisite"
	// RelationshipDuplicate 常量表示知识点之间存在重复。
	RelationshipDuplicate RelationshipType = "duplicate"
	// RelationshipSimilar 常量表示知识点之间内容相似。
	RelationshipSimilar RelationshipType = "similar"
	// RelationshipSplitFrom 常量表示知识点从另一个拆分而来。
	RelationshipSplitFrom RelationshipType = "split_from"
	// RelationshipMergedFrom 常量表示知识点由多个合并而来。
	RelationshipMergedFrom RelationshipType = "merged_from"
	// RelationshipSupports 常量表示知识点支持另一个知识点。
	RelationshipSupports RelationshipType = "supports"
	// RelationshipContradicts 常量表示知识点与另一个相矛盾。
	RelationshipContradicts RelationshipType = "contradicts"
)

// KnowledgeRelationship 表示知识点之间的关联关系。
type KnowledgeRelationship struct {
	ID          ID
	WorkspaceID ID
	SourceID    ID
	TargetID    ID
	Type        RelationshipType
	Label       string
	Weight      float64
	SourceType  string
	Archived    bool
}

// KnowledgeFilter 定义知识点的查询过滤条件。
type KnowledgeFilter struct {
	WorkspaceID       ID
	Query             string
	ApprovalStatus    ApprovalStatus
	Tag               string
	IncludeRejected   bool
	IncludeArchived   bool
	IncludeUnapproved bool
}

// KnowledgePoint 表示一个知识点。
type KnowledgePoint struct {
	ID                 ID             `json:"id"`
	LearnerWorkspaceID ID             `json:"learnerWorkspaceId"`
	SourceMaterialID   ID             `json:"sourceMaterialId,omitempty"`
	MaterialVersionID  ID             `json:"materialVersionId,omitempty"`
	Content            string         `json:"content"`
	Summary            string         `json:"summary,omitempty"`
	Notes              string         `json:"notes,omitempty"`
	ApprovalStatus     ApprovalStatus `json:"approvalStatus"`
	CreationSource     CreationSource `json:"creationSource"`
	DuplicateGroupID   ID             `json:"duplicateGroupId,omitempty"`
	GraphLabel         string         `json:"graphLabel,omitempty"`
	ApprovedAt         *time.Time     `json:"approvedAt,omitempty"`
	RejectedAt         *time.Time     `json:"rejectedAt,omitempty"`
	Tags               []Tag          `json:"tags,omitempty"`
}

// Approve 审批通过知识点，记录通过时间。
func (k KnowledgePoint) Approve(now time.Time) KnowledgePoint {
	k.ApprovalStatus = KnowledgeApproved
	k.ApprovedAt = &now
	k.RejectedAt = nil
	return k
}
// Reject 驳回知识点，记录驳回时间。
func (k KnowledgePoint) Reject(now time.Time) KnowledgePoint {
	k.ApprovalStatus = KnowledgeRejected
	k.RejectedAt = &now
	k.ApprovedAt = nil
	return k
}
// MarkNeedsReview 将知识点标记为需要进一步审核。
func (k KnowledgePoint) MarkNeedsReview() KnowledgePoint {
	k.ApprovalStatus = KnowledgeNeedsReview
	return k
}
