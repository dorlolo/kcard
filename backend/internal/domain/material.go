package domain

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strings"
)

type SourceType string

const (
	SourceFile    SourceType = "file"
	SourceWebPage SourceType = "web_page"
	SourceText    SourceType = "text"
)

type ProcessingStatus string

const (
	MaterialDraft       ProcessingStatus = "draft"
	MaterialQueued      ProcessingStatus = "queued"
	MaterialProcessing  ProcessingStatus = "processing"
	MaterialNeedsReview ProcessingStatus = "needs_review"
	MaterialProcessed   ProcessingStatus = "processed"
	MaterialFailed      ProcessingStatus = "failed"
)

type DuplicateStatus string

const (
	DuplicateUnchecked DuplicateStatus = "unchecked"
	DuplicatePossible  DuplicateStatus = "possible_duplicate"
	DuplicateConfirmed DuplicateStatus = "confirmed_duplicate"
	DuplicateUnique    DuplicateStatus = "unique"
)

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

type MaterialVersion struct {
	ID               ID     `json:"id"`
	SourceMaterialID ID     `json:"sourceMaterialId"`
	VersionNumber    int    `json:"versionNumber"`
	ContentDigest    string `json:"contentDigest"`
	ContentLocation  string `json:"contentLocation"`
	Summary          string `json:"summary"`
	CreatedByAction  string `json:"createdByAction"`
}

type Tag struct {
	ID    ID     `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color,omitempty"`
}

func NewTextMaterial(workspaceID ID, title, text string, tags []Tag) (SourceMaterial, error) {
	if strings.TrimSpace(text) == "" {
		return SourceMaterial{}, errors.New("text material cannot be empty")
	}
	if strings.TrimSpace(title) == "" {
		title = "Untitled text material"
	}
	return SourceMaterial{LearnerWorkspaceID: workspaceID, SourceType: SourceText, Title: title, OriginalLocation: "pasted text", ContentDigest: DigestText(text), ProcessingStatus: MaterialDraft, DuplicateStatus: DuplicateUnchecked, Tags: tags}, nil
}

func DigestText(text string) string {
	sum := sha256.Sum256([]byte(strings.TrimSpace(text)))
	return hex.EncodeToString(sum[:])
}

func (m SourceMaterial) Queue(duplicate DuplicateStatus) SourceMaterial {
	m.ProcessingStatus = MaterialQueued
	m.DuplicateStatus = duplicate
	return m
}
func (m SourceMaterial) Fail(reason string) SourceMaterial {
	m.ProcessingStatus = MaterialFailed
	m.FailureReason = reason
	return m
}
func (m SourceMaterial) ReadyForReview() SourceMaterial {
	m.ProcessingStatus = MaterialNeedsReview
	return m
}
