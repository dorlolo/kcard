// Package domain 提供卡片和卡组相关类型。
package domain

// CardStatus 表示卡片的状态。
type CardStatus string

const (
	// CardDraft 常量表示卡片为草稿状态。
	CardDraft CardStatus = "draft"
	// CardActive 常量表示卡片为活跃状态。
	CardActive CardStatus = "active"
	// CardArchived 常量表示卡片已归档。
	CardArchived CardStatus = "archived"
	// CardDeleted 常量表示卡片已删除。
	CardDeleted CardStatus = "deleted"
)

// Deck 表示一个卡组。
type Deck struct {
	ID                 ID             `json:"id"`
	LearnerWorkspaceID ID             `json:"learnerWorkspaceId"`
	Name               string         `json:"name"`
	Description        string         `json:"description"`
	CreationSource     CreationSource `json:"creationSource"`
	Status             CardStatus     `json:"status"`
	Tags               []Tag          `json:"tags,omitempty"`
}
// Card 表示一张复习卡片。
type Card struct {
	ID                 ID         `json:"id"`
	LearnerWorkspaceID ID         `json:"learnerWorkspaceId"`
	FrontPrompt        string     `json:"frontPrompt"`
	BackAnswer         string     `json:"backAnswer"`
	Explanation        string     `json:"explanation"`
	LearnerNotes       string     `json:"learnerNotes"`
	Difficulty         string     `json:"difficulty"`
	Status             CardStatus `json:"status"`
	ReviewStatus       string     `json:"reviewStatus"`
	DeckIDs            []ID       `json:"deckIds"`
	Tags               []Tag      `json:"tags,omitempty"`
}
// CardSourceLink 表示卡片与资料之间的关联。
type CardSourceLink struct {
	ID               ID
	CardID           ID
	KnowledgePointID ID
	SourceMaterialID ID
	SourceQuote      string
}
// AIDraft 表示 AI 生成的草稿内容。
type AIDraft struct {
	ID        ID
	DraftType string
	Status    string
	Payload   map[string]any
}
