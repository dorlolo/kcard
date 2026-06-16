// Package domain 提供领域模型的核心类型定义。
package domain

import "time"

// ID 表示领域实体唯一标识符。
type ID string

// WorkspaceScoped 表示属于某个工作区的实体。
type WorkspaceScoped struct {
	LearnerWorkspaceID ID `json:"learnerWorkspaceId"`
}

// Timestamps 包含创建时间和更新时间。
type Timestamps struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// ArchiveState 表示归档状态，包含归档时间。
type ArchiveState struct {
	ArchivedAt *time.Time `json:"archivedAt,omitempty"`
}

// IsArchived 判断当前实体是否已归档。
func (a ArchiveState) IsArchived() bool { return a.ArchivedAt != nil }
