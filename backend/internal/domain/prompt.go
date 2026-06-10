package domain

type PromptPreset struct {
	ID                 ID
	LearnerWorkspaceID ID
	Name               string
	Purpose            string
	PromptText         string
	IsDefault          bool
	VersionNumber      int
}
type PromptSnapshot struct {
	ID              ID
	PromptPresetID  ID
	Purpose         string
	PromptText      string
	ModelID         string
	SchemaVersion   string
	CreatedForJobID ID
}
