// Package domain 提供提示词预设相关类型。
package domain

// PromptPreset 表示用户定义的提示词预设。
type PromptPreset struct {
	ID                 ID
	LearnerWorkspaceID ID
	Name               string
	Purpose            string
	PromptText         string
	IsDefault          bool
	VersionNumber      int
}
// PromptSnapshot 表示提示词在某个时刻的快照。
type PromptSnapshot struct {
	ID              ID
	PromptPresetID  ID
	Purpose         string
	PromptText      string
	ModelID         string
	SchemaVersion   string
	CreatedForJobID ID
}
