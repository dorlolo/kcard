// Package domain 提供资料素材相关领域类型。
package domain

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strings"
)

// SourceType 表示资料来源类型。
type SourceType string

const (
	// SourceFile 常量表示文件来源。
	SourceFile SourceType = "file"
	// SourceWebPage 常量表示网页来源。
	SourceWebPage SourceType = "web_page"
	// SourceText 常量表示文本来源。
	SourceText SourceType = "text"
)

// ProcessingStatus 表示资料处理状态。
type ProcessingStatus string

const (
	// MaterialDraft 常量表示资料处于草稿状态。
	MaterialDraft ProcessingStatus = "draft"
	// MaterialQueued 常量表示资料已加入处理队列。
	MaterialQueued ProcessingStatus = "queued"
	// MaterialProcessing 常量表示资料正在处理中。
	MaterialProcessing ProcessingStatus = "processing"
	// MaterialNeedsReview 常量表示资料需要人工审核。
	MaterialNeedsReview ProcessingStatus = "needs_review"
	// MaterialProcessed 常量表示资料已处理完成。
	MaterialProcessed ProcessingStatus = "processed"
	// MaterialFailed 常量表示资料处理失败。
	MaterialFailed ProcessingStatus = "failed"
)

// DuplicateStatus 表示资料去重检查状态。
type DuplicateStatus string

const (
	// DuplicateUnchecked 常量表示尚未进行去重检查。
	DuplicateUnchecked DuplicateStatus = "unchecked"
	// DuplicatePossible 常量表示可能存在重复。
	DuplicatePossible DuplicateStatus = "possible_duplicate"
	// DuplicateConfirmed 常量表示已确认重复。
	DuplicateConfirmed DuplicateStatus = "confirmed_duplicate"
	// DuplicateUnique 常量表示确认为唯一资料。
	DuplicateUnique DuplicateStatus = "unique"
)

// SourceMaterial 表示一份来源资料。
type SourceMaterial struct {
	ID                 ID               `json:"id"`
	LearnerWorkspaceID ID               `json:"learnerWorkspaceId"`
	SourceType         SourceType       `json:"sourceType"`
	Title              string           `json:"title"`
	OriginalLocation   string           `json:"originalLocation"`
	ContentDigest      string           `json:"contentDigest"`
	ContentSummary     string           `json:"contentSummary"`
	ProcessingStatus   ProcessingStatus `json:"processingStatus"`
	FailureReason      string           `json:"failureReason,omitempty"`
	DuplicateStatus    DuplicateStatus  `json:"duplicateStatus"`
	Tags               []Tag            `json:"tags,omitempty"`
	Timestamps
	ArchiveState
}

// MaterialVersion 表示资料的历史版本。
type MaterialVersion struct {
	ID               ID     `json:"id"`
	SourceMaterialID ID     `json:"sourceMaterialId"`
	VersionNumber    int    `json:"versionNumber"`
	ContentDigest    string `json:"contentDigest"`
	ContentLocation  string `json:"contentLocation"`
	Summary          string `json:"summary"`
	CreatedByAction  string `json:"createdByAction"`
}

// Tag 表示资料的标签。
type Tag struct {
	ID    ID     `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color,omitempty"`
}

// NewTextMaterial 创建一份新的文本来源资料。
func NewTextMaterial(workspaceID ID, title, text string, tags []Tag) (SourceMaterial, error) {
	if strings.TrimSpace(text) == "" {
		return SourceMaterial{}, errors.New("text material cannot be empty")
	}
	if strings.TrimSpace(title) == "" {
		title = "Untitled text material"
	}
	return SourceMaterial{LearnerWorkspaceID: workspaceID, SourceType: SourceText, Title: title, OriginalLocation: "pasted text", ContentDigest: DigestText(text), ProcessingStatus: MaterialDraft, DuplicateStatus: DuplicateUnchecked, Tags: tags}, nil
}

// DigestText 返回文本内容的 SHA256 摘要。
func DigestText(text string) string {
	sum := sha256.Sum256([]byte(strings.TrimSpace(text)))
	return hex.EncodeToString(sum[:])
}

// Queue 将资料标记为已加入处理队列。
func (m SourceMaterial) Queue(duplicate DuplicateStatus) SourceMaterial {
	m.ProcessingStatus = MaterialQueued
	m.DuplicateStatus = duplicate
	return m
}
// Fail 将资料标记为处理失败，并记录失败原因。
func (m SourceMaterial) Fail(reason string) SourceMaterial {
	m.ProcessingStatus = MaterialFailed
	m.FailureReason = reason
	return m
}
// ReadyForReview 将资料标记为待人工审核。
func (m SourceMaterial) ReadyForReview() SourceMaterial {
	m.ProcessingStatus = MaterialNeedsReview
	return m
}
