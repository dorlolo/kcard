package domain

type CardStatus string

const (
	CardDraft    CardStatus = "draft"
	CardActive   CardStatus = "active"
	CardArchived CardStatus = "archived"
	CardDeleted  CardStatus = "deleted"
)

type Deck struct {
	ID                 ID             `json:"id"`
	LearnerWorkspaceID ID             `json:"learnerWorkspaceId"`
	Name               string         `json:"name"`
	Description        string         `json:"description"`
	CreationSource     CreationSource `json:"creationSource"`
	Status             CardStatus     `json:"status"`
	Tags               []Tag          `json:"tags,omitempty"`
}
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
type CardSourceLink struct {
	ID               ID
	CardID           ID
	KnowledgePointID ID
	SourceMaterialID ID
	SourceQuote      string
}
type AIDraft struct {
	ID        ID
	DraftType string
	Status    string
	Payload   map[string]any
}
