package domain

type PrivacyState string

const PrivacyPrivate PrivacyState = "private"

type LearnerWorkspace struct {
	ID                          ID           `json:"id"`
	DisplayName                 string       `json:"displayName"`
	OwnerIdentity               string       `json:"ownerIdentity"`
	DefaultReviewCapacityPerDay int          `json:"defaultReviewCapacityPerDay"`
	DefaultReviewGradingStyle   string       `json:"defaultReviewGradingStyle"`
	PrivacyState                PrivacyState `json:"privacyState"`
	Timestamps
}

type LearnerPreference struct {
	ID                   ID     `json:"id"`
	LearnerWorkspaceID   ID     `json:"learnerWorkspaceId"`
	DailyCapacityDefault int    `json:"dailyCapacityDefault"`
	ReviewGradingStyle   string `json:"reviewGradingStyle"`
	Timezone             string `json:"timezone"`
	VisualThemePaletteID ID     `json:"visualThemePaletteId"`
}
