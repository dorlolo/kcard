// Package domain 提供复习会话相关类型。
package domain

import "time"

// ReviewSession 表示一次复习会话。
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
// ReviewResult 表示单张卡片的复习结果。
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
