package service

import (
	"time"
)

type ReviewScheduler struct{}

func (ReviewScheduler) NextDue(result string, reviewedAt time.Time) time.Time {
	switch result {
	case "again", "incorrect":
		return reviewedAt.Add(10 * time.Minute)
	case "hard":
		return reviewedAt.Add(24 * time.Hour)
	case "easy":
		return reviewedAt.Add(7 * 24 * time.Hour)
	default:
		return reviewedAt.Add(3 * 24 * time.Hour)
	}
}
