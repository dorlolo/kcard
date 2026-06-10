package domain

import "time"

type ID string

type WorkspaceScoped struct {
	LearnerWorkspaceID ID `json:"learnerWorkspaceId"`
}

type Timestamps struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ArchiveState struct {
	ArchivedAt *time.Time `json:"archivedAt,omitempty"`
}

func (a ArchiveState) IsArchived() bool { return a.ArchivedAt != nil }
