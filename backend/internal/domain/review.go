package domain

import "time"

type ReviewSession struct {
	ID                 ID
	LearnerWorkspaceID ID
	SessionType        string
	DeckID             ID
	ReviewPlanID       ID
	Status             string
	StartedAt          time.Time
	Summary            string
}
type ReviewResult struct {
	ID                ID
	ReviewSessionID   ID
	CardID            ID
	Result            string
	Confidence        int
	ElapsedResponseMS int
	ReviewedAt        time.Time
	NextDueAt         time.Time
	ScheduleReason    string
}
