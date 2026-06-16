// Package domain 提供学习者工作区和偏好相关类型。
package domain

// PrivacyState 表示工作区隐私状态。
type PrivacyState string

// PrivacyPrivate 常量表示工作区为私有状态。
const PrivacyPrivate PrivacyState = "private"

// LearnerWorkspace 表示学习者工作区。
type LearnerWorkspace struct {
	ID                          ID           `json:"id"`
	DisplayName                 string       `json:"displayName"`
	OwnerIdentity               string       `json:"ownerIdentity"`
	DefaultReviewCapacityPerDay int          `json:"defaultReviewCapacityPerDay"`
	DefaultReviewGradingStyle   string       `json:"defaultReviewGradingStyle"`
	PrivacyState                PrivacyState `json:"privacyState"`
	Timestamps
}

// LearnerPreference 表示学习者偏好设置。
type LearnerPreference struct {
	ID                   ID     `json:"id"`
	LearnerWorkspaceID   ID     `json:"learnerWorkspaceId"`
	DailyCapacityDefault int    `json:"dailyCapacityDefault"`
	ReviewGradingStyle   string `json:"reviewGradingStyle"`
	Timezone             string `json:"timezone"`
	VisualThemePaletteID ID     `json:"visualThemePaletteId"`
}
