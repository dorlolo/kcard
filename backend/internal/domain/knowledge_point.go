package domain

import "time"

type ApprovalStatus string

const (
	KnowledgeDraft       ApprovalStatus = "draft"
	KnowledgeApproved    ApprovalStatus = "approved"
	KnowledgeRejected    ApprovalStatus = "rejected"
	KnowledgeNeedsReview ApprovalStatus = "needs_review"
)

type CreationSource string

const (
	CreationAIGenerated CreationSource = "ai_generated"
	CreationManual      CreationSource = "manual"
	CreationImported    CreationSource = "imported"
)

type RelationshipType string

const (
	RelationshipRelated      RelationshipType = "related"
	RelationshipPrerequisite RelationshipType = "prerequisite"
	RelationshipDuplicate    RelationshipType = "duplicate"
	RelationshipSimilar      RelationshipType = "similar"
	RelationshipSplitFrom    RelationshipType = "split_from"
	RelationshipMergedFrom   RelationshipType = "merged_from"
	RelationshipSupports     RelationshipType = "supports"
	RelationshipContradicts  RelationshipType = "contradicts"
)

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

type KnowledgeFilter struct {
	WorkspaceID       ID
	Query             string
	ApprovalStatus    ApprovalStatus
	Tag               string
	IncludeRejected   bool
	IncludeArchived   bool
	IncludeUnapproved bool
}

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

func (k KnowledgePoint) Approve(now time.Time) KnowledgePoint {
	k.ApprovalStatus = KnowledgeApproved
	k.ApprovedAt = &now
	k.RejectedAt = nil
	return k
}
func (k KnowledgePoint) Reject(now time.Time) KnowledgePoint {
	k.ApprovalStatus = KnowledgeRejected
	k.RejectedAt = &now
	k.ApprovedAt = nil
	return k
}
func (k KnowledgePoint) MarkNeedsReview() KnowledgePoint {
	k.ApprovalStatus = KnowledgeNeedsReview
	return k
}
